package main

import (
	"fmt"
	"log"
	"main/pubsub"

	"github.com/fsnotify/fsnotify"
)

// subscriber
type Printer struct {
	Name string
}

func (p Printer) OnMessage(message string) {
	fmt.Printf("Printer %s. Message: %s\n", p.Name, message)
}

var _ pubsub.Subscriber = Printer{}

// publisher
type FolderWatcher struct {
	FolderPath string
	pubsub     pubsub.PubSub
	topicName  string
}

func (fw FolderWatcher) Watch() {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Add(fw.FolderPath)
	if err != nil {
		log.Fatal(err)
	}

	// Start listening for events.
	go func() {
		defer watcher.Close()

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				var message string
				switch {
				case event.Op&fsnotify.Write == fsnotify.Write:
					message = fmt.Sprintf("Write:  %s: %s", event.Op, event.Name)
				case event.Op&fsnotify.Create == fsnotify.Create:
					message = fmt.Sprintf("Create: %s: %s", event.Op, event.Name)
				case event.Op&fsnotify.Remove == fsnotify.Remove:
					message = fmt.Sprintf("Remove: %s: %s", event.Op, event.Name)
				case event.Op&fsnotify.Rename == fsnotify.Rename:
					message = fmt.Sprintf("Rename: %s: %s", event.Op, event.Name)
				case event.Op&fsnotify.Chmod == fsnotify.Chmod:
					message = fmt.Sprintf("Chmod:  %s: %s", event.Op, event.Name)
				}

				fw.pubsub.Publish(fw.topicName, message) // send message to subscribers

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	log.Println("Watching folder:", fw.FolderPath)
}

func main() {
	fmt.Println("Yoooo G!")

	topicName := "watcher"
	topicName2 := "watcher2"
	pubsub := pubsub.GetNewPubSub(topicName, topicName2)

	folderWatcher := FolderWatcher{FolderPath: "test", pubsub: pubsub, topicName: topicName}
	pubsub.Subscribe(topicName, Printer{Name: "OfficeRoom1"})

	folderWatcher2 := FolderWatcher{FolderPath: "test2", pubsub: pubsub, topicName: topicName2}
	pubsub.Subscribe(topicName2, Printer{Name: "OfficeRoom2"})

	folderWatcher.Watch()
	folderWatcher2.Watch()

	<-make(chan struct{})
}
