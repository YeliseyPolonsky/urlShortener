package link

import (
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Link string `json:"url"`
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string) *Link {
	return &Link{
		Link: url,
		Hash: RandStringRunes(5),
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxwzABCDEFGHIJKLMNOPQRSTUVWXWZ")

func RandStringRunes(n int) string {
	runes := make([]rune, n)
	for i := range runes {
		runes[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(runes)
}
