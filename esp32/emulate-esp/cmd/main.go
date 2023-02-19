package main

import (
	"fmt"
	"math/rand"

	"github.com/Mayhul-Jindal/cleanX/esp32/emulate-esp/pkg"
)

// "fmt"
// "log"
// "net/http"
// "time"
// "github.com/Mayhul-Jindal/cleanX/esp32/emulate-esp/pkg"

func main(){
	// testing buffered channel
	bufferedCh := make(chan string, 3) // this actual works in a asynchronous manner

	for i := 0; i < 3; i++ {
		go pkg.BufferedChanelTest(fmt.Sprint(rand.Intn(100)) , bufferedCh)
	}

	close(bufferedCh)

	for i := 0; i < 3; i++ {
		r := <- bufferedCh
		fmt.Println(r)
	}

	// testing go-routines
	// startP := time.Now()
	// list := []string{"https://amazon.com", "https://google.com", "https://facebook.com"}
	// for _, url := range list{
	// 	start := time.Now()
	// 	if resp, err := http.Get(url); err != nil{
	// 		log.Fatal(err)
	// 	}else{
	// 		fmt.Println(time.Since(start).Round(time.Millisecond))
	// 		resp.Body.Close()
	// 	}
	// }
	// fmt.Printf("\nnormal %s\n\n", time.Since(startP).Round(time.Millisecond))

	// // -------

	// startP = time.Now()
	// results := make(chan time.Duration)
	// list = []string{"https://amazon.com", "https://google.com", "https://facebook.com"}
	// for _, url := range list{
	// 	go pkg.ParallelGet(url, results)
	// }
	// for range list{
	// 	r := <- results
	// 	if r == 0{
	// 		log.Println(0)
	// 	}else{
	// 		log.Println(r)
	// 	}
	// }
	// fmt.Printf("goroutines %s",time.Since(startP).Round(time.Millisecond))

	// copied pub sub code is working
	// pkg.PubSubClientTest();	
}