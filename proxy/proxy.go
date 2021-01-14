package main

// import (
// 	"io"
// 	"net/http"
// 	"net/url"
// )

// func main() {
// 	http.ListenAndServe("0.0.0.0:8080", HTTPHandler(http.DefaultServeMux, "8080"))
// }

// //HTTPHandler handles all request routes
// func HTTPHandler(handler http.Handler, port string) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		handleHTTP(w, r, port)
// 	})
// }

// func handleHTTP(w http.ResponseWriter, req *http.Request, port string) {
// 	println("###Incoming HTTP request to: " + req.Method + " " + req.Host + req.URL.Path + "?" + req.URL.RawQuery)
// 	//Request url needs to contain http scheme(an open request)
// 	//Sends request to RedGesher(They communicate with an open connection HTTP)
// 	req.URL, _ = url.Parse("http://" + "192.168.1.10:8080" + "/" + req.Host + req.URL.Path + "?" + req.URL.RawQuery)
// 	req.Header.Set("X-Forwarded-Host", req.Host+":"+port)
// 	req.Header.Set("X-Forwarded-For", req.RemoteAddr)
// 	println("###Sending HTTP request to: " + req.Method + " " + req.URL.String())
// 	resp, err := http.DefaultTransport.RoundTrip(req)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusServiceUnavailable)
// 		return
// 	}
// 	defer resp.Body.Close()
// 	copyHeader(w.Header(), resp.Header)
// 	w.WriteHeader(resp.StatusCode)
// 	io.Copy(w, resp.Body)
// }

// func copyHeader(dst, src http.Header) {
// 	for k, vv := range src {
// 		for _, v := range vv {
// 			dst.Add(k, v)
// 		}
// 	}
// }
