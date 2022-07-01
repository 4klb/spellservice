package routes

import (
	"main/internal/services/api"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func SetUpRoutes(text api.Text, log *zap.SugaredLogger, v *viper.Viper) error {
	r := gin.Default()
	port := v.GetString("server.port")
	if port != "" {
		os.Setenv("PORT", port) // обходим порт по умолчанию
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"texts": text.Texts,
		})
	})

	return r.Run()
}
