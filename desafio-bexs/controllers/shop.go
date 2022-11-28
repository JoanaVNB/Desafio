package controllers

import (
	"context"
	"desafio/models"
	"log"
	"net/http"
	"os"
	"strconv"
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"github.com/google/uuid"
	"errors"
	"github.com/go-playground/validator/v10"
	"sort"
)

func conectShopCollection() (*firestore.CollectionRef, error){
	_ = os.Setenv("FIRESTORE_EMULATOR_HOST", "firestore:9091")

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "desafio-c0479")
	if err != nil {
		log.Println(err)
	}
	shopCollection := client.Collection("Shop")
	return shopCollection, err
}

func nameRegistered(name string , c *gin.Context) bool{
	shopCollection, _ := conectShopCollection()	

	iter := shopCollection.Where("name", "==", name).Documents(c)
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
			c.JSON(http.StatusConflict, "Nome da loja já foi cadastrada.")
			return true
		}
	}	
	c.JSON(http.StatusOK, "Nome da loja pode ser criada")
			return false
}	

func Create(c *gin.Context){
	var s models.Shop
	var ve validator.ValidationErrors
	shopCollection, err := conectShopCollection()
	

	if err := c.ShouldBindJSON(&s); err != nil {
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

	if nameRegistered(s.Name, c) == false{
			s.ID = uuid.NewString()
			_, err = shopCollection.Doc(s.ID).Create(c, s)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro ao criar na collection": err.Error()})
			return
			}
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
		
		if err := doc.DataTo(&s); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
		"erro ao extrair da struct User": err.Error()})
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
		"erro ao extrair da struct Shop": err.Error()})
		return
	}
	s.ID = doc.Ref.ID
	c.JSON(http.StatusOK, s)
}

func ReadByName(c *gin.Context){
	var s models.Shop
	shopCollection, _ := conectShopCollection()

	name:= c.Params.ByName("name")

	iter := shopCollection.Where("name", "==", name).Documents(c)
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
		"erro ao extrair da struct Shop": err.Error()})
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

	iter := shopCollection.Where("score", ">=", score).Documents(c)
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
		return
	}

	iter := shopCollection.Where("price", "<=", price).Documents(c)
	for {
		doc, err := iter.Next()
		if err == iterator.Done{
			break
		}
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro ao encontrar preço na collection": err.Error()})
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

func UpdateScore(c *gin.Context){
	shopCollection, _ := conectShopCollection()
	
	givenID := c.Params.ByName("id")
	score, err:= strconv.ParseFloat(c.Param("score"), 64)
		_, err = shopCollection.Doc(givenID).Update(c, []firestore.Update{{Path: "score", Value: score}})
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"erro ao atualizar": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "score atualizado")
}

func UpdatePrice(c *gin.Context){
	var p models.PriceUpdated

	shopCollection, _ := conectShopCollection()
	givenID := c.Params.ByName("id")

	if err := c.ShouldBindJSON(&p); err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{
			"erro ao converter": err.Error()})
		return
	}
		_, err := shopCollection.Doc(givenID).Update(c, []firestore.Update{{Path: "price", Value: p.NewPrice}})
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"erro ao atualizar": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "preço atualizado")
}

func UpdateName(c *gin.Context){
	var n models.NameUpdated

	shopCollection, _ := conectShopCollection()
	givenID := c.Params.ByName("id")

	if err := c.ShouldBindJSON(&n); err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}

	_, err := shopCollection.Doc(givenID).Update(c, []firestore.Update{{Path: "name", Value: n.NewName}})
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"erro ao atualizar": err.Error()})
		return
	}
	c.JSON(http.StatusOK,  "nome atualizado")
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

func ListScores(c *gin.Context) (map[string]float64){
	var s models.Shop
	shopCollection, _:= conectShopCollection()

	scores := make(map[string]float64)

	iter := shopCollection.Documents(c)
	for {
		doc, err := iter.Next()
		if err == iterator.Done{
			break
		}
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"erro": err.Error()})
				return nil
		}
		
		if err := doc.DataTo(&s); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
		"erro ao extrair dados da struct Shop": err.Error()})
		return nil
		}
		scores[s.Name] = s.Score
	}
	return scores
}

func Ranking(c *gin.Context){
	list := ListScores(c)
	
	keys := make([]string, 0, len(list))

	for k := range list{
		keys = append(keys, k)
	}
	
	sort.SliceStable(keys, func(i, j int) bool{
		return list[keys[i]] > list[keys[j]]
	})

	c.JSON(http.StatusOK, keys)
}
