package interactive

import (
	"bufio"
	"fmt"
	"github.com/badkaktus/gorocket"
	"log"
	"os"
	"strings"
)

func interactive(s string) (input string) {
	fmt.Printf("%v:", s)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		log.Fatal("Error reading input", err)
	}
	return input
}

func Getdata() (login gorocket.LoginPayload, message gorocket.Message, url string) {
	
	username := interactive("username")
	password := interactive("password")
	url = interactive("url")
	channel := interactive("channel")
	text := interactive("text")

	login = gorocket.LoginPayload{
		User:     username,
		Password: password,
	}

message = gorocket.Message{
		Channel: channel,
		Text:    text,
	}

	return login, message, url

}
