package usecase

import (
	"faceSwapper/internal/dto"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (S *BotService) targetHandler(msg *tgbotapi.Message) error {
	chatID := msg.Chat.ID
	file, err := S.targetURL(msg)
	if err != nil {
		return err
	}
	S.action.AddTarget(chatID, file)
	S.action.SetIdleState(chatID)
	return nil
}

func (S *BotService) targetURL(msg *tgbotapi.Message) (string, error) {
	var file string
	if n := len(msg.Photo); n > 0 {
		file = msg.Photo[n-1].FileID
	} else if msg.Video != nil {
		file = msg.Video.FileID
	} else {
		return "", dto.ErrNotFoundForTarget
	}
	url, err := S.bot.GetFileDirectURL(file)
	if err != nil {
		return "", err
	}
	return url, nil
}
