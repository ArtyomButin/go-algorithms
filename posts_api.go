package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const postUrl = "https://jsonplaceholder.typicode.com/posts"
type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func getPosts()[]Post {
	resp, err := http.Get(postUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		fmt.Println(err)
	}
	return posts
}
