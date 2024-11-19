package operatorRoutes

import (
	controller "github.com/Bigthugboy/TourWithUs/internals/infrastructure/adapter/input/controller/operatorController"
	"github.com/Bigthugboy/TourWithUs/internals/infrastructure/middleware"
	"github.com/gin-gonic/gin"
)

func OperatorRoute(r *gin.Engine, service *controller.OperatorController) {
	r.POST("/register/operator", service.RegisterOperator)
	r.GET("/login/operator", service.LoginOperator)

	protectedRouter := r.Group("/api/auth")
	protectedRouter.Use(middleware.Authenticate())
	{
		protectedRouter.POST("/create/tour", service.CreateTour)
	}
}
