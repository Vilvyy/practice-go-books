package models

type Book struct {
	UUID      string  `json:"uuid"`
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	Publisher string  `json:"publisher"`
	Price     float64 `json:"price"`
}

type CreateBookRequest struct {
	Title     string  `json:"title" binding:"required"`
	Author    string  `json:"author" binding:"required"`
	Publisher string  `json:"publisher" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
}

type CreateBookResponse struct {
	UUID   string `json:"uuid"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
