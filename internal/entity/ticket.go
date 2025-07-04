package entity

type TicketType uint8

const (
	CANARY_TEST = iota
	AB_TEST
)

type TicketStatus uint8

const (
	PENDING = iota
	RUNNING
	COMPLETE
)

type Ticket struct {
	Base
	Name      string       `gorm:"column:name"`
	ImageName string       `gorm:"column:image_name"`
	Type      TicketType   `gorm:"column:type"`
	Status    TicketStatus `gorm:"column:status"`
}

func (Ticket) TableName() string {
	return "tickets"
}
