package test

import (
	"fmt"
	"testing"

	"github.com/procore/gaia"
)

func TestNewConfig(t *testing.T) {
	config := gaia.NewConfig()
	if config.Net.Host != "localhost" {
		t.Errorf("config.Net.Host = %s, expected 'localhost'", config.Net.Host)
	}
}

func TestNewConfigWithChanges(t *testing.T) {
	config := gaia.NewConfig()
	config.User.Name = "elastic"
	config.User.Password = "psswrd"
	config.Pretty = true
	config.Debug = true
	config.Net.Host = "127.0.0.1"
	if config.Net.Host != "127.0.0.1" {
		t.Errorf("config.Net.Host == %s, expected 127.0.0.1", config.Net.Host)
	}

	if config.User.Name != "elastic" || config.User.Password != "psswrd" {
		t.Errorf("config.User.Name is incorrect")
	}

	if !config.Debug {
		t.Errorf("config.Debug == %t, expected true", config.Debug)
	}
}

func ExampleNewConfig() {
	config := gaia.NewConfig()
	fmt.Println(config.Net.Host)
	fmt.Println(config.Net.Port)
	fmt.Println(config.Pretty)
	fmt.Println(config.Debug)
	// Output:
	// localhost
	// 9200
	// false
	// false
}
