package main

import (
	"bufio"
	"context"
	"os"

	"github.com/lenra-io/counter/internal/counter"
	"github.com/lenra-io/counter/pkg/lenra"
	"github.com/sirupsen/logrus"
)

func init() {
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	// LOG_LEVEL not set, let's default to debug
	if !ok {
		lvl = "debug"
	}
	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.DebugLevel
	}
	logrus.SetLevel(ll)
}

func main() {
	var request []byte

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		request = scanner.Bytes()
	}

	manifest := &counter.Manifest{}
	lenra.Serve(context.Background(), manifest, request)
}
