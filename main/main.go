package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Starship struct {
	ID              string `bson:"-"`
	Name            string `bson:"name"`
	Model           string `bson:"model"`
	CostInCredits   int64  `bson:"costInCredits"`
	Identificadores [2]int `bson:identificadores`
}

func main() {
	host := "localhost"
	port := 27017

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port))
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connections
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Congratulations, you're already connected to MongoDB!")

	collection := client.Database("swapi").Collection("starships")

	deathStar := Starship{
		Name:          "Death Star",
		Model:         "DS-1 Orbital Battle Station",
		CostInCredits: 1000000000000,
	}

	insertResult, err := collection.InsertOne(context.TODO(), deathStar)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Death Star had been inserted: ", insertResult.InsertedID)

	executor := Starship{
		Name:            "Executor",
		Model:           "Executor-class star dreadnought",
		CostInCredits:   1143350000,
		Identificadores: [2]int{2, 5},
	}
	millenniumFalcon := Starship{
		Name:          "Millennium Falcon",
		Model:         "YT-1300 light freighter",
		CostInCredits: 100000,
	}

	starships := []interface{}{executor, millenniumFalcon}

	insertManyResult, err := collection.InsertMany(context.TODO(), starships)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple starships: ", insertManyResult.InsertedIDs)

	//Para buscar
	var result Starship

	filter := bson.D{{"name", "Executor"}}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starship found: %v\n", result)
}
