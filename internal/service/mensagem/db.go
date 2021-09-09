package mensagem

import (
	"log"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Contatos struct {
	Numero   string `json:"telefone_numero" db:"Numero"`
	Status   string `json:"mensagem" db:"Status"`
	Datahora string `json:"datahora" db:"Datahora"`
}

var db *sqlx.DB

func init() {
	if db == nil {
		db, _ = sqlx.Connect("sqlite3", "lista.db")
	}

}

func PostStatus(numero, status string) (resposta string, err error) {

	dataHoraAgora := dataHora()

	query, err := db.Prepare("UPDATE contatos SET Status = ?, Datahora = ? WHERE Numero = ?")

	if err != nil {
		log.Println("Erro db.Prepare: ", err.Error())
		return
	}

	tx, err := db.Begin()

	if err != nil {
		log.Println("Erro db.Begin: ", err.Error())
		return
	}

	_, err = tx.Stmt(query).Exec(status, dataHoraAgora, numero)

	if err != nil {
		log.Println("Erro tx.Stmt: ", err.Error())
		tx.Rollback()
		return
	} else {
		tx.Commit()
	}

	return
}

func PostCadastro(numero string) (resposta string, err error) {

	dataHoraAgora := dataHora()

	numero = strings.Replace(numero, " ", "", -1)
	numero = strings.Replace(numero, "(", "", -1)
	numero = strings.Replace(numero, ")", "", -1)
	numero = strings.Replace(numero, "-", "", -1)

	query, err := db.Prepare("INSERT INTO contatos(Numero, Status, Datahora) VALUES(?,?,?)")

	if err != nil {
		log.Println("Erro db.Prepare: ", err.Error())
		return
	}

	tx, err := db.Begin()

	if err != nil {
		log.Println("Erro db.Begin: ", err.Error())
		return
	}

	_, err = tx.Stmt(query).Exec(numero, "", dataHoraAgora)

	if err != nil {
		log.Println("Erro tx.Stmt: ", err.Error())
		tx.Rollback()
		return
	} else {
		tx.Commit()
	}

	return
}

func GetContatos(texto string) (contatos []string, err error) {

	query := `SELECT distinct(Numero) FROM contatos WHERE LOWER(Status) != '` + strings.ToLower(texto) + `'`

	err = db.Select(&contatos, query)

	return
}

func GetConsulta() (contatos []Contatos, err error) {

	query := `SELECT distinct(Numero) as Numero, Status, Datahora FROM contatos`

	err = db.Select(&contatos, query)

	return
}

func dataHora() string {

	t := time.Now()

	data := t.Format("2006-01-02 15:04:05")

	return data
}
