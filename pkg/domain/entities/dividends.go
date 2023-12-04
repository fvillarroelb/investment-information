package entities

import (
	"encoding/json"
	"fmt"
)

// Dividend es la estructura que representa un dividendo según tu JSON
type Dividend struct {
	DescripVC string  `json:"descrip_vc"`
	FecLim    string  `json:"fec_lim"`
	FecPago   string  `json:"fec_pago"`
	Moneda    string  `json:"moneda"`
	Nemo      string  `json:"nemo"`
	NumAccAnt int     `json:"num_acc_ant"`
	NumAccDer int     `json:"num_acc_der"`
	NumAccNue int     `json:"num_acc_nue"`
	PreAntVC  float64 `json:"pre_ant_vc"`
	PreExVC   float64 `json:"pre_ex_vc"`
	ValAcc    float64 `json:"val_acc"`
}

// DividendsResponse es la estructura que representa la respuesta completa
type DividendsResponse struct {
	ListaResult []Dividend `json:"listaResult"`
}

// ParseDividens intenta deserializar el JSON en una estructura MarketsResponse y retorna un []interface{}.
func ParseDividends(jsonData string) ([]interface{}, error) {
	var response DividendsResponse

	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON Markets: %v", err)
	}
	//fmt.Println(response)

	// Convertir la estructura en un slice de interfaces
	var resultInterfaces []interface{}
	for _, marketInfo := range response.ListaResult {
		resultInterfaces = append(resultInterfaces, marketInfo)
	}

	return resultInterfaces, nil
}
