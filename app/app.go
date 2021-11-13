package app

import (
	"GoDatabaseAssignment/config"
	"GoDatabaseAssignment/pairs"
	"GoDatabaseAssignment/records"
	"log"
	"net/http"
)

/*var routes = []route{
	newRoute("GET", "/", home),
	newRoute("GET", "/contact", contact),
	newRoute("GET", "/api/widgets", apiGetWidgets)
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

func Serve(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
}*/

func Init() {
	addr := config.Get().Servers.AppServer
	log.Printf("App Address : %s", addr)

	http.HandleFunc("/records", records.Records)
	http.HandleFunc("/config", pairs.KeyValuePairs)
	http.ListenAndServe(addr, nil)

}
