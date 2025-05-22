package model

type Category struct {
	Id   string `json:"id,omitempty" firestore:"id,omitempty"`
	Name string `json:"name,omitempty" firestore:"name,omitempty"`
}
