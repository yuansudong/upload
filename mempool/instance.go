package mempool

var _Inst *Engine = _NewEngine()

// Get 获取内存池实例
func Get() *Engine {
	return _Inst
}
