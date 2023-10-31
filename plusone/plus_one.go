package plusone

import "time"

//go:generate mockgen -source=plus_one.go -destination=./mock/plus_one_mock.go

type Car struct {
	VIN             string
	Model           string
	FabricationYear time.Time
	Impounded       bool
}

type DbRepo interface {
	GetNumber() int
	FindCarByVIN(vin string) Car
	UpdateCar(Car)
}

type SomethingThatTakesFunctionality interface {
	DoSomething(vin string, doSthWith func(registrationNumber string))
}

type PlusDataFromDB struct {
	DbRepo    DbRepo
	Something SomethingThatTakesFunctionality
}

func (service PlusDataFromDB) AddWithDataFromDB(x int) int {
	return x + service.DbRepo.GetNumber()
}

func (service PlusDataFromDB) ImpoundCar(vin string) {
	car := service.DbRepo.FindCarByVIN(vin)
	car.Impounded = true
	service.DbRepo.UpdateCar(car)
}

func (service PlusDataFromDB) SendFunctionality(vin string) string {
	var regNumber string
	service.Something.DoSomething(vin, func(registrationNumber string) {
		regNumber = registrationNumber
	})
	return regNumber
}
