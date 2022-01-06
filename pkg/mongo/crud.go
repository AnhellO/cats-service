package mongo

import (
	"fmt"

	"github.com/AnhellO/cats-service/pkg/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func CreateCat(b entities.Cat) (interface{}, error) {
	result, err := CatsCollection.InsertOne(Ctx, b)
	if err != nil {
		return "", err
	}

	return result.InsertedID, err
}

func GetCat(id string) (entities.Cat, error) {
	var b entities.Cat
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return b, err
	}

	err = CatsCollection.FindOne(Ctx, bson.M{"_id": objectId}).Decode(&b)
	if err != nil {
		return b, err
	}
	fmt.Println(objectId)

	return b, nil
}

func GetCats() ([]entities.Cat, error) {
	var cats []entities.Cat

	cursor, err := CatsCollection.Find(Ctx, bson.M{})
	if err != nil {
		return cats, err
	}

	defer cursor.Close(Ctx)
	for cursor.Next(Ctx) {
		var cat entities.Cat
		if err = cursor.Decode(&cat); err != nil {
			return cats, err
		}

		cats = append(cats, cat)
	}

	return cats, nil
}

func UpdateCat(id primitive.ObjectID, values map[string]interface{}) error {
	set := bson.M{}
	for key, value := range values {
		set[key] = value
	}

	_, err := CatsCollection.UpdateOne(
		Ctx,
		bson.M{"_id": id},
		bson.M{"$set": set},
	)

	return err
}

func DeleteCat(id primitive.ObjectID) error {
	_, err := CatsCollection.DeleteOne(Ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
