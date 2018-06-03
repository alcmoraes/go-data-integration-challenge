package types

import "gopkg.in/mgo.v2/bson"

type Company struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Name    string        `bson:"name,omitempty" form:"name" json:"name" binding:"required"`
	Zip     string        `bson:"zip,omitempty" form:"zip" json:"zip" binding:"required"`
	Website string        `bson:"website,omitempty" form:"website" json:"website"`
}
