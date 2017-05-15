package gin_resources

import (
	"gopkg.in/gin-gonic/gin.v1"
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
