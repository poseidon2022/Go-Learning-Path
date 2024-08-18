package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"blog_filtration_feature/repository"
	"blog_filtration_feature/controllers"
	"blog_filtration_feature/usecase"
)

func Setup(router *gin.Engine, dataBase *mongo.Database) {
	filtrationRepo := repository.NewFiltrationRepo(dataBase, "blogs")
	filtrationController := &controllers.FiltrationController {
		FiltrationUseCase : usecase.NewFiltrationUseCase(filtrationRepo),
	}

	router.GET("/filterblog", filtrationController.FilterBlog())
}