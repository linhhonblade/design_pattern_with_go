package observer

type publisher interface {
	Subcribe(s subscriber)
	Unsubscribe(s subscriber)
	NotifyAll()
	PushThread(t thread)
}

type thread struct {
	title  string
	author string
}

func NewThread(title string, author string) thread {
	return thread{
		title:  title,
		author: author,
	}
}

type Channel struct {
	name          string
	subscribeList []subscriber
	latestThread  thread
}

func NewChannel(name string) publisher {
	return &Channel{
		name: name,
	}
}

func (c *Channel) Subcribe(s subscriber) {
	c.subscribeList = append(c.subscribeList, s)
}

func removeFromSlice(subscribeList []subscriber, s subscriber) []subscriber {
	length := len(subscribeList)
	for i, v := range subscribeList {
		if v.getId() == s.getId() {
			// Move the one that we want to remove to the end of slice
			// Then return the slice excluding the last item
			subscribeList[length-1], subscribeList[i] = subscribeList[i], subscribeList[length-1]
			return subscribeList[:length-2]
		}
	}
	return subscribeList
}

func (c *Channel) Unsubscribe(s subscriber) {
	c.subscribeList = removeFromSlice(c.subscribeList, s)
}

func (c *Channel) NotifyAll() {
	for _, s := range c.subscribeList {
		s.update(c.latestThread.title)
	}
}

func (c *Channel) PushThread(t thread) {
	c.latestThread = t
	c.NotifyAll()
}
