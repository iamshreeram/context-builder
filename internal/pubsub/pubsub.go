package pubsub

type PubSub interface {
	Subscribe(subject string, callback func(data []byte))
	Publish(subject string, data []byte)
	Close()
}
