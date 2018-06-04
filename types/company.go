package types

import "github.com/globalsign/mgo/bson"

// Company defaul struct
type Company struct {
	ID      bson.ObjectId `bson:"_id,omitempty" json:"id" csv:"-"`
	Name    string        `bson:"name,omitempty" form:"name" csv:"name" json:"name" binding:"required"`
	Zip     string        `bson:"zip,omitempty" form:"zip" csv:"zip" json:"zip" binding:"required"`
	Website string        `bson:"website,omitempty" form:"website" csv:"website" json:"website"`
}

// CompanyUploadQuery schema for uploading CSV files
type CompanyUploadQuery struct {
	Persist bool `bson:"persist" form:"persist" json:"persist"`
}
