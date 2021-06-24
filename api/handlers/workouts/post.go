package workouts

import (
	"fmt"
	"net/http"

	"github.com/fanfit/feed/api/views"
	"github.com/fanfit/feed/models/workouts/repository"
	"github.com/fanfit/feed/models/workouts/service"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new Workout
// @Description Add a new workout to workout table
// @Tags Workouts
// @Accept  json
// @Produce  json
// @Param input body repository.Workout true "Details of the new Workout"
// @Success 202 {object} repository.Workout	"ok"
// @Failure 500 {object} views.ErrView	"ok"
// @Security ApiKeyAuth
// @Router /v1/workouts [post]
func post(service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input repository.Workout
		if err := c.ShouldBind(&input); err != nil {
			views.Wrap(err, c)
			return
		}
		fmt.Println("About to create")
		newWorkout, err := service.CreateWorkout(c.Request.Context(), input)
		if err != nil {
			views.Wrap(err, c)
			return
		}
		c.JSON(http.StatusAccepted, newWorkout)
	}
}
