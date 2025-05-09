package config

type Config struct {
	UserName string `json:"user_name"`
	IntroText string `json:"intro_text"`
	ProjectList []Project `json:"project_list"`
}

type Project struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

