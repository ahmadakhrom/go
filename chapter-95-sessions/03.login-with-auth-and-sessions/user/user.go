package user

//Users Struct
type Users struct {
	Username string
	Password string
	Role     string
}

//DetailCookie variable
var Cookie = map[string]string{}

//M map string
type M = map[string]interface{}

//DetailUser variable
var DetailUser = map[string]Users{}
