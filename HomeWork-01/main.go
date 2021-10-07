package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(":80", nil)
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", rootHandler)
	// mux.HandleFunc("/healthz", healthz)
	// mux.HandleFunc("/debug/pprof/", pprof.Index)
	// mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	// mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	// mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	// err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	// use url localhost:80/?user=xxx to get user parameter
	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	//to write string message body to response
	io.WriteString(w, "===================Details of the http request header:============\n")
	//all request header message write to response

	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
		//can write into response header,but cannot check in website dev Mode
		w.Header().Set(k, respHeader(v))
	}
}

func respHeader(value []string) string {
	var result string
	for _, v := range value {
		result += v
	}
	if result != "" {
		fmt.Println(result)
	} else {
		fmt.Println("null")
	}

	return result
}
