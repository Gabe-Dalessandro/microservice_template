package repository

import (
	"context"
	"database/sql"
	"fmt"
)

// Repository is used by the service to communicate with the underlying database
type Repository interface {
	// CREATES
	CreateWorkout(context.Context, Workout) (Workout, error)
	// DELETES
	DeleteWorkout(context.Context, int32) error
	// GETS
	GetWorkoutByID(context.Context, int32) (Workout, error)
}

type repository struct {
	queries *Queries
	db      *sql.DB
}

// Gets workout by the id
func (repo *repository) GetWorkoutByID(ctx context.Context, id int32) (Workout, error) {
	response, err := repo.queries.GetWorkoutByID(ctx, id)
	if err != nil {
		fmt.Print(err)
	}

	return response, err
}

// Create a new workout
func (repo *repository) CreateWorkout(ctx context.Context, workout Workout) (Workout, error) {
	newWorkout, err := repo.queries.CreateNewWorkout(ctx, CreateNewWorkoutParams{
		WorkoutName: workout.WorkoutName,
		WorkoutType: workout.WorkoutType,
	})
	if err != nil {
		fmt.Print(err)
	}
	return newWorkout, err
}

// Deletes a workout according to its ID
func (repo *repository) DeleteWorkout(ctx context.Context, id int32) error {
	return nil
}

func NewUserStore(db *sql.DB) Repository {
	return &repository{
		queries: New(db),
		db:      db,
	}
}
