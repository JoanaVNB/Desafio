package controllers

import (
	"context"
	"desafio/models"
	"log"
	"net/http"
	"os"
	"fmt"
	"strconv"
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	//"google.golang.org/api/option"
	//"firebase.google.com/go/v4"
	"github.com/google/uuid"
)

func conectShopCollection() (*firestore.CollectionRef, error){
	_ = os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:9091")

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "desafio-c0479") //NÃO É CLIENT
	if err != nil {
		log.Println(err)
	}
	shopCollection := client.Collection("Shop")
	return shopCollection, err
}

func Create(c *gin.Context){
	var s models.Shop
	shopCollection, err := conectShopCollection()
	
	if err := c.ShouldBindJSON(&s); err != nil { //converte os dados recebidos em JSON para bites para alocar dentro da struct User(u)
		c.JSON(http.StatusBadRequest, gin.H{
			"erro em converter byte para json": err.Error()})
		return
	}
	
	s.ID = uuid.NewString()

	_, err = shopCollection.Doc(s.ID).Create(c, s)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro ao criar na collection": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, s)
}

func ListAll(c *gin.Context){
	var s models.Shop
	shopCollection, _:= conectShopCollection()

	iter := shopCollection.Documents(c)
	for {
		doc, err := iter.Next()
		if err == iterator.Done{
			break
		}
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro ao encontrar nome na collection": err.Error()})
				return
		}
		
		if err := doc.DataTo(&s); err != nil {//we can extract the document's data into a value of type Shop
			c.JSON(http.StatusInternalServerError, gin.H{
		"erro ao passar dado para struct User": err.Error()})
		return
		}
	
	c.JSON(http.StatusOK,  s)
	}
}

func ReadByID(c *gin.Context){
	var s models.Shop
	shopCollection, _:= conectShopCollection()

	givenID, _:= c.Params.Get("id")
	doc, err := shopCollection.Doc(givenID).Get(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
		"erro ao encontrar id na collection": err.Error()})
		return
	}

	if err := doc.DataTo(&s); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
		"erro ao passar dado para struct Shop": err.Error()})
		return
	}
	s.ID = doc.Ref.ID
	c.JSON(http.StatusOK, s)
}

func ReadByName(c *gin.Context){
	var s models.Shop
	shopCollection, _ := conectShopCollection()

	name:= c.Params.ByName("name")
	fmt.Println("loja:", name)

	//na coleção Shop, procurar por name igual a name
	iter := shopCollection.Where("Name", "==", name).Documents(c)
	for {
		doc, err := iter.Next()
		if err == iterator.Done{
			break
		}
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro ao encontrar nome na collection": err.Error()})
				return
		}
			
		if err := doc.DataTo(&s); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
		"erro ao passar dado para struct User": err.Error()})
		return
		}
	
	c.JSON(http.StatusOK,  s)
	}
}

func ReadByScore(c *gin.Context){
	var s models.Shop
	shopCollection, _ := conectShopCollection()

	score, err:= strconv.ParseFloat(c.Param("score"), 64)
	if err != nil{
		c.JSON(http.StatusBadRequest, "erro ao converter para float64")
	}
	fmt.Println("nota:", score)

	iter := shopCollection.Where("Score", ">=", score).Documents(c)
	for {
		doc, err := iter.Next()
		if err == iterator.Done{
			break
		}
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro ao encontrar nota na collection": err.Error()})
				return
		}
		
		if err := doc.DataTo(&s); err != nil {
			c.JSON(http.StatusInternalServerError, 	"erro extrair da struct Shop")
		return
		}
	
	c.JSON(http.StatusOK, s)
	}
}

func ReadByPrice(c *gin.Context){
	var s models.Shop
	shopCollection, _ := conectShopCollection()

	price, err:= strconv.ParseFloat(c.Param("price"), 64)
	if err != nil{
		c.JSON(http.StatusBadRequest, "erro ao converter para float64")
	}
	fmt.Println("preço:", price)

	iter := shopCollection.Where("Price", "<=", price).Documents(c)
	for {
		doc, err := iter.Next()
		if err == iterator.Done{
			break
		}
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro ao encontrar nota na collection": err.Error()})
			return
		}
		if err := doc.DataTo(&s); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro extrair da struct Shop": "tente novamente"})
			return
		}
	
	c.JSON(http.StatusOK, s)
	}
}

func Update(c *gin.Context){
	var s models.Shop
	shopCollection, _ := conectShopCollection()

	givenID := c.Params.ByName("id")
	doc, err := shopCollection.Doc(givenID).Get(c.Request.Context())
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"erro ao retornar documento": err.Error()})
		return
	}
	
	if err:= c.ShouldBindJSON(&s); err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"erro ao converter para struct": err.Error()})
		return
	}
	s.ID = uuid.NewString()
	s.ID = doc.Ref.ID
	shopCollection.Doc(givenID).Set(c, s)
	c.JSON(http.StatusOK, s)
}

func UpdateNoteToZero(c *gin.Context){
	shopCollection, _ := conectShopCollection()

	givenID := c.Params.ByName("id")
	_, err := shopCollection.Doc(givenID).Update(c, []firestore.Update{{Path: "Score", Value: 0}})
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro ao atualizar": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "zerado")
}

func UpdatePrice(c *gin.Context){
	var s models.Shop
	shopCollection, _ := conectShopCollection()

	givenID := c.Params.ByName("id")
	doc, err := shopCollection.Doc(givenID).Get(c.Request.Context())
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"erro ao retornar documento": err.Error()})
		return
	}
	
	if err:= c.ShouldBindJSON(&s.Price); err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"erro ao converter para struct": err.Error()})//"erro ao converter para struct": "json: cannot unmarshal object into Go value of type float64"
		return
	}
	s.ID = uuid.NewString()
	s.ID = doc.Ref.ID
	shopCollection.Doc(givenID).Set(c, s)
	c.JSON(http.StatusOK, gin.H{
		"preço atualizado": s})
}

func UpdateName(c *gin.Context){
	var s models.Shop
	shopCollection, _ := conectShopCollection()

	givenID := c.Params.ByName("id")
	doc, err := shopCollection.Doc(givenID).Get(c.Request.Context())
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"erro ao retornar documento": err.Error()})
		return
	}
	
	if err:= c.ShouldBindJSON(&s.Name); err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"erro ao converter para struct": err.Error()})//"json: cannot unmarshal object into Go value of type string"
		return
	}
	s.ID = uuid.NewString()
	s.ID = doc.Ref.ID
	shopCollection.Doc(givenID).Set(c, s)
	c.JSON(http.StatusOK, gin.H{
		"preço atualizado": s})
}

func Delete(c *gin.Context){
	shopCollection, _ := conectShopCollection()

	givenID := c.Params.ByName("id")
	_, err := shopCollection.Doc(givenID).Delete(c); if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"erro ao deletar": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "apagado")
}
