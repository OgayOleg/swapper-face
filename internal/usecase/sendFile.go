package usecase

import (
	"faceSwapper/internal/dto"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (S *BotService) sendFile(msg *tgbotapi.Message) error {
	chatID := msg.Chat.ID
	// chat action status
	if len(msg.Photo) != 0 { // photo type
		S.bot.Send(tgbotapi.NewMessage(chatID, "Target photo added successfully!"))
		S.bot.Send(tgbotapi.NewChatAction(chatID, tgbotapi.ChatUploadPhoto))
	} else if msg.Video != nil { // video type
		S.bot.Send(tgbotapi.NewMessage(chatID, "Target video added successfully!"))
		S.bot.Send(tgbotapi.NewChatAction(chatID, tgbotapi.ChatUploadVideo))
	} else { // unsupported another type
		return dto.ErrNotFoundForTarget
	}

	// get url
	id, err := S.wmodel.GetID(S.action.Get(chatID))
	if err != nil {
		return err
	}
	file, err := S.wmodel.FetchResourceWithRetry(id)
	if err != nil {
		return err
	}

	// send result
	if len(msg.Photo) != 0 {
		_, err := S.bot.Send(tgbotapi.NewPhoto(chatID, tgbotapi.FileURL(file)))
		if err != nil {
			return err
		}
	} else {
		_, err := S.bot.Send(tgbotapi.NewVideo(chatID, tgbotapi.FileURL(file)))
		if err != nil {
			return err
		}
	}
	return nil
}
