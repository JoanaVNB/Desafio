package models

type User struct {
	name 		string		`json: "name, omitempty"`
	e-mail 		string		`json:"email,omitempty"`
	password string		`json:"password,omitempty"`
}

type Shop struct{
	ID				uint64		`json:"id,omitempty"`
	name		string		`json:"nameomitempty"`
	flavors		Flavors		`json:"flavors,omitempty"`
	score		float64		`json:"score,omitempty"`
	price		float64		`json:"price,omitempty"`
	link			string		`json:"link,omitempty"`
}

type Flavors struct{
	flavorOne	string		`json:"flavorOne,omitempty"`
	flavorTwo	string		`json:"flavorTwo"`
	flavorThree	string		`json:"flavorThree"`
	flavorFour	string		`json:"flavorFour"`
	flavorFive	string		`json:"flavorFive"`
}