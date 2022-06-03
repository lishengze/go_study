package main

import "fmt"

type iGun interface {
	GetName() string
	GetPower() int
	SetName(string)
	SetPower(int)
}

type Gun struct {
	Name  string
	Power int
}

func (gun *Gun) GetName() string {
	return gun.Name
}

func (gun *Gun) GetPower() int {
	return gun.Power
}

func (gun *Gun) SetName(name string) {
	gun.Name = name
}

func (gun *Gun) SetPower(power int) {
	gun.Power = power
}

type AK47 struct {
	Gun
}

// func GetName(gun *AK47) string {
// 	return gun.Gun Name
// }

// func GetPower(gun *AK47) int {
// 	return gun.Power
// }

// func SetName(gun *AK47, name string) {
// 	gun.Name = name
// }

// func SetPower(gun *AK47, power int) {
// 	gun.Power = power
// }

type Musket struct {
	Gun
}

func NewMusket() iGun {
	return &Musket{
		Gun: Gun{
			Name:  "Musket",
			Power: 10,
		},
	}
}

func NewAk47() iGun {
	return &AK47{
		Gun: Gun{
			Name:  "Ak47",
			Power: 5,
		},
	}
}

func get_gun(gun_type string) (iGun, error) {
	if gun_type == "AK47" {
		return NewAk47(), nil
	} else if gun_type == "Musket" {
		return NewMusket(), nil
	} else {
		return nil, fmt.Errorf("Type is Error")
	}
}

func PrintDetails(gun iGun) {
	fmt.Printf("Name: %s, Power: %d\n", gun.GetName(), gun.GetPower())
}

// func main() {
// 	ak47, err := get_gun("AK47")

// 	if err == nil {
// 		PrintDetails(ak47)
// 	}

// 	musket, err := get_gun("Musket")

// 	if err == nil {
// 		PrintDetails(musket)
// 	}
// }
