package models

type Shop struct {
	ID      string  `json:"id" firestore:"id"`
	Name    string  `json:"name" firestore:"name" binding:"required"`
	Flavors [6]string `json:"flavors" firestore:"flavors"`
	Score   float64 `json:"score" firestore:"score" binding:"required"`
	Price   float64 `json:"price" firestore:"price"`
	Link    string  `json:"link" firestore:"link"`
	Favorite	bool  `json:"favorite" firestore:"favorite"`
}

type NameUpdated struct{
	NewName	string	`json:"newname" firestore:"newname"`
}

type PriceUpdated struct{
	NewPrice	float64	`json:"newprice" firestore:"newprice"`
}


