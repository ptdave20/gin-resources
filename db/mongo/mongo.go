package mongo

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"log"
)

func MongoDBHandler(dbURL string, dbName string) gin.HandlerFunc {
	session, err := mgo.Dial(dbURL)
	if err != nil {
		log.Fatal(err)
	}
	return func(c *gin.Context) {
		clone := session.Clone()
		defer clone.Close()
		c.Set("db", clone.DB(dbName))
		c.Next()

	}
}

func GetMongoDB(c *gin.Context) *mgo.Database {
	return c.MustGet("db").(*mgo.Database)
}