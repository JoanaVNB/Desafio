package controllers

import (
	"context"
	"desafio/models"
	"log"
	"net/http"
	"os"
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"google.golang.org/api/iterator"
	"errors"
)

func conectUserCollection() (*firestore.CollectionRef, error){
	_ = os.Setenv("FIRESTORE_EMULATOR_HOST", "firestore:9091")

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "desafio-c0479")
	if err != nil {
		log.Println(err)
	}
	usersCollection := client.Collection("Users")
	return usersCollection, err
}

func emailRegistered(email string , c *gin.Context) (bool) {
	usersCollection, _ := conectUserCollection()	

	iter := usersCollection.Where("email", "==", email).Documents(c)
	for {
		doc, err := iter.Next()
		if err == iterator.Done{
			break
		}
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"não conseguiu percorrer coleção": err.Error()})
		}
		if doc != nil{
			c.JSON(http.StatusConflict, "E-mail foi cadastrado.")
			return true
		}
	}	
	c.JSON(http.StatusOK, "Conta pode ser criada")
			return false
}	

func CreateUser(c  *gin.Context){
	var u models.User
	var ve validator.ValidationErrors
	usersCollection, err :=conectUserCollection()
	

	if err := c.ShouldBindJSON(&u); err != nil {
		if errors.As(err, &ve){
			out := make([]models.ErrorMsg, len(ve))
			for i, fe := range ve{
				out[i] = models.ErrorMsg{fe.Field(), models.GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": out})
		}
		return 
	}

	if emailRegistered(u.Email, c) == false{
			u.ID = uuid.NewString()
			_, err = usersCollection.Doc(u.ID).Create(c, u)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro ao criar na collection": err.Error()})
			return
			}
		}
		c.JSON(http.StatusCreated, u)
}

func FindUser(c *gin.Context){
	var u models.User
	usersCollection, err := conectUserCollection()

	givenID:= c.Params.ByName("id")
	doc, err := usersCollection.Doc(givenID).Get(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
		"erro ao encontrar id na collection": err.Error()})
		return
	}

	if err := doc.DataTo(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
		"erro ao extrair dado da struct User": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}

func Login(c *gin.Context){
	var u models.User
	var l models.Login
	usersCollection, _ := conectUserCollection()	
	
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro ao extrair dado da struct Login": err.Error()})
			return
	}	

	iter := usersCollection.Where("email", "==", l.Email).Documents(c)
	for {
		doc, err := iter.Next()
		if err == iterator.Done{
			break
		}
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"não conseguiu percorrer coleção": err.Error()})
		}
	
	if err := doc.DataTo(&u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
		"erro ao extrair dado da struct User": err.Error()})
		return
	}
}

if u.Password != l.Password{
		c.JSON(http.StatusBadRequest, "Senha incorreta")
		return
	}
	c.JSON(http.StatusAccepted,  "Usuário autorizado")
}
