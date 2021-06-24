package workouts_test

import (
	"net/http"
	"testing"

	handler "github.com/fanfit/feed/api/handlers"
	"github.com/fanfit/feed/api/handlers/workouts"
	"github.com/fanfit/feed/models/workouts/service"
	"github.com/fanfit/feed/testutil"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func TestHandler_InvalidPathCalled_StatusNotFound(t *testing.T) {
	mockService := mockService{}

	router := setupRouter(mockService)
	response, cleanup := testutil.PerformRequest(router, "GET", "/v1/invalid-path/", "")
	defer cleanup()

	testutil.ValidateResponse(t, response, testutil.GenerateError("", "query-parameter-todo", "Path not found", http.StatusNotFound), http.StatusNotFound)
}

func setupRouter(s service.Service) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.NoRoute(handler.NoRoute)
	group := r.Group("/v1")
	workouts.Routes(group, s)
	return r
}
