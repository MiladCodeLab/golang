package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// first_name,last_name,email,birthdate
type User struct {
	Firstname string `csv:"first_name"`
	Lastname  string `csv:"last_name"`
	Email     string `csv:"email"`
	Birthdate string `csv:"birthdate"`
}

func main() {
	// define user channel
	userChan := make(chan User)
	usersAbove30 := int64(0)

	wg := sync.WaitGroup{}

	wg.Go(func() {
		// TODO: 1- read from multiple data sources
		// read csv file and log file line by line at same time
		// send parsed user data to userChan
		file, err := os.Open("storage/user.csv")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		reader := csv.NewReader(file)
		reader.Read() // remove the first line
		for line, err := reader.Read(); err != io.EOF; line, err = reader.Read() {
			var u User
			u.Firstname = line[0]
			u.Lastname = line[1]
			u.Email = line[2]
			u.Birthdate = line[3]
			userChan <- u
		}
		close(userChan)
	})

	wg.Go(func() {
		// TODO: 2- read users on userChan
		// calculate age
		// count if older than 30
		for user := range userChan {
			//fmt.Println(user)
			// calculate above 30 years old users
			//1997/10/25
			b := strings.ReplaceAll(user.Birthdate, "/", "-") + " 00:00:00"
			t, err := time.Parse(time.DateTime, b)
			if err != nil {
				panic(b)
			}
			if t.Before(time.Now().Add(-time.Hour * 24 * 365 * 30)) {
				fmt.Printf("above 30 years: %s %s %s %s\n", user.Firstname, user.Lastname, user.Email, user.Birthdate)
				atomic.AddInt64(&usersAbove30, 1)
			}
		}

	})

	// TODO: 3- fix deadlock issue if happened
	wg.Wait()
	fmt.Println("users above 30: ", usersAbove30)
}
