package models

import "encoding/json"

type Policy struct {
	ID        []byte          `json:"id" bson:"id"`
	Name      string          `json:"name" bson:"name"`
	RoleId    []byte          `json:"roleId" bson:"roleId"`
	Actions   json.RawMessage `json:"actions" bson:"actions"` // json with slice actions Ids
	CreatedAt int64           `json:"createdAt" bson:"createdAt"`
	UpdatedAt int64           `json:"updatedAt" bson:"updatedAt"`
}

/*

Actions value:

{
	"actions":[] // slice of actionIds
}

*/
