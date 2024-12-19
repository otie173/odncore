# Server Configuration Documentation

This document describes all available configuration options for the game server.

## Configuration Fields

| Field | Type | Description |
|-------|------|-------------|
| server_name | string | Name of the server that will be displayed to players |
| server_description | string | Description of the server that will be displayed to players |
| address | string | Server binding address in format "ip:port" (e.g. "0.0.0.0:8080") |
| max_player | number | Maximum number of players that can connect simultaneously |
| discord_webhook_enabled | bool | Enables/disables Discord webhook integration |
| discord_webhook_url | string | Discord webhook URL for server notifications |
| discord_webhook_name | string | Custom name for Discord webhook messages |

## Example Configuration
```json
{
    "server_name": "My Awesome Game Server",
    "server_description": "A fun place to build and play together!",
    "address": "0.0.0.0:8080",
    "max_player": 16,
    "discord_webhook_enabled": true,
    "discord_webhook_url": "https://discord.com/api/webhooks/123456789/abcdef...",
    "discord_webhook_name": "Server Status Bot"
}
