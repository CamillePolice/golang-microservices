package domain

//We will be using mongo that’s why we use a bson.ObjectId

type Auth struct {
	Email    string `json:"email"  bson:"email"`
	Password string `json:"password" bson:"password"`
}