package handler

import (
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintf(w, "<h1>The time is: %s </h1>", t.Format("15:04"))
}
