package main

import (
	csvreader "WhatToCook/csvreader"
	databasereader "WhatToCook/databasereader"
	"fmt"
)

func main() {
	fmt.Println("Initiate the CSV reader..")
	food := csvreader.ReadCSV()
	databasereader.ReadDatabaseAndDiff()
	fmt.Println(food)
	//Read out what is in the database 	
}