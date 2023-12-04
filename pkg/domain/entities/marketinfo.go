package entities

import (
	"encoding/json"
	"fmt"
)

// MarketInfo representa la información sobre un mercado según tu JSON
type MarketInfo struct {
	Mercado   string `json:"MERCADO"`
	Monto     int    `json:"MONTO"`
	NNegocios int    `json:"n_NEGOCIOS"`
}

// MarketsResponse es la estructura que representa la respuesta completa
type MarketsResponse struct {
	ListaResult []MarketInfo `json:"listaResult"`
}

// ParseMarkets intenta deserializar el JSON en una estructura MarketsResponse y retorna un []interface{}.
func ParseMarkets(jsonData string) ([]interface{}, error) {
	var response MarketsResponse

	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON Markets: %v", err)
	}

	// Convertir la estructura en un slice de interfaces
	var resultInterfaces []interface{}
	for _, marketInfo := range response.ListaResult {
		resultInterfaces = append(resultInterfaces, marketInfo)
	}

	return resultInterfaces, nil
}
