package main

import "github.com/Mayhul-Jindal/cleanX/esp32/emulate-esp/pkg"

func main() {
	//* yaha aaj ka
	pkg.Semantic()
	//TODO: testing pipelines

	//* testing go-routines
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
	// -------------------------------------------------------------------------------------
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

	//* copied pub sub code is working
	// pkg.PubSubClientTest();
}
