// This contains the collections struct that holds the names of the collections in the database.

package types

type Collections struct {
	Database string
	Users    string
	Messages string
}

var ProximityChat = Collections{
	Database: "ProximityChat",
	Users:    "Users",
	Messages: "Messages",
}
