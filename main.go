package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"flag"
)

func getIP(w http.ResponseWriter, r *http.Request) {
	ipr := regexp.MustCompile(`^(?:(?P<ipv4>[^:]+)|\[(?P<ipv6>.+)\]):(?P<port>\d+)$`)
	ip := ipr.FindStringSubmatch(r.RemoteAddr)
	if ip[1] == "" {
		// ip[1] is empty when client uses IPv6
		fmt.Fprintf(w, "%s\n", ip[2])
	} else {
		fmt.Fprintf(w, "%s\n", ip[1])
	}
	
}

func main() {
	listen := flag.String("l", ":8080", "Listening address")

	http.HandleFunc("/", getIP)
	log.Printf("Serving on %s", *listen)
	log.Fatal(http.ListenAndServe(*listen, nil))
}
