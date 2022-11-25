package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"cloud.google.com/go/firestore"
	"os"
	"context"
	"log"
	"desafio/models"
	"github.com/google/uuid"
)

func conectDataBase() (*firestore.CollectionRef, error){
	_ = os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:9091")

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "desafio-c0479")
	if err != nil {
		log.Println(err)
	}
	usersCollection := client.Collection("Users")
	return usersCollection, err
}

func   CreateUser(c  *gin.Context){
	var u models.User
	usersCollection, err := conectDataBase()
	
	if err := c.ShouldBindJSON(&u); err != nil { //converte os dados recebidos em JSON para bites para alocar dentro da struct User(u)
		c.JSON(http.StatusBadRequest, gin.H{
			"erro em converter byte para json": err.Error()})
		return
	}

	u.ID = uuid.NewString()

	_, err = usersCollection.Doc(u.ID).Create(c, u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro ao criar na collection": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, u)//converte bytes para json
}

func FindUser(c *gin.Context){
	var u models.User
	usersCollection, err := conectDataBase()

	givenID:= c.Params.ByName("id")
	doc, err := usersCollection.Doc(givenID).Get(c.Request.Context())//gera um documento, este doc é somente para buscar id no documento
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
		"erro ao encontrar id na collection": err.Error()})
		return
	}

	if err := doc.DataTo(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
		"erro ao passar dado para struct User": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Sucesso": u})
}


//func  para testar conexão
func Teste (c  *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message" : "testado",
	})
}