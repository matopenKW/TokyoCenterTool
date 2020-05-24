package domain

type Carfare struct {
	SeqNo      int `gorm:"AUTO_INCREMENT"`
	UserId     string
	Date       string
	StartPoint string
	EndPoint   string
	Price      int
}
