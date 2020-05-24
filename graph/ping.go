package graph

import (
	"context"

	"github.com/haydenwoodhead/duplico/graph/model"
)

func (r *queryResolver) ping(_ context.Context) (*model.Pong, error) {
	return &model.Pong{Pong: true}, nil
}
