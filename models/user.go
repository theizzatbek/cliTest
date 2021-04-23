package models

type User struct {
	ID          int    `json:"ID" gorm:"primary_key"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func (u *User) GetById() error {
	return db.Table("users").Where("id = ?", u.ID).First(u).Error
}
