package models

import "cliTest/models/postgres"

type Product struct {
	ID    int    `json:"ID" gorm:"primary_key"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (p *Product) GetById() error {
	return postgres.GetDB().Table("products").Where("id = ?", p.ID).First(p).Error
}
