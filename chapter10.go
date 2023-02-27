package main

import ( "github.com/go-telegram-bot-api/telegram-bot-api" "log" "strings" )

func main() { bot, err := tgbotapi.NewBotAPI("") if err != nil { log.Fatal(err) }

u := tgbotapi.NewUpdate(0)
u.Timeout = 60

updates, err := bot.GetUpdatesChan(u)

for update := range updates {
    if update.Message == nil {
        continue
    }

    if strings.Contains(update.Message.Text, "плохо") {
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Сообщение удалено!")
        msg.ReplyToMessageID = update.Message.MessageID
        bot.Send(msg)

        delConf := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
        bot.Send(delConf)
    }
}
}