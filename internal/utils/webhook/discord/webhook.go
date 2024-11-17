package discord

import (
	"fmt"

	"github.com/gtuk/discordwebhook"
	"github.com/otie173/odncore/internal/utils/config"
)

var (
	webhookEnabled bool
	webhookName    string
	webhookURL     string
)

func InitDiscord() {
	cfg := config.GetConfig()
	webhookEnabled = cfg.DiscordWebhookEnabled
	webhookURL = cfg.DiscordWebhookURL
	webhookName = cfg.DiscordWebhookName
}

func WebhookEnabled() bool {
	if webhookEnabled {
		return true
	}
	return false
}

func SendMessage(content string) error {
	message := discordwebhook.Message{
		Username: &webhookName,
		Content:  &content,
	}

	if err := discordwebhook.SendMessage(webhookURL, message); err != nil {
		return err
	}
	return nil
}

func PlayerMessage(name, text string) error {
	content := fmt.Sprintf("**%s** %s", name, text)
	message := discordwebhook.Message{
		Username: &webhookName,
		Content:  &content,
	}

	if err := discordwebhook.SendMessage(webhookURL, message); err != nil {
		return err
	}
	return nil
}
