package server

import (
	"os"

	"github.com/olahol/melody"
	"github.com/otie173/odncore/core/game/world"
	"github.com/otie173/odncore/utils/filesystem"
	"github.com/otie173/odncore/utils/logger"
	"github.com/vmihailenco/msgpack/v5"
)

type Request struct {
	Name     string
	Action   byte
	Texture  uint32
	X        float32
	Y        float32
	Passable bool
}

func handleRequest(opcode byte, data []byte) {
	switch opcode {
	case RECEIVE_WORLD:
		ReceiveWorld(data)
	case ADD_BLOCK:
		var reqData Request
		if err := msgpack.Unmarshal(data, &reqData); err != nil {
			logger.Error("Error unmarshalling data:", err)
		}
		world.AddBlock(reqData.Texture, reqData.X, reqData.Y, reqData.Passable)
	case REMOVE_BLOCK:
		world.RemoveBlock(float32(data[1]), float32(data[2]))
	case RECEIVE_ID:
		world.LoadIdNetwork(data)
	}
}

func SetupReadHandler() {
	websocket.HandleMessageBinary(func(s *melody.Session, b []byte) {
		handleRequest(b[0], b[1:])
	})
}

func AskWorld(session *melody.Session) error {
	if err := session.WriteBinary([]byte{SEND_WORLD}); err != nil {
		return err
	}
	return nil
}

func ReceiveWorld(data []byte) error {
	if err := world.ByteToFile(data); err != nil {
		return err
	}
	world.IsWorldWaiting = false
	return nil
}

func SendWorld(session *melody.Session) error {
	world.Save()

	worldData, err := os.ReadFile(filesystem.WORLD_DIR_PATH + "world.odn")
	if err != nil {
		return err
	}

	data := append([]byte{RECEIVE_WORLD}, worldData...)

	if err = session.WriteBinary(data); err != nil {
		return err
	}

	return nil
}

func AskId(session *melody.Session) error {
	if err := session.WriteBinary([]byte{SEND_ID}); err != nil {
		return err
	}
	return nil
}

func SendToClients(sender *melody.Session, msg []byte) error {
	return websocket.BroadcastFilter(msg, func(session *melody.Session) bool {
		return sender != session
	})
}
