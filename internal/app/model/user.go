package model

import (
    "github.com/google/uuid"
)

type User struct {
    ID        string
    Token     string
    Confirmed bool
}

func NewUser(id string) *User {
    return &User{
        ID:    id,
        Token: uuid.New().String(),
    }
}
