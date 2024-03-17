package cron

import (
	"fmt"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/robfig/cron/v3"
)

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
	job.AddFunc("1 * * * * *", cmd)
	job.Start()
	wg.Wait()
}
