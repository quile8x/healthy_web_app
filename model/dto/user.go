package dto

import (
	"encoding/json"
)

type LoginDto struct {
	user_name string `json: "user_name"`
	password  string `json: "password"`
}

func NewLoginDto() *LoginDto {
	return &LoginDto{}
}

func (l *LoginDto) ToString() (string, error) {
	bytes, err := json.Marshal(l)
	return string(bytes), err
}
