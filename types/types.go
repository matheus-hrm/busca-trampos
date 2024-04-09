package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserbyID(id int) (*User, error)
	CreateUser(User)
}

type RegisterUserPayload struct {
	FirstName string 	`json:"firstName" validate="required"`
	LastName  string 	`json:"lastName" validate="required"`
	Email     string 	`json:"email" validate="required,email"`
	Password  string 	`json:"password" validate="required,min=3,max=24"`
}	
	

type User struct {
	ID        int    	`json:"id"`
	FirstName string 	`json:"firstName"`
	LastName  string 	`json:"LastName"`
	Email     string 	`json:"email"`
	Password  string 	`json:"-"`
	CreatedAt time.Time `json:"CreatedAt"`
}	

type Post struct {
	PostID			int
	Title 			string
	Description		string
	Address			string
	Salary 			float32
	Status			string 
	CreatedAt 		time.Time 
}