package main

import (
	"encoding/json"
	"fmt"
	"log"

)

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
    Timestamp string `json:"timestamp"`
}

type query struct {
	Recentchanges []change `json:"recentchanges"`
    //Other string `json:"other"`
}

type answer struct {
    Query query `json:"query"`
}


func main() {

     text := `{"batchcomplete":"","continue":{"rccontinue":"20191108103855|1203480232","continue":"-||"},"query":{"recentchanges":[{"type":"edit","ns":0,"title":"Philippe Starck","pageid":156436,"revid":925180693,"old_revid":925180580,"rcid":1203480282,"timestamp":"2019-11-08T10:39:10Z"},{"type":"categorize","ns":14,"title":"Category:Wikipedia books (user books)","pageid":62285509,"revid":925180691,"old_revid":0,"rcid":1203480281,"timestamp":"2019-11-08T10:39:09Z"},{"type":"edit","ns":0,"title":"Adamsville, Alabama","pageid":104761,"revid":925180692,"old_revid":923888589,"rcid":1203480280,"timestamp":"2019-11-08T10:39:09Z"},{"type":"new","ns":2,"title":"User:Xpicto/Books/Hierarchical clustering in bioinformatics","pageid":62285509,"revid":925180691,"old_revid":0,"rcid":1203480279,"timestamp":"2019-11-08T10:39:09Z"},{"type":"edit","ns":0,"title":"List of Indian people by net worth","pageid":61238433,"revid":925180689,"old_revid":925180464,"rcid":1203480277,"timestamp":"2019-11-08T10:39:09Z"},{"type":"edit","ns":109,"title":"Book talk:Lebanon","pageid":26973325,"revid":925180690,"old_revid":922632856,"rcid":1203480278,"timestamp":"2019-11-08T10:39:08Z"},{"type":"edit","ns":0,"title":"List of Riverdale episodes","pageid":53276188,"revid":925180688,"old_revid":925114639,"rcid":1203480273,"timestamp":"2019-11-08T10:39:08Z"},{"type":"edit","ns":0,"title":"Muammar Gaddafi","pageid":53029,"revid":925180694,"old_revid":924473200,"rcid":1203480283,"timestamp":"2019-11-08T10:39:07Z"},{"type":"edit","ns":0,"title":"Fakenham","pageid":710113,"revid":925180687,"old_revid":907104023,"rcid":1203480271,"timestamp":"2019-11-08T10:39:07Z"},{"type":"edit","ns":0,"title":"Trace3","pageid":50722791,"revid":925180685,"old_revid":906133645,"rcid":1203480266,"timestamp":"2019-11-08T10:39:07Z"},{"type":"edit","ns":4,"title":"Wikipedia:Articles for deletion/Dr. Nick","pageid":62285208,"revid":925180683,"old_revid":925174677,"rcid":1203480260,"timestamp":"2019-11-08T10:39:06Z"},{"type":"edit","ns":0,"title":"Chechnya","pageid":6095,"revid":925180686,"old_revid":925180591,"rcid":1203480270,"timestamp":"2019-11-08T10:39:05Z"},{"type":"categorize","ns":14,"title":"Category:CS1 errors: missing periodical","pageid":4778196,"revid":925180681,"old_revid":917830032,"rcid":1203480263,"timestamp":"2019-11-08T10:39:05Z"},{"type":"categorize","ns":14,"title":"Category:Pages with citations using unnamed parameters","pageid":4778196,"revid":925180681,"old_revid":917830032,"rcid":1203480262,"timestamp":"2019-11-08T10:39:05Z"},{"type":"log","ns":2,"title":"User:Tipturhassan","pageid":0,"revid":0,"old_revid":0,"rcid":1203480258,"timestamp":"2019-11-08T10:39:05Z"},{"type":"edit","ns":0,"title":"Mental chronometry","pageid":4778196,"revid":925180681,"old_revid":917830032,"rcid":1203480256,"timestamp":"2019-11-08T10:39:05Z"},{"type":"edit","ns":0,"title":"Nalanda","pageid":200282,"revid":925180684,"old_revid":925180387,"rcid":1203480264,"timestamp":"2019-11-08T10:39:04Z"},{"type":"log","ns":2,"title":"User:Naran Chaudhary","pageid":0,"revid":0,"old_revid":0,"rcid":1203480255,"timestamp":"2019-11-08T10:39:04Z"},{"type":"edit","ns":0,"title":"Thailand national under-20 football team","pageid":18196044,"revid":925180679,"old_revid":925166656,"rcid":1203480252,"timestamp":"2019-11-08T10:39:03Z"},{"type":"edit","ns":0,"title":"Headroom (photographic framing)","pageid":23381809,"revid":925180678,"old_revid":919764248,"rcid":1203480248,"timestamp":"2019-11-08T10:39:03Z"},{"type":"edit","ns":0,"title":"Diego Costa","pageid":10157131,"revid":925180682,"old_revid":923419442,"rcid":1203480257,"timestamp":"2019-11-08T10:39:02Z"},{"type":"edit","ns":0,"title":"Abraham de Moivre","pageid":57327,"revid":925180680,"old_revid":917975826,"rcid":1203480254,"timestamp":"2019-11-08T10:39:02Z"},{"type":"edit","ns":0,"title":"2019 Liga 1","pageid":59160739,"revid":925180677,"old_revid":925180625,"rcid":1203480247,"timestamp":"2019-11-08T10:39:02Z"},{"type":"edit","ns":0,"title":"PAOK FC","pageid":887670,"revid":925180676,"old_revid":925172379,"rcid":1203480246,"timestamp":"2019-11-08T10:39:02Z"},{"type":"log","ns":1,"title":"Talk:Black Legend/talk","pageid":55250000,"revid":925180674,"old_revid":0,"rcid":1203480275,"timestamp":"2019-11-08T10:39:01Z"},{"type":"log","ns":1,"title":"Talk:Black Legend/Archive index","pageid":32761929,"revid":925180672,"old_revid":0,"rcid":1203480274,"timestamp":"2019-11-08T10:39:01Z"},{"type":"log","ns":1,"title":"Talk:Black Legend/Archive 3","pageid":56507906,"revid":925180670,"old_revid":0,"rcid":1203480272,"timestamp":"2019-11-08T10:39:01Z"},{"type":"log","ns":1,"title":"Talk:Black Legend/Archive 2011","pageid":32760708,"revid":925180669,"old_revid":0,"rcid":1203480269,"timestamp":"2019-11-08T10:39:01Z"},{"type":"categorize","ns":14,"title":"Category:Articles needing translation from Finnish Wikipedia","pageid":47160787,"revid":925180673,"old_revid":758595111,"rcid":1203480251,"timestamp":"2019-11-08T10:39:01Z"},{"type":"categorize","ns":14,"title":"Category:All articles to be expanded","pageid":47160787,"revid":925180673,"old_revid":758595111,"rcid":1203480250,"timestamp":"2019-11-08T10:39:01Z"},{"type":"categorize","ns":14,"title":"Category:Articles to be expanded","pageid":47160787,"revid":925180673,"old_revid":758595111,"rcid":1203480249,"timestamp":"2019-11-08T10:39:01Z"},{"type":"edit","ns":0,"title":"2015\u20132016 SWIFT banking hack","pageid":50660269,"revid":925180675,"old_revid":919423602,"rcid":1203480245,"timestamp":"2019-11-08T10:39:01Z"},{"type":"edit","ns":0,"title":"Haarajoki","pageid":47160787,"revid":925180673,"old_revid":758595111,"rcid":1203480244,"timestamp":"2019-11-08T10:39:01Z"},{"type":"edit","ns":0,"title":"Melksham Oak Community School","pageid":30782962,"revid":925180671,"old_revid":915294017,"rcid":1203480243,"timestamp":"2019-11-08T10:39:01Z"},{"type":"log","ns":1,"title":"Talk:Black Legend/Archive 2010","pageid":32761898,"revid":925180668,"old_revid":0,"rcid":1203480268,"timestamp":"2019-11-08T10:39:00Z"},{"type":"log","ns":1,"title":"Talk:Black Legend/Archive 2009","pageid":32760704,"revid":925180667,"old_revid":0,"rcid":1203480267,"timestamp":"2019-11-08T10:39:00Z"},{"type":"log","ns":1,"title":"Talk:Black Legend/Archive 2","pageid":32760707,"revid":925180666,"old_revid":0,"rcid":1203480265,"timestamp":"2019-11-08T10:39:00Z"},{"type":"log","ns":1,"title":"Talk:Black Legend/Archive 1","pageid":26049623,"revid":925180665,"old_revid":0,"rcid":1203480261,"timestamp":"2019-11-08T10:39:00Z"},{"type":"log","ns":1,"title":"Talk:Black Legend","pageid":186868,"revid":925180660,"old_revid":0,"rcid":1203480259,"timestamp":"2019-11-08T10:38:59Z"},{"type":"categorize","ns":14,"title":"Category:Indian symbols by state or union territory","pageid":37161824,"revid":925180663,"old_revid":924231554,"rcid":1203480242,"timestamp":"2019-11-08T10:38:58Z"},{"type":"categorize","ns":14,"title":"Category:Lists of Indian state symbols","pageid":37161824,"revid":925180663,"old_revid":924231554,"rcid":1203480241,"timestamp":"2019-11-08T10:38:58Z"},{"type":"edit","ns":1,"title":"Talk:G\u00f6del's incompleteness theorems/History","pageid":27898236,"revid":925180664,"old_revid":789382733,"rcid":1203480240,"timestamp":"2019-11-08T10:38:58Z"},{"type":"edit","ns":0,"title":"List of Tamil Nadu state symbols","pageid":37161824,"revid":925180663,"old_revid":924231554,"rcid":1203480239,"timestamp":"2019-11-08T10:38:58Z"},{"type":"edit","ns":2,"title":"User:SSSB/common.js","pageid":60468388,"revid":925180662,"old_revid":925180582,"rcid":1203480238,"timestamp":"2019-11-08T10:38:58Z"},{"type":"log","ns":0,"title":"Black Legend","pageid":186824,"revid":925180654,"old_revid":0,"rcid":1203480253,"timestamp":"2019-11-08T10:38:56Z"},{"type":"edit","ns":0,"title":"Dendi Santoso","pageid":29506928,"revid":925180661,"old_revid":925126929,"rcid":1203480236,"timestamp":"2019-11-08T10:38:56Z"},{"type":"edit","ns":0,"title":"Zemra","pageid":50652428,"revid":925180658,"old_revid":850321909,"rcid":1203480235,"timestamp":"2019-11-08T10:38:56Z"},{"type":"edit","ns":3,"title":"User talk:Noxchii729393929","pageid":62285508,"revid":925180657,"old_revid":925180614,"rcid":1203480234,"timestamp":"2019-11-08T10:38:56Z"},{"type":"log","ns":0,"title":"Black legend","pageid":189116,"revid":0,"old_revid":0,"rcid":1203480276,"timestamp":"2019-11-08T10:38:55Z"},{"type":"edit","ns":0,"title":"Special Forces of India","pageid":262705,"revid":925180659,"old_revid":925180281,"rcid":1203480237,"timestamp":"2019-11-08T10:38:55Z"}]}}`
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

    // url := "https://en.wikipedia.org/w/api.php?action=query&list=recentchanges&rclimit=50&format=json"

	// spaceClient := http.Client{
	// 	Timeout: time.Second * 2, // Maximum of 2 secs
	// }

	// req, err := http.NewRequest(http.MethodGet, url, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// //req.Header.Set("User-Agent", "spacecount-tutorial")

	// res, getErr := spaceClient.Do(req)
	// if getErr != nil {
	// 	log.Fatal(getErr)
	// }

	// body, readErr := ioutil.ReadAll(res.Body)
	// if readErr != nil {
	// 	log.Fatal(readErr)
	// }

	body := []byte(text)

	people1 := query{}
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(people1)
}