package function

import (
	"testing"
)

func TestNewBaseStruct(t *testing.T) {
	var quote ISuper
	var direct ISuper

	quote = NewQuote(2, NewBaseStruct("world"))
	direct = NewDirect(2, NewBaseStruct("world"))

	t.Logf("quote: val[%d], stc[%s-%p]\n", quote.getValue(), quote.getStruct().val, quote.getStruct())
	t.Logf("direct: val[%d], stc[%s-%p]\n", direct.getValue(), direct.getStruct().val, direct.getStruct())

	quote.selfMultiplying()
	direct.selfMultiplying()
	quote.setStructVal("hello")
	direct.setStructVal("hello")
	t.Logf("quote: val[%d], stc[%s-%p]\n", quote.getValue(), quote.getStruct().val, quote.getStruct())
	t.Logf("direct: val[%d], stc[%s-%p]\n", direct.getValue(), direct.getStruct().val, direct.getStruct())
}
