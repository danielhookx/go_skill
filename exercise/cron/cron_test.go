package cron

import (
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/danielhookx/xcron"
	"github.com/robfig/cron/v3"
)

func TestSecondJob1(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	job := cron.New(
		cron.WithLogger(cron.VerbosePrintfLogger(
			log.New(os.Stdout, "cron: ", log.LstdFlags),
		)),
	)
	job.Schedule(Every(1*time.Second), cron.FuncJob(func() {
		fmt.Println("do some things")
		wg.Done()
	}))
	job.Start()
	wg.Wait()
}

func TestSecondJob(t *testing.T) {
	job := cron.New(
		cron.WithLogger(cron.VerbosePrintfLogger(
			log.New(os.Stdout, "cron: ", log.LstdFlags),
		)),
	)

	//job := cron.New(cron.WithLogger(cron.DefaultLogger), cron.WithSeconds(), cron.WithChain())

	i := 0
	wg := sync.WaitGroup{}
	wg.Add(1)
	cmd := func() {
		if i > 10 {
			wg.Done()
		}
		fmt.Printf("do some things %d \n", i)
		i++
	}
	//job.Schedule(Every(500*time.Millisecond), cron.FuncJob(cmd))
	//job.Schedule("", cron.FuncJob())
	job.Schedule(Every(1*time.Second), cron.FuncJob(cmd))
	job.Start()
	wg.Wait()
}

func BenchmarkNewCron(b *testing.B) {
	b.Run("robfig cron", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				job := cron.New()
				job.Schedule(Every(1*time.Second), cron.FuncJob(func() {}))
				job.Start()
			}
		})
	})

	b.Run("danielhookx xcron", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				job := xcron.NewCron()
				job.Schedule(Every(1*time.Second), cron.FuncJob(func() {}))
				job.Start()
			}
		})
	})
}

func BenchmarkCronSchedule(b *testing.B) {
	b.Run("robfig cron", func(b *testing.B) {
		job := cron.New()
		job.Start()
		wg := sync.WaitGroup{}
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				wg.Add(1)
				once := sync.Once{}

				job.Schedule(Every(1*time.Millisecond), cron.FuncJob(func() {
					once.Do(func() {
						wg.Done()
					})
				}))
			}
		})
		wg.Wait()
	})

	b.Run("danielhookx xcron", func(b *testing.B) {
		job := xcron.NewCron()
		job.Start()
		wg := sync.WaitGroup{}
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				wg.Add(1)
				once := sync.Once{}
				job.Schedule(Every(1*time.Millisecond), cron.FuncJob(func() {
					once.Do(func() {
						wg.Done()
					})
				}))
			}
		})
		wg.Wait()
	})
}
