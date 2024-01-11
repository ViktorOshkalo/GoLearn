package observer

type Observer interface {
	OnMessage(message string)
}

type ObserverManager struct {
	observers []Observer
}

func (om *ObserverManager) Register(subscriber Observer) {
	om.observers = append(om.observers, subscriber)
}

func (om *ObserverManager) Unregister(subscriber Observer) {
	substitute := []Observer{}
	for _, s := range om.observers {
		if subscriber != s {
			substitute = append(substitute, s)
		}
	}
	om.observers = substitute
}

func (om ObserverManager) NotifyAll(message string) {
	for _, s := range om.observers {
		s.OnMessage(message)
	}
}

func (om ObserverManager) NotifyOthers(sender Observer, message string) {
	for _, o := range om.observers {
		if o != sender {
			o.OnMessage(message)
		}
	}
}
