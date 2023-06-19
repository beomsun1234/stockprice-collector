package domain

import (
	"strconv"
	"time"
)

type Token struct {
	AccessToken string
	IssuedAt    string
	ExpiresIn   string
}

func NewToken() *Token {
	return &Token{}
}

func (t *Token) BuildAccessToken(accessToken string) *Token {
	t.AccessToken = accessToken
	return t
}
func (t *Token) BuildIssuedAt(issuedAt string) *Token {
	t.IssuedAt = issuedAt
	return t
}
func (t *Token) BuildExpiresIn(expiresIn string) *Token {
	t.ExpiresIn = expiresIn
	return t
}

func (t *Token) IsTokenExpired() bool {
	//todo string date -> date
	if t.AccessToken == "" || t.ExpiresIn == "" || t.IssuedAt == "" {
		return true
	}
	tokenIssedAt, _ := time.Parse("2006-01-02 15:04:05", t.IssuedAt)
	expiresIn, _ := strconv.ParseInt(t.ExpiresIn, 10, 64)
	expiredAt := tokenIssedAt.Add(time.Second * time.Duration(expiresIn))
	now := time.Now()
	if expiredAt.Before(now) || expiredAt.Equal(now) {
		return true
	}
	return false
}
