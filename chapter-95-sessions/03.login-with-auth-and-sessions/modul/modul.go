package modul

import (
	"net/http"
	"time"

	"chapter-95-sessions/03.login-with-auth-and-sessions/user"

	"github.com/google/uuid"
)

//GetCookie func
func GetCookie(w http.ResponseWriter, r *http.Request, cookie string) user.Users {
	var u user.Users
	c, err := r.Cookie(cookie)
	if err != nil {
		d := uuid.New()
		c = &http.Cookie{
			Name:     "_uid",
			Value:    d.String(),
			HttpOnly: true,
			Expires:  time.Now().Add(time.Hour * 1), //1 year
		}
		http.SetCookie(w, c)
		return u
	}

	if val, status := user.Cookie[c.Value]; status == true {
		u = user.DetailUser[val]
	}
	return u
}

//GetSession func
func GetSession(r *http.Request, cookie string) error {
	_, err := r.Cookie(cookie)
	if err != nil {
		return err
	}

	return nil
	//check value cookie whats that is username already on DB,
	//if ok, if user on index position direct to dashboard, if on certain page direct to there page
}
