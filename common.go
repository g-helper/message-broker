package message_broker

import (
	"errors"
	"fmt"

	"github.com/nats-io/nats.go"
)

// NatsResponse ..
type NatsResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []byte `json:"data"`
}

type ServiceDesc struct {
	Queues []QueueDesc
}

// Register callback handler
func (s ServiceDesc) Register(queue interface{}) error {
	if len(s.Queues) == 0 {
		return errors.New(fmt.Sprintf("Server [staff] not declare queue"))
	}
	for _, q := range s.Queues {
		QueueSubscribe(q.Subject, q.Worker, q.Handle(queue))
	}
	return nil
}

type QueueDesc struct {
	Subject string
	Worker  string
	Handle  MethodHandle
}

// GetSubject get subject for nats request
func (q QueueDesc) GetSubject(serverName string) string {
	return fmt.Sprintf("%s:%s", serverName, q.Subject)
}

type MethodHandle func(queue interface{}) nats.MsgHandler
