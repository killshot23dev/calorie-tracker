package models

import(
	"go-mongodb.org/mongo-driver/bson/primitive"

)

type Entry struct{
	ID			primitive.ObjectId `bson:"id"`
	Dish	 	*string 		   `json:"dish"`  //we put a dish is what we are passing to golang and 
	//id is generated so we pass the reference to the dish to generate the proper ID
	fat 		*float64		   `json:"fat"`
	Calories    *float64		   `json:"calories"`
	Ingredients  *string 		   `json:"ingredients"`
}
