package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"
)

var users []*User

// I still have to see how to avoid this global variable, but for now, lets leave it
var logger *log.Logger = log.New(os.Stdout, "Simple Api - ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)

func main() {
	mux := http.NewServeMux()
	port, pOk := os.LookupEnv("APP_PORT")

	users = append(users, NewUser(&CreateUserDto{Name: "Mitchel Hashimoto", Username: "mitchellh"}))

	if !pOk {
		port = "8080"
	}

	RegisterRoutes(mux)

	server := http.Server{
		Addr:         "127.0.0.1:" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	logger.Println("Starting application on address: ", server.Addr)

	server.ListenAndServe()
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		resp, err := json.Marshal(NewApi())

		if err != nil {
			bytes, err := json.Marshal(NewApiError("Internal Server Error", err, http.StatusInternalServerError))
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(bytes)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	})

	mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			resp, err := json.Marshal(users)

			if err != nil {
				bytes, _ := json.Marshal(NewApiError("Internal Server Error", err, http.StatusInternalServerError))

				logger.Println(err.Error())

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(bytes)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(resp)
		} else if r.Method == "POST" {
			if r.Header.Get("Content-Type") != "application/json" {
				apiErr := NewApiError("Unsupported Media Type", errors.New("The only supported Content-Type is application/json"), http.StatusUnsupportedMediaType)
				bytes, _ := json.Marshal(apiErr)

				logger.Println(apiErr.Message)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(apiErr.Status)
				w.Write(bytes)
				return
			}

			var dto CreateUserDto
			err := json.NewDecoder(r.Body).Decode(&dto)

			if err != nil {
				apiErr := NewApiError("Bad Request", err, http.StatusBadRequest)
				bytes, _ := json.Marshal(apiErr)

				logger.Println(apiErr.Message)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(apiErr.Status)
				w.Write(bytes)
			}

			u := NewUser(&dto)
			users = append(users, u)

			logger.Println("A new user was added")

			resp, _ := json.Marshal(u)

			w.Header().Set("Content-Type", "application")
			w.WriteHeader(http.StatusCreated)
			w.Write(resp)
		}
	})
}
