package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func fortune(c chan string) {
	fileasBytes, err := ioutil.ReadFile("Fortunes.txt")
	if err != nil {
		log.Fatalln("Error reading file: ", err)

	}
	fortunes := strings.Split(string(fileasBytes), "%%")

	for {
		_ = <-c
		rand.Seed(time.Now().Unix())
		randomNum := rand.Intn(len(fortunes))
		fortune := fortunes[randomNum]
		fmt.Println(fortune)
		fmt.Println("Would you like another fortune?: YES or NO")
	}
}

func main() {
	var str string
	ch := make(chan string)
	go fortune(ch)
	fmt.Println("Would you like a fortune?: YES or NO")

	for {
		_, err := fmt.Scan(&str)
		if err != nil {
			log.Fatalln(err)
		}
		if strings.ToLower(str) == "yes" {
			ch <- ""

		} else if strings.ToLower(str) == "no" {
			os.Exit(-1)
		} else {
			fmt.Println("Would you like a fortune?: YES or NO")
			continue
		}
	}

}
