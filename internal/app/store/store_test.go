package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=ec2-99-80-170-190.eu-west-1.compute.amazonaws.com dbname=deve2jnoij3f0f user=ntbigkedlteqiw password=e5e000375a5a4d59e84d1c47c69841f927a728af24c3b2ff726687076e372d5b"
	}

	os.Exit(m.Run())
}
