package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/aioooi/goMicroservicesGCP/entity"
)

func GetShoutHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(entity.AllShouts())
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusFound)
		rw.Write(data)
	}
}

func PostShoutHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var shout entity.Shout
		err = json.Unmarshal(data, &shout)
		if err != nil {
			rw.WriteHeader(http.StatusExpectationFailed)
			rw.Write([]byte("Invalid data format"))
			return
		}

		if shout.Author == "" || shout.Message == "" {
			rw.WriteHeader(http.StatusExpectationFailed)
			rw.Write([]byte("Empty author or message field"))
			return
		}
		entity.PostShout(shout.Author, shout.Message)
		rw.WriteHeader(http.StatusOK)
	}
}
