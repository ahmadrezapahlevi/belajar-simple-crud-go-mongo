package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// ConnectDB menyambungkan ke MongoDB Atlas
func ConnectDB() {
	// Ganti dengan URI milik kamu
	uri := ""

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal("Error connecting to MongoDB Atlas:", err)
	}

	// Ping untuk pastikan koneksi
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Ping to MongoDB Atlas failed:", err)
	}

	DB = client.Database("Saranaku") // sesuai nama database kamu
	log.Println("Connected to MongoDB Atlas!")
}

// GetCollection mengembalikan collection dari database
func GetCollection(collectionName string) *mongo.Collection {
	if DB == nil {
		log.Fatal("Database belum terhubung. Panggil ConnectDB() dulu sebelum GetCollection().")
	}
	return DB.Collection(collectionName)
}
