package csvreader

import (
	datatypes "WhatToCook/modals"
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSV() []datatypes.Food{ 
	fmt.Println("Reading the CSV file..")
	csvFile, err := os.Open("Menu.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(csvFile)
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var foods []datatypes.Food
	for line, each := range csvData {
		if (line >= 1){
			food := datatypes.Food{
				Meal: each[0],
				Cuisine: each[1],
				Name: each[2],
			}
			foods = append(foods, food)
		}
	}
	csvFile.Close()
	fmt.Println(foods)
	return foods
}