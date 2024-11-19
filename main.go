package main

import (
	"github.com/Bigthugboy/TourWithUs/internals/domain/services/Operator"
	"github.com/Bigthugboy/TourWithUs/internals/domain/services/tour"
	"github.com/Bigthugboy/TourWithUs/internals/domain/services/tourist"
	dto "github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/operatorDto"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/dto/touristDto"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/input/controller"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/input/controller/operatorController"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/input/routes"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/input/routes/operatorRoutes"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/output/keycloakAdapter"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/output/persistence/Db/OperatorDb"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/output/persistence/Db/query"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/output/persistence/Db/tourDb"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/output/persistence/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {

	db, err := utils.NewDatabaseConnection()
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer db.Close()
	logrus.Info("---------> STARTED TOUR WITH US SERVER <--------------")

	db.AutoMigrate(&touristDto.TouristObject{})
	db.AutoMigrate(&dto.OperatorDto{})

	database := query.NewTourDB(db)
	operatorDb := OperatorDb.NewOperatorDb(db)
	tourdb := tourDb.NewTourDB(db)

	adapter := keycloakAdapter.KeycloakAdapter{}
	touristService := tourist.NewTourist(database, &adapter)
	touristController := controller.NewTouristController(touristService)

	tourUseCase := tour.NewTour(tourdb)
	usecase := Operator.NewOperatorService(operatorDb, tourUseCase)
	operatorController := operatorController.NewController(usecase)

	r := gin.Default()
	routes.SetupRoutes(r, touristController)
	operatorRoutes.OperatorRoute(r, operatorController)
	logrus.Info("=================Running TOUR WITH US SERVER================")

	if err := r.Run("localhost:9090"); err != nil {
		log.Fatalf("Could not start server: %v", err)

	}

}
