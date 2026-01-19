package usecase

import (
	"faceSwapper/internal/dto"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// starts the bot service
func (S *BotService) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := S.bot.GetUpdatesChan(u)
	for update := range updates {
		S.handleUpdate(update)
	}
}

// processes request and send response
func (S *BotService) handleUpdate(update tgbotapi.Update) {

	chatID := update.Message.Chat.ID
	state := S.action.State(chatID)
	msg := update.Message
	// logic for handling states flags
	switch state {
	case dto.StateWaitingFace:
		if err := S.faceHandler(msg); err != nil {
			S.sendErr(msg, err)
		}
	case dto.StateWaitingTarget:
		if err := S.targetHandler(msg); err != nil {
			S.sendErr(msg, err)
		}
		if err := S.sendFile(msg); err != nil {
			S.sendErr(msg, err)
		}
	case dto.StateIdle:
		S.commandHandler(msg)
	}

}
