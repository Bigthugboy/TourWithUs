package routes

import (
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/input/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, service *controller.TouristController) {
	r.POST("/register", service.RegisterTourist)
	r.POST("/login", service.LoginTourist)

}
