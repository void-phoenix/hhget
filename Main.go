package main

import (
	"fmt"
)

const(
	vacanciesBasePath = "https://api.hh.ru/vacancies?search_field=name&search_field=description&specialization=1&period=1&"
)


func main() {

	vacanciesResult, err := GetVacanciesResult(vacanciesBasePath+"text=java")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(vacanciesResult.Page)
	fmt.Println(vacanciesResult.Pages)
	fmt.Println(vacanciesResult.Found)
	fmt.Println(vacanciesResult.PerPage)

}
