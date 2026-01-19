package usecase

import (
	"errors"
	"faceSwapper/internal/dto"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// send error message to user
func (S *BotService) sendErr(msg *tgbotapi.Message, err error) {
	var line string
	if err == dto.ErrNotFoundForFace || err == dto.ErrNotFoundForTarget || errors.Is(err, dto.ErrStatusFailed) || errors.Is(err, dto.ErrStatusCanceled) {
		line = fmt.Sprintln("Error: ", err)
	} else {
		line = fmt.Sprintln("Error: server error")
	}
	msg_output := tgbotapi.NewMessage(msg.Chat.ID, line)
	fmt.Println("@", msg.From.UserName, " ", err)
	S.bot.Send(msg_output)
}
