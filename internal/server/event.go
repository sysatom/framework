package server

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/sysatom/framework/pkg/flog"
)

// push instruct
func onInstructPushEventHandler(msg *message.Message) error {
	flog.Debug("[event] on event %+v %+v", msg.UUID, msg.Metadata)

	return nil
}
