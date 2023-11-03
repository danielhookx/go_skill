package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/logx"
	"k8s.io/apimachinery/pkg/util/json"
)

func TestMarshalBytes(t *testing.T) {
	type Text struct {
		Content []byte `json:"content"`
	}

	text := &Text{
		Content: []byte("你好"),
	}
	data, err := json.Marshal(text)
	assert.Nil(t, err)
	t.Log(string(data))
}

func TestPrint(t *testing.T) {
	logx.Info("hello")
}
