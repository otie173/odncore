package discord

import (
	"fmt"

	"github.com/gtuk/discordwebhook"
	"github.com/otie173/odncore/internal/utils/config"
	"github.com/otie173/odncore/internal/utils/logger"
)

var (
	webhookEnabled bool
	webhookName    string
	webhookURL     string
)

func InitDiscord() {
	webhookURL = config.Cfg.DiscordWebhook
	webhookName = config.Cfg.DiscordWebhookName

	if webhookURL != "" && webhookName != "" {
		webhookEnabled = true
	}
}

func WebhookEnabled() bool {
	if webhookEnabled {
		return true
	}
	return false
}

func SendMessage(content string) {
	message := discordwebhook.Message{
		Username: &webhookName,
		Content:  &content,
	}

	if err := discordwebhook.SendMessage(webhookURL, message); err != nil {
		logger.Error(err)
	}
}

func PlayerMessage(text, name string) {
	content := fmt.Sprintf("%s **%s**", text, name)
	message := discordwebhook.Message{
		Username: &webhookName,
		Content:  &content,
	}

	if err := discordwebhook.SendMessage(webhookURL, message); err != nil {
		logger.Error(err)
	}
}
