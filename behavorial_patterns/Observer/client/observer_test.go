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

		if len(testChannel.GetSubsribeList()) != 3 {
			t.Fail()
		}
	})

	// Unsubscribe test
	// Test case: Unsubcribe testSubcriber2
	t.Run("Unsubscribe", func(t *testing.T) {
		testChannel.Unsubscribe(testSubscriber2)
		if len(testChannel.GetSubsribeList()) != 2 {
			t.Error("The size of subscriberList is not the expected. 3 != %d\n" + len(testChannel.subscribeList))
		}
	})
}
