package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/proskehy/flux-webhook-receiver/pkg/handlers/git"
	"github.com/proskehy/flux-webhook-receiver/pkg/handlers/image"
)

type Server struct {
	GitHandler   git.Handler
	ImageHandler image.Handler
	lock         uint32
}

func NewServer() *Server {
	return &Server{}
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
			s.GitHandler.GitSync(b, r.Header)
			atomic.StoreUint32(&s.lock, 0)
		}()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Webhook received"))
	} else {
		w.WriteHeader(http.StatusLocked)
		w.Write([]byte("Already processing Git sync"))
	}
}

func (s *Server) ImageSync(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading request body: %s", err)
		return
	}
	go func() {
		s.ImageHandler.ImageSync(b, r.Header)
	}()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received"))
}
