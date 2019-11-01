package test

import (
	"testing"

	"github.com/procore/gaia"
)

func TestNewClient(t *testing.T) {
	config := gaia.NewConfig()
	client := gaia.NewClient(config)
	if client.Config.Net.Host != "localhost" {
		t.Errorf("client.Config.Net.Host == %s, expected localhost", client.Config.Net.Host)
	}
}

func TestNewClientWithConfig(t *testing.T) {
	config := gaia.NewConfig()
	config.User.Name = "elastic"
	config.User.Password = "psswrd"
	config.Pretty = true
	config.Debug = true
	config.Net.Host = "127.0.0.1"
	client := gaia.NewClient(config)
	if client.Config.Net.Host != "127.0.0.1" {
		t.Errorf("client.Config.Net.Host == %s, expected 127.0.0.1", client.Config.Net.Host)
	}

	if client.Config.User.Name != "elastic" || client.Config.User.Password != "psswrd" {
		t.Errorf("client.Config.User.Name is incorrect")
	}

	if !client.Config.Debug {
		t.Errorf("client.Config.Debug == %t, expected true", client.Config.Debug)
	}
}
