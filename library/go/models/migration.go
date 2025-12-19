package models

type Migration struct {
	ID         []byte `json:"id" bson:"id"`
	EntryPoint string `json:"username" bson:"entryPoint"`
	Action     string `json:"password" bson:"action"`
}
