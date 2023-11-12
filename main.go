package main

import (
	csvreader "WhatToCook/csvreader"
	"fmt"
)

func main() {
	fmt.Println("Initiate the CSV reader..")
	food := csvreader.ReadCSV()
	fmt.Println(food)
	//Read out what is in the database 	
}