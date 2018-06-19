// Log all received data

package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func main() {
	addr := flag.String("addr", "", "The address to listen to; default is \"\" (all interfaces).")
	port := flag.Int("port", 8000, "The port to listen on.")

	flag.Parse()

	src := *addr + ":" + strconv.Itoa(*port)
	fmt.Printf("Listening on %s.\n", src)
	err := http.ListenAndServe(src, http.HandlerFunc(logRequest))
	if err != nil {
		log.Fatal(err)
	}
}

func logRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	buf, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Print("Error ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("\n%s", hex.Dump(buf))
}
