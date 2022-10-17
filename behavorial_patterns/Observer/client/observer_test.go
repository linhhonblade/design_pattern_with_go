package main

import (
	"testing"

	observer "github.com/linhhonblade/observer"
)

func TestObserver(t *testing.T) {
	//Setup
	channelName := "Mailovemisa"
	testChannel := observer.NewChannel(channelName)
	testSubscriber1 := &observer.Subscriber{Name: "subscriber1", Id: "sub1"}
	testSubscriber2 := &observer.Subscriber{Name: "subscriber2", Id: "sub2"}
	testSubscriber3 := &observer.Subscriber{Name: "subscriber3", Id: "sub3"}
	// Subcribe test
	//Test case: Add 3 subscribers to the testChannel
	t.Run("Subscribe", func(t *testing.T) {
		testChannel.Subcribe(testSubscriber1)
		testChannel.Subcribe(testSubscriber2)
		testChannel.Subcribe(testSubscriber3)

		if len(testChannel.(*observer.Channel).SubscribeList) != 3 {
			t.Fail()
		}
	})

	// Unsubscribe test
	// Test case: Unsubcribe testSubcriber2
	t.Run("Unsubscribe", func(t *testing.T) {
		testChannel.Unsubscribe(testSubscriber2)
		if len(testChannel.(*observer.Channel).SubscribeList) != 2 {
			t.Errorf("The size of subscriberList is not the expected. 2 != %d\n", len(testChannel.(*observer.Channel).SubscribeList))
		}
	})

	// NotifyAll test
	t.Run("NotifyAll", func(t *testing.T) {
		// Check all the newesttPost is empty before calling NotifyAll
		for _, subscriber := range testChannel.(*observer.Channel).SubscribeList {
			if subscriber.(*observer.Subscriber).NewestPost != "" {
				t.Errorf("The subscriber newestPost field wasn't empty: %s\n", subscriber.(*observer.Subscriber).NewestPost)
			}
		}
		// Call PushThread and check if the newestPost is updated or not
		testChannel.PushThread(observer.NewThread("MeoMeoMeo", "mailovemisa"))
		for _, subscriber := range testChannel.(*observer.Channel).SubscribeList {
			if subscriber.(*observer.Subscriber).NewestPost != "MeoMeoMeo" {
				t.Errorf("The subscriber newestPost field is not the expected. MeoMeoMeo != %s\n", subscriber.(*observer.Subscriber).NewestPost)
			}
		}
	})
}
