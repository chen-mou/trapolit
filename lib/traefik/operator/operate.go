package operator

type ProviderType string

const (
	FILE = "file"
	ETCD = "ETCD"
)

type Provider struct {
	Type ProviderType
	Name string
}

type Service struct {
}

type Router struct {
	Path          string
	Name          string
	TargetService string
}

// Operator 对于不同的Provider可能会有不同的添加实现方法
// 现在默认提供etcd 与 file两种方式
type Operator interface {
	AddProvider()
	AddService(name string, targetUrl []string)
	AddRouter(name, originUrl, serviceName string)
	DelProvider(name string)
	DelRouter(name string)
	DelService(name string)
	Flush() error // 将修改提交到目标中
}
