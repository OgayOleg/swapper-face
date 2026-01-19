package usecase

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// commandHandler handles incoming commands from the Telegram bot.
func (S *BotService) commandHandler(msg *tgbotapi.Message) {
	cmd := msg.Command()
	chatID := msg.Chat.ID
	// commands processing
	switch cmd {
	case "start":
		S.bot.Send(tgbotapi.NewMessage(chatID, "Hello!"))

	case "faceswap":
		S.action.SetWaitingFaceState(chatID)
		S.bot.Send(tgbotapi.NewMessage(chatID, "Send a photo for face target"))
	}
}
