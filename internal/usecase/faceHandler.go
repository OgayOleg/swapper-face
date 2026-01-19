package usecase

import (
	"faceSwapper/internal/dto"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// faceHandler handles the face swapping process.
func (S *BotService) faceHandler(msg *tgbotapi.Message) error {
	chatID := msg.Chat.ID
	link, err := S.faceURL(msg)
	if err != nil {
		return err
	}
	S.action.AddFace(chatID, link)
	S.bot.Send(tgbotapi.NewMessage(chatID, "Face added successfully!"))
	S.bot.Send(tgbotapi.NewMessage(chatID, "Send a photo or video for face swapping"))
	S.action.SetWaitingTargetState(chatID)
	return nil
}

// get url photo user face for swap
func (S *BotService) faceURL(msg *tgbotapi.Message) (string, error) {
	files := msg.Photo
	if len(files) == 0 {
		return "", dto.ErrNotFoundForFace
	}
	photo := files[len(files)-1]

	if link, err := S.bot.GetFileDirectURL(photo.FileID); err != nil {
		return "", err
	} else {
		return link, nil
	}
}
