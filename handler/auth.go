package handler

import (
	"file-store-server/db"
	"log"
	"net/http"
)

// HTTPInterceptor HTTP 请求拦截器
func HTTPInterceptor(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			err := r.ParseForm()
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusForbidden)
				return
			}

			username := r.FormValue("username")
			token := r.FormValue("token")
			ok := db.IsTokenValid(username, token)
			if !ok {
				log.Println("Token is not validated")
				w.WriteHeader(http.StatusForbidden)
				return
			}
			h(w, r)
		})
}
