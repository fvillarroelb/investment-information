package adapter

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository es la implementación de QueryRepository
type Repository struct {
	db                             *Database
	collectionDividends            *mongo.Collection
	collectionFinancialInformation *mongo.Collection
}

// NewRepository crea una nueva instancia de Repository
func NewRepository(db *Database, collectionDividends *mongo.Collection, collectionFinancialInformation *mongo.Collection) *Repository {
	//clientInversiones :=
	return &Repository{
		db:                             db,
		collectionDividends:            collectionDividends,
		collectionFinancialInformation: collectionFinancialInformation,
	}
}

// createUniqueIndex crea un índice único en la colección
func (repo *Repository) CreateUniqueIndexDividends() error {
	// Crea un índice único en los campos relevantes (Ejemplo: Descripcion y Monto)
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "nemo", Value: 1},
			{Key: "descrip_vc", Value: 1},
			{Key: "fec_lim", Value: 1},
			{Key: "fec_pago", Value: 1},

			// Agrega otros campos según sea necesario
		},
		Options: options.Index().SetUnique(true),
	}

	_, err := repo.collectionDividends.Indexes().CreateOne(context.Background(), indexModel)
	return err
}

func (repo *Repository) InsertManyDividends(dividendos []interface{}) (*mongo.InsertManyResult, error) {
	//fmt.Println(dividendos)
	// Opciones de inserción con 'ordered' establecido en false
	options := options.InsertMany().SetOrdered(false)
	result, err := repo.collectionDividends.InsertMany(context.Background(), dividendos, options)
	if err != nil {
		fmt.Println("Error insert many:", err)
		return result, err
	}
	// Imprime los IDs de los documentos insertados
	//	fmt.Println("Documentos insertados correctamente. IDs:", result)
	return result, err
}

func (repo *Repository) InsertManyFinancialInformation(informacion []interface{}) (*mongo.InsertManyResult, error) {

	result, err := repo.collectionFinancialInformation.InsertMany(context.Background(), informacion)
	if err != nil {
		fmt.Println("Error insert many:", err)
		return result, err
	}
	// Imprime los IDs de los documentos insertados
	fmt.Println("Documentos Resumen Informacion Financiera insertados  correctamente. IDs:", result)
	return result, err
}

func (repo *Repository) DropCollectionDividends() error {

	err := repo.collectionDividends.Drop(context.Background())
	if err != nil {
		fmt.Println("Error al Drop Collection Dividends:", err)
		return err
	}
	fmt.Println("Drop Collection Dividends:", err)
	return err
}

func (repo *Repository) DropCollectionFinancialInformation() error {

	err := repo.collectionFinancialInformation.Drop(context.Background())
	if err != nil {
		fmt.Println("Error al Drop Collection informacion financiera:", err)
		return err
	}
	fmt.Println("Drop Collection informacion financiera:", err)
	return err
}
