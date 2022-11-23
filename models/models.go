package models

type User struct {
	ID				string		`json:"-"		firestore:"-"`
	name 		string		`json: "name, omitempty"	firestore:"id"` 
	email 		string		`json:"email, omitempty"	firestore:"email"`
	password string		`json:"password, omitempty"		firestore:"password"`
}

type Shop struct{
	ID				string		`json:"id"		firestore:"id"`
	name		string		`json:"name, omitempty "	firestore:"name"`
	flavors		Flavors		`json:"flavors, omitempty"	firestore:"flavors"`
	score		float64		`json:"score, omitempty"	firestore:'score"`
	price		float64		`json:"price, omitempty"	firestore:'"price"`
	link			string		`json:"link, omitempty"		firestore:link"`
}

type Flavors struct{
	flavorOne	string		`json:"flavorOne, omitempty"	firestore:"flavor_one"`
	flavorTwo	string		`json:"flavorTwo"	firestore:"flavor_two"`
	flavorThree	string		`json:"flavorThree	firestore:"flavor_three""`
	flavorFour	string		`json:"flavorFour"	firestore:"flavor_four"`
	flavorFive	string		`json:"flavorFive"	firestore:"flavor_five"`
}