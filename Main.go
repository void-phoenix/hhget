package main

import (
	"fmt"
	"strconv"
)

const(
	vacanciesBasePath = "https://api.hh.ru/vacancies?search_field=name&search_field=description&specialization=1&period=1"
)

func GetAllItems(path string) []*Vacancy {
	var vacancies []*Vacancy
	count := 0
	for {
		vacanciesResult, err := GetVacanciesResult(path + "&page=" + strconv.Itoa(count))
		if err != nil {
			fmt.Println("Error getting response... ", err.Error())
			break
		}
		vacancies = append(vacancies, ParseVacanciesResult(vacanciesResult)...)

		if vacanciesResult.Page == vacanciesResult.Pages || vacanciesResult.Pages == count {
			break
		}
		count++
	}
	return vacancies
}


func main() {

	vacancies := GetAllItems(vacanciesBasePath+"&text=java")
	for _, item := range vacancies{
		fmt.Println(item.Name)
	}

	fmt.Println("Total: ", len(vacancies))

}
