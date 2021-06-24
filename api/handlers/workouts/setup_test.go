package workouts_test

import (
	"context"

	"github.com/fanfit/feed/models/workouts/repository"
)

type mockService struct {
	GetWorkoutByIDResponse repository.Workout
	CreateWorkoutResponse  repository.Workout

	GetByIDErr error
	CreateErr  error
	DeleteErr  error
}

func (m mockService) GetWorkoutByID(ctx context.Context, id int32) (repository.Workout, error) {
	return m.GetWorkoutByIDResponse, m.GetByIDErr
}

func (m mockService) CreateWorkout(ctx context.Context, user repository.Workout) (repository.Workout, error) {
	return m.CreateWorkoutResponse, m.CreateErr
}

func (m mockService) DeleteWorkout(ctx context.Context, id int32) error {
	return m.DeleteErr
}
