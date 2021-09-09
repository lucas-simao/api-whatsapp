package main

import (
	"net/http"

	"bitbucket.org/lucassimao01/api-whatsapp/internal/handler/mensagem"
	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "online")
	})

	e.GET("/consulta", mensagem.Consulta)
	e.POST("/zap", mensagem.Mensagem)
	e.POST("/cadastro", mensagem.Cadastro)

	e.Logger.Fatal(e.Start(":9000"))

}
