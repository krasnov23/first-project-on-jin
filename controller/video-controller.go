package controller

import (
	"github.com/gin-gonic/gin"
	"my-first-project/entity"
	"my-first-project/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video // прописываем context для доступа к данным приходящим к нам с реквеста
}

type controller struct {
	service service.VideoService // Передаем наш сервис в контроллер, создает наш видеоСервис
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video) // преобразует json в сущность
	c.service.Save(video)
	return video
}
