package models

type UserTwitterDetails struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type UserTwitterBase struct {
	Data UserTwitterDetails `json:"data"`
}
