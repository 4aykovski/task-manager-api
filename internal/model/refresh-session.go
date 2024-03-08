package model

import "time"

type RefreshSession struct {
	Token       string
	ExpiresIn   time.Time
	Fingerprint string
	UserId      string
}
