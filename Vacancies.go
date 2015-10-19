package main
import (
	"encoding/json"
	"net/http"
	"fmt"
	"io/ioutil"
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
		fmt.Println("Error parsing body:")
		fmt.Println(ioutil.ReadAll(rawResult.Body))
		return nil, err
	}
	return result, nil
}

