package pubsub

import "fmt"

type PubSub struct {
	topics map[string]*Topic
}

func GetNewPubSub(topicNames ...string) PubSub {
	ps := PubSub{topics: make(map[string]*Topic)}
	for _, name := range topicNames {
		ps.topics[name] = &Topic{}
	}
	return ps
}

func (pb *PubSub) Subscribe(topicName string, subscriber Subscriber) {
	topic := pb.topics[topicName]
	topic.AddSubscriber(subscriber)
}

func (pb *PubSub) Unsubscribe(topicName string, subscriber Subscriber) {
	topic := pb.topics[topicName]
	topic.RemoveSubscriber(subscriber)
}

func (pb PubSub) Publish(topicName, message string) {
	topic, exists := pb.topics[topicName]
	if !exists {
		fmt.Printf("Topic not exists: %s\n", topicName)
	}
	topic.Notify(message)
}
