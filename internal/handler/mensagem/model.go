package mensagem

type Zap struct {
	Texto  string `json:"mensagem" form:"mensagem"`
	Numero string `json:"telefone_whatsapp" form:"telefone_whatsapp"`
}
