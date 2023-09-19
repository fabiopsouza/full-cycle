package balance

import "net/http"

type Handler interface {
	List(w http.ResponseWriter, r *http.Request)
}
