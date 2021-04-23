package models

type Product struct {
	ID    int    `json:"ID" gorm:"primary_key"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (p *Product) GetById() error {
	return db.Table("products").Where("id = ?", p.ID).First(p).Error
}
