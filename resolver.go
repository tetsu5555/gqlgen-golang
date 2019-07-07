//go:generate go run github.com/99designs/gqlgen

package gqlgen_todos

import (
	context "context"
	"errors"
	"fmt"
	"math/rand"
)

type Resolver struct {
	todos []*Todo
	users []*User
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (*Todo, error) {
	var user *User
	for _, u := range r.users {
		if u.ID == input.UserID {
			user = u
		}
	}
	if user == nil {
		return &Todo{}, errors.New("no user found")
	}

	todo := &Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUser) (*User, error) {
	user := &User{
		ID:       fmt.Sprintf("T%d", rand.Int()),
		Name:     input.Name,
		Password: input.Password,
	}
	r.users = append(r.users, user)
	return user, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	return r.users, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return &User{}, errors.New("user not found")
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *Todo) (*User, error) {
	return &User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}
