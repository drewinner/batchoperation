package batchoperation

import (
	"errors"
	"time"
)

const (
	DefaultSize  = 10
	DefaultDelay = 5
)

type Server struct {
	InQueue   chan interface{}    //监听的队列 -- 消息输入队列
	OpQueue   []interface{}       //消息队列、最终需要进行操作的队列
	Size      int                 //消息队列的最大长度
	Delay     time.Duration       //延迟时间--最长等待时间
	Operation IFaceBatchOperation //操作句柄
	IsAsync   bool                //同步还是异步操作 true异步、false同步
}

func NewServer(inQueue chan interface{}, delay time.Duration, size int, isAsync bool, operation IFaceBatchOperation) (s *Server, e error) {
	if inQueue == nil {
		return nil, errors.New("inQueue nil")
	}
	if size == 0 {
		size = DefaultSize
	}
	if delay == 0 {
		delay = DefaultDelay
	}
	return &Server{
		InQueue:   inQueue,
		OpQueue:   make([]interface{}, 0),
		Size:      size,
		Delay:     delay,
		Operation: operation,
		IsAsync:   isAsync,
	}, nil
}

func (s Server) Start() {
	delayTimer := time.NewTimer(0)
	if !delayTimer.Stop() {
		<-delayTimer.C
	}
	defer delayTimer.Stop()
	for {
		select {
		case msg := <-s.InQueue:
			s.OpQueue = append(s.OpQueue, msg)
			if len(s.OpQueue) != s.Size {
				if len(s.OpQueue) == 1 {
					delayTimer.Reset(s.Delay)
				}
				break
			}
			if err := s.Operation.BatchProcessor(s.IsAsync, s.OpQueue); err != nil {
				s.Operation.ErrorHandler(err, s.OpQueue)
			}
			if !delayTimer.Stop() {
				<-delayTimer.C
			}
			s.OpQueue = make([]interface{}, 0)
		case <-delayTimer.C:
			if err := s.Operation.BatchProcessor(s.IsAsync, s.OpQueue); err != nil {
				s.Operation.ErrorHandler(err, s.OpQueue)
			}
			s.OpQueue = make([]interface{}, 0)
		}
	}
}
