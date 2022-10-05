package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/lenra-io/counter/internal/counter/requests"
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
	var input []byte

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Bytes()
	}

	output, marshal := requests.HandleRootRequest(input)

	if marshal {
		json_output, err := json.Marshal(output)
		if err != nil {
			logrus.Errorf("Internal response return is malformed: %v", output)
		}
		fmt.Printf("%s", string(json_output))
	} else {
		fmt.Printf("%s", output)
	}
}
