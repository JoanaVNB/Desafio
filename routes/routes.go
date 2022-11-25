package routes

import(
	"github.com/gin-gonic/gin"
	"desafio/controllers"
)

func HandleRequest(){
	r := gin.Default()

	r.GET("/", controllers.Teste)
	r.POST("/user", controllers.CreateUser)
	r.GET("/:id", controllers.FindUser)
	r.POST("/pizzarias", controllers.Create)
	r.GET("/pizzarias", controllers.ListAll)
	r.GET("/pizzarias/:id", controllers.ReadByID) 
	r.GET("/pizzarias/name/:name", controllers.ReadByName) 
	r.GET("/pizzarias/score/:score", controllers.ReadByScore)
	r.GET("/pizzarias/price/:price", controllers.ReadByPrice)
	r.PUT("/pizzarias/:id", controllers.Update)
	r.PUT("/pizzarias/zerar/:id", controllers.UpdateNoteToZero)
	//r.PATCH("/pizzarias/name/:id", controllers.UpdateName)//com erro
	//r.PUT("/pizzarias/price/:id", controllers.UpdatePrice)//com erro
	r.DELETE("/pizzarias/:id", controllers.Delete)

	r.Run(":5500")
}	

//TO DO
	//funções Uptades que atualizam campos especificos
	//validar campos
	//não repetir nome da loja e email do usuario
	//função middleware autenticar
	//fazer login com validação de e-mail e senha
	//fazer ranking de pizzarias

