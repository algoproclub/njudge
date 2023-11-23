package njudge

import (
	"context"
	"errors"
	"time"
)

type Tag struct {
	ID   int
	Name string
}

func NewTag(name string) *Tag {
	return &Tag{Name: name}
}

var (
	ErrorTagNotFound = errors.New("njudge: user not found")
)

type Tags interface {
	Get(ctx context.Context, ID int) (*Tag, error)
	GetAll(ctx context.Context) ([]Tag, error)
	Insert(ctx context.Context, p Tag) (*Tag, error)
	Delete(ctx context.Context, ID int) error
	Update(ctx context.Context, p Tag) error
}

type ProblemTag struct {
	ID        int
	ProblemID int
	Tag       Tag
	UserID    int
	Added     time.Time
}
