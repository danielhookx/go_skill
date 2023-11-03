package channel

import "testing"

func TestCloseChannel(t *testing.T) {
	ch := make(chan int, 20)

	//fill channel
	for i := 1; i <= 20; i++ {
		ch <- i
	}

	close(ch)
	//
	//for i:=1; i<=22;i++{
	//	v := <-ch
	//	t.Log(v)
	//}

	for v := range ch {
		t.Log(v)
	}
}
