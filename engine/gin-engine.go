package engine

import "github.com/gin-gonic/gin"

var engine *gin.Engine

//CORSMiddleware  is
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")

		if c.Request.Method == "OPTIONS" {
			// fmt.Println("options")
			c.JSON(200, nil)
			return
		}

		c.Next()
	}
}

func init() {
	engine = gin.Default()
	engine.Use(CORSMiddleware())
}

//GetAdmin is
func GetAdmin() *gin.RouterGroup {
	return GetEngine().Group("/admin", gin.BasicAuth(gin.Accounts{
		"yoedihar": "2006",
	}))
}

//GetEngine is
func GetEngine() *gin.Engine {
	if engine == nil {
		engine = gin.Default()
	}
	return engine
}
