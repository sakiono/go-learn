package main

import(
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"net/url"
)

func main(){
	resp, err := http.Get("http://localhost:8080?p=Gopher")

	if err != nil{
		log.Fatal(err)
	}

	defer resp.Body.Close()

	/*
	var p Person

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&p); err != nil{
		log.Fatal(err)
	}
	fmt.Println(p)

	 */

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

/*
//疑問
json使ってdecodeの場合は、どうやるんだろう？
 */