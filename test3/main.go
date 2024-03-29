package main

import (
	"fmt"
)

type Student struct {
	First_Name string
	Last_Name  string
	DOB        string
}

type Address struct {
	House_No int
	Village  string
	District string
	State    string
}

type SocialProfile struct {
	Facebook  string
	Instagram string
}

type ContactDetails struct {
	Mobile_No int
	Gmail_Id  string
}

type School struct {
	School_Name string
	Class_Name  int
}

type All_Data struct {
	Student        Student
	Address        Address
	SocialProfile  SocialProfile
	ContactDetails ContactDetails
	School         School
}

func main() {
	student := Student{
		First_Name: "Pawan",
		Last_Name:  "Sharma",
		DOB:        "03/08/2003",
	}

	address := Address{
		House_No: 123,
		Village:  "Kheri",
		District: "Rohtak",
		State:    "Haryana",
	}

	socialProfile := SocialProfile{
		Facebook:  "facebook",
		Instagram: "instagram",
	}

	contactDetails := ContactDetails{
		Mobile_No: 7988323110,
		Gmail_Id:  "pawan@gmail.com",
	}

	school := School{
		School_Name: "govt",
		Class_Name:  10,
	}

	allData := []All_Data{
		{
			Student:        student,
			Address:        address,
			SocialProfile:  socialProfile,
			ContactDetails: contactDetails,
			School:         school,
		},
	}

	fmt.Println(allData)

}
