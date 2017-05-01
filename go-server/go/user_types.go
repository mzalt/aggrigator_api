package 

import (
	"net/http"
)

type User_types struct {

}

func GetTypes(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

