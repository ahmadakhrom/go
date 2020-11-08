package modul

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Fatalf("error: %s ", err)
		return
	}
}
