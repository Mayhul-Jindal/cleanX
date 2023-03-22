package learningGO

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type meraObject struct {
	Transcript string `json:"transcript"`
	Img string		`json:"img"`
}

func HttpTest(num int){
	hn := fmt.Sprintf("https://xkcd.com/%d/info.0.json", num)

	resp, err := http.Get(hn)
	if err != nil{
		fmt.Fprintf(os.Stderr, "Cannot read %v", err);
		os.Exit(-1)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK{
		var mera meraObject
		body , err := ioutil.ReadAll(resp.Body);
		if err != nil{
			fmt.Fprintf(os.Stderr, "Cannot read %v", err);
			os.Exit(-1)
		}

		fmt.Println(string(body))

		err = json.Unmarshal(body, &mera)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot unmarshal JSON: %v", err)
			os.Exit(-1)
		}

		fmt.Println()
		fmt.Printf("%+v", mera)
	}

}