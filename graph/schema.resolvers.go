package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4"

	"github.com/ip75/media.server.go/graph/generated"
	"github.com/ip75/media.server.go/graph/model"
)

func (r *mutationResolver) SignIn(ctx context.Context, email string, password string) (*model.SignInResponse, error) {

	if email == "" || password == "" {
		return nil, errors.New("empty email or password")
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	defer conn.Close(context.Background())

	var userID int64
	var userEmail, userPassword string

	err := conn.QueryRow(context.Background(), "SELECT id, email, password FROM public.user WHERE email = $email", email).Scan(&userID, &userEmail, &userPassword)
	if err != nil {
		return nil, err
	}

	accessDuration, _ := time.ParseDuration(os.Getenv("JWT_ACCESS_TOKEN_EXPIRATION_HOURS"))
	const accessToken = createToken(userID, accessDuration, true)
	refreshDuration, _ := time.ParseDuration(os.Getenv("JWT_REFRESH_TOKEN_EXPIRATION_HOURS"))
	const refreshToken = createToken(userID, refreshDuration, false)

	if userPassword == generateHash(password) {

		var count int
		err := conn.QueryRow(context.Background(), "SELECT count(*) as count FROM user_token WHERE user_id = $(id) ORDER BY created ASC", userID).Scan(&count)
		if err != nil {
			return nil, err
		}

		_, err = conn.Exec(context.Background(),
			`DELETE FROM user_token
				WHERE ctid IN (
					SELECT ctid
					FROM user_token
					ORDER BY timestamp
					LIMIT $1 - $2
				);`, count, os.Getenv("COUNT_OF_SESSIONS"))

		if err != nil {
			return nil, err
		}

		_, err = conn.Exec(context.Background(),
			`INSERT INTO user_token (user_id, user_agent, token, created, expires)
				VALUES($1, $2, $3, $4, $5));`, userID, "userAgent", refreshToken.token, refreshToken.iat, refreshToken.exp)

		if err != nil {
			return nil, err
		}
	}

	return &model.SignInResponse{
		accessToken.token,
		refreshToken.token,
		refreshToken.exp_ms,
		accessToken.exp_ms,
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, user *model.UserInput) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditUser(ctx context.Context, userID *int, user *model.ModifyUserInput) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddMediaToHisory(ctx context.Context, mediaID *int) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddMediaToFavorite(ctx context.Context, mediaID *int) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMediaFromFavorite(ctx context.Context, mediaID *int) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCollectionFromFavorite(ctx context.Context, collectionID *int) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddCollectionToHistory(ctx context.Context, collectionID *int) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddCollectionToFavorite(ctx context.Context, collectionID *int) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMedia(ctx context.Context, input *model.MediaInput) (*model.Media, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateMedia(ctx context.Context, id int, input *model.MediaInput) (*model.Media, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetMedia(ctx context.Context, id *int, withData bool) (*model.Media, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetMediaFiltered(ctx context.Context, filter *model.MediaFilterInput, startPage *int, pageSize *int) ([]*model.Media, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCollections(ctx context.Context, input *model.GetCollectionsInput) ([]*model.Collection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUsers(ctx context.Context, startPage *int, pageSize *int) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MediaAdded(ctx context.Context) (<-chan *int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) UserUpdated(ctx context.Context) (<-chan *int, error) {
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
