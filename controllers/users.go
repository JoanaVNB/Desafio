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

func  CreateUser(c  *gin.Context){
	_ = os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:9090")

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "desafio")
	if err != nil {
		log.Println(err)
	}

	usersCollection := client.Collection("Users")

	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil { //torna os dados recebidos em JSON para bites para alocar dentro da struct User(u)
		c.String(http.StatusBadRequest, gin.H{
			"erro no bind": err.Error()})
		return
	}

	givenID, _ := c.Params.Get("id")
	u.ID = uuid.NewString()

	doc, err := usersCollection.Doc(givenID).Get(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro ao add na collection": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, u)//converte bytes para json

	if err := doc.DataTo(&u); err != nil {
		c.String(http.StatusInternalServerError, gin.H{
			"erro ao DataTo": err.Error()})///?????
		return
	}

	
	u.ID = doc.Ref.ID
		c.JSON(http.StatusOK, u)
}

func  Teste(c  *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message" : "testado",
	})
}