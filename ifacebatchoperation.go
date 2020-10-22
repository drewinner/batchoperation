package batchoperation

type IFaceBatchOperation interface {
	//批量处理的方法
	BatchProcessor(isAsync bool, msg []interface{}) (err error)
	//错误处理
	ErrorHandler(err error, msg []interface{})
}
