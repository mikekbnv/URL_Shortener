package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	//go:embed "index.html"

	f   embed.FS
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis:6379"),
		Password: os.Getenv("DB_PASS"),
		DB:       0,
	})
	ctx = context.Background()

	//urls = map[string]string{}
	//appConfig config.Config
)

func shorten(url string) string {
	var k string

	for {
		k = big.NewInt(int64(10000000 + rand.Intn(1000000))).Text(62)
		exist, err := rdb.Exists(ctx, k).Result()
		if err != nil {
			log.Fatal(err)
		}
		if exist != 1 {
			err = rdb.Set(ctx, k, url, 0).Err()
			if err != nil {
				log.Fatal(err)
			}
			break
		}
		// if _, ok := urls[k]; !ok {
		// 	urls[k] = url
		// 	break
		// }
	}
	return k
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			log.Printf("ERROR: %v", err)
			fmt.Fprintf(w, "error: %v", err)
			return
		}

		url := strings.TrimSpace(r.PostForm.Get("url"))
		if url == "" {
			http.Error(w, "Invalid URL", http.StatusNotAcceptable)
			return
		}

		log.Printf("INFO: shortening %s", url)
		fmt.Fprint(w, "http://"+r.Host+"/"+shorten(url))
	default:
		if r.URL.Path == "/" {
			index, err := f.ReadFile("index.html")
			if err != nil {
				log.Println(err)
			}
			fmt.Fprint(w, string(index))
			return
		}

		k := strings.Trim(r.URL.Path, "/")
		val, err := rdb.Get(ctx, k).Result()
		if err != nil {
			if err != redis.Nil {
				log.Fatal(err)
			} else {
				log.Printf("ERROR: %s not found", k)
				http.Error(w, "Not Found", http.StatusNotFound)
			}
		}
		//if url, ok := urls[k]; ok {
		log.Printf("INFO: redirecting %s -> %s\n", k, val)
		http.Redirect(w, r, val, http.StatusPermanentRedirect)
		return
		//}

		// log.Printf("ERROR: %s not found", k)
		// http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func main() {
	_, err := rdb.Ping(ctx).Result()
	fmt.Println("test")
	if err != nil {
		log.Fatal(err)
	}
	
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/", http.HandlerFunc(handler))
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), mux))
}
