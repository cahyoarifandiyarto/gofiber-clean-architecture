package entity

type User struct {
	ID       uint64 `gorm:"column:id;primaryKey"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func (u *User) TableName() string {
	return "users"
}
