package server

import (
	"log"

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

func handleRequest(session *melody.Session, opcode byte, data []byte) error {
	switch opcode {
	case blockPacket:
		if err := msgpack.Unmarshal(data, &packet); err != nil {
			return err
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
			sendBlockPacket(session, blockRemove, typeconv.GetFloat32(packet["X"]), typeconv.GetFloat32(packet["Y"]), nil)
		}
	case playerPacket:
		if err := msgpack.Unmarshal(data, &packet); err != nil {
			return err
		}

		switch typeconv.GetByte(packet["Action"]) {
		case playerMove:
			log.Println(packet)
		case playerAdd:
			player.Add(
				session.Request.Header.Get("Session-Nickname"),
				typeconv.GetFloat32(packet["X"]),
				typeconv.GetFloat32(packet["Y"]),
				typeconv.GetFloat32(packet["TargetX"]),
				typeconv.GetFloat32(packet["TargetY"]),
			)
		case playerList:
			sendPlayersList()
		}
	}
	return nil
}

func SetupReadHandler() {
	websocket.HandleMessageBinary(func(s *melody.Session, b []byte) {
		if err := handleRequest(s, b[0], b[1:]); err != nil {
			logger.Errorf("Error with handle request from client: %v", err)
		}
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

	return websocket.BroadcastBinaryFilter(append([]byte{blockPacket}, binaryPacket...), func(session *melody.Session) bool {
		return sender != session
	})
}

func sendPlayersList() {
	list := player.GetList()
	data, err := msgpack.Marshal(&list)
	if err != nil {
		logger.Error("Error with unmarshal players list: ", err)
	}
	websocket.BroadcastBinary(append([]byte{playerPacket, playerList}, data...))
}
