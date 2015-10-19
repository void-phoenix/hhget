package main
import (
	"encoding/json"
	"net/http"
)

type VacanciesResult struct {
	Found int `json:"found"`
	PerPage int `json:"per_page"`
	Page int  `json:"page"`
	Pages int `json:"pages"`
}

func GetVacanciesResult(path string) (x *VacanciesResult, err error) {

	result := new(VacanciesResult)

	rawResult, err := http.Get(path)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(rawResult.Body).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

