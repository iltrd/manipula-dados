package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Open() error {
	// obtém as variáveis de ambiente necessárias para se conectar ao banco de dados
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	// cria a string de conexão ao banco de dados
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// abre uma conexão com o banco de dados
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	// verifica se a conexão banco de dados é bem-sucedida
err = db.Ping()
if err != nil {
	return err
}

// cria a tabela no banco de dados, se ainda não existir
_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS data (
		id SERIAL PRIMARY KEY,
		cpf TEXT NOT NULL UNIQUE,
		private BOOLEAN NOT NULL,
		incompleto BOOLEAN NOT NULL,
		ultima_compra TEXT NOT NULL,
		ticket_medio FLOAT NOT NULL,
		ticket_ult_compra FLOAT NOT NULL,
		loja_frequente TEXT NOT NULL,
		loja_ult_compra TEXT NOT NULL
	)
`)
if err != nil {
	return err
}

return nil
}

func Close() {
// fecha a conexão com o banco de dados
db.Close()
}

func InsertData(data model.Data) error {
// insere os dados no banco de dados
_, err := db.Exec( INSERT INTO data (cpf, private, incompleto, ultima_compra, ticket_medio, ticket_ult_compra, loja_frequente, loja_ult_compra) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) , data.CPF, data.Private, data.Incompleto, data.UltimaCompra, data.TicketMedio, data.TicketUltComp, data.LojaFrequente, data.LojaUltCompra)
if err != nil {
return err
}
return nil
}