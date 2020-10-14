package phones

import (
	"fmt"
	"sort"
)

type Contact struct {
	Name  string
	Phone int64
}

type Contacts []Contact

func (c Contacts) Len() int {
	return len(c)
}

func (c Contacts) Less(i, j int) bool {
	return c[i].Name < c[j].Name
}

func (c Contacts) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func Process() {

	items := Contacts{
		{"Иван", 7921123456},
		{"Pavel", 7551234567},
		{"Ольга", 7658744567},
		{"Alina", 7658444947},
		{"Алексей", 7718481947},
	}

	sort.Sort(items)

	fmt.Println(items)
}
