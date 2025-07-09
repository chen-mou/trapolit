package entity

type Process struct {
	Base
	Name string
}

type ProcessNode struct {
	Base
	ProcessId uint64 `gorm:"column:process_id"`
	Env       string
	Image     string
}
