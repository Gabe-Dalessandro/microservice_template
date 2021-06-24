// Code generated by sqlc. DO NOT EDIT.

package repository

import (
	"context"
)

type Querier interface {
	CreateNewWorkout(ctx context.Context, arg CreateNewWorkoutParams) (Workout, error)
	DeleteWorkout(ctx context.Context, id int32) (Workout, error)
	GetWorkoutByID(ctx context.Context, id int32) (Workout, error)
}

var _ Querier = (*Queries)(nil)
