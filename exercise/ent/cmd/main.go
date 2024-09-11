package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	data "github.com/danielhookx/go_skill/exercise/ent"
	"github.com/danielhookx/go_skill/exercise/ent/ent"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	client := data.NewEntClient()

	var done = make(chan struct{})

	for i := 0; i < 10; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					_, err := client.User.Query().All(context.Background())
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					insert(client)
				}
			}
		}()
	}

	// init signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-sc
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			close(done)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func insert(client *ent.Client) {
	for i := 0; i < 20; i++ {
		_, err := client.User.Create().
			SetName(randStringBytes(5)).
			SetNickName(randStringBytes(7)).
			SetF1(randStringBytes(200)).
			SetF2(randStringBytes(200)).
			SetF3(randStringBytes(200)).
			SetF4(randStringBytes(200)).
			SetF5(randStringBytes(200)).
			SetF6(randStringBytes(200)).
			SetF7(randStringBytes(200)).
			SetF8(randStringBytes(200)).
			SetF9(randStringBytes(200)).
			SetF10(randStringBytes(200)).
			SetF11(randStringBytes(200)).
			SetF12(randStringBytes(200)).
			SetF13(randStringBytes(200)).
			SetF14(randStringBytes(200)).
			SetF15(randStringBytes(200)).
			SetF16(randStringBytes(200)).
			SetF17(randStringBytes(200)).
			SetF18(randStringBytes(200)).
			SetF19(randStringBytes(200)).
			SetF20(randStringBytes(200)).
			SetF21(randStringBytes(200)).
			SetF22(randStringBytes(200)).
			Save(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = client.Class.Create().
			SetName(randStringBytes(5)).
			SetClassName(randStringBytes(7)).
			SetF1(randStringBytes(200)).
			SetF2(randStringBytes(200)).
			SetF3(randStringBytes(200)).
			SetF4(randStringBytes(200)).
			SetF5(randStringBytes(200)).
			SetF6(randStringBytes(200)).
			SetF7(randStringBytes(200)).
			SetF8(randStringBytes(200)).
			SetF9(randStringBytes(200)).
			SetF10(randStringBytes(200)).
			SetF11(randStringBytes(200)).
			SetF12(randStringBytes(200)).
			SetF13(randStringBytes(200)).
			SetF14(randStringBytes(200)).
			SetF15(randStringBytes(200)).
			SetF16(randStringBytes(200)).
			SetF17(randStringBytes(200)).
			SetF18(randStringBytes(200)).
			SetF19(randStringBytes(200)).
			SetF20(randStringBytes(200)).
			SetF21(randStringBytes(200)).
			SetF22(randStringBytes(200)).
			Save(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func txInsert(client *ent.Client) {
	// time.Sleep(5 * time.Millisecond)
	tx, err := client.Tx(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 20; i++ {
		_, err = tx.User.Create().
			SetName(randStringBytes(5)).
			SetNickName(randStringBytes(7)).
			SetF1(randStringBytes(200)).
			SetF2(randStringBytes(200)).
			SetF3(randStringBytes(200)).
			SetF4(randStringBytes(200)).
			SetF5(randStringBytes(200)).
			SetF6(randStringBytes(200)).
			SetF7(randStringBytes(200)).
			SetF8(randStringBytes(200)).
			SetF9(randStringBytes(200)).
			SetF10(randStringBytes(200)).
			SetF11(randStringBytes(200)).
			SetF12(randStringBytes(200)).
			SetF13(randStringBytes(200)).
			SetF14(randStringBytes(200)).
			SetF15(randStringBytes(200)).
			SetF16(randStringBytes(200)).
			SetF17(randStringBytes(200)).
			SetF18(randStringBytes(200)).
			SetF19(randStringBytes(200)).
			SetF20(randStringBytes(200)).
			SetF21(randStringBytes(200)).
			SetF22(randStringBytes(200)).
			Save(context.Background())
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
			return
		}

		_, err = tx.Class.Create().
			SetName(randStringBytes(5)).
			SetClassName(randStringBytes(7)).
			SetF1(randStringBytes(200)).
			SetF2(randStringBytes(200)).
			SetF3(randStringBytes(200)).
			SetF4(randStringBytes(200)).
			SetF5(randStringBytes(200)).
			SetF6(randStringBytes(200)).
			SetF7(randStringBytes(200)).
			SetF8(randStringBytes(200)).
			SetF9(randStringBytes(200)).
			SetF10(randStringBytes(200)).
			SetF11(randStringBytes(200)).
			SetF12(randStringBytes(200)).
			SetF13(randStringBytes(200)).
			SetF14(randStringBytes(200)).
			SetF15(randStringBytes(200)).
			SetF16(randStringBytes(200)).
			SetF17(randStringBytes(200)).
			SetF18(randStringBytes(200)).
			SetF19(randStringBytes(200)).
			SetF20(randStringBytes(200)).
			SetF21(randStringBytes(200)).
			SetF22(randStringBytes(200)).
			Save(context.Background())
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
			return
		}
	}

	tx.Commit()
}

func txInsertTiny(client *ent.Client) {
	time.Sleep(50 * time.Millisecond)
	tx, err := client.Tx(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	// for i := 0; i < 20; i++ {
	_, err = tx.User.Create().
		SetName(randStringBytes(5)).
		SetNickName(randStringBytes(7)).
		Save(context.Background())
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return
	}

	_, err = tx.Class.Create().
		SetName(randStringBytes(5)).
		SetClassName(randStringBytes(7)).
		Save(context.Background())
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return
	}
	// }
	tx.Commit()
}
