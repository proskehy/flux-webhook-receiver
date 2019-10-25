package main

import (
	"log"
	"net/http"
	"sync/atomic"
	"io/ioutil"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
)

type Message struct {
	Ref string `json:"ref"`
	Repository struct {
		Url string `json:"url"`
	} `json:"repository"`
}

type Config struct {
	Secret string
}

type Server struct {
	myHandler *MyHandler
	lock      uint32
	Config *Config
}

type MyHandler struct{}

func (h *MyHandler) Sync(m Message) {
	//fmt.Println("hi")
	log.Println(m)
}

func (s *Server) Hello(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()		
	var m Message	
//	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
//		log.Println(err)
//		return
//	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("error reading")
		log.Println(err)
		return 
	}
	if err := json.Unmarshal(b, &m); err != nil {
		log.Println("error unmarshalling")
		log.Println(err)
		return
	}
	log.Println("header verification:")
	log.Println(verifySignature(r.Header.Get("X-Hub-Signature"), b, s.Config.Secret))
	
	if atomic.CompareAndSwapUint32(&s.lock, 0, 1) {
		go func() {
			s.myHandler.Sync(m)
			atomic.StoreUint32(&s.lock, 0)
		}()
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func verifySignature(signature string, payload []byte, secret string) bool {
	if len(signature) == 0 {
		return false
	}
	mac := hmac.New(sha1.New, []byte(secret))
	_, _ = mac.Write(payload)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))

	if !hmac.Equal([]byte(signature[5:]), []byte(expectedMAC)) {
		return false
	}
	return true
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	h := &MyHandler{}
	c := &Config{Secret: "test"}
	s := &Server{myHandler: h, Config: c}
	http.HandleFunc("/", s.Hello)
	log.Fatal(http.ListenAndServe(":3031", nil))
}
