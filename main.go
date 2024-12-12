package main

import (
	"fmt"
	"github.com/green-api/whatsapp-chatbot-golang"
	"os"
)

func main() {
	bot := whatsapp_chatbot_golang.NewBot(os.Getenv("GREEN_API_ID_INSTANCE"), os.Getenv("GREEN_API_TOKEN_INSTANCE"))

	file, err := os.OpenFile("phone_number.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	bot.IncomingMessageHandler(func(message *whatsapp_chatbot_golang.Notification) {
		phoneNumber, err := message.Sender()
		if err != nil {
			fmt.Println(err.Error())
		}
		chatName := message.Body["senderData"].(map[string]interface{})["chatName"]
		content := phoneNumber + "(" + chatName.(string) + ")"
		_, err = file.WriteString(content + "\n")

		message.AnswerWithText("Добро пожаловать в риелторское агенство \"Мечта\", в скорем времени с вами свяжутся наши агенты, а пока можете расписать причину вашего обращения и пожелания")
	})

	bot.StartReceivingNotifications()
}
