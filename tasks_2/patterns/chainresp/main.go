package main

import "github.com/vlasove/materials/tasks_2/patterns/chainresp/hospital"

func main() {
	cashierHandler := &hospital.Cashier{}
	medicalHandler := hospital.NewMedical(cashierHandler)
	doctorHandler := hospital.NewDoctor(medicalHandler)
	receptionHandler := hospital.NewReception(doctorHandler)

	patient := hospital.NewPatient("Vasya")
	receptionHandler.Execute(patient)
}
