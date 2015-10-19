package main
import (
	"encoding/json"
	"time"
)

type Vacancy struct {
	Salary *Salary `json:"salary`
	Snippet json.RawMessage `json:"snippet`
	Archived bool `json:"archived`
	Name string `json:"name`
	Url string `json:"url`
	CreatedAt time.Time `json:"created_at`
	Employer *Employer `json:"employer`
	PublishedAt time.Time `json:"published_at`
	Id string `json:"id`
}

type Salary struct{
	To float64 `json:"to`
	From float64 `json:"from`
	Currency string `json:"currency`
}

type Employer struct {
	Name string `json:"name"`
	Id string `json:"id"`
}

func ParseVacanciesResult(v *VacanciesResult) []*Vacancy {
	var result []*Vacancy
	for _, item := range v.Items {
		result = append(result, item)
	}
	return result
}