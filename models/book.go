package models

type Book struct {
	Id     string  `json:"id" gorm:"primaryKey"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Price  float32 `json:"price"`
}

type GetReq struct{}

type FindBookAllResp struct {
	Books []Book `json:"books"`
}
