package mempool

import (
	"sync"
	"upload/constant"
)

// Engine 用于描述一个引擎
type Engine struct {
	_BfPOOL *sync.Pool
}

// _NewEngine 用于新建立一个mempool的管理者
func _NewEngine() *Engine {
	mInst := new(Engine)
	return mInst._InitPool()
}

func (e *Engine) _InitPool() *Engine {
	e._BfPOOL = &sync.Pool{
		New: func() interface{} {
			return make([]byte, constant.MempoolSize)
		},
	}
	return e
}

// GetBytes 从字节池中获得
func (e *Engine) GetBytes() []byte {
	return e._BfPOOL.Get().([]byte)
}

// PutBytes 放回到字节池
func (e *Engine) PutBytes(body []byte) {
	if int64(len(body)) != constant.MempoolSize {
		return
	}
	e._BfPOOL.Put(body)
}
