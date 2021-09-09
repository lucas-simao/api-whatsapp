# Processo para testar API

1 - Instalar adb no seu PC

2 - Habilitar modo desenvolvedor no celular e habilitar USB debugging

3 - Ter um celular Android com whatsapp instalado

4 - Habilitar no whatsapp a opção de enviar mensagem com a tecla enter em Configurações -> Conversas -> Enviar com Enter

5 - go run api-whatsapp.go

# Rotas da API

# Cadastro - Cadastrar contato
http://localhost:9000/cadastro
curl --location --request POST 'localhost:9000/cadastro' --header 'Content-Type: application/json' --data-raw '{ "telefone_whatsapp": "13996403088"}'

# Consulta - Consultar numeros cadastrados
http://localhost:9000/consulta
curl --location --request GET 'localhost:9000/consulta'

# Zap - Enviar mensagem "Existe uma regra para não permitir enviar a mesma mensagem para a mesma pessoa"
http://localhost:9000/zap
curl --location --request POST 'localhost:9000/zap' --header 'Content-Type: application/json' --data-raw '{"mensagem": "Teste de envio"}'
