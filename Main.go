package main

import (
	"fmt"
)


func main() {
	vacancies := GetAllItemsFor("java")
	for _, item := range vacancies{
		fmt.Println(item.Name)
	}

	fmt.Println("Total: ", len(vacancies))
}
