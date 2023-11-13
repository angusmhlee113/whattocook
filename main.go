package main

import (
	csvreader "WhatToCook/csvreader"
	envreader "WhatToCook/envreader"
	"fmt"
)

func main() {
	fmt.Println("Initiate the CSV reader..")
	food := csvreader.ReadCSV()
	envreader.ReadEnv()
	fmt.Println(food)
	//Read out what is in the database 	
}