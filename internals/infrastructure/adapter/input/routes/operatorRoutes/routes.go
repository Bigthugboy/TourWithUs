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
		protectedRouter.PUT("/update/tour/:id", service.UpdateTour)
		protectedRouter.DELETE("/delete/tour/:id", service.DeleteTour)
		protectedRouter.GET("getAll/tours", service.GetTourList)
		protectedRouter.GET("view/tour/:id", service.ViewTourDetails)
		protectedRouter.GET("get/particular/tour/ByOperator/:id/:tourId", service.GetTourByAOperator)
		protectedRouter.GET("getAll/tour/byOperator", service.GetToursByAOperator)
		protectedRouter.PUT("changeTourAvailability", service.ChangeTourAvailability)

	}
}
