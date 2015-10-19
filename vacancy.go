package main
import (
	"encoding/json"
)

type Vacancy struct {
	Salary json.RawMessage `json:"salary`
	Snippet json.RawMessage `json:"snippet`
	Archived json.RawMessage `json:"archived`
	Name string `json:"name`
	Url string `json:"url`
	CreatedAt json.RawMessage `json:"created_at`
	Employer json.RawMessage `json:"employer`
	PublishedAt json.RawMessage `json:"published_at`
	Id json.RawMessage `json:"id`
}

func ParseVacanciesResult(v *VacanciesResult) []*Vacancy {
	var result []*Vacancy
	for _, item := range v.Items {
		result = append(result, item)
	}
	return result
}