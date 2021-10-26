package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

const baseUrl = "https://jsonplaceholder.typicode.com/"

type userPost struct {
	User User
	Post []Post
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	start := time.Now()
	users := make([]User, 0)
	posts := make([]Post, 0)
	userPosts := make([]Post, 0)
	result := make([]userPost, 0)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := getResponse(userUrl, &users); err != nil {
			panic(err)
		}
	}()
	go func() {
		defer wg.Done()
		if err := getResponse(postUrl, &posts); err != nil {
			panic(err)
		}
	}()
	wg.Wait()
	for _, user := range users {
		for _, post := range posts {
			if user.Id == post.UserId {
				userPosts = append(userPosts, post)
			}
		}
		result = append(result, userPost{User: user, Post: userPosts})
		userPosts = []Post{}
	}
	duration := time.Since(start)
	fmt.Println(duration)
}

func getResponse(uri string, T interface{}) error {
	client := &http.Client{}
	r, err := client.Get(baseUrl + uri)
	if err != nil {
		return err
	}
	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, &T); err != nil {
		return err
	}
	return nil
}
