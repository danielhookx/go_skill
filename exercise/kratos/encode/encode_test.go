package encode_test

import (
	"testing"

	user "github.com/danielhookx/go_skill/exercise/kratos/encode/type"
	"github.com/go-kratos/kratos/v2/encoding"
	_ "github.com/go-kratos/kratos/v2/encoding/json"
	_ "github.com/go-kratos/kratos/v2/encoding/proto"
	"golang.org/x/exp/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func randomFields(n int) []*user.Field {
	ret := make([]*user.Field, n)
	for i := 0; i < n; i++ {
		ret[i] = randomField()
	}
	return ret
}

func randomField() *user.Field {
	return &user.Field{
		F1: randStringBytes(20),
		F2: int32(rand.Intn(100)),
		F3: randStringBytes(20),
		F4: randStringBytes(20),
		F5: int32(rand.Intn(100)),
		F6: int32(rand.Intn(100)),
		F7: randStringBytes(20),
	}
}

var testUser *user.User

func init() {
	testUser = &user.User{
		Field1:  randStringBytes(20),
		Field2:  randStringBytes(20),
		Field3:  randStringBytes(20),
		Field4:  randStringBytes(20),
		Field5:  int32(rand.Intn(100)),
		Field6:  int32(rand.Intn(100)),
		Field7:  randStringBytes(20),
		Field8:  randStringBytes(20),
		Field9:  int32(rand.Intn(100)),
		Field10: int32(rand.Intn(100)),
		Field11: int32(rand.Intn(100)),
		Field12: randStringBytes(20),
		Fs:      randomFields(20),
	}
}

func BenchmarkMarshal(b *testing.B) {
	b.Run("Json", func(b *testing.B) {
		in := &user.User{}
		contentType := "json"
		_, err := encoding.GetCodec(contentType).Marshal(in)
		if err != nil {
			b.FailNow()
		}
	})
	b.Run("Proto", func(b *testing.B) {
		in := &user.User{}
		contentType := "proto"
		_, err := encoding.GetCodec(contentType).Marshal(in)
		if err != nil {
			b.FailNow()
		}
	})

	b.Run("Json-long", func(b *testing.B) {
		contentType := "json"
		_, err := encoding.GetCodec(contentType).Marshal(testUser)
		if err != nil {
			b.FailNow()
		}
	})

	b.Run("Proto-long", func(b *testing.B) {
		contentType := "proto"
		_, err := encoding.GetCodec(contentType).Marshal(testUser)
		if err != nil {
			b.FailNow()
		}
	})
}

func BenchmarkMarshalAndUmmarshal(b *testing.B) {
	b.Run("Json", func(b *testing.B) {
		in := &user.User{}
		contentType := "json"
		data, err := encoding.GetCodec(contentType).Marshal(in)
		if err != nil {
			b.FailNow()
		}

		out := &user.User{}
		err = encoding.GetCodec(contentType).Unmarshal(data, out)
		if err != nil {
			b.FailNow()
		}
	})

	b.Run("Proto", func(b *testing.B) {
		in := &user.User{}
		contentType := "proto"
		data, err := encoding.GetCodec(contentType).Marshal(in)
		if err != nil {
			b.FailNow()
		}

		out := &user.User{}
		err = encoding.GetCodec(contentType).Unmarshal(data, out)
		if err != nil {
			b.FailNow()
		}
	})

	b.Run("Json-long", func(b *testing.B) {
		contentType := "json"
		data, err := encoding.GetCodec(contentType).Marshal(testUser)
		if err != nil {
			b.FailNow()
		}

		out := &user.User{}
		err = encoding.GetCodec(contentType).Unmarshal(data, out)
		if err != nil {
			b.FailNow()
		}
	})

	b.Run("Proto-long", func(b *testing.B) {
		contentType := "proto"
		data, err := encoding.GetCodec(contentType).Marshal(testUser)
		if err != nil {
			b.FailNow()
		}

		out := &user.User{}
		err = encoding.GetCodec(contentType).Unmarshal(data, out)
		if err != nil {
			b.FailNow()
		}
	})
}
