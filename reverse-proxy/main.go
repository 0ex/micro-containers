/**
 * micro reverse proxy
 *
 * vim:ts=4:noet:
 **/
package main

import (
	"os"
	"fmt"
	"flag"
	"log"
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var port = flag.Int("p", 80, "port")

func main() {
	var prog = os.Args[0]
	flag.Parse()
	
	if len(flag.Args()) != 1 {
		fmt.Printf("usage: %v [-port 80] [target-url]\n", prog)
		os.Exit(1)
	}

	target, err := url.Parse(flag.Args()[0])
	if err != nil {
		panic(err)
	}

	var listen = fmt.Sprintf(":%d", *port)
	log.Printf("proxying %s to %s", listen, target)

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify : true},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		r.Header.Set("X-Forwarded-Host", r.Host)
		r.Host = target.Host

		proxy.ServeHTTP(w, r)
	})

	err = http.ListenAndServe(listen, nil)
	if err != nil {
		panic(err)
	}
}

