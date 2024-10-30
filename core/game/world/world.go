package world

import (
	"os"

	"github.com/otie173/odncore/utils/filesystem"
	"github.com/otie173/odncore/utils/logger"
)

type Rectangle struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

type PixelFormat int32

type Texture2D struct {
	ID      uint32
	Width   uint32
	Height  uint32
	Mipmaps int32
	Format  PixelFormat
}

type Block struct {
	img      Texture2D
	rec      Rectangle
	passable bool
}

var (
	world          map[Rectangle]Block
	id             map[int]Texture2D
	IsWorldWaiting bool
	IsWorldLoaded  bool
	IsIdWaiting    bool
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
	id = make(map[int]Texture2D, MAX_ID)

	dirs := []string{filesystem.WORLD_DIR_PATH}

	for _, path := range dirs {
		if !filesystem.DirExists(path) {
			err := os.Mkdir(path, 0755)
			if err != nil {
				logger.Error("Error creating directory: ", err)
			}
		}
	}

	if !filesystem.FileExists("world.odn") {
		IsWorldWaiting = true
	}
	IsIdWaiting = true
}

func AddBlock(img uint32, x, y float32, passable bool) {
	block := Block{
		img:      Texture2D{ID: img, Width: uint32(TILE_SIZE), Height: uint32(TILE_SIZE), Mipmaps: 1, Format: 7},
		rec:      Rectangle{x * TILE_SIZE, y * TILE_SIZE, TILE_SIZE, TILE_SIZE},
		passable: passable,
	}
	world[block.rec] = block

	logger.Infof("Игрок поставил блок ID: %d на позиции X: %.0f, Y: %.0f и его поле Passable: %t\n", img, x, y, passable)
}

func RemoveBlock(x, y float32) {
	delete(world, Rectangle{x * TILE_SIZE, y * TILE_SIZE, TILE_SIZE, TILE_SIZE})

	logger.Infof("Игрок удалил блок на позиции X: %.0f, Y: %.0f\n", x, y)
}
