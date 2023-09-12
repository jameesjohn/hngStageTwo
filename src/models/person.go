package models

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type Person struct {
	Id        uint64    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" `
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PersonModel struct {
	Db *gorm.DB
}

func (p PersonModel) Create(person Person) (*Person, error) {
	result := p.Db.Create(&person)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &person, nil
}

func (p PersonModel) Update(person Person) (*Person, error) {
	result := p.Db.Model(&Person{}).Where(&Person{Id: person.Id}).Updates(person)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &person, nil
}

func (p PersonModel) GetAll() ([]Person, error) {
	var persons []Person
	result := p.Db.Model(&Person{}).Limit(30).Find(&persons)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return persons, nil
}

func (p PersonModel) Delete(personId uint64) error {
	result := p.Db.Delete(&Person{Id: personId})

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return result.Error
}

func (p PersonModel) Find(personId uint64) (*Person, error) {
	var person Person
	result := p.Db.Where(&Person{Id: personId}).First(&person)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &person, nil
}
