package service

import (
	"context"
	"fmt"

	"github.com/fanfit/feed/models/workouts/repository"
)

// Service receives commands from handlers and forwards them to the repository
type Service interface {
	// CREATES
	CreateWorkout(context.Context, repository.Workout) (repository.Workout, error)
	// DELETES
	DeleteWorkout(context.Context, int32) error
	// GETS
	GetWorkoutByID(context.Context, int32) (repository.Workout, error)
}

type service struct {
	repository repository.Repository
}

// New creates a service instance with the repository passed
func New(repository repository.Repository) Service {
	return &service{repository: repository}
}

func (service *service) CreateWorkout(ctx context.Context, inputWorkout repository.Workout) (repository.Workout, error) {
	fmt.Print("Going into repo")
	return service.repository.CreateWorkout(ctx, inputWorkout)
}
func (service *service) GetWorkoutByID(ctx context.Context, id int32) (repository.Workout, error) {
	return service.repository.GetWorkoutByID(ctx, id)
}

func (service *service) DeleteWorkout(ctx context.Context, id int32) error {
	return service.repository.DeleteWorkout(ctx, id)
}
