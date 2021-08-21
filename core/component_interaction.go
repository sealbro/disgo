package core

import (
	"context"

	"github.com/DisgoOrg/disgo/discord"
)

type ComponentInteraction struct {
	*Interaction
	Message *Message
	Data    *ComponentInteractionData
}

// DeferUpdate replies to the ComponentInteraction with discord.InteractionResponseTypeDeferredUpdateMessage and cancels the loading state
func (i *ComponentInteraction) DeferUpdate(ctx context.Context) error {
	return i.Respond(ctx, discord.InteractionResponseTypeDeferredUpdateMessage, nil)
}

// Update replies to the ComponentInteraction with discord.InteractionResponseTypeUpdateMessage & MessageUpdate which edits the original Message
func (i *ComponentInteraction) Update(ctx context.Context, messageUpdate discord.MessageUpdate) error {
	return i.Respond(ctx, discord.InteractionResponseTypeUpdateMessage, messageUpdate)
}

// CustomID returns the Custom ID of the ComponentInteraction
func (i *ComponentInteraction) CustomID() string {
	return i.Data.CustomID
}

// ComponentType returns the ComponentType of a Component
func (i *ComponentInteraction) ComponentType() discord.ComponentType {
	return i.Data.ComponentType
}

// Component returns the Component which issued this ComponentInteraction. nil for ephemeral Message(s)
func (i *ComponentInteraction) Component() Component {
	return i.Message.ComponentByID(i.CustomID())
}

type ComponentInteractionData struct {
	*InteractionData
}
