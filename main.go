package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"flag"
)

func getIP(w http.ResponseWriter, r *http.Request) {
	var client_ip string

	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		client_ip = forwarded
	} else {
		ipr := regexp.MustCompile(`^(?:(?P<ipv4>[^:]+)|\[(?P<ipv6>.+)\]):(?P<port>\d+)$`)
		ip := ipr.FindStringSubmatch(r.RemoteAddr)
	
		if ip[1] == "" {
			// ip[1] is empty when client uses IPv6
			client_ip = ip[2]
		} else {
			client_ip = ip[1]
		}
	}
	fmt.Fprintf(w, "%s\n", client_ip)
}

func main() {
	listen := flag.String("l", ":8080", "Listening address")

	http.HandleFunc("/", getIP)
	log.Printf("Serving on %s", *listen)
	log.Fatal(http.ListenAndServe(*listen, nil))
}
