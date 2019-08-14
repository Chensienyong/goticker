package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Event struct {
	TimeToShow  time.Time
	TimeEndShow time.Time
	Name        string
}

func main() {
	router := httprouter.New()
	router.GET("/healthz", Healthz)
	log.Println("Listening at port 1337")
	go func() {
		log.Fatal(http.ListenAndServe(":1337", router))
	}()

	now := time.Now()
	events := []Event{
		Event{TimeToShow: now.Add(2 * time.Second), TimeEndShow: now.Add(4 * time.Second), Name: "First"},
		Event{TimeToShow: now.Add(6 * time.Second), TimeEndShow: now.Add(9 * time.Second), Name: "Second"},
		Event{TimeToShow: now.Add(10 * time.Second), TimeEndShow: now.Add(11 * time.Second), Name: "Third"},
		Event{TimeToShow: now.Add(15 * time.Second), TimeEndShow: now.Add(18 * time.Second), Name: "Fourth"},
		Event{TimeToShow: now.Add(20 * time.Second), TimeEndShow: now.Add(22 * time.Second), Name: "Last"},
	}

	ticker := time.NewTicker(time.Second)
	go func() {
		for t := range ticker.C {
			now := time.Now()
			if len(events) > 0 {
				if t.Truncate(time.Second).Sub(events[0].TimeToShow.Truncate(time.Second)) >= 0 && t.Truncate(time.Second).Sub(events[0].TimeEndShow.Truncate(time.Second)) <= 0 {
					fmt.Println(events[0].Name)
				} else {
					fmt.Println("Time: ", t.Truncate(time.Second))
				}
				if t.Truncate(time.Second).Sub(events[0].TimeEndShow.Truncate(time.Second)) >= 0 {
					events = events[1:]
				}
			} else {
				fmt.Println("Nothing")
			}
			fmt.Println(time.Now().Sub(now))
		}
	}()
	select {}
}

func Healthz(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "ok")
}
