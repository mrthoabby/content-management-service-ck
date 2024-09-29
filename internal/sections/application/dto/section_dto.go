package dto

type SectionDTO struct {
	ID    string    `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Pages []PageDTO `json:"pages,omitempty"`
}
