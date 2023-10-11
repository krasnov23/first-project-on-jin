package main

import (
	"github.com/gin-gonic/gin"
	"my-first-project/controller"
	"my-first-project/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	// По умолчанию возвращается экземпляр движка с уже подключенными регистратором и промежуточным программным
	// обеспечением для восстановления, включая логгер и восстановление после любой паники и возврат 500 ошибки если такая была
	server := gin.Default()

	server.GET("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(context *gin.Context) {
		context.JSON(200, videoController.Save(context))
	})

	server.Run(":8080")
}
