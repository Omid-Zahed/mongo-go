package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func viperEnvVariable(key string) string {
	viper.SetConfigFile("mongo/.env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

func main() {
	first_connect()
}

func first_connect() {

	host := viperEnvVariable("MONGO_HOST")
	port := viperEnvVariable("MONGO_PORT")
	user := viperEnvVariable("MONGO_USER")
	password := viperEnvVariable("MONGO_PASS")

	client, err := mongo.NewClient(options.Client().
		ApplyURI("mongodb://" + user + ":" + password + "@" + host + ":" + port))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	defer client.Disconnect(ctx)
}
