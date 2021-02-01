package main

import (
	"fmt"
)

// Man is man
type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

// NewMan construct Man
func NewMan(a, b, c string, d, e int) Man {
	return Man{
		Name:     a,
		LastName: b,
		Age:      d,
		Gender:   c,
		Crimes:   e,
	}
}

func main() {

	bankRobbers := make(map[string]Man)

	bankRobbers["Jesse"] = NewMan("Jesse", "James", "male", 32, 40)
	bankRobbers["Robert"] = NewMan("Robert", "Parker", "male", 34, 35)
	bankRobbers["Clyde"] = NewMan("Clyde", "Barrow", "male", 24, 30)
	bankRobbers["Bonnie"] = NewMan("Bonnie", "Parker", "female", 22, 25)
	bankRobbers["Charles"] = NewMan("Charles", "Floyd", "male", 29, 20)
	bankRobbers["John"] = NewMan("John", "Dillinger", "male", 30, 19)
	bankRobbers["Lester"] = NewMan("Lester", "Gillis", "male", 26, 18)
	bankRobbers["William"] = NewMan("William", "Sutton", "male", 46, 110)
	bankRobbers["Patricia"] = NewMan("Patricia", "Hearst", "female", 21, 8)

	var suspects [3][]string

	//gangsters
	suspects[0] = []string{"Jesse", "Robert", "Clyde", "Bonnie", "Charles"}
	suspects[1] = []string{"Charles", "John", "Lester", "William", "Patricia"}

	//civilian
	suspects[2] = []string{"Joe", "Richard", "Cody", "Bennie", "Charlotta", "Johnie", "Leopold", "Wolfgang", "Patric"}

	for i := 0; i < len(suspects); i++ {
		challenge(suspects[i], bankRobbers)
	}

}

func challenge(suspects []string, challengers map[string]Man) {

	var winner string

	fmt.Println("")
	fmt.Print("Список имён: ")
	fmt.Println(fmt.Sprintf("%q\n", suspects))

	for _, name := range suspects {

		_, ok := challengers[name]
		if !ok {
			continue
		}

		if challengers[name].Crimes > challengers[winner].Crimes {
			winner = name
		}

	}
	if winner == "" {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым")
		fmt.Println("")
		return
	}
	fmt.Printf("Подозреваемый: %s, %s\n", challengers[winner].Name, challengers[winner].LastName)
	fmt.Printf("Возраст:	%d\n", challengers[winner].Age)
	fmt.Printf("Пол:		%s\n", challengers[winner].Gender)
	fmt.Printf("Преступлений:	%d\n", challengers[winner].Crimes)
	fmt.Println("")

}
