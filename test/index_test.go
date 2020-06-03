package test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/procore/gaia"
)

func TestIndexAliasGet(t *testing.T) {
	responseBody := `{"contacts_2_1_0":{"aliases":{"contacts":{}}}}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, responseBody)
	}))
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	config := gaia.NewConfig()
	config.Net.Host = u.Hostname()
	config.Net.Port = u.Port()
	client := gaia.NewClient(config)
	resp := client.IndexAliasesGet("contacts")
	if strings.TrimRight(resp, "\n") != responseBody {
		t.Errorf("index alias get returned incorrect response")
	}
}
