package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
  dbURI := os.Getenv("DB_URI")
  ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
  defer cancel()
  client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
  if err != nil {
    log.Fatalf("Unable to connect to mongodb\n%v\n", err)
  }
  return client
}

func DisconnectDB(c *mongo.Client) {
  ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
  defer cancel()
  if err := c.Disconnect(ctx); err != nil {
    log.Fatalf("Unable to disconnect\n%v\n", err)
  }
}
