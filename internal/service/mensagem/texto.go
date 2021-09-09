package mensagem

import (
	"log"
	"net/url"
)

func PostMensagem(texto string) (mensagem string, err error) {

	t := &url.URL{Path: texto}
	textoEncodedString := t.String()

	telefones, err := GetContatos(texto)

	if err != nil {
		log.Println("Error select contatos: ", err.Error())
	}

	for _, v := range telefones {
		gatilho := "https://api.whatsapp.com/send?phone=55" + v + "&text=" + textoEncodedString
		PostReceber(gatilho, v, texto)
	}

	return "Mensagem Enviadas", nil
}
