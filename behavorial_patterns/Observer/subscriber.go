package observer

import "fmt"

type subscriber interface {
	update(string)
	getId() string
}

type Subscriber struct {
	Name string
	Id   string
}

func (s *Subscriber) update(title string) {
	fmt.Printf("Subcriber %s receives notification about new thread: %s\n", s.Name, title)
}

func (s *Subscriber) getId() string {
	return s.Id
}
