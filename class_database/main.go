package main

import (
	"fmt"
	"strconv"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ClassMember struct {
	ID         int
	Name       string
	Occupation string
	Address    string
	Reason     string
}

func main() {
	// get input
	var getInput string
	fmt.Print("Enter name or ID: ")
	fmt.Scan(&getInput)

	// check if it is nameInput or idInput variable
	var nameInput string
	idInput, err := strconv.Atoi(getInput)
	if err != nil {
		nameInput = cases.Title(language.Und).String(getInput)
	}

	// set flag
	nameFound := false
	idFound := false

	// declare array of structs
	kelasGolang := []ClassMember{
		{1, "Annisa", "Karyawan Swasta", "Jakarta", "Alasan Annisa"},
		{2, "Rahmah", "BUMN", "Jakarta", "Alasan Rahmah"},
		{3, "Najya", "Banking", "Singapore", "Alasan Najya"},
		{4, "Farida", "Startup", "Bali", "Alasan Farida"},
	}

	// check if input matches any value in structs
	for _, member := range kelasGolang {
		if member.Name == nameInput || member.ID == idInput {
			nameFound = true
			idFound = true
			fmt.Println("ID        :", member.ID)
			fmt.Println("Nama      :", member.Name)
			fmt.Println("Pekerjaan :", member.Occupation)
			fmt.Println("Alamat    :", member.Address)
			fmt.Println("Alasan    :", member.Reason)
		}
	}

	if !nameFound || !idFound {
		fmt.Println("Data dengan nama/absen tsb tidak tersedia.")
	}
}
