package main

import (
	"fmt"
	"sync"
	//"io/ioutil"
)

//go get golang.org/x/net/html

func main() {
	var wg sync.WaitGroup
	fmt.Print("Crawler")
	manager := URLManager{remain: make(chan string, 3000), visited: make(map[string]bool), wg: wg, entry: "https://www.wikipedia.org/"}
	manager.AddListener(5)
	manager.Start()
}
