1. 作用是监听特定channel、根据延迟时间、设定条数、批量处理
    1. 达到指定延迟时间、条数无论是否达到指定条数、都会调用业务函数
    2. 达到指定条数、无论是否到达延迟时间、都会调用业务函数
    3. 业务函数内的slice为interface类型、具体业务具体处理
2. module
     `github.com/drewinner/batchoperation`
2. go get github.com/drewinner/batchoperation

3. 测试用例：
```go
type BTest struct {
	batchoperation.BaseBatchOperation
}

func (b *BTest) BatchProcessor(isAsync bool, msg []interface{}) (err error) {
    fmt.Println("获取批量数据。业务逻辑...")
	fmt.Println("bTst", msg)
	return nil
}
func (b *BTest) ErrorHandler(err error, msg []interface{}) {
	fmt.Println("错误处理逻辑....")
}

func main() {
    //1. 初始化监听队列
	inQueue1 := make(chan interface{}, 0)
    //2. 创建对象
	b1 := &BTest{}
    //3. 创建服务参数：
        //1. 参数1 队列、作为producer
        //2. 参数2 延迟时间
        //3. 参数3 指定大小、调用处理逻辑
        //4. 是否为异步处理 true异步 false非异步
        //5. interface对象
	s, _ := batchoperation.NewServer(inQueue1, 5*time.Second, 1000, false, b1)
    //4. 启动服务
	go s.Start()
    //5. 测试生产内容到inQueue1中
	go func() {
		for i := 0; i < 10000; i++ {
			inQueue1 <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()
	select {}
}
```