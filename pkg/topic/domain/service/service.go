package service

import (
	"context"
)

type Service interface {
	DuplicatedName(ctx context.Context, name string) (bool, error)
}
