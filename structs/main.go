package main

import "fmt"

/*
 * Creating struct is like creating a datatype of your own;
 * here we have a Company struct that contains basic information of the Company
 * that we will use to extend, change and update it.
*
*/
type Company struct {
	Name    string
	Address string
	CEO     Employee
}

type Employee struct {
	Name  string
	Phone int64
	Id    int8
}

/*
 * SO can also create methods to update our structs,
 * here we will change the CEO of the Company
 */
func (c *Company) SetCEO(employee *Employee) {
	c.CEO = *employee
}

func main() {
	Abdusalam := Employee{
		Name:  "Abdusalam",
		Phone: 25199000000,
		Id:    1,
	}

	Google := Company{
		Name:    "Google",
		Address: "CA, San francisco",
		CEO:     Abdusalam,
	}

	fmt.Println(Google)

	Mahamed := Employee{
		Name:  "Mahamed",
		Phone: 251111111111111,
		Id:    2,
	}

	Google.SetCEO(&Mahamed)

	fmt.Println(Google)
	maimaimaimaimaimaimainnnnnnn
}
