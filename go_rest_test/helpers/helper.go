package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Catch(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func respondWIthError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"message": msg})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//func Logger() http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		fmt.Println(time.Now(), r.Method, r.URL)
//		router.Serve
//	})
//}