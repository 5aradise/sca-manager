package models

type Cat struct {
	ID                int32  `json:"id"`
	Name              string `json:"name"`
	YearsOfExperience int32  `json:"years_of_experience"`
	Breed             string `json:"breed"`
	Salary            Money  `json:"salary"`
}
