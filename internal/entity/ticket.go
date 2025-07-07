package entity

type TicketType uint8

const (
	NORMAL = iota
	CANARY_TEST
	AB_TEST
)

type TicketStatus uint8

const (
	PENDING = iota
	RUNNING
	COMPLETE
)

type TicketEnv uint8

const (
	DEV  = iota //开发环境
	STG         //测试环境
	PROD        //生产环境
)

type Ticket struct {
	Base
	Name      string       `gorm:"column:name"`       //容器名
	ImageName string       `gorm:"column:image_name"` // 镜像名
	Type      TicketType   `gorm:"column:type"`       // 发布类型 AB 灰度 普通
	Status    TicketStatus `gorm:"column:status"`     //运行状态
	Tag       string       `gorm:"column:tag"`        // 版本号
	Env       TicketEnv    `gorm:"column:env"`        //发布环境
}

func (Ticket) TableName() string {
	return "tickets"
}
