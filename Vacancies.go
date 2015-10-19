package main
import (
	"encoding/json"
	"net/http"
	"fmt"
	"strconv"
)

const(
	vacanciesBasePath = "https://api.hh.ru/vacancies?search_field=name&search_field=description&specialization=1&period=1"
)

type VacanciesResult struct {
	Found int `json:"found"`
	PerPage int `json:"per_page"`
	Page int  `json:"page"`
	Pages int `json:"pages"`
	Items []*Vacancy `json:items`
}

func GetVacanciesResult(path string) (x *VacanciesResult, err error) {

	result := new(VacanciesResult)

	rawResult, err := http.Get(path)
	if err != nil {
		fmt.Println("Error getting result for path: ", path)
		return nil, err
	}

	err = json.NewDecoder(rawResult.Body).Decode(result)
	if err != nil {
		fmt.Println("Error parsing body...")
		return nil, err
	}
	return result, nil
}

func GetAllItemsFor(path string) []*Vacancy {
	path = vacanciesBasePath+"&text=" + path
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

