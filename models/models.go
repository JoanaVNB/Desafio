package models

type User struct {
	ID				string		`json:"id"		firestore:"-"`
	Name 		string		`json: "name"	firestore:"id"` 
	Email 		string		`json:"email"	firestore:"email"`
	Password string		`json:"password"		firestore:"password"`
}

type Shop struct{
	ID				string		`json:"id"		firestore:"id"`
	Name		string		`json:"name, omitempty "	firestore:"name"`
	Flavors		Flavors		`json:"flavors, omitempty"	firestore:"flavors"`
	Score		float64		`json:"score, omitempty"	firestore:'score"`
	Price		float64		`json:"price, omitempty"	firestore:'"price"`
	Link			string		`json:"link, omitempty"		firestore:'link"`
}

type Flavors struct{
	FlavorOne	string		`json:"flavorOne, omitempty"	firestore:"flavor_one"`
	FlavorTwo	string		`json:"flavorTwo"	firestore:"flavor_two"`
	FlavorThree	string		`json:"flavorThree	firestore:"flavor_three""`
	FlavorFour	string		`json:"flavorFour"	firestore:"flavor_four"`
	FlavorFive	string		`json:"flavorFive"	firestore:"flavor_five"`
}