package workouts

import (
	"net/http"
	"strconv"

	"github.com/fanfit/feed/api/views"
	"github.com/fanfit/feed/models/workouts/service"

	"github.com/gin-gonic/gin"
)

// @Summary Get Workout by ID
// @Tags Workouts
// @Description Gets a workout by its ID
// @Accept  json
// @Produce  json
// @Param id path string true "User Email ID"
// @Success 200 {object} repository.Workout	"ok"
// @Success 404 {object} views.ErrView
// @Success 500 {object} views.ErrView
// @Security ApiKeyAuth
// @Router /v1/workouts/{id} [get]
func getWorkoutByID(service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		workoutIDStr := c.Param("id")
		workoutID, _ := strconv.Atoi(workoutIDStr)
		resource, err := service.GetWorkoutByID(c.Request.Context(), int32(workoutID))
		if err != nil {
			views.Wrap(err, c)
			return
		}
		c.JSON(http.StatusOK, resource)
	}
}
