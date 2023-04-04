package main

import (
	"bufio"
	"log"

	// "context"
	"fmt"
	"os"
	// openai "github.com/sashabaranov/go-openai"
)

func main() {
	// Load config from file or env
	err := loadConfig()
	if err != nil {
		panic(err)
	}
	log.Printf("Finished loading custom configs.")

	// start a demo on console to handle user input and loop
	fmt.Print("Input your prompt >")
	reader := bufio.NewReader(os.Stdin)
	userText, _ := reader.ReadString('\n')
	userText = userText[:len(userText)-1]

	log.Printf("The user input text handle to be:\n%s", userText)

	// client := openai.NewClientWithConfig(ClientCfg)
	// resp, err := client.CreateChatCompletion(
	// 	context.Background(),
	// 	openai.ChatCompletionRequest{
	// 		Model: openai.GPT3Dot5Turbo,
	// 		Messages: []openai.ChatCompletionMessage{
	// 			{
	// 				Role:    openai.ChatMessageRoleUser,
	// 				Content: "Hello!",
	// 			},
	// 		},
	// 	},
	// )

	// if err != nil {
	// 	fmt.Printf("ChatCompletion error: %v\n", err)
	// 	return
	// }

	// fmt.Println(resp.Choices[0].Message.Content)

}
