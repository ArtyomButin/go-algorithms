package main

import (
	"sync"
)

type userPost struct {
	User User
	Post []Post
}

func main() {
	users := getUsers()
	posts := getPosts()
	c := make(chan []Post)
	var wg sync.WaitGroup
	var result []userPost
	var userPosts []Post

	for _, user := range users {
		wg.Add(1)
		go func(userId int) {
			defer wg.Done()
			findPosts(userId, &posts, c)
		}(user.Id)
		userPosts = append(userPosts, <-c...)
		result = append(result, userPost{User: user, Post: userPosts})
		userPosts = []Post{}
	}
	wg.Wait()
	//jsonRes, _ := json.Marshal(result)
	//fmt.Println(string(jsonRes))
}

func findPosts(userId int, outerPosts *[]Post, c chan []Post) {
	var posts []Post
	for _, post := range *outerPosts {
		if userId == post.UserId {
			posts = append(posts, post)
		}
	}
	c <- posts
}
