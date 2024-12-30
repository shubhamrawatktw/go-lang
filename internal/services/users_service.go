package services

import (
	"context"

	"github.com/shubhamrwtktw/lic/internal/config/db"
	"github.com/shubhamrwtktw/lic/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser (user models.User)(*mongo.InsertOneResult, error){

	userCollection := db.DB.Collection.Users

	result,err :=userCollection.InsertOne(context.TODO(),user)

	 if err != nil {
		return nil, err
	 }

	 return result,nil

}