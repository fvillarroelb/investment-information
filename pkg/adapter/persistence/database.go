package adapter

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client        *mongo.Client
	mainDatabase  string
	collectionMap map[string]*mongo.Collection
}

// NewDatabase crea una nueva instancia de Database y establece la conexión a MongoDB.
func NewDatabase() (*Database, error) {

	// Lógica para inicializar la conexión a MongoDB
	clientOptions := options.Client().ApplyURI(os.Getenv("DATABASE_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, fmt.Errorf("Error al conectar a MongoDB: %v", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("Error al hacer ping a MongoDB: %v", err)
	}

	fmt.Println("Conexión a MongoDB establecida")

	return &Database{
		client:        client,
		mainDatabase:  os.Getenv("DATABASE_NAME"),
		collectionMap: make(map[string]*mongo.Collection),
	}, nil
}

// Close cierra la conexión a MongoDB
func (db *Database) Close() {
	if db.client != nil {
		err := db.client.Disconnect(context.Background())
		if err != nil {
			log.Printf("Error al cerrar la conexión a MongoDB: %v", err)
		} else {
			fmt.Println("Conexión a MongoDB cerrada")
		}
	}
}

// GetCollection devuelve una referencia a una colección específica.
func (db *Database) GetCollection(collectionName string) (*mongo.Collection, error) {

	// Verificar si ya tenemos una referencia a la colección
	if collection, ok := db.collectionMap[collectionName]; ok {
		return collection, nil
	}

	collectionReference, err := CollectionExists(db.client, db.mainDatabase, collectionName)

	if collectionReference {
		// Si no existe, obtener una referencia y almacenarla en el mapa
		collection := db.client.Database(db.mainDatabase).Collection(collectionName)
		db.collectionMap[collectionName] = collection
		return collection, nil
	} else {
		// Si no existe, obtener una referencia y almacenarla en el mapa
		err := db.client.Database(db.mainDatabase).CreateCollection(context.Background(), collectionName)
		if err != nil {
			fmt.Println("Error creando collection ", err)
		}
		// Si no existe, obtener una referencia y almacenarla en el mapa
		collection := db.client.Database(db.mainDatabase).Collection(collectionName)
		db.collectionMap[collectionName] = collection
		return collection, nil
	}

	return nil, err
}

// collectionExists verifica si la colección ya existe en la base de datos
func CollectionExists(client *mongo.Client, databaseName, collectionName string) (bool, error) {
	names, err := client.Database(databaseName).ListCollectionNames(context.Background(), bson.D{})
	if err != nil {
		return false, err
	}

	for _, name := range names {
		if name == collectionName {
			return true, nil
		}
	}

	return false, nil
}

// createCollection intenta crear la colección en la base de datos
func CreateCollection(client *mongo.Client, databaseName, collectionName string) error {
	// Puedes personalizar las opciones de creación según tus necesidades
	options := &options.CreateCollectionOptions{}
	return client.Database(databaseName).CreateCollection(context.Background(), collectionName, options)
}
