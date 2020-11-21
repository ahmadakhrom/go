package modul

import "log"

//CheckErr func
func CheckErr(err error) {
	if err != nil {
		log.Fatalf("error: %s ", err)
		return
	}
}
