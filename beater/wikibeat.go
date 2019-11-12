package beater

import (
	"fmt"
	"time"
	"log"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/aflamant/wikibeat/config"

	"net/http"
	"encoding/json"
	"io/ioutil"
)

// wikibeat configuration.
type wikibeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

// New creates an instance of wikibeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &wikibeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

//JSON structures 
type change struct {
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


// Run starts wikibeat.
func (bt *wikibeat) Run(b *beat.Beat) error {
	logp.Info("wikibeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)

	url := "https://en.wikipedia.org/w/api.php?action=query&list=recentchanges&rclimit=10&format=json&rcprop=user|userid|comment|flags|timestamp|title|ids|sizes"

	APIClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	
	req, err := http.NewRequest(http.MethodGet, url, nil)
		
    if err != nil {
	    log.Fatal(err)
    }
	
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		res, getErr := APIClient.Do(req)
			if getErr != nil {
			log.Fatal(getErr)
		}

		body, readErr := ioutil.ReadAll(res.Body)
			if readErr != nil {
			log.Fatal(readErr)
		}

		answer1 := answer{}
		json.Unmarshal(body, &answer1)

		// rc := answer1.Query.Recentchanges
		for _, rc := range answer1.Query.Recentchanges {
			// fmt.Printf("Object : %+v\n", rc)

			event := beat.Event{
				Timestamp: time.Now(),
				Fields: common.MapStr{
					"type":    		b.Info.Name,
					"edittype": 	rc.Type,
					"namespace": 	rc.Ns,
					"title": 		rc.Title,
					"pageid":		rc.Pageid,
					"revid": 		rc.Revid,
					"oldrevid": 	rc.Oldrevid,
					"rcid": 		rc.Rcid,
					"wikiuser": 	rc.User,
					"userid": 		rc.Userid,
					"oldlen": 		rc.Oldlen,
					"newlen": 		rc.Newlen,
					"date": 		rc.Date,
					"comment": 		rc.Comment,
					"bot": 			(rc.Bot != nil),
					"anonymous": 	(rc.Anon != nil),
					"new": 			(rc.New != nil),
					"minor": 		(rc.Minor != nil),
				},
			}
			bt.client.Publish(event)
			logp.Info("Event sent")
		}
	}
}

// Stop stops wikibeat.
func (bt *wikibeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
