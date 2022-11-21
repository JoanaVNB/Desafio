package controllers

/* r.POST("/login", controllers.Login) // com e-mail e senha
r.POST("/hamburgueria", controllers.Create) //criar loja com: nome, nota, sabores, link para pedido e preço
r.GET("/hamburgueria/:name", controllers.ReadByName) 
r.GET("/hamburgueria/:score", controllers.ReadByScore)
r.GET("/hamburgueria/:price", controllers.ReadByPrice)
r.PUT("/hamburgueria/:name", controllers.Update)//Atualiza qualquer item
r.DELETE("/hamburgueria/:name", controllers.Delete) */
//(s models.Shop)

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func  Create(c  *gin.Context){
	c.JSON (http.StatusOK , gin.H {
		"message" : "criar loja" ,
	})
}