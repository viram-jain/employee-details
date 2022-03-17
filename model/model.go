package model

type Employee struct {
	EmpID   string
	Name    string
	Age     int
	Manager string
	Position    string
}

type Employees []Employee

var EmployeeJSON = Employees{
	Employee{
		"KIN/00396",
		"Viram Jain",
		22,
		"Gokul Palanisamy",
		"Software Engineer Professional",
	},
	Employee{
		"KIN/00391",
		"Vikas T Shankar",
		22,
		"Madhusudhan Aithal",
		"Software Engineer Professional",
	},
	Employee{
		"KIN/00393",
		"Deepen Shrestha",
		22,
		"Gokul Palanisamy",
		"Software Engineer Professional",
	},
}
