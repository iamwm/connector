package mongo

import (
	"context"
	"github.com/iamwm/connector/exception"
	"github.com/iamwm/connector/store"
	"github.com/iamwm/connector/util/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func InitMongoHandler() {
	var globalStore = store.STORE
	log.Println("START INIT MONGO HANDLER")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoConfig, err := conf.InitMongoConfig("")
	if err != nil {
		log.Println(err)
		log.Println(exception.ErrMongoConfig)
		return
	}

	cs := "mongodb://" + mongoConfig.Host + ":" + mongoConfig.Port
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cs))

	if err != nil {
		log.Println("Mongo connect failed")
		return
	}
	//defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println("mongo connect failed")
		return
	}

	globalStore.Mongo = client
	log.Println("MONGO HANDLER INIT COMPLETE")
}
