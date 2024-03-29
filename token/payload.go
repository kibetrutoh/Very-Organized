package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken        = errors.New("token is invalid")
	ErrExpiredToken        = errors.New("token is expired")
	ErrBlacklistedToken    = errors.New("token is blacklisted")
	ErrInternalServerError = errors.New("something went wrong")
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    int       `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func TokenPayload(userID int, duration time.Time) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		UserID:    userID,
		IssuedAt:  time.Now().UTC(),
		ExpiresAt: duration,
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().UTC().After(payload.ExpiresAt) {
		return ErrExpiredToken
	}
	return nil
}
