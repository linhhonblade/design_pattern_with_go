package main

import (
	observer "github.com/linhhonblade/observer"
)

func main() {
	mailovemisaChannel := observer.NewChannel("Mailovemisa Channel")
	mailovemisaChannel.PushThread(observer.NewThread("You are sad girl", "Don't be sad girl"))
	mailovemisaChannel.NotifyAll() // Should do nothing cuz no one register this channel yet

	meomaimaimeo := &observer.Subscriber{Name: "Maimeo", Id: "meomaimaimeo"}
	maidien := &observer.Subscriber{Name: "Maidien", Id: "maidien"}

	// Firstly, memaimaimeo subscribe this channel
	mailovemisaChannel.Subcribe(meomaimaimeo)

	// New thread, meomaimaimeo should receive a notification
	mailovemisaChannel.PushThread(observer.NewThread("You are mad girl", "Just be mad girl"))

	// meomaimaimeo then incite maidien to join this gang too
	mailovemisaChannel.Subcribe(maidien)

	// New thread, they both get notification
	mailovemisaChannel.PushThread(observer.NewThread("You deserve the world", "Just be mad girl"))
}
