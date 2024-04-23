package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserbyID(id int) (*User, error)
	CreateUser(User) error 
}

type PostsStore interface {
	GetPosts() ([]Post, error)
	GetPostByID(id int) (*Post, error)
	CreatePost(Post) error
	UpdatePost(Post) error
	DeletePost(id int) error
}
type RegisterUserPayload struct {
	FirstName 		string 			`json:"firstName" validate:"required"`
	LastName  		string 			`json:"lastName" validate:"required"`
	Email     		string 			`json:"email" validate:"required,email"`
	Password  		string 			`json:"password" validate:"required,min=3,max=24"`
}	

type LoginUserPayload struct {
	Email    		string 			`json:"email" validate:"required,email"`
	Password 		string 			`json:"password" validate:"required"`
}

type User struct {
	ID        		int    			`json:"id"`
	FirstName 		string 			`json:"firstName"`
	LastName  		string 			`json:"LastName"`
	Email     		string 			`json:"email"`
	Password  		string 			`json:"-"`
	CreatedAt 		time.Time 		`json:"CreatedAt"`
}	

type Post struct {
	ID			    int				`json:"id"`
	Title 			string  		`json:"title"`
	Description		string 			`json:"description"`
	Address			string			`json:"address"`
	Salary 			float64 		`json:"salary"`
	Status			string 			`json:"status"`
	CreatedAt 		time.Time 		`json:"createdAt"`
}