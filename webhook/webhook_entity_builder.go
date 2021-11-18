package webhook

import (
	"fmt"

	"github.com/DisgoOrg/disgo/discord"
)

type EntityBuilder interface {
	WebhookClient() *Client
	CreateMessage(message discord.Message) *Message
	CreateWebhook(webhook discord.Webhook) *Webhook
}

func NewEntityBuilder(webhookClient *Client) EntityBuilder {
	return &entityBuilderImpl{
		webhookClient: webhookClient,
	}
}

type entityBuilderImpl struct {
	webhookClient *Client
}

func (b *entityBuilderImpl) WebhookClient() *Client {
	return b.webhookClient
}

func (b *entityBuilderImpl) CreateMessage(message discord.Message) *Message {
	return &Message{
		Message:       message,
		WebhookClient: b.WebhookClient(),
	}
}

func (b *entityBuilderImpl) CreateWebhook(webhook discord.Webhook) *Webhook {
	if w, ok := webhook.(discord.IncomingWebhook); ok {
		return &Webhook{
			IncomingWebhook:       w,
			WebhookClient: b.WebhookClient(),
		}
	}
	panic(fmt.Sprintf("invalid webhook type %d received", webhook.Type()))
}
