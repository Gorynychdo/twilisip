package model

type Phone struct {
    Number string `json:"number" binding:"required,numeric"`
}
