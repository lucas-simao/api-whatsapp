package mensagem

import (
	echo "github.com/labstack/echo/v4"

	"bitbucket.org/lucassimao01/api-whatsapp/internal/service/mensagem"
)

func Mensagem(c echo.Context) error {
	p := new(Zap)

	if err := c.Bind(p); err != nil {
		return err
	}

	mensagem, err := mensagem.PostMensagem(p.Texto)

	if err != nil {
		return c.JSON(422, err.Error())
	}

	return c.JSON(200, mensagem)
}

func Cadastro(c echo.Context) error {
	p := new(Zap)

	if err := c.Bind(p); err != nil {
		return err
	}

	mensagem, err := mensagem.PostCadastro(p.Numero)

	if err != nil {
		return c.JSON(422, err.Error())
	}

	return c.JSON(200, mensagem)
}

func Consulta(c echo.Context) error {

	contatos, err := mensagem.GetConsulta()

	if err != nil {
		return c.JSON(422, err.Error())
	}

	return c.JSON(200, contatos)
}
