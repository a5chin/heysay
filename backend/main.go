package main

import (
	"fmt"
	"heysay/config"
	"heysay/controller"
	"heysay/docs"
	"heysay/entity"
	"heysay/infrastructure/driver"
	"heysay/infrastructure/middleware"
	"heysay/infrastructure/repogitory"
	"heysay/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title		heysay
// @description	heysay
// @version		1.0
func main() {
	// Load config
	conf := config.Load()
	db := driver.NewDB(conf)
	tokenDriver := driver.NewTokenDriver(conf)

	userRepogitory := repogitory.NewUserRepogitory(tokenDriver)
	roleRepogitory := repogitory.NewRoleRepogitory(tokenDriver)

	userUseCase := usecase.NewUserUseCase(userRepogitory)
	roleUseCase := usecase.NewRoleUseCase(roleRepogitory)

	userController := controller.NewUserController(userUseCase)
	roleController := controller.NewRoleController(roleUseCase)

	// Setup webserver
	app := gin.Default()
	app.Use(middleware.Transaction(db))

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "It works")
	})

	api := app.Group("/api/v1")

	userRouter := api.Group("/users")
	userRouter.GET("/", handleResponse(userController.GetUsers))
	userRouter.GET("/:userId", handleResponse(userController.GetUserByID))

	roleRouter := api.Group("/roles")
	roleRouter.GET("/", handleResponse(roleController.GeteRoles))
	roleRouter.POST("/", handleResponse(roleController.CreateRole))

	runApp(app, conf)
}

func runApp(app *gin.Engine, conf *config.Config) {
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", conf.Port)
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	app.Run(fmt.Sprintf("%s:%d", conf.Hostname, conf.Port))

	log.Println(fmt.Sprintf("http://localhost:%d", conf.Port))
	log.Println(fmt.Sprintf("http://localhost:%d/swagger/index.html", conf.Port))
}

func handleResponse(f func(ctx *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := f(c)
		if err != nil {
			e, ok := err.(*entity.Error)
			if ok {
				c.JSON(e.Code, entity.ErrorResponse{Message: err.Error()})
			} else {
				c.JSON(http.StatusInternalServerError, entity.ErrorResponse{Message: err.Error()})
			}
			c.Abort()
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}
