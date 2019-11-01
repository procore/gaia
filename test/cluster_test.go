package test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/procore/gaia"
)

func TestClusterAllocationExplain(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"error":{"root_cause":[{"type":"index_not_found_exception","reason":"no such index","resource.type":"index_expression","resource.id":"_cluster","index_uuid":"_na_","index":"_cluster"}]}`)
	}))
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	config := gaia.NewConfig()
	fmt.Println(ts.URL)
	config.Net.Host = u.Hostname()
	config.Net.Port = u.Port()
	client := gaia.NewClient(config)
	resp := client.AllocationExplain()
	if resp == "" {
		t.Errorf("_cluster/allocation/explain returned no response")
	}
}
