package constant

const (
	// KB 1KB等于1024B
	KB int64 = 1024
	// MB 1MB等于1024KB
	MB int64 = 1024 * KB
	// GB 1GB等于1024MB
	GB int64 = 1024 * MB
	// TB 1TB等于1024GB
	TB int64 = 1024 * GB
	// PB 1PB等于1024TB
	PB int64 = 1024 * TB
)

const (
	// MempoolSize 定义内存池的大小为4MB
	MempoolSize = 8 * MB
)

const (
	// EmptyString 空字符串
	EmptyString = ""
	// ErrNotFile 这个请求格式非文件
	ErrNotFile = "The Style Isn't File"
)
