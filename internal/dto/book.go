package dto

type Genre struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type Author struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Year   string `json:"year"`
	Author Author `json:"author"`
	Genre  Genre  `json:"genre"`
}
