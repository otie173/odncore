package player

var (
	players map[string]Player
)

type Player struct {
	nickname  string
	inventory Inventory
}

type Inventory struct {
	X                float32 `json:"x"`
	Y                float32 `json:"y"`
	Health           int     `json:"health"`
	WoodCount        int     `json:"wood"`
	StoneCount       int     `json:"stone"`
	MetalCount       int     `json:"metal"`
	PickaxeOpen      bool    `json:"pickaxe_open"`
	AxeOpen          bool    `json:"axe_open"`
	ShovelOpen       bool    `json:"shovel_open"`
	WallOpen         bool    `json:"wall_open"`
	WallWindowOpen   bool    `json:"wall_window_open"`
	FloorOpen        bool    `json:"floor_open"`
	DoorOpen         bool    `json:"door_open"`
	DoorOpenOpen     bool    `json:"door_open_open"`
	ChestOpen        bool    `json:"chest_open"`
	WallCount        int     `json:"wall_count"`
	WallWindowCount  int     `json:"wall_window_count"`
	FloorCount       int     `json:"floor_count"`
	DoorCount        int     `json:"door_count"`
	ChestCount       int     `json:"chest_count"`
	DoorOpenCount    int     `json:"door_open_count"`
	BigBarrelOpen    bool    `json:"big_barrel_open"`
	BookshelfOpen    bool    `json:"bookshelf_open"`
	ChairOpen        bool    `json:"chair_open"`
	ClosetOpen       bool    `json:"closet_open"`
	Fence1Open       bool    `json:"fence1_open"`
	Fence2Open       bool    `json:"fence2_open"`
	Floor2Open       bool    `json:"floor2_open"`
	Floor4Open       bool    `json:"floor4_open"`
	LampOpen         bool    `json:"lamp_open"`
	ShelfOpen        bool    `json:"shelf_open"`
	SignOpen         bool    `json:"sign_open"`
	SmallBarrelOpen  bool    `json:"small_barrel_open"`
	TableOpen        bool    `json:"table_open"`
	TrashOpen        bool    `json:"trash_open"`
	LootboxOpen      bool    `json:"lootbox_open"`
	TombstoneOpen    bool    `json:"tombstone_open"`
	SaplingOpen      bool    `json:"sapling_open"`
	SeedOpen         bool    `json:"seed_open"`
	CabbageOpen      bool    `json:"cabbage_open"`
	BigBarrelCount   int     `json:"big_barrel_count"`
	BookshelfCount   int     `json:"bookshelf_count"`
	ChairCount       int     `json:"chair_count"`
	ClosetCount      int     `json:"closet_count"`
	Fence1Count      int     `json:"fence1_count"`
	Fence2Count      int     `json:"fence2_count"`
	Floor2Count      int     `json:"floor2_count"`
	Floor4Count      int     `json:"floor4_count"`
	LampCount        int     `json:"lamp_count"`
	ShelfCount       int     `json:"shelf_count"`
	SignCount        int     `json:"sign_count"`
	SmallBarrelCount int     `json:"small_barrel_count"`
	TableCount       int     `json:"table_count"`
	TrashCount       int     `json:"trash_count"`
	LootboxCount     int     `json:"lootbox_count"`
	TombstoneCount   int     `json:"tombstone_count"`
	SaplingCount     int     `json:"sapling_count"`
	SeedCount        int     `json:"seed_count"`
	CabaggeCount     int     `json:"cabbage_count"`
}
