package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Read users from a text file
func parseUsersFile(p string) []user {
	users := []user{}
	contents, err := os.ReadFile(p)

	if err != nil {
		log.Fatal(err)
	}

	commaSplit := strings.Split(string(contents), ",")

	for i := 0; i < len(commaSplit)-1; i++ {
		colonSplit := strings.Split(commaSplit[i], ":")
		u := user{
			Id:    colonSplit[0],
			Name:  colonSplit[1],
			Email: colonSplit[2],
		}
		users = append(users, u)
	}

	return users
}

func main() {
	users := make([]user, 0)
	users = parseUsersFile("./users.txt")

	r := gin.Default()
	r.LoadHTMLGlob("static/*")

	//Display home page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//Get all users
	r.GET("/users", func(c *gin.Context) {
		if len(users) > 0 {
			c.JSON(http.StatusOK, gin.H{
				"users": users,
			})
		} else {
			c.JSON(http.StatusNotFound, "No users found")
		}
	})

	//Get a user by Id
	r.GET("/user", func(c *gin.Context) {
		id := c.Query("id")
		exists := false

		//Check if query param was given and display user or error
		if id != "" {
			for _, u := range users {
				if u.Id == id {
					exists = true
					c.JSON(http.StatusOK, u)
				}
			}

			if !exists {
				c.JSON(http.StatusNotFound, "User not found")
			}

		} else {
			c.JSON(http.StatusNotFound, "No user id specified")
		}
	})

	//Create a user - doesnt actually save to text file
	r.POST("/user", func(c *gin.Context) {
		name := c.Query("name")
		email := c.Query("email")
		id := strconv.Itoa(len(users) + 1)
		newUser := user{
			Id:    id,
			Name:  name,
			Email: email,
		}

		users = append(users, newUser)

		c.JSON(http.StatusOK, "User created")
	})

	//Update a user - doesnt save to text file and doesnt update id
	r.PUT("/user", func(c *gin.Context) {
		exists := false
		id := c.Query("id")
		name := c.Query("name")
		email := c.Query("email")

		for _, user := range users {
			if user.Id == id {
				exists = true
				index, err := strconv.Atoi(id)

				if err == nil {
					users[index-1].Name = name
					users[index-1].Email = email
					c.JSON(http.StatusOK, "User updated")
				}
			}
		}

		if !exists {
			c.JSON(http.StatusNotFound, "User not found")
		}
	})

	r.Run()
}
