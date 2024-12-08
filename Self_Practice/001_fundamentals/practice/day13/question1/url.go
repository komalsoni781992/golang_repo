package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

/*
q1. Create a slice with 3 random urls

	Create a function doGetRequest()
	doGetRequest:
	    It spins up 2 goroutines
	    1st goroutines do get request and put the url,resp,err inside one single channel
	    //1st goroutine spins up n number of goroutines for n number of urls (fanout pattern)
	    2nd goroutines fetch the values from the channel and perform following operations
	        -check err
	        -read body
	        -check if status code above 299
	        -and print url resp.Status
*/
type URLInfo struct {
	url  string
	resp *http.Response
	err  error
}

func main() {
	urls := []string{"https://www.google.com", "https://www.geeksforgeeks.org/", "https://leetcode.com/"}
	wg := new(sync.WaitGroup)
	c := make(chan URLInfo)
	for _, u := range urls {
		doGetRequest(u, wg, &c)
	}
	wg.Wait()
}

func doGetRequest(url string, wg *sync.WaitGroup, c *chan URLInfo) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Storing: ", url)
		res, err := http.Get(url)
		if err != nil {
			log.Println(err)
			return
		}
		*c <- URLInfo{url, res, err}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Fetching: ", url)
		u := <-*c
		if u.err != nil {
			log.Println(u.err)
			return
		}
		defer u.resp.Body.Close()
		bytes, err := io.ReadAll(u.resp.Body)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(bytes))
		if u.resp.StatusCode > 299 {
			fmt.Println("Status code is laarge: ", u.resp.StatusCode)
		}
		fmt.Println("Status is: ", url, u.resp.Status)
	}()
}
