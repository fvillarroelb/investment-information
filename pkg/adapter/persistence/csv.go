package adapter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	entity "fv.io/investment-information/pkg/domain/entities"
)

// CSVCreator es una estructura que encapsula la lógica para crear archivos CSV.
type CSVCreator struct {
	FileName string
}

// NewCSVCreator crea una nueva instancia de CSVCreator.
func NewCSVCreator(fileName string) *CSVCreator {
	return &CSVCreator{FileName: fileName}
}

// CreateCSV crea un archivo CSV con datos proporcionados.
func (c *CSVCreator) CreateCSVDividends(moves string) error {

	// Crear el archivo CSV
	file, err := os.Create(c.FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Crear un escritor CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escribir el encabezado
	header := []string{"Nemo", "Fecha Limite", "Fecha Pago", "Moneda", "Valor Dividendo", "Descripcion"}
	if err := writer.Write(header); err != nil {
		return err
	}

	var dividendStruct entity.DividendsResponse

	err = json.Unmarshal([]byte(moves), &dividendStruct)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON Markets: %v", err)
	}

	for _, dividend := range dividendStruct.ListaResult {
		record := []string{
			dividend.Nemo,
			dividend.FecLim,
			dividend.FecPago,
			dividend.Moneda,
			strconv.FormatFloat(dividend.ValAcc, 'f', -1, 64),
			dividend.DescripVC,
		}

		if err := writer.Write(record); err != nil {
			return err
		}
	}

	fmt.Println("Archivo CSV creado exitosamente:", c.FileName)

	return nil
}

func interfaceSliceToStringSlice(input []interface{}) [][]string {
	var result [][]string

	for _, item := range input {
		switch v := item.(type) {
		case string:
			result = append(result, []string{v})
		case int:
			result = append(result, []string{strconv.Itoa(v)})
		case float64:
			result = append(result, []string{strconv.FormatFloat(v, 'f', -1, 64)})
		default:
			// Handle other types as needed
			result = append(result, []string{fmt.Sprintf("%v", v)})
		}
	}

	return result
}
