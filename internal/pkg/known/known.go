package known

const (
	// XRequestIDKey 用来定义 Gin 上下文中的键，代表请求的 uuid.
	XRequestIDKey = "X-Request-ID"

	// XEmailKey 用来定义 Gin 上下文的键，代表请求的所有者.
	XEmailKey = "X-Email"
)
