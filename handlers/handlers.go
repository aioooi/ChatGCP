package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/aioooi/goMicroservicesGCP/entity"
)

func GetShoutHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// TODO get data

		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusFound)
		rw.Write([]byte("Hello, World :)\n"))
		// TODO rw.Write(data)
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

		// TODO write data
	}
}
