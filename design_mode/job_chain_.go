package main

import "fmt"

type patient struct {
	Name  string
	Stage string
}

type worker interface {
	execute(*patient)
	setNext(worker)
}

type baseDoctor struct {
	name       string
	nextWorker worker
}

func (doctor *baseDoctor) execute(cur_patient *patient) {
	cur_patient.Stage = doctor.name
	fmt.Printf("Patient.Name: %s, Patient.Stage: %s \n", cur_patient.Name, cur_patient.Stage)

	if doctor.nextWorker != nil {
		doctor.nextWorker.execute(cur_patient)
	} else {
		fmt.Println("Work Over!")
	}
}

func (doctor *baseDoctor) setNext(next_worker worker) {
	doctor.nextWorker = next_worker
}

type DoctorA struct {
	baseDoctor
}

func (doctor *DoctorA) execute(cur_patient *patient) {
	cur_patient.Stage = doctor.name
	fmt.Printf("DoctorA Patient.Name: %s, Patient.Stage: %s \n", cur_patient.Name, cur_patient.Stage)

	if doctor.nextWorker != nil {
		doctor.nextWorker.execute(cur_patient)
	} else {
		fmt.Println("")
	}
}

// func (doctor *DoctorA) setNext(next_worker worker) {
// 	doctor.nextWorker = next_worker
// }

type DoctorB struct {
	baseDoctor
}

func (doctor *DoctorB) execute(cur_patient *patient) {
	cur_patient.Stage = doctor.name

	fmt.Printf("DoctorB Patient.Name: %s, Patient.Stage: %s \n", cur_patient.Name, cur_patient.Stage)

	if doctor.nextWorker != nil {
		doctor.nextWorker.execute(cur_patient)
	} else {
		fmt.Println("")
	}
}

// func main() {
// 	patient_ := patient{
// 		Name: "Tom",
// 	}

// 	doctor_a := DoctorA{
// 		baseDoctor: baseDoctor{
// 			name: "DoctorA",
// 		},
// 	}

// 	doctor_b := DoctorB{
// 		baseDoctor: baseDoctor{
// 			name: "DoctorB",
// 		},
// 	}

// 	doctor_a.setNext(&doctor_b)

// 	doctor_a.execute(&patient_)
// }
