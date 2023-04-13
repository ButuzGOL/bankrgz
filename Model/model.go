package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type BankBranch struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Address  string             `json:"address"`
	Number   int16              `json:"number"`
	Phone    string             `json:"phone"`
	District string             `json:"district"`
}
