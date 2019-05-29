package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func timer() func() time.Duration {
	start := time.Now()
	return func() time.Duration {
		return time.Now().Sub(start)
	}
}

func callAPI(d int) {
	t := timer()
	resp, err := http.Get("http://softwarekonjon.netlify.com")
	if err != nil {
		log.Println(err)
		return
	}
	p, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	log.Printf("%d: %s in %dms\n", d, p, t().Nanoseconds()/1000000)
}

func main() {
	t := timer()
	wg := &sync.WaitGroup{}
	semaphore := make(chan struct{}, 50)
	for d := 0; d < 10000; d++ {
		d := d
		semaphore <- struct{}{}
		wg.Add(1)
		go func() {
			callAPI(d)
			wg.Done()
			<-semaphore
		}()
	}
	wg.Wait()
	log.Printf("All Done in %.0fs", t().Seconds())
}