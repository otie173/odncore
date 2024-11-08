package world

import (
	"log"
	"os"

	"github.com/otie173/odncore/internal/utils/filesystem"
	"github.com/otie173/odncore/internal/utils/logger"
	"github.com/vmihailenco/msgpack/v5"
)

const (
	BLOCK_BITS = 7
	BLOCK_MASK = (1 << BLOCK_BITS) - 1
)

func Save() {
	blocks := make([]byte, (WORLD_SIZE+1)*(WORLD_SIZE+1))
	index := 0
	for y := -WORLD_SIZE / 2; y <= WORLD_SIZE/2; y++ {
		for x := -WORLD_SIZE / 2; x <= WORLD_SIZE/2; x++ {
			rect := Rectangle{
				X:      float32(x) * TILE_SIZE,
				Y:      float32(y) * TILE_SIZE,
				Width:  TILE_SIZE,
				Height: TILE_SIZE,
			}

			if block, exists := world[rect]; exists {
				var textureID byte
				for id, texture := range id {
					if block.img == texture {
						textureID = byte(id)
						break
					}
				}
				if textureID >= (1 << BLOCK_BITS) {
					logger.Warnf("Предупреждение: ID блока %d превышает максимальное значение %d", textureID, (1<<BLOCK_BITS)-1)
					textureID = 0 // или какое-то другое значение по умолчанию
				}
				blocks[index] = textureID + 1 // Добавляем 1, чтобы 0 означало пустой блок
			} else {
				blocks[index] = 0 // Пустой блок
			}
			index++
		}
	}

	data := make([]byte, (len(blocks)*BLOCK_BITS+7)/8)
	for i, block := range blocks {
		bitIndex := i * BLOCK_BITS
		byteIndex := bitIndex / 8
		bitOffset := uint(bitIndex % 8)
		data[byteIndex] |= byte(block << bitOffset)
		if bitOffset > 8-BLOCK_BITS && byteIndex+1 < len(data) {
			data[byteIndex+1] |= byte(block >> (8 - bitOffset))
		}
	}

	err := os.WriteFile(filesystem.WORLD_DIR_PATH+"world.odn", data, 0644)
	if err != nil {
		logger.Fatalf("Не удалось сохранить мир: %v", err)
	}
	logger.Info("Мир успешно сохранен")
}

func Load() {
	data, err := os.ReadFile(filesystem.WORLD_DIR_PATH + "world.odn")
	if err != nil {
		log.Printf("Ошибка при чтении файла: %v", err)
	}

	blocks := make([]byte, (WORLD_SIZE+1)*(WORLD_SIZE+1))
	for i := range blocks {
		bitIndex := i * BLOCK_BITS
		byteIndex := bitIndex / 8
		bitOffset := uint(bitIndex % 8)
		if byteIndex+1 < len(data) {
			blocks[i] = byte((uint16(data[byteIndex]) | uint16(data[byteIndex+1])<<8) >> bitOffset & BLOCK_MASK)
		} else if byteIndex < len(data) {
			blocks[i] = byte(uint16(data[byteIndex]) >> bitOffset & BLOCK_MASK)
		} else {
			blocks[i] = 0
		}
	}

	loadedWorld := make(map[Rectangle]Block)
	index := 0
	for y := -WORLD_SIZE / 2; y <= WORLD_SIZE/2; y++ {
		for x := -WORLD_SIZE / 2; x <= WORLD_SIZE/2; x++ {
			textureID := int(blocks[index]) - 1 // Вычитаем 1, чтобы вернуться к оригинальному ID
			if textureID >= 0 {                 // Загружаем только непустые блоки
				rect := Rectangle{
					X:      float32(x) * TILE_SIZE,
					Y:      float32(y) * TILE_SIZE,
					Width:  TILE_SIZE,
					Height: TILE_SIZE,
				}

				passable := false
				passableBlocks := []int{DOOR, GRASS1, GRASS2, GRASS3, GRASS4, GRASS5, GRASS6, FLOOR, FLOOR2, FLOOR4, DOOROPEN}
				for _, block := range passableBlocks {
					if textureID == block {
						passable = true
						break
					}
				}
				loadedWorld[rect] = Block{img: id[textureID], rec: rect, passable: passable}
			}
			index++
		}
	}
	world = loadedWorld
	IsWorldWaiting = false
}

func SaveId() {
	data, err := msgpack.Marshal(&id)
	if err != nil {
		logger.Error("Error with saving id list: ", err)
	}

	os.WriteFile(filesystem.WORLD_DIR_PATH+"id.odn", data, 0644)
	logger.Info("Id list saved succesfully")
}

func LoadIdNetwork(data []byte) {
	if err := msgpack.Unmarshal(data, &id); err != nil {
		logger.Error("Error with loading id list from network: ", err)
	}
	IsIdWaiting = false
	logger.Info("Id list loaded succesfully")
}

func LoadIdFile() {
	data, err := os.ReadFile(filesystem.WORLD_DIR_PATH + "id.odn")
	if err != nil {
		logger.Error("Error with loading id list from file: ", err)
	}

	if err := msgpack.Unmarshal(data, &id); err != nil {
		logger.Error("Error with loading id list from file: ", err)
	}
	IsIdWaiting = false
	logger.Info("Id list loaded succesfully")
}

func ByteToFile(data []byte) error {
	err := os.WriteFile(filesystem.WORLD_DIR_PATH+"world.odn", data, 0644)
	if err != nil {
		return err
	}
	Load()

	return nil
}
