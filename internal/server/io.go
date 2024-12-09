package server

import (
	"github.com/olahol/melody"
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
		return err
	}
	return nil
}
