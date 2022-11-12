package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

var (
	listen string
)

func main() {
	flag.StringVar(&listen, "listen", "0.0.0.0:8080", "Listen address for HTTP server")
	flag.Parse()

	s, err := newServer()
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		// helper print out available URLs
		rw.Write([]byte("/streams/v1/index.json\n/streams/v1/com.ubuntu.cloud-released-imagemetadata.json\n"))
	}))

	http.Handle("/streams/v1/index.json", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ubuntuImages, err := s.getUbuntuImages()

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			log.Printf("Failed to retrieve ubuntu images: %v\n", err)
			return
		}

		index := s.makeIndexFromImages(ubuntuImages)
		b, err := json.Marshal(index)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			log.Printf("Failed to marshal index: %v\n", err)
			return
		}

		rw.Write(b)
	}))

	http.Handle("/streams/v1/com.ubuntu.cloud-released-imagemetadata.json", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ubuntuImages, err := s.getUbuntuImages()

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			log.Printf("Failed to retrieve ubuntu images: %v\n", err)
			return
		}

		metadata := s.makeMetadataFromImages(ubuntuImages)
		b, err := json.Marshal(metadata)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			log.Printf("Failed to marshal ubuntu images metadata: %v\n", err)
			return
		}

		rw.Write(b)
	}))

	log.Printf("Listening on %s\n", listen)
	if err := http.ListenAndServe(listen, nil); err != nil {
		panic(err)
	}
}
