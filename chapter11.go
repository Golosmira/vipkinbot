package main

import ( "log" "strconv"

"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() { bot, err := tgbotapi.NewBotAPI("5837275423:AAGLDXrqUmlKJl5crn4ifMKE6VBwQP5l4Us") if err != nil { log.Panic(err) }

// bot.Debug = true

log.Printf("Authorized on account %s", bot.Self.UserName)

u := tgbotapi.NewUpdate(0)
u.Timeout = 60

updates, err := bot.GetUpdatesChan(u)

for update := range updates {
	if update.Message == nil { // ignore any non-Message Updates
		continue
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	reply := "Your message: " + update.Message.Text + " - has been received with ID: " + strconv.Itoa(update.Message.MessageID)

	msg.Text = reply
	bot.Send(msg)
}
}