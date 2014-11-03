package broadcastr

import (
	"testing"
)

func ExecBroadcaster(b Broadcaster, numSubscribers, numMessages int) {
	doneReading := make(chan struct{}, numSubscribers)
	listening := make(chan struct{}, numSubscribers)

	for i := 0; i < numSubscribers; i++ {
		go func() {
			me := b.Subscribe()
			listening <- struct{}{}
			me.ReadSync()
			doneReading <- struct{}{}
		}()
	}

	for i := 0; i < numSubscribers; i++ {
		<-listening
	}

	for i := 0; i < numMessages; i++ {
		go b.Send(i)
	}

	for i := 0; i < numSubscribers; i++ {
		<-doneReading
	}
}

//func BenchmarkLinkedBroadcaster1Subscribers5Messages(b *testing.B) {
	//for i := 0; i < b.N; i++ {
		//b := NewLinkedBroadcaster()
		//ExecBroadcaster(b, 1, 5)
	//}
//}

//func BenchmarkAsyncBroadcaster1Subscribers5Messages(b *testing.B) {
	//for i := 0; i < b.N; i++ {
		//b := NewAsyncBroadcaster()
		//ExecBroadcaster(b, 1, 5)
	//}
//}

//func BenchmarkSyncBroadcaster1Subscribers5Messages(b *testing.B) {
	//for i := 0; i < b.N; i++ {
		//b := NewSyncBroadcaster()
		//ExecBroadcaster(b, 1, 5)
	//}
//}

func BenchmarkLinkedBroadcaster2Subscribers10Messages(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := NewLinkedBroadcaster()
		ExecBroadcaster(b, 2, 10)
	}
}

func BenchmarkAsyncBroadcaster2Subscribers10Messages(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := NewAsyncBroadcaster()
		ExecBroadcaster(b, 2, 10)
	}
}

func BenchmarkSyncBroadcaster2Subscribers10Messages(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := NewSyncBroadcaster()
		ExecBroadcaster(b, 2, 10)
	}
}

//func BenchmarkLinkedBroadcaster200Subscribers10Messages(b *testing.B) {
	//for i := 0; i < b.N; i++ {
		//b := NewLinkedBroadcaster()
		//ExecBroadcaster(b, 200, 10)
	//}
//}

//func BenchmarkAsyncBroadcaster200Subscribers10Messages(b *testing.B) {
	//for i := 0; i < b.N; i++ {
		//b := NewAsyncBroadcaster()
		//ExecBroadcaster(b, 200, 10)
	//}
//}

//func BenchmarkSyncBroadcaster200Subscribers10Messages(b *testing.B) {
	//for i := 0; i < b.N; i++ {
		//b := NewSyncBroadcaster()
		//ExecBroadcaster(b, 200, 10)
	//}
//}
