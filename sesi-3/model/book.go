package model

type Book struct {
	ID     string `json:"id" example:"b1"`
	Title  string `json:"title" example:"A Place Called Home"`
	Author string `json:"author" example:"Preeti Shenoy"`
	Desc   string `json:"desc" example:"Book 1"`
}
