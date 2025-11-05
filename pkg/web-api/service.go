package web_api

import (
	"context"
	"calendar/internal/token"
	repo "calendar/pkg/account/repository"
	"calendar/pkg/calendar/repository"
)

type Service interface {
	AddEvent(ctx context.Context, event repository.Event) (string, error)
	ListEvent(ctx context.Context, userId uint64) ([]repository.Event, error)
	DeleteEvent(ctx context.Context, eventId string, userId uint64) error

	SignUp(ctx context.Context, user repo.User) (uint64, token.Token, error)
	Login(ctx context.Context, user repo.User) (uint64, token.Token, error)
	Logout(ctx context.Context, token token.Token) error
}
