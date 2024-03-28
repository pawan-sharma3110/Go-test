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
	var first mobile

	first.CompanyName = "Moto"
	first.LaunchYear = 2022
	first.Color = "Lite Blue"
	first.ScreenType = "OLED"
	first.ScreenSize = 6.7
	first.Ram = 8.0
	first.Storage = 128.0
	first.Price = 1299.99
	first.Name = "edge 40 neo"

	var second mobile
	second.CompanyName = "Apple"
	second.LaunchYear = 2022
	second.Color = "Lite Blue"
	second.ScreenType = "OLED"
	second.ScreenSize = 5.7
	second.Ram = 8.0
	second.Storage = 128.0
	second.Price = 1299.99
	second.Name = "Iphone 13"

	var third mobile
	third.CompanyName = "Samsung"
	third.LaunchYear = 2023
	third.Color = " Blue"
	third.ScreenType = "OLED"
	third.ScreenSize = 6.7
	third.Ram = 8.0
	third.Storage = 256.0
	third.Price = 1299.99
	third.Name = "S23"

	var forth mobile
	forth.CompanyName = "MI"
	forth.LaunchYear = 2022
	forth.Color = "Black"
	forth.ScreenType = "OLED"
	forth.ScreenSize = 6.7
	forth.Ram = 8.0
	forth.Storage = 128.0
	forth.Price = 1299.99
	forth.Name = "Prime 8"

	var fifth mobile
	fifth.CompanyName = "Vivo"
	fifth.LaunchYear = 2022
	fifth.Color = "Blue"
	fifth.ScreenType = "OLED"
	fifth.ScreenSize = 6.7
	fifth.Ram = 8.0
	fifth.Storage = 128.0
	fifth.Price = 1299.99
	fifth.Name = "V27"

	fmt.Println(first)

	data, err := json.MarshalIndent(first, "   ", "")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(data))
	data2, err := json.MarshalIndent(second, "   ", "")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(data2))
	data3, err := json.MarshalIndent(third, "   ", "")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(data3))
	data4, err := json.MarshalIndent(forth, "   ", "")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(data4))
	data5, err := json.MarshalIndent(fifth, "   ", "")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(data5))

	// data, err := json.Marshal(first)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(string(data))
}
