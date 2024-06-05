package pubsub

import (
	"github.com/nats-io/nats.go"
)

type PubSub interface {
	Subscribe(subject string, callback func(data []byte))
	Publish(subject string, data []byte)
	Close()
}

type NatsPubSub struct {
	nc *nats.Conn
}

func NewNatsPubSub() (*NatsPubSub, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}

	return &NatsPubSub{nc: nc}, nil
}

func (np *NatsPubSub) Subscribe(subject string, callback func(data []byte)) {
	np.nc.Subscribe(subject, func(msg *nats.Msg) {
		callback(msg.Data)
	})
}

func (np *NatsPubSub) Publish(subject string, data []byte) {
	np.nc.Publish(subject, data)
}

func (np *NatsPubSub) Close() {
	np.nc.Close()
}
