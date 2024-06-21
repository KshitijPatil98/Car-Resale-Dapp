package models

type Car struct {
	Name        string `json:"name" binding:"required"`
	Id          int    `json:"id" binding:"required"`
	Owner       string `json:"owner" binding:"required"`
	InsuranceId string `json:"insuranceId" binding:"required"`
	Miles       int    `json:"miles" binding:"required"`
}

var cars []Car = []Car{}

func (c Car) Register() {

	cars = append(cars, c)

}

func GetAll() []Car {
	return cars
}
