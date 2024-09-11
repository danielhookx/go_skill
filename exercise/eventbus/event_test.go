package eventbus_test

import (
	"sync"
	"testing"

	"github.com/danielhookx/eventbus"
	"github.com/luobote55/EventBus"
)

func BenchmarkSubPub(b *testing.B) {
	b.Run("Async EventBus", func(b *testing.B) {
		e := EventBus.NewEventBus()
		topic := "testpub1"

		b.RunParallel(func(pb *testing.PB) {
			var wg sync.WaitGroup
			wg.Add(2)

			go func() {
				defer wg.Done()
				for pb.Next() {
					e.SubscribeAsync(topic, func(name string) {}, false)
				}
			}()

			go func() {
				defer wg.Done()
				for pb.Next() {
					e.Publish(topic, "jack")
				}
			}()

			wg.Wait()
		})
	})
	b.Run("Async eventbus", func(b *testing.B) {
		e := eventbus.New()
		topic := "testpub1"
		b.RunParallel(func(pb *testing.PB) {
			var wg sync.WaitGroup
			wg.Add(2)

			go func() {
				defer wg.Done()
				for pb.Next() {
					e.Subscribe(topic, func(name string) {})
				}
			}()

			go func() {
				defer wg.Done()
				for pb.Next() {
					e.Publish(topic, "jack")
				}
			}()

			wg.Wait()
		})
	})
}

func BenchmarkSubPubSync(b *testing.B) {

	b.Run("Sync EventBus", func(b *testing.B) {
		e := EventBus.NewEventBus()
		topic := "testpub1"

		b.RunParallel(func(pb *testing.PB) {
			var wg sync.WaitGroup
			wg.Add(2)

			go func() {
				defer wg.Done()
				for pb.Next() {
					e.Subscribe(topic, func(name string) {})
				}
			}()

			go func() {
				defer wg.Done()
				for pb.Next() {
					e.Publish(topic, "jack")
				}
			}()

			wg.Wait()
		})

	})

	b.Run("Sync eventbus", func(b *testing.B) {
		e := eventbus.New()
		topic := "testpub1"

		b.RunParallel(func(pb *testing.PB) {
			var wg sync.WaitGroup
			wg.Add(2)

			go func() {
				defer wg.Done()
				for pb.Next() {
					e.Subscribe(topic, func(name string) {})
				}
			}()

			go func() {
				defer wg.Done()
				for pb.Next() {
					e.Publish(topic, "jack")
				}
			}()

			wg.Wait()
		})

	})

}
