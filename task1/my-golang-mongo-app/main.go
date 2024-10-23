package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var countNumber int = 0
var countLock sync.Mutex
var client *mongo.Client

// Kết nối đến MongoDB
func connectToMongo() {
	var err error
	mongoURI := "mongodb://0.0.0.0:27017"
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI)); err != nil {
		log.Fatal(err)
		return
	} else if client == nil {
		log.Fatal("CLient is nil")
		return
	}
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal("Không thể kết nối đến MongoDB", err)
		return
	}
	fmt.Println("Đã kết nối đến MongoDB!")
}

func main() {
	connectToMongo()
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Endpoint tăng countNumber khi gọi GET /
	r.GET("/", func(c *gin.Context) {
		countLock.Lock()
		defer countLock.Unlock()

		countNumber += 1
		collection := client.Database("counterDB").Collection("counters")
		// can phai add document vao collection
		filter := bson.D{{"name", "requestCounter"}}
		update := bson.D{{"$set", bson.D{{"count", countNumber}}}}
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}

		c.String(http.StatusOK, "Counter đã được tăng = " + string(countNumber))
	})

	// Endpoint trả về giá trị countNumber khi gọi GET /count
	r.GET("/count", func(c *gin.Context) {
		var result bson.M
		collection := client.Database("counterDB").Collection("counters")
		err := collection.FindOne(context.TODO(), bson.D{{"name", "requestCounter"}}).Decode(&result)
		if err != nil {
			log.Fatal(err)
			// c.JSON(http.StatusOK, gin.H{"count": count})
			return
		}
		count := result["count"].(int32)
		c.JSON(http.StatusOK, gin.H{"count": count})
	})

	// Chạy server trên cổng 5000
	r.Run(":5000")
}