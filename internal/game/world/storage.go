package world

import (
	"encoding/json"
	"os"

	"github.com/otie173/odncore/internal/utils/filesystem"
	"github.com/otie173/odncore/internal/utils/logger"
	"github.com/vmihailenco/msgpack/v5"
)

const (
	BLOCK_BITS = 7
	BLOCK_MASK = (1 << BLOCK_BITS) - 1
)

func Save() error {
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
					textureID = 0
				}
				blocks[index] = textureID + 1
			} else {
				blocks[index] = 0
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
		return err
	}
	return nil
}

func Load() error {
	data, err := os.ReadFile(filesystem.WORLD_DIR_PATH + "world.odn")
	if err != nil {
		return err
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
			textureID := int(blocks[index]) - 1
			if textureID >= 0 {
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
	return nil
}

func SaveId() error {
	data, err := msgpack.Marshal(&id)
	if err != nil {
		return err
	}

	os.WriteFile(filesystem.WORLD_DIR_PATH+"id.odn", data, 0644)
	return nil
}

func LoadIdNetwork(data []byte) error {
	if err := msgpack.Unmarshal(data, &id); err != nil {
		return err
	}
	IsIdWaiting = false
	return nil
}

func LoadIdFile() error {
	data, err := os.ReadFile(filesystem.WORLD_DIR_PATH + "id.odn")
	if err != nil {
		return err
	}

	if err := msgpack.Unmarshal(data, &id); err != nil {
		return err
	}
	IsIdWaiting = false
	return nil
}

func ByteToFile(data []byte) error {
	err := os.WriteFile(filesystem.WORLD_DIR_PATH+"world.odn", data, 0644)
	if err != nil {
		return err
	}
	if err := Load(); err != nil {
		return err
	}
	return nil
}

func SaveInfo() error {
	jsonData, err := json.Marshal(&worldInfo)
	if err != nil {
		return err
	}

	os.WriteFile(filesystem.WORLD_DIR_PATH+"world_info.json", jsonData, 0644)
	return nil
}

func LoadInfoNetwork(data WorldInfo) {
	worldInfo = data
	IsWorldInfoWaiting = false
}

func LoadInfoFile() error {
	jsonData, err := os.ReadFile(filesystem.WORLD_DIR_PATH + "world_info.json")
	if err != nil {
		logger.Error("Failed to read world info from file: ", err)
	}

	err = json.Unmarshal(jsonData, &worldInfo)
	if err != nil {
		logger.Error("Failed to unmarshal json data to worldInfo variable: ", err)
	}
	IsWorldInfoWaiting = false
	return nil
}
