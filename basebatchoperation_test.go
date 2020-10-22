package batchoperation

import (
	"fmt"
	"testing"
	"time"
)

type BTest struct {
	BaseBatchOperation
}

func (b *BTest) BatchProcessor(isAsync bool, msg []interface{}) (err error) {
	fmt.Println("bTst", msg)
	return nil
}
func TestNewBatchOperation(t *testing.T) {
	inQueue1 := make(chan interface{}, 0)
	b1 := &BTest{}
	s, _ := NewServer(inQueue1, 5*time.Second, 1000, false, b1)
	go s.Start()

	go func() {
		for i := 0; i < 10000; i++ {
			inQueue1 <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()
	select {}
}
