package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ip75/media.server.go/graph/generated"
	"github.com/ip75/media.server.go/graph/model"
)

func (r *mutationResolver) CreateMedia(ctx context.Context, input *model.MediaInput) (*model.Media, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMedia(ctx context.Context, id int, input *model.MediaInput) (*model.Media, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetMedia(ctx context.Context, id *int, withData bool) (*model.Media, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MediaAdded(ctx context.Context) (<-chan *model.Media, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
