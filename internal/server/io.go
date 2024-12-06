package server

import (
	"github.com/olahol/melody"
	"github.com/otie173/odncore/internal/game/player"
	"github.com/otie173/odncore/internal/game/world"
	"github.com/otie173/odncore/internal/utils/logger"
	"github.com/otie173/odncore/internal/utils/typeconv"
	"github.com/vmihailenco/msgpack/v5"
)

var (
	packet map[string]interface{}
)

func handleRequest(session *melody.Session, opcode byte, data []byte) {
	switch opcode {
	case blockPacket:
		if err := msgpack.Unmarshal(data, &packet); err != nil {
			logger.Error("Failed to unmarshal block packet: ", err)
		}

		switch typeconv.GetByte(packet["Action"]) {
		case blockAdd:
			world.AddBlock(
				typeconv.GetUint32(packet["Texture"]),
				typeconv.GetFloat32(packet["X"]),
				typeconv.GetFloat32(packet["Y"]),
				typeconv.GetBool(packet["Passable"]),
			)
			sendBlockPacket(session, blockAdd, typeconv.GetFloat32(packet["X"]), typeconv.GetFloat32(packet["Y"]), typeconv.GetPtrUint32(packet["Texture"]))

		case blockRemove:
			world.RemoveBlock(
				typeconv.GetFloat32(packet["X"]),
				typeconv.GetFloat32(packet["Y"]),
			)
			if err := sendBlockPacket(session, blockRemove, typeconv.GetFloat32(packet["X"]), typeconv.GetFloat32(packet["Y"]), nil); err != nil {
				logger.Error("Failed to send block packet: ", err)
			}
		}
	case playerPacket:
		if err := msgpack.Unmarshal(data, &packet); err != nil {
			logger.Error("Failed to unmarshal player packet: ", err)
		}

		switch typeconv.GetByte(packet["Action"]) {
		case playerMove:
			if err := sendPlayerMove(data); err != nil {
				logger.Error("Failed to send player move packet: ", err)
			}
		case playerAdd:
			player.Add(
				session.Request.Header.Get("Session-Nickname"),
				typeconv.GetFloat32(packet["X"]),
				typeconv.GetFloat32(packet["Y"]),
				typeconv.GetFloat32(packet["TargetX"]),
				typeconv.GetFloat32(packet["TargetY"]),
			)

			if err := sendPlayersList(); err != nil {
				logger.Errorf("Failed to send updated players list: %v", err)
			}
		}
	}
}

func InitHandler() {
	websocket.HandleMessageBinary(func(s *melody.Session, b []byte) {
		handleRequest(s, b[0], b[1:])
	})
}

func sendBlockPacket(sender *melody.Session, action byte, x, y float32, texture *uint32) error {
	packet := map[string]interface{}{
		"Action": action,
		"X":      x,
		"Y":      y,
	}
	if action == blockAdd && texture != nil {
		packet["Texture"] = *texture
	}

	binaryPacket, err := msgpack.Marshal(&packet)
	if err != nil {
		return err
	}

	if err := websocket.BroadcastBinaryFilter(append([]byte{blockPacket}, binaryPacket...), func(session *melody.Session) bool {
		return sender != session
	}); err != nil {
		logger.Error("Error with broadcast block packet: ", err)
		return err
	}
	return nil
}

func sendPlayersList() error {
	list := player.GetList()
	data, err := msgpack.Marshal(&list)
	if err != nil {
		return err
	}
	if err := websocket.BroadcastBinary(append([]byte{playerPacket, playerList}, data...)); err != nil {
		return err
	}
	return nil
}

func sendPlayerMove(data []byte) error {
	if err := websocket.BroadcastBinary(append([]byte{playerPacket, playerMove}, data...)); err != nil {
		logger.Error("Failed to broadcast player move data: ", err)
		return err
	}
	return nil
}
