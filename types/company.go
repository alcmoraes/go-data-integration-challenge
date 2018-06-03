package types

import "github.com/globalsign/mgo/bson"

type Company struct {
	ID      bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name    string        `bson:"name,omitempty" form:"name" json:"name" binding:"required"`
	Zip     string        `bson:"zip,omitempty" form:"zip" json:"zip" binding:"required"`
	Website string        `bson:"website,omitempty" form:"website" json:"website"`
}

type Companies struct {
	Companies []Company `bson:"companies,omitempty" json:"companies"`
}
