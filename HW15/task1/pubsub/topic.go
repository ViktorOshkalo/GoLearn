package pubsub

type Subscriber interface {
	OnMessage(message string)
}

type Topic struct {
	subscribers []Subscriber
}

func (tp *Topic) AddSubscriber(subscriber Subscriber) {
	tp.subscribers = append(tp.subscribers, subscriber)
}

func (tp *Topic) RemoveSubscriber(subscriber Subscriber) {
	substitute := []Subscriber{}
	for _, s := range tp.subscribers {
		if subscriber != s {
			substitute = append(substitute, s)
		}
	}
	tp.subscribers = substitute
}

func (tp Topic) Notify(message string) {
	for _, s := range tp.subscribers {
		s.OnMessage(message)
	}
}
