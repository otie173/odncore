package world

import (
	"log"
	"os"
)

const (
	BLOCK_BITS = 7
	BLOCK_MASK = (1 << BLOCK_BITS) - 1
)

func WorldExists() bool {
	_, err := os.Stat("world.odn")
	return !os.IsNotExist(err)
}

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
					if ServerTexture(block.img) == texture {
						textureID = byte(id)
						break
					}
				}
				if textureID >= (1 << BLOCK_BITS) {
					textureID = 0
				}
				blocks[index] = textureID
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

	err := os.WriteFile("world.odn", data, 0644)
	if err != nil {
		log.Fatalf("Failed to save world: %v", err)
	} else {
		log.Println("World saved successfully")
	}
}

func Load() {
	data, err := os.ReadFile("world.odn")
	if err != nil {
		log.Printf("Failed to load world: %v", err)
		return
	}
	log.Println("World loaded successfully")

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

	loadedWorld := make(map[Rectangle]Block, WORLD_SIZE*WORLD_SIZE)
	index := 0
	for y := -WORLD_SIZE / 2; y <= WORLD_SIZE/2; y++ {
		for x := -WORLD_SIZE / 2; x <= WORLD_SIZE/2; x++ {
			textureID := blocks[index]
			if textureID > 0 { // Загружаем только непустые блоки
				rect := Rectangle{
					X:      float32(x) * TILE_SIZE,
					Y:      float32(y) * TILE_SIZE,
					Width:  TILE_SIZE,
					Height: TILE_SIZE,
				}

				passable := false
				passableBlocks := []int{DOOR, GRASS1, GRASS2, GRASS3, GRASS4, GRASS5, GRASS6, FLOOR, FLOOR2, FLOOR4, DOOROPEN}
				for _, block := range passableBlocks {
					if textureID == byte(block) {
						passable = true
						break
					}
				}

				loadedWorld[rect] = Block{img: textureID, rec: rect, passable: passable}
			}
			index++
		}
	}

	world = loadedWorld
}
