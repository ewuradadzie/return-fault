package main

/* 
This script returns the fault name and description of a fault code. 
The fault codes and their corresponding names and descriptions are stored in a csv file.
Make sure the csv file is in the same folder as this script.
Enter the code number as a parameter in the function returnFault(code), and it will return its name and description.
*/

// Importing supporting packages

import (
	"fmt"
	"strconv"
	"encoding/csv"
	"os"
)

// Main function

func main() {

	returnFault(code) // Enter the code number here

	}

// returnFault function

func returnFault(code int) {

	// 1. Extract the csv file

	csvFile, err := os.Open("./faultCodes.csv") // Open the file
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}
	reader := csv.NewReader(csvFile) // Read the file

	codes, _ := reader.ReadAll() // Extract the file

	// 2. Loop through the file

	lng := len(codes)
	for  i := 1; i < lng; i++ {
		faultCode := codes[i][0]
		faultName := codes[i][1]
		faultDesc := codes[i][2]

		ifaultCode, err := strconv.Atoi(faultCode) // Converting str to int
		if err != nil {
			fmt.Println(ifaultCode)
		}

	// 3. Determine name and description

		if ifaultCode == code {
			fmt.Println("Fault Code: ", faultCode)
			fmt.Println("Fault Name: ", faultName)
			fmt.Println("Description: ", faultDesc)

			// Continue or quit the process

			fmt.Println("Press 1 to enter another code, or any character to quit.")

			var s string
			fmt.Scanln(&s)
			if s == "1" {
				returnFault()
			} else {
				fmt.Println("Goodbye.")
				break
			}
		}
	}
}
