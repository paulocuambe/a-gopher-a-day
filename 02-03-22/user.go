package main

import "time"

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserDto struct {
	Name     string `json:"name"`
	Username string `json:"username"`
}

func NewUser(dto *CreateUserDto) *User {
	return &User{
		Id:        int(time.Now().Unix()),
		Name:      dto.Name,
		Username:  dto.Username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
