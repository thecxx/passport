package service

import (
	"context"
)

type UserIdService struct {
}

// New user id service.
func NewUserIdService() *UserIdService {
	return new(UserIdService)
}

// Allocate a new user ID.
func (u *UserIdService) AllocUserId(ctx context.Context) (string, error) {
	return "", nil
}

// Check if the specified user ID exists.
func (u *UserIdService) ProbeUserId(ctx context.Context, uid string) (bool, error) {
	return true, nil
}
