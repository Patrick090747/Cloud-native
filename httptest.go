package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

func HttpsServer(port int) {
	log.SetPrefix("info:")
	log.SetFlags(log.Ldate | log.Llongfile)
	http.HandleFunc("/", HttpAccess)
	http.HandleFunc("healthz", HealthzCheck)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err)
	}

}
func HealthzCheck(w http.ResponseWriter, r *http.Request) {
	HeaalthzCode := "200"
	w.Write([]byte(HeaalthzCode))
}
func HttpAccess(w http.ResponseWriter, r *http.Request) {
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			log.Printf("%s=%s", k, v[0])
			w.Header().Set(k, v[0])
		}
	}
	r.ParseForm()
	if len(r.Form) > 0 {
		for k, v := range r.Form {
			log.Printf("%s=%s", k, v[0])
		}
	}

	log.Printf("\nVersionENV:", os.Getenv("VERSION"))
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("err:", err)
	}
	if net.ParseIP(ip) != nil {
		fmt.Printf("ip=%s", ip)
		log.Println(ip)
	}
	fmt.Printf("status code is %s", http.StatusOK)
	log.Println(http.StatusOK)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}
func main() {
	HttpsServer(8080)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
