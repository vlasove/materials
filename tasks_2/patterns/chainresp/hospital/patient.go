package hospital

type Patient struct {
	Name              string
	RegistrationDone  bool
	DoctorCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
}

func NewPatient(name string) *Patient {
	return &Patient{
		Name: name,
	}
}
