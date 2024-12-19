# Network Packets Documentation

This documentation describes network packets that can be both sent to and received from the game server. Since the server broadcasts incoming packets to all players, the same packet structure is used for both incoming and outgoing communication. Server-specific outgoing packets are not documented as the game is not open-source.

## Packets

### Block Addition Packet
Used to add a new block to the game world.

| Field | Type | Description |
|-------|------|-------------|
| Action | byte | Packet opcode (blockAdd) |
| Texture | uint32 | Block texture ID |
| X | float32 | Block X position |
| Y | float32 | Block Y position |
| Passable | bool | Whether entities can pass through |

### Block Removal Packet
Used to remove a block from the game world.

| Field | Type | Description |
|-------|------|-------------|
| Action | byte | Packet opcode (blockRemove) |
| X | float32 | Block X position |
| Y | float32 | Block Y position |
