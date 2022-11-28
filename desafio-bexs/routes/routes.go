package routes

import(
	"github.com/gin-gonic/gin"
	"desafio/controllers"
)

func HandleRequest(){
	r := gin.Default()

	r.POST("/user", controllers.CreateUser)
	r.GET("/:id", controllers.FindUser)
	r.POST("/login", controllers.Login)

	r.POST("/pizzarias", controllers.Create)
	r.GET("/pizzarias", controllers.ListAll)
	r.GET("/pizzarias/:id", controllers.ReadByID) 
	r.GET("/pizzarias/name/:name", controllers.ReadByName) 
	r.GET("/pizzarias/score/:score", controllers.ReadByScore)
	r.GET("/pizzarias/price/:price", controllers.ReadByPrice)
	r.PUT("/pizzarias/:id", controllers.Update)
	r.PUT("/pizzarias/:id/score/:score", controllers.UpdateScore)//por URL
	r.PUT("/pizzarias/:id/name", controllers.UpdateName)//por JSON
	r.PUT("/pizzarias/:id/price", controllers.UpdatePrice)//por JSON
	r.DELETE("/pizzarias/:id", controllers.Delete)
	r.GET("/pizzarias/ranking", controllers.Ranking)

	r.Run(":5500")
}	
