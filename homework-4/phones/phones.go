package phones

import "fmt"

type Contact struct {
	name    string
	surname string
}

type phones []int

func (p phones) ViewList() {
	for i, phone := range p {
		fmt.Printf("\t %v) %v \n", i, phone)
	}
}

func Process() {
	addressBook := make(map[Contact]phones)

	contact1 := Contact{"Миша", "Пашин"}
	contact2 := Contact{"Вася", "Никитин"}

	addressBook[contact1] = phones{78293467382}
	addressBook[contact2] = phones{89167253764, 89635437382}

	for contact, ph := range addressBook {
		fmt.Println(contact.name, contact.surname)
		ph.ViewList()
	}
}
