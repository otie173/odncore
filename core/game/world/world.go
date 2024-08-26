package world

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

func init() {
	world = make(map[Rectangle]Block, WORLD_SIZE*WORLD_SIZE)
	id = make(map[int]ServerTexture, MAX_ID)
	for i := WALL; i <= SEED2BIG; i++ {
		id[i] = ServerTexture(i)
	}

	if WorldExists() {
		// TODO: отдавать клиенту этот мир при подключении к серверу
		//loadWorldFile()
	} else {
		// TODO: ждать пока клиент сгенерирует у себя мир и отошлет этот мир на сервер (тут наверное надо менять счетчик isWorldWaiting, который будет давать понять другим клиентам, что первый подключившийся клиент не сгенерировал мир и им нужно чуть повисеть в очереди)
		IsWorldWaiting = true
	}
}