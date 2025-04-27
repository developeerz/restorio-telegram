package telegram

import "context"

type Cache interface {
	PutUser(ctx context.Context, telegram string, userJSON []byte) error
	PutVerificationCode(ctx context.Context, telegram string, code int) error
	GetUser(ctx context.Context, telegram string) ([]byte, error)
	GetVerificationCode(ctx context.Context, telegram string) (int, error)
}
