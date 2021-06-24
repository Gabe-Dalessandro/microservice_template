CREATE TABLE workouts (
    id              SERIAL     PRIMARY KEY,
    workout_name    TEXT       NOT NULL,
    workout_type    TEXT        NOT NULL
);

INSERT INTO workouts(workout_name, workout_type)
VALUES('workout 1', 'Strength');

INSERT INTO workouts(workout_name, workout_type)
VALUES('workout 2', 'Core');

INSERT INTO workouts(workout_name, workout_type)
VALUES('workout 3', 'Lower Body');