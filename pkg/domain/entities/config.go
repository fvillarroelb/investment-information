package entities

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	BolsaComercio struct {
		Chile struct {
			URLDividendos        string `json:"urlDividendos"`
			URLGetMercadoResumen string `json:"urlGetMercadoResumen"`
		} `json:"chile"`
	} `json:"bolsaComercio"`
	GoogleSheets struct {
		SpreadsheetDividends            string `json:"spreadsheetDividends"`
		SpreadsheetFinancialInformation string `json:"spreadsheetFinancialInformation"`
	} `json:"googleSheets"`
	DataBase struct {
		CollectionDividends            string `json:"collectionDividends"`
		CollectionFinancialInformation string `json:"collectionFinancialInformation"`
	} `json:"dataBase"`
}

// LoadConfig carga la configuración desde un archivo JSON
func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error abriendo el archivo: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("error decodificando el archivo JSON: %v", err)
	}

	return &config, nil
}
