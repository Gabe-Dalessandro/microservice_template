-- name: GetWorkoutByID :one
SELECT * FROM workouts 
WHERE id = $1;

-- name: CreateNewWorkout :one
INSERT INTO workouts(workout_name, workout_type)
VALUES(
    $1,
    $2
)
RETURNING *;

-- name: DeleteWorkout :one
DELETE FROM workouts
WHERE id = $1
RETURNING *;