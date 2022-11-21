package router

func HandleRequest(){
	r := gin.Default()

	r.POST("/login", controllers.Login) // com e-mail e senha
	r.POST("/hamburgueria", controllers.Create) //criar loja com: nome, nota, sabores, link para pedido e preço
	r.GET("/hamburgueria/:name", controllers.ReadByName) 
	r.GET("/hamburgueria/:score", controllers.ReadByScore)
	r.GET("/hamburgueria/:price", controllers.ReadByPrice) //irá procurar pelo preço até o limite definido
	r.PUT("/hamburgueria/:name", controllers.Update)//Atualiza qualquer item
	r.DELETE("/hamburgueria/:name", controllers.Delete) //Deleta loja

	r.Run(:5000)
	//se der tempo, fazer função autenticar
}