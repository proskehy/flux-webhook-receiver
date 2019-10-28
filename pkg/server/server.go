package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/proskehy/flux-webhook-receiver/pkg/handlers"
)

type Server struct {
	Handler handlers.Handler
	lock    uint32
}

func NewServer(h handlers.Handler) *Server {
	return &Server{
		Handler: h,
	}
}

func (s *Server) GitSync(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %s", err)
		return
	}

	if atomic.CompareAndSwapUint32(&s.lock, 0, 1) {
		go func() {
			s.Handler.GitSync(b, r.Header)
			atomic.StoreUint32(&s.lock, 0)
		}()
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
