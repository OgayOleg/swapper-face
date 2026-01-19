package usecase

import (
	"faceSwapper/internal/dto"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Wmodel interface {
	GetID(target, source string) (string, error)
	FetchResourceWithRetry(id string) (string, error)
}

type BotService struct {
	action dto.Action
	bot    *tgbotapi.BotAPI
	wmodel Wmodel
}

func New(bot *tgbotapi.BotAPI, wmodel Wmodel) *BotService {
	return &BotService{
		action: dto.NewAction(),
		bot:    bot,
		wmodel: wmodel,
	}
}
