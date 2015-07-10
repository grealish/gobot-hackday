package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	addresses := []string{}
	for _, res := range regexp.MustCompile(`Address\: (.+)`).FindAllStringSubmatch(string(bytes), -1) {
		addresses = append(addresses, res[1])
	}
	ssids := []string{}
	for _, res := range regexp.MustCompile(`ESSID\:\"(.+)\"`).FindAllStringSubmatch(string(bytes), -1) {
		ssids = append(ssids, res[1])
	}

	if len(addresses) != len(ssids) {
		log.Fatal("parse fail: addresses and ssids count mismatch")
	}

	for i := range addresses {
		fmt.Println(addresses[i])
		fmt.Println(ssids[i])
	}
}
