package csv

import (
	"encoding/csv"
	"os"

	"github.com/seu-usuario/nome-do-repo/pkg/model"
)

func ParseCSV(filename string) ([]model.Data, error) {
	// abre o arquivo CSV/TXT
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// lê o conteúdo do arquivo usando o pacote CSV da biblioteca padrão
	reader := csv.NewReader(file)
	reader.Comma = ';' // define o separador de colunas

	// lê todas as linhas do arquivo e converte em uma slice de struct
	var data []model.Data
	for {
		row, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" { // atinge o final do arquivo
				break
			}
			return nil, err
		}
		datum := model.Data{
			CPF:           row[0],
			Private:       row[1] == "1",
			Incompleto:    row[2] == "1",
			UltimaCompra:  row[3],
			TicketMedio:   parseFloat(row[4]),
			TicketUltComp: parseFloat(row[5]),
			LojaFrequente: row[6],
			LojaUltCompra: row[7],
		}
		data = append(data, datum)
	}

	return data, nil
}

func parseFloat(str string) float64 {
	// converte uma string em um float64
	// se a string estiver vazia, retorna 0.0
	if str == "" {
		return 0.0
	}
	val, _ := strconv.ParseFloat(str, 64)
	return val
}