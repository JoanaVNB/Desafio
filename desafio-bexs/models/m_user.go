package models

type User struct {
	ID       string `json:"id" firestore:"id"`
	Name     string `json:"name" firestore:"name" binding:"required"`
	Email    string `json:"email" firestore:"email" binding:"required,email"`
	Password string `json:"password" firestore:"password" binding:"required"`
}

type Login struct{
	Email    string `json:"email"  firestore:"email" `
	Password string `json:"password" firestore:"password"`
}


