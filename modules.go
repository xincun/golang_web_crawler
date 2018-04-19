package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type URLManager struct {
	mux     sync.Mutex
	remain  chan string
	visited map[string]bool
	wg      sync.WaitGroup
	entry   string
}

func (man *URLManager) html_traverse(n *html.Node) {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {

				man.AddURL(a.Val, 3)
				//fmt.Println(a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		man.html_traverse(c)
	}
}

func (man *URLManager) crawl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	} else {
		//body, _ := ioutil.ReadAll(resp.Body)

		doc, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Print(err)
		} else {
			man.html_traverse(doc)
		}

	}
}

func (m *URLManager) AddURL(url string, timeout int) {
	counter := 0
	fmt.Println("Add")
	if m.visited[url] == true {
		fmt.Println("Already visited!")
	} else {
		for {
			select {
			case m.remain <- url:
				fmt.Println("Url added: ", url)
				return
			default:
				time.Sleep(1 * time.Millisecond)
				counter += 1
				if counter > timeout {
					return
				}

			}

		}

	}

}

func (m *URLManager) Visit() {

	for {
		select {
		case url := <-m.remain:
			m.mux.Lock()
			fmt.Println("Visited: ", url)
			m.crawl(url)
			m.visited[url] = true
			m.mux.Unlock()
		default:
			fmt.Println("Channel empty, waiting for incoming urls")
			time.Sleep(50 * time.Millisecond)

		}
	}
	m.wg.Done()
}

func (m *URLManager) AddListener(number int) {
	m.wg.Add(number)
	for i := 0; i < number; i++ {
		go m.Visit()
	}

}

func (m *URLManager) Start() {
	m.AddURL(m.entry, 3)
	m.wg.Wait()
}
