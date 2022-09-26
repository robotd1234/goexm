package main

import "fmt"

type Company struct {
	CompanyName string
	CompanyAddr string
}

type Staff struct {
	Name      string
	Age       int
	Positions string
	Gender    string
	Company
}

func (person *Staff) FmtProfile() {

	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("Sex: %s\n", person.Gender)

}

func main() {

	var p int
	var v *int
	v = &p
	*v = 11
	fmt.Println(*v)
	fmt.Printf("v is %d\n", v)

	companyA := Company{
		CompanyName: "Microsoft",
		CompanyAddr: "上海市紫月路888号",
	}

	staffA := Staff{
		Name:      "张三",
		Age:       18,
		Positions: "上海",
		Gender:    "男",
		Company:   companyA,
	}

	fmt.Println(staffA.Name, staffA.CompanyName)
	fmt.Println(staffA.Name, staffA.Company.CompanyName)
}
