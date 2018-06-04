package types

import "github.com/globalsign/mgo/bson"

type Company struct {
	ID      bson.ObjectId `bson:"_id,omitempty" json:"id" csv:"-"`
	Name    string        `bson:"name,omitempty" form:"name" csv:"name" json:"name" binding:"required"`
	Zip     string        `bson:"zip,omitempty" form:"zip" csv:"zip" json:"zip" binding:"required"`
	Website string        `bson:"website,omitempty" form:"website" csv:"website" json:"website"`
	Persist bool          `bson:"persist" form:"persist" json:"persist"`
}

type CompanyUploadQuery struct {
	Persist bool `bson:"persist" form:"persist" json:"persist"`
}
