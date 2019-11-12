package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"
	"io/ioutil"
	"time"
)

//JSON structures 
type change struct {
	//Id string `json:"id"`
	//Body string `json:"body"`
	Type string `json:"type"`
	Ns int `json:"ns"`
	Title string `json:"title"`
	Pageid int `json:"pageid"`
	Revid int `json:"revid"`
	Oldrevid int `json:"old_revid"`
	Rcid int `json:"rcid"`
	User string `json:"user"`
	Userid int `json:"userid"`
	Oldlen int `json:"oldlen`
	Newlen int `json:"newlen`
	Date string `json:"timestamp"`
	Comment string `json:"comment"`

	//flags
		Bot *string `json:"bot,omitempty"`
		Minor *string `json:"minor,omitempty`
		Anon *string `json:"anon,omitempty`
		New *string `json:"new,omitempty`
}

type query struct {
	Recentchanges []change `json:"recentchanges"`
    //Other string `json:"other"`
}

type answer struct {
    Query query `json:"query"`
}

func printRecentChanges(changes []change) {
	for _, element := range changes {
		fmt.Println("------------------------------------------")
		fmt.Println("Title :\t\t", element.Title)
		fmt.Println("User :\t\t", element.User)
		fmt.Println("Type :\t\t", element.Type)
		fmt.Println("Namespace :\t", element.Ns)
		fmt.Println("Timestamp :\t", element.Date)
		fmt.Println("RevID :\t\t", element.Revid)
		fmt.Println("Flags :")
		fmt.Println("\tBot :\t\t", (element.Bot != nil))
		fmt.Println("\tMinor edit :\t", (element.Minor != nil))
		fmt.Println("\tAnonymous :\t", (element.Anon != nil))
		fmt.Println("\tNew entry :\t", (element.New != nil))
	}
}

func main() {

    //  text := `{"batchcomplete":"","continue":{"rccontinue":"20191108103855|1203480232","continue":"-||"},"query":{"recentchanges":[{"type":"edit","ns":0,"title":"Philippe Starck","pageid":156436,"revid":925180693,"old_revid":925180580,"rcid":1203480282,"timestamp":"2019-11-08T10:39:10Z"},{"type":"categorize","ns":14,"title":"Category:Wikipedia books (user books)","pageid":62285509,"revid":925180691,"old_revid":0,"rcid":1203480281,"timestamp":"2019-11-08T10:39:09Z"}]}}`
    // //var result map[string]interface{}

    // textBytes := []byte(text)

    // answer1 := answer{}
	// jsonErr1 := json.Unmarshal(textBytes, &answer1)
    // //jsonErr2 := json.Unmarshal(query1, &change1)
    // //query1 := result["query"].(map[string]interface{})
	// if jsonErr1 != nil  {
	// 	log.Fatal(jsonErr1)
	// }
	// //fmt.Println(textBytes)
    // for i := 0; i<len(answer1.Query.Recentchanges); i++ {
    //     fmt.Println(answer1.Query.Recentchanges[i].Timestamp)
    // }

    url := "https://en.wikipedia.org/w/api.php?action=query&list=recentchanges&rclimit=10&format=json&rcprop=user|userid|comment|flags|timestamp|title|ids|sizes"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	//req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// body := []byte(text)

	people1 := answer{}
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// printRecentChanges(people1.Query.Recentchanges)
	fmt.Printf("%+v\n", people1.Query.Recentchanges[0])
}