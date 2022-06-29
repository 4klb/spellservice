package routes

import (
	"net/http"

	"main/internal/services/api"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetUpRoutes(text api.Text, log *zap.SugaredLogger) error {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		responces, err := api.GetResponce(text.Texts, log)
		if err != nil {
			log.Debug(err)
			return
		}
		api.Replace(responces, text.Texts)
		c.JSON(http.StatusOK, gin.H{
			"texts": text.Texts,
		})
	})
	return r.Run()
}
