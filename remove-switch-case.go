package main

// Removing the switch like below statement using a dictionary
// switch {

// 	// @bot search me HMAC
// 	case strings.Contains(message, "search me"):
// 		query := strings.Split(message, "search me ")[1]
// 		return webSearch(query), "html"

// 	// @bot thesaurus me challenge
// 	case strings.Contains(message, "thesaurus me"):
// 		query := strings.Split(message, "thesaurus me ")[1]
// 		return synonyms(query), "html"

// Solution: Use map[string]command similar to how the net/http package registers handlers.

import (
	"errors"
	"fmt"
)

type BotFunc func(string) (string, error)

type BotMap map[string]BotFunc

var Bot = BotMap{}

func (b BotMap) RegisterCommand(command string, f BotFunc) error {
	if _, exists := b[command]; exists {
		return errors.New("command already exists")
	}
	b[command] = f
	return nil
}

func (b BotMap) Execute(statement string) (string, error) {
	// Add the parsing logic, this is just an example
	command := statement[:9]
	query := statement[10:]

	return b.ExecuteQuery(command, query)
}

func (b BotMap) ExecuteQuery(command, query string) (string, error) {
	if com, exists := b[command]; exists {
		return com(query)
	}
	return "", errors.New("command doesn't exist")

}

func runBot() {
	err := Bot.RegisterCommand("search me", func(query string) (string, error) {
		fmt.Println("search", query)
		return "searched", nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = Bot.RegisterCommand("thesaurus me", func(query string) (string, error) {
		fmt.Println("thesaurus", query)
		return "thesaurused", nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := Bot.Execute("search me please")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
