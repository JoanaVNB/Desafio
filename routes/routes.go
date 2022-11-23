package routes

import(
	"github.com/gin-gonic/gin"
	"desafio/controllers"
)

func HandleRequest(){
	r := gin.Default()

	//r.POST("/login", controllers.Login) // com e-mail e senha
	r.POST("/user", controllers.CreateUser)
	r.GET("/:id", controllers.FindUser)
	r.GET("/", controllers.Teste)
	//r.POST("/pizzarias", controllers.Create) //criar loja com: nome, nota, sabores, link para pedido e preço
	//r.GET("/pizzarias/:name", controllers.ReadByName) 
	//r.GET("/pizzarias/:score", controllers.ReadByScore)
	//r.GET("/pizzarias/:price", controllers.ReadByPrice) //irá procurar pelo preço até o limite definido
	//r.PUT("/pizzarias/:name", controllers.Update)//Atualiza qualquer item
	//r.DELETE("/pizzarias/:name", controllers.Delete) //Deleta loja

	r.Run()
	//se der tempo, fazer função autenticar
}