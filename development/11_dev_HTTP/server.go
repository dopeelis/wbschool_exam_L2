package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var userCount int = 1
var usersStore *userStore = &userStore{uStore: make(map[string]User)}

type User struct {
	ID         string                       `json:"user id"`
	Name       string                       `json:"user name"`
	EventCount int                          `json:"user event count"`
	EventStore map[string]map[string]string `json:"-"`
}

type userStore struct {
	uStore map[string]User
	sync.RWMutex
}

type Event struct {
	UserName  string `json:"user"`
	ID        string `json:"event id"`
	Title     string `json:"title"`
	Date      string `json:"date"`
	TimeStart string `json:"time start"`
	TimeEnd   string `json:"time end"`
}

func newUser(name string) *User {
	u := &User{ID: strconv.Itoa(userCount), Name: name, EventCount: 0}
	u.EventStore = make(map[string]map[string]string)
	usersStore.addUserToMap(u)
	userCount++
	return u
}

func newEvent(u *User, title string, date string, timeStart string, tineEnd string) *Event {
	u.EventCount = u.EventCount + 1
	e := &Event{UserName: u.Name, ID: strconv.Itoa(u.EventCount), Title: title, Date: date, TimeStart: timeStart, TimeEnd: tineEnd}

	u.addEvent(e)

	return e
}

func (u *User) addEvent(e *Event) {
	mu := new(sync.RWMutex)
	mu.Lock()
	u.EventStore[e.ID] = map[string]string{
		"title":      e.Title,
		"date":       e.Date,
		"time start": e.TimeStart,
		"time end":   e.TimeEnd,
	}
	mu.Unlock()
}

func (uS *userStore) addUserToMap(u *User) {
	uS.Lock()
	uS.uStore[u.ID] = *u
	uS.Unlock()
}

func main() {
	Bob := newUser("Bob")
	Lisa := newUser("Lisa")

	newEvent(Lisa, "Egor's DoB", "2021-10-24", "00:00", "23:59")
	newEvent(Bob, "Daily meeting", "2021-09-23", "12:50", "13:30")

	// newEvent(Lisa, "WB exam", "2021-10-01", "18:00", "21:00")

	http.HandleFunc("/", middleware(queryPar))

	host := ":8000"
	err := http.ListenAndServe(host, nil)
	if err != nil {
		log.Println("Error: 'listen and serve at host'", host, err.Error())
	}
}

func queryPar(rw http.ResponseWriter, r *http.Request) {
	if r.FormValue("user_id") != "" {
		if _, ok := usersStore.uStore[r.FormValue("user_id")]; ok {

			user := usersStore.uStore[r.FormValue("user_id")]
			if r.FormValue("date") == "" {
				_, err := rw.Write([]byte(fmt.Sprintf("Hi, %s! This is your awesome calendar\n", user.Name)))
				if err != nil {
					log.Println("Error: can't write text on user{id} page", err.Error())
				}
				rw.Header().Set("Content-Type", "application/json")
				rw.WriteHeader(http.StatusOK)
				jsonBytes, err := json.Marshal(user)
				if err != nil {
					log.Println("can't display user info")
				}
				rw.Write([]byte(fmt.Sprintln("\nUser info: ")))
				rw.Write(jsonBytes)
				rw.Write([]byte(fmt.Sprintln(" ")))
			}

			if r.FormValue("date") != "" {
				count := len(user.EventStore)
				for k, v := range user.EventStore {
					if v["date"] == r.FormValue("date") {
						_, err := rw.Write([]byte(fmt.Sprintf("Hi, %s!\n\nThis is events for day %s:\n\n", user.Name, r.FormValue("date"))))
						if err != nil {
							log.Println("Error: can't write text on user{id} & event{date} page", err.Error())
						}
						count--
						rw.Header().Set("Content-Type", "application/json")
						rw.WriteHeader(http.StatusOK)
						data := user.EventStore[k]
						jsonBytes, err := json.Marshal(data)
						if err != nil {
							log.Println("can't display data")
						}
						rw.Write(jsonBytes)
					} else {
						continue
					}
				}
				if count == len(user.EventStore) {
					rw.WriteHeader(http.StatusNotFound)
					rw.Write([]byte(`"error":"not found"`))
				}
			}

		} else {
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte(`"error":"not found"`))
		}
	} else {
		homePage(rw, r)
	}

}

func homePage(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)

	_, err := rw.Write([]byte("Hi! This is your awesome calendar"))
	if err != nil {
		log.Println("Error: can't write []byte on home page", err.Error())
	}
}

func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Request info: ", r)
		next(rw, r)
	}
}
