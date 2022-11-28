package middle

import "desafio/controllers"

func Autenticar() gin.HandlerFunc {
	return func (c *gin.Context) {
		controllers.Login(c)
		c.Next()
	}
}