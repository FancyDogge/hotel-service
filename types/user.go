package types

type User struct {
	ID        string `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string `bson:"username,omitempty" json:"username,omitempty"`
	FirstName string `bson:"firstName,omitempty" json:"firstName,omitempty"`
	LastName  string `bson:"lastName,omitempty" json:"lastName,omitempty"`
}
