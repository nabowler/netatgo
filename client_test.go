package netatgo_test

import (
	"context"
	"os"
	"testing"

	"github.com/nabowler/netatgo"
)

func mustCCClient(t *testing.T, scopes []netatgo.Scope) netatgo.Client {
	if testing.Short() {
		t.Skip("skipping because the short flag is set")
	}

	cfg := netatgo.ClientCredentialsConfig{
		ClientID:     os.Getenv("NETATGO_TEST_CLIENT_ID"),
		ClientSecret: os.Getenv("NETATGO_TEST_CLIENT_SECRET"),
		Username:     os.Getenv("NETATGO_TEST_USERNAME"),
		Password:     os.Getenv("NETATGO_TEST_PASSWORD"),
		Scopes:       scopes,
	}

	if cfg.ClientID == "" || cfg.ClientSecret == "" || cfg.Password == "" || cfg.Username == "" {
		t.Fatal("necessary config is missing")
	}

	return netatgo.NewClientCredentialsClient(cfg)
}

func TestReadStation(t *testing.T) {
	for _, tc := range []struct {
		name   string
		scopes []netatgo.Scope
	}{
		{name: "default scope"},
		{name: "explicit scope", scopes: []netatgo.Scope{netatgo.ReadStation}},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			client := mustCCClient(t, tc.scopes)

			resp, err := client.GetStationData(context.Background(), "", false)
			if err != nil {
				t.Error(err)
			}
			t.Logf("Status: %s  Response: %#v", resp.Status, resp)
		})
	}
}
