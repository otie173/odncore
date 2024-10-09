package world

import (
	"log"
)

type Rectangle struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

type Block struct {
	img      byte
	rec      Rectangle
	passable bool
}

type ServerTexture int

var (
	world          map[Rectangle]Block
	id             map[int]ServerTexture
	IsWorldWaiting bool
)

const (
	TILE_SIZE  float32 = 10.0
	WORLD_SIZE int     = 320
	MAX_ID     byte    = 127
)

const (
	WALL = iota
	WALLWINDOW
	FLOOR
	DOOR
	CHEST
	SMALL_TREE
	NORMAL_TREE
	BIG_TREE
	STONE1
	STONE2
	STONE3
	STONE4
	BIGSTONE1
	BIGSTONE2
	BIGSTONE3
	BIGSTONE4
	BIGSTONE5
	GRASS1
	GRASS2
	GRASS3
	GRASS4
	GRASS5
	GRASS6
	BARRIER
	LOOTBOX
	BONES1
	BONES2
	BONES3
	BONES4
	BONES5
	PICKAXE
	AXE
	SHOVEL
	DOOROPEN
	BIGBARREL
	BOOKSHELF
	CHAIR
	CLOSET
	FENCE1
	FENCE2
	FLOOR2
	FLOOR4
	LAMP
	SHELF
	SIGN
	SMALLBARREL
	TABLE
	TOMBSTONE
	TRASH
	STAIRSDOWN
	STAIRSUP
	SAPLING
	SEED1NORMAL
	SEED1BIG
	SEED2SMALL
	SEED2NORMAL
	SEED2BIG
)

func InitWorld() {
	world = make(map[Rectangle]Block, WORLD_SIZE*WORLD_SIZE)
	id = make(map[int]ServerTexture, MAX_ID)
	for i := WALL; i <= SEED2BIG; i++ {
		id[i] = ServerTexture(i)
	}

	if !WorldExists() {
		IsWorldWaiting = true
	}
}

func AddBlock(img byte, x, y float32, passable bool) {
	block := Block{
		img:      img,
		rec:      Rectangle{x * TILE_SIZE, y * TILE_SIZE, TILE_SIZE, TILE_SIZE},
		passable: passable,
	}
	world[block.rec] = block

	log.Printf("Игрок поставил блок ID: %d на позиции X: %.0f, Y: %.0f и его поле Passable: %t\n", img, x, y, passable)
	Save()

	log.Println(x, y)
}

func RemoveBlock(x, y float32) {
	delete(world, Rectangle{x * TILE_SIZE, y * TILE_SIZE, TILE_SIZE, TILE_SIZE})

	log.Printf("Игрок удалил блок на позиции X: %.0f, Y: %.0f\n", x, y)
	Save()
}
