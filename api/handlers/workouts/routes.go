package workouts

import (
	"github.com/fanfit/feed/models/workouts/service"
	"github.com/gin-gonic/gin"
)

// Routes sets up resource specific routes on the engine instance
func Routes(r *gin.RouterGroup, service service.Service) {
	router := r.Group("/workouts")
	router.GET("/:id", getWorkoutByID(service))
	router.POST("/", post(service))
}
