package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type mobile struct {
	CompanyName string      `json:"company_name"`
	LaunchYear  int         `json:"launch_year"`
	Color       string      `json:"color"`
	ScreenType  string      `json:"screen_type"`
	ScreenSize  float32     `json:"screen_size"`
	Ram         float32     `json:"ram"`
	Storage     float64     `json:"storage"`
	Price       float64     `json:"price"`
	Name        interface{} `json:"name"`
}

func main() {
	allMobile := []mobile{

		{CompanyName: "Moto",
			LaunchYear: 2022,
			Color:      "Lite Blue",
			ScreenType: "OLED",
			ScreenSize: 6.7,
			Ram:        8.0,
			Storage:    128.0,
			Price:      1299.99,
			Name:       "edge 40 neo"},

		{
			CompanyName: "Apple",
			LaunchYear:  2022,
			Color:       "Lite Blue",
			ScreenType:  "OLED",
			ScreenSize:  5.7,
			Ram:         8.0,
			Storage:     128.0,
			Price:       1299.99,
			Name:        "Iphone 13",
		},

		{
			CompanyName: "Samsung",
			LaunchYear:  2023,
			Color:       "Blue",
			ScreenType:  "OLED",
			ScreenSize:  6.7,
			Ram:         8.0,
			Storage:     256.0,
			Price:       1299.99,
			Name:        "S23",
		},

		{
			CompanyName: "MI",
			LaunchYear:  2022,
			Color:       "Black",
			ScreenType:  "OLED",
			ScreenSize:  6.7,
			Ram:         8.0,
			Storage:     128.0,
			Price:       1999.99,
			Name:        "Prime 8",
		},

		{
			CompanyName: "MI",
			LaunchYear:  2022,
			Color:       "Blue",
			ScreenType:  "OLED",
			ScreenSize:  6.7,
			Ram:         8.0,
			Storage:     128.0,
			Price:       1299.99,
			Name:        "V27",
		},
	}
	data, err := json.MarshalIndent(allMobile, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
