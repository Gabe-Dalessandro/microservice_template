// @title Fan fit feed service
// @version 0.1.0

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @description ## Feed
// @description Gives us details about workouts
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
//
// @BasePath /
package main

import (
	"fmt"
	"os"
	"time"

	// API Routes
	workoutHandlers "github.com/fanfit/feed/api/handlers/workouts"

	// Tags
	// Workout Tag
	workoutRepository "github.com/fanfit/feed/models/workouts/repository"
	workoutServicePackage "github.com/fanfit/feed/models/workouts/service"

	"github.com/fanfit/feed/api/middleware"
	"github.com/fanfit/feed/server"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{
		TimestampFormat: time.StampMilli,
		FullTimestamp:   true,
	}

	db, err := server.CreatePostGresConnection(log, os.Getenv("DB_URL"))
	if err != nil {
		fmt.Print("Something went wrong!" + err.Error())
	}

	// Instantiate service for each tag
	workoutStore := workoutRepository.NewUserStore(db)
	workoutService := workoutServicePackage.New(workoutStore)

	// Initialize the middleware and routes
	engine := gin.Default()
	router := server.GenerateRouter(engine)

	// Set routes for each tag
	router.Use(middleware.VerifyToken)
	workoutHandlers.Routes(router, workoutService)

	server.Orchestrate(engine, db)
}
