package server

import (
	"log"
	"os"

	"github.com/olahol/melody"
	"github.com/otie173/odncore/core/game/world"
	"github.com/vmihailenco/msgpack/v5"
)

type Request struct {
	Name     string
	Action   byte
	Texture  byte
	X        float32
	Y        float32
	Passable bool
}

func handleRequest(opcode byte, data []byte) {
	switch opcode {
	case ADD_BLOCK:
		var reqData Request
		if err := msgpack.Unmarshal(data, &reqData); err != nil {
			log.Println("Error unmarshalling data:", err)
		}
		log.Printf("Received block: %+v", reqData) // Логируем полученные данные
		world.AddBlock(reqData.Texture, reqData.X, reqData.Y, reqData.Passable)

	case REMOVE_BLOCK:
		world.RemoveBlock(float32(data[1]), float32(data[2]))
	}
}

func SetupReadHandler() {
	websocket.HandleMessageBinary(func(s *melody.Session, b []byte) {
		if world.IsWorldWaiting {
			world.IsWorldWaiting = false

			if err := world.ByteToFile(b); err != nil {
				log.Println("Error: ", err)
			}
		}

		handleRequest(b[0], b[1:])
	})
}

func ReceiveWorld(session *melody.Session) error {
	if err := session.WriteBinary([]byte{SEND_WORLD}); err != nil {
		return err
	}

	return nil
}

func SendWorld(session *melody.Session) error {
	log.Println("Мир отправлен")

	worldData, err := os.ReadFile("world.odn")
	if err != nil {
		return err
	}

	data := append([]byte{RECEIVE_WORLD}, worldData...)

	if err = session.WriteBinary(data); err != nil {
		return err
	}

	return nil
}

func SendToClients(sender *melody.Session, msg []byte) error {
	return websocket.BroadcastFilter(msg, func(session *melody.Session) bool {
		return sender != session
	})
}
