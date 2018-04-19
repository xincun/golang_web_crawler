# golang_web_crawler

The idea is dead simple. 

## URLManager
This is the core part of the crawler, it manages the urls and fetch new nodes.

* mux: mutex for concurrency
* remain: Urls to be visited
* visited: visited urls
* wg: WaitGroup to manage go routines
* entry: Entry point

The idea is to create a channel, add goroutines and listen. When new urls discovered, send to channel.
