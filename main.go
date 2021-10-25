package main

type userPost struct {
	User User
	Post []Post
}

func main() {
	users := getUsers()
	posts := getPosts()
	var usersPosts []userPost
	var userPosts []Post
	for _, user := range users {
		for _, post := range posts {
			if user.Id == post.UserId {
				userPosts = append(userPosts, post)
			}
		}
		usersPosts = append(usersPosts, userPost{User: user, Post: userPosts})
		userPosts = []Post{}
	}
}
