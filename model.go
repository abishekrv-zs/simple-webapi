package main

type department struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type employee struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	PhoneNumber string     `json:"phoneNumber"`
	Dept        department `json:"department"`
}

var employees = []employee{
	{
		Id:          "1",
		Name:        "Abishek",
		PhoneNumber: "1234567890",
		Dept: department{
			Id:   "1",
			Name: "Software",
		},
	},
	{
		Id:          "2",
		Name:        "Kavin",
		PhoneNumber: "1234567891",
		Dept: department{
			Id:   "1",
			Name: "Software",
		},
	},
	{
		Id:          "3",
		Name:        "Kiren",
		PhoneNumber: "1234567892",
		Dept: department{
			Id:   "2",
			Name: "Finance",
		},
	},
	{
		Id:          "4",
		Name:        "Sujith",
		PhoneNumber: "1234567893",
		Dept: department{
			Id:   "3",
			Name: "Admin",
		},
	},
}
