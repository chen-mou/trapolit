package entity

type Role string

const (
	VIEWER = "view"
	WRITER = "write"
	SUPER  = "super"
)

type User struct {
	Base
	Username string
	Password string
	Role     Role //权限只有三个 读 写 都有
}

func (*User) TableName() string {
	return "user"
}
