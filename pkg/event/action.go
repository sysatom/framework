package event

import (
	"github.com/sysatom/framework/pkg/types"
)

func BotEventFire(ctx types.Context, eventName string, param types.KV) error {
	return PublishMessage(ctx.Context(), types.BotRunEvent, types.BotEvent{
		EventName: eventName,
		Uid:       ctx.AsUser.String(),
		Topic:     ctx.Topic,
		Param:     param,
	})
}
