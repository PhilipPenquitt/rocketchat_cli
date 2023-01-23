package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"rocketchat_cli/interactive"

	"github.com/badkaktus/gorocket"
)

var username string
var password string
var url string
var channel string
var text string
var prompt bool
var help bool

func init() {
	flag.StringVar(&username, "user", "", "Name of your User")
	flag.StringVar(&password, "password", "", "Your Password")
	flag.StringVar(&channel, "channel", "general", "Name of the channel to post to")
	flag.StringVar(&url, "url", "rocket.chat", "Url of your Rocketchat Instanz")
	flag.StringVar(&text, "text", "Lorem Ipsum", "The Text you want to post")
	flag.BoolVar(&prompt, "prompt", false, "Interactive Input")
	flag.BoolVar(&help, "help", false, "help")
	flag.Parse()
}

func main() {
	if help {
		flag.Usage()
		os.Exit(0)
	} else if prompt {
		lg, mg, url := interactive.Getdata()
		fmt.Println(lg, mg, url)
		os.Exit(0)
	} else if !prompt && (username == "" || password == "") {
		flag.Usage()
		os.Exit(1)
	}

	client := gorocket.NewClient(url)
	login := gorocket.LoginPayload{
		User:     username,
		Password: password,
	}

	lg, _ := client.Login(&login)
	// error Wert wird nicht korrekt zurückgegeben, daher wird Status geprüft

	if lg.Status != "success" {
		log.Fatal("Login Fehlgeschlagen mit folgenden Fehler: \n", lg.Message)
		os.Exit(1)
	}

	// post a message
	str := gorocket.Message{
		Channel: channel,
		Text:    text,
	}

	msg, err := client.PostMessage(&str)
	if err != nil {
		fmt.Printf("Error: %+v", err)
	}
	fmt.Printf("Message was posted %t", msg.Success)
}
