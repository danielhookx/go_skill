package data

import (
	"context"
	"os"
	"sync"
	"testing"
	"time"

	"math/rand"

	"github.com/danielhookx/go_skill/exercise/ent/ent"
)

var client *ent.Client

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	client = NewEntClient()
	os.Exit(m.Run())
}

func TestNewEntClient(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(client *ent.Client) {
		defer wg.Done()
		for i := 0; i < 50; i++ {
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
				t.Log(err)
			}
		}

	}(client)

	go func(client *ent.Client) {
		defer wg.Done()
		for i := 0; i < 50; i++ {
			_, err := client.Class.Create().
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
				t.Log(err)
			}
		}
	}(client)

	wg.Wait()
	t.Log("success")
}

func BenchmarkQueryUser(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := client.User.Query().All(context.Background())
			if err != nil {
				b.Fail()
			}
		}
	})
}

func BenchmarkCreateUser(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
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
				b.Fail()
			}
		}
	})
}

func BenchmarkCreateUserTx(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			tx, err := client.Tx(context.Background())
			if err != nil {
				b.FailNow()
			}
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
				b.FailNow()
			}
			tx.Commit()
		}
	})
}

func BenchmarkCreateUserTxParallelQuery(b *testing.B) {
	b.Run("group", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				tx, err := client.Tx(context.Background())
				if err != nil {
					b.FailNow()
				}
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
					b.FailNow()
				}
				tx.Commit()
			}
		})
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				_, err := client.User.Query().All(context.Background())
				if err != nil {
					b.FailNow()
				}
			}
		})
	})
}

func BenchmarkCreateUser2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			for pb.Next() {
				_, err := client.User.Create().
					SetName(randStringBytes(5)).
					SetNickName(randStringBytes(7)).
					Save(context.Background())
				if err != nil {
					b.Log(err)
				}
			}
		}()

		go func() {
			defer wg.Done()
			for pb.Next() {
				_, err := client.Class.Create().
					SetName(randStringBytes(5)).
					SetClassName(randStringBytes(7)).
					Save(context.Background())
				if err != nil {
					b.Log(err)
				}
			}
		}()

		wg.Wait()
	})
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
