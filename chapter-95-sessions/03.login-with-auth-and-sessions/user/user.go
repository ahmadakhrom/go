package user

//Users Struct
type Users struct {
	Username string
	Password string
}

//DetailCookie variable
var DetailCookie = map[string]string{}

//DetailUser variable
var DetailUser = map[string]Users{}
