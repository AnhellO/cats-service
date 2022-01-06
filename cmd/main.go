package main

import (
	"fmt"
	"log"

	"github.com/AnhellO/cats-service/pkg/entities"
	"github.com/AnhellO/cats-service/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// Connect to MongoDB
	err := mongo.Setup()
	if err != nil {
		log.Fatal(err)
	}

	// Insert a record
	res, err := mongo.CreateCat(entities.Cat{
		Name:  "Greta",
		Breed: "Maine Coon",
		Age:   1,
		Image: "https://example.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	id := res.(primitive.ObjectID).Hex()
	fmt.Println("Inserted element: ", id)

	// Get that record
	object, err := mongo.GetCat(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Got element %+v\n", object)

	// Insert another record
	res, err = mongo.CreateCat(entities.Cat{Name: "Milka",
		Breed: "Cowie",
		Age:   1,
		Image: "https://example.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	id2 := res.(primitive.ObjectID).Hex()
	fmt.Println("Inserted element: ", id2)

	// Get all records
	cats, err := mongo.GetCats()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", cats)

	// Update records
	for _, c := range cats {
		err = mongo.UpdateCat(c.ID, map[string]interface{}{"image": "https://updated.example.com"})
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Updated all records!")

	// Delete all records
	for _, c := range cats {
		err = mongo.DeleteCat(c.ID)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Deleted all records!")
}
