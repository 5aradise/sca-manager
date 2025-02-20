package models

type Mission struct {
	ID          int32    `json:"id"`
	CatID       int32    `json:"cat_id,omitempty"`
	IsCompleted bool     `json:"is_completed"`
	Targets     []Target `json:"targets"`
}

type Target struct {
	ID          int32  `json:"-"`
	Name        string `json:"name"`
	Country     string `json:"country"`
	Notes       string `json:"notes,omitempty"`
	IsCompleted bool   `json:"is_completed"`
}
