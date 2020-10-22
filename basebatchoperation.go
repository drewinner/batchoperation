package batchoperation

type BaseBatchOperation struct {
}

func (b *BaseBatchOperation) BatchProcessor(isAsync bool, msg []interface{}) (err error) {
	return nil
}

func (b *BaseBatchOperation) ErrorHandler(err error, msg []interface{}) {
	return
}
