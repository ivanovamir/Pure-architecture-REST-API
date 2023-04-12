package dto

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Books     []Book `json:"books"`
}

type RegisterUser struct {
	Name string `json:"name"`
}

type SuccessRegister struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
