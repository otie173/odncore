package server

import (
	"github.com/olahol/melody"
	"github.com/otie173/odncore/internal/game/world"
	"github.com/otie173/odncore/internal/utils/logger"
	"github.com/otie173/odncore/internal/utils/typeconv"
	"github.com/vmihailenco/msgpack/v5"
)

func handleRequest(opcode byte, data []byte) {
	switch opcode {
	case BLOCK_PACKET:
		var packet map[string]interface{}
		if err := msgpack.Unmarshal(data, &packet); err != nil {
			logger.Error("Error unmarshalling data:", err)
		}

		switch typeconv.GetByte(packet["Action"]) {
		case ADD_BLOCK:
			world.AddBlock(
				typeconv.GetUint32(packet["Texture"]),
				typeconv.GetFloat32(packet["X"]),
				typeconv.GetFloat32(packet["Y"]),
				typeconv.GetBool(packet["Passable"]),
			)
		case REMOVE_BLOCK:
			world.RemoveBlock(
				typeconv.GetFloat32(packet["X"]),
				typeconv.GetFloat32(packet["Y"]),
			)
		}
	}
}

func SetupReadHandler() {
	websocket.HandleMessageBinary(func(s *melody.Session, b []byte) {
		handleRequest(b[0], b[1:])
	})
}

func SendToClients(sender *melody.Session, msg []byte) error {
	return websocket.BroadcastFilter(msg, func(session *melody.Session) bool {
		return sender != session
	})
}
