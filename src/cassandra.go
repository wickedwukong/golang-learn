package main

import (
	"fmt"
	"log"
	"tux21b.org/v1/gocql"
	"math/rand"
	"time"
)

type User struct {
	Id		gocql.UUID
    Name    string
}

type CRUD func(cassandra *gocql.Session)

func insertUser(name string)(CRUD) {
	return func(cassandra *gocql.Session) {
		if err := cassandra.Query(`INSERT INTO users (id, name) VALUES (?, ?)`, gocql.TimeUUID(), name).Exec(); err != nil {
	        log.Fatal(err)
	    }   
    }     
}

var printUsers = func(cassandra *gocql.Session) {
		allUsers := cassandra.Query(`select id, name from users`).Iter()
		user := User{}
		for allUsers.Scan(&user.Id, &user.Name) {
			fmt.Printf("User: %v\n", user)
		}
	}


type WithSession func(crud CRUD)

func createSession()(WithSession) {
	cluster := gocql.NewCluster("172.17.0.187")	
	cluster.Keyspace = "golang"
	cluster.Consistency = gocql.Quorum

	return func(crud CRUD) {
		cassandra, _ := cluster.CreateSession()
		crud(cassandra)
		defer cassandra.Close();
	}
}

func randName(n int8)(string) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
    random := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
    // rand.Seed( time.Now().UTC().UnixNano())
	for i := range b {
		b[i] = letters[random.Intn(len(letters))]
	}
	return string(b)
}

func main() {
    withSession := createSession()

    // fmt.Printf("%T", withSession)

    withSession(insertUser(randName(7)))

    withSession(printUsers)
}