package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var err error
	var ipv4 []byte
	var ipv6 []byte

	if ipv4, err = getIPv4(); err != nil {
		fmt.Printf("Error getting IPv4: %s\n", err.Error())
		os.Exit(-1)
	}

	if ipv6, err = getIPv6(); err != nil {
		fmt.Printf("Error getting IPv6: %s\n", err.Error())
		os.Exit(-2)
	}

	fmt.Printf("\n")
	fmt.Printf("IPv4: %s\n", ipv4)
	fmt.Printf("IPv6: %s\n", ipv6)
	fmt.Printf("\n")
}

func getIPv4() (result []byte, err error) {
	url := "https://api.ipify.org?format=text"
	return get(url)
}

func getIPv6() (result []byte, err error) {
	url := "https://api6.ipify.org?format=text"
	return get(url)
}

func get(url string) (result []byte, err error) {
	var response *http.Response

	if response, err = http.Get(url); err != nil {
		return
	}

	defer response.Body.Close()
	result, err = ioutil.ReadAll(response.Body)
	return
}
