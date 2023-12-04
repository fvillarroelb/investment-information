package main

import (
	"fmt"
	"log"
	"os"

	adapter "fv.io/investment-information/pkg/adapter/persistence"
	"fv.io/investment-information/pkg/domain/entities"
	entity "fv.io/investment-information/pkg/domain/entities"
	utils "fv.io/investment-information/pkg/domain/utils"
	repository "fv.io/investment-information/pkg/service"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var repo *adapter.Repository

func main() {
	// cargar config.json
	configFilename := "config.json"

	// Carga la configuración desde el archivo
	appConfig, err := entity.LoadConfig(configFilename)
	if err != nil {
		fmt.Printf("Error cargando la configuración: %v\n", err)
		return
	}

	//Leer configuraciones
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error Obtener variables de entorno	:", err)
	}

	flagDb, _ := utils.StringToBool(os.Getenv("DATABASE_FLAG"))
	var db *adapter.Database
	var collectionDividends, collectionFinancialInformation *mongo.Collection
	if flagDb {
		//0 QUERY INVESTMENT
		// Crear una instancia de Database
		db, err = adapter.NewDatabase()
		if err != nil {
			fmt.Println("Error al crear la conexión a MongoDB:", err)
			return
		}
		defer db.Close()

		collectionDividends, err = db.GetCollection(appConfig.DataBase.CollectionDividends)
		if err != nil {
			fmt.Println("Error GetCollection Dividendos:", err)
			return
		}
		collectionFinancialInformation, err = db.GetCollection(appConfig.DataBase.CollectionFinancialInformation)
		if err != nil {
			fmt.Println("Error GetCollection Resumen:", err)
			return
		}

	}

	//CALL HTTPS
	utmResumeString, err := repository.GetUTMercados(appConfig.BolsaComercio.Chile.URLGetMercadoResumen)
	if err != nil {
		//FATAL error
	}

	reqDividends := entity.RequestMarketAll{
		InitDate: os.Getenv("BOLSA_COMERCIO_FECHA_INI"),
		LastDate: os.Getenv("BOLSA_COMERCIO_FECHA_FIN"),
		Nemo:     ""}
	dividendsStrings, err := repository.GetDividends(appConfig.BolsaComercio.Chile.URLDividendos, reqDividends)
	if err != nil {
		//FATAL error
	}

	saveInformation(appConfig, dividendsStrings, utmResumeString, flagDb, db, collectionDividends, collectionFinancialInformation)
	//fmt.Println(dividendsStrings, utmResumeString)
}

func saveInformation(appConfig *entity.Config, dividends string, resumen string, flagDb bool, db *adapter.Database, collectionDividends *mongo.Collection, collectionFinancialInformation *mongo.Collection) {

	//database

	fmt.Println("Parse dividend")
	dividenInterface, err := entities.ParseDividends(dividends)
	if err != nil {
		fmt.Println("Error unmarshaling JSON Dividend:", err)
		return
	}
	fmt.Println("Parse market")
	resumenInterface, err := entities.ParseMarkets(resumen)
	if err != nil {
		fmt.Println("Error unmarshaling JSON Resumen:", err)
		return
	}

	if flagDb {
		// Crear una instancia de Repository
		repo = adapter.NewRepository(db, collectionDividends, collectionFinancialInformation)

		repo.DropCollectionDividends()
		repo.DropCollectionFinancialInformation()
		//	repo.CreateUniqueIndexDividends()

		fmt.Println("Seccion Insert dividends")
		repo.InsertManyDividends(dividenInterface)
		fmt.Println("Seccion Insert resumen")
		repo.InsertManyFinancialInformation(resumenInterface)

	}

	flagGoogleSheet, _ := utils.StringToBool(os.Getenv("GOOGLESHEETS_FLAG"))
	if flagGoogleSheet {

		// Crear un nuevo cliente de Google Sheets
		googleClient, err := adapter.NewGoogleDriveClient(os.Getenv("GOOGLESHEETS_CREDENTIALS_PATH"))
		if err != nil {
			log.Fatalf("Error al crear el cliente de Google Sheets: %v", err)
		}

		// Añadir la nueva pestaña a la hoja de cálculo
		err = googleClient.AddSheet(os.Getenv("GOOGLESHEETS_SPREADSHEETS_ID"), appConfig.GoogleSheets.SpreadsheetDividends)
		if err != nil {
			fmt.Println(err)
		}

		// Añadir datos a la pestaña existente
		err = googleClient.AddData(os.Getenv("GOOGLESHEETS_SPREADSHEETS_ID"), appConfig.GoogleSheets.SpreadsheetDividends, dividends)
		if err != nil {
			log.Fatalf("Error al añadir datos a la pestaña: %v", err)
		}
		/*
			// Añadir la nueva pestaña a la hoja de cálculo
			err = googleClient.AddSheet(os.Getenv("GOOGLESHEETS_SPREADSHEETS_ID"), appConfig.GoogleSheets.SpreadsheetFinancialInformation)
			if err != nil {
				fmt.Println(err)
			}

			// Añadir datos a la pestaña existente
			err = googleClient.AddData(os.Getenv("GOOGLESHEETS_SPREADSHEETS_ID"), appConfig.GoogleSheets.SpreadsheetFinancialInformation, resumeInterfaceArray)
			if err != nil {
				log.Fatalf("Error al añadir datos a la pestaña: %v", err)
			}*/

		fmt.Println("Datos añadidos exitosamente.")
	}

	flagCsv, _ := utils.StringToBool(os.Getenv("CSV_FLAG"))
	if flagCsv {
		//5 GENERAR CSV

		csvService := adapter.NewCSVCreator("./output/" + os.Getenv("FILE_CSV_NAME") + "-dividendos.csv")
		csvService.CreateCSVDividends(dividends)
		//csvServiceResume := adapter.NewCSVCreator("./output/" + os.Getenv("FILE_CSV_NAME") + "-informacion-financiera.csv")
		//csvServiceResume.CreateCSV(resumen)

	}
}
