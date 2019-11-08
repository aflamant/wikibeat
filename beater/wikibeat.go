package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/aflamant/wikibeat/config"

	"net/http"
	"encoding/json"
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

// Run starts wikibeat.
func (bt *wikibeat) Run(b *beat.Beat) error {
	logp.Info("wikibeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	ticker := time.NewTicker(bt.config.Period)
	counter := 1

	url := "https://en.wikipedia.org/w/api.php?action=query&list=recentchanges&rclimit=50&format=json"

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

		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type":    b.Info.Name,
				"counter": counter,
			},
		}
		bt.client.Publish(event)
		logp.Info("Event sent")
		counter++
	}
}

// Stop stops wikibeat.
func (bt *wikibeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
