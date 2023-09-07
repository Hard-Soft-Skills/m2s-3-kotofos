package main

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"os"
	"os/signal"
)

type BotApp struct {
	bot   *bot.Bot
	queue chan VoiceTask
}

func NewBotApp(token string) *BotApp {
	app := BotApp{}
	opts := []bot.Option{
		bot.WithDefaultHandler(app.defaultHandler),
	}

	b, err := bot.New(token, opts...)
	if nil != err {
		panic(err)
	}
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, app.startHandler)
	app.bot = b
	app.queue = make(chan VoiceTask, 100)

	return &app
}

func (app *BotApp) Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go app.processVoiceMessage(app.queue)

	println("Starting bot")

	// will block until SIGINT
	app.bot.Start(ctx)
}

func (app *BotApp) defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message.Voice != nil {
		app.voiceMessageHandler(ctx, b, update)
		return
	}
}

func (app *BotApp) startHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Welcome to the bot! send voice message to get started",
	})
	if err != nil {
		log.Printf("Failed to send start message: %v", err)
	}

	user := User{Username: "", Language: "en"}
	_, err = CreateUser(user)
	if err != nil {
		panic("Failed to create user")
	}
}

type VoiceTask struct {
	UserID int64
	ChatID int64
	ctx    context.Context
	FileID string
}

func (app *BotApp) voiceMessageHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	//user, err := FindUserByUsername(update.Message.From.Username)
	//if err != nil {
	//	//todo send: user not found. please start bot first
	//	panic("Failed to find user")
	//}

	// todo
	// check user preferences
	// select processing method
	// create processing task

	//lang := user.Language

	task := VoiceTask{
		ctx:    ctx,
		UserID: update.Message.From.ID,
		FileID: update.Message.Voice.FileID,
		ChatID: update.Message.Chat.ID,
	}
	app.queue <- task

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "processing",
	})
	if err != nil {
		log.Printf("Failed to acknowledge voice message: %v", err)
	}
}

func (app *BotApp) processVoiceMessage(queue chan VoiceTask) {
	for {
		task, ok := <-queue
		if !ok {
			fmt.Println("Channel closed, exiting worker.")
			return
		}
		fmt.Printf("Received task: %d", task)
		_, err := app.bot.SendMessage(task.ctx, &bot.SendMessageParams{
			ChatID: task.ChatID,
			Text:   "done",
		})
		// todo process voice message
		_, err = app.bot.SendVoice(task.ctx, &bot.SendVoiceParams{
			ChatID: task.ChatID,
			Voice:  &models.InputFileString{Data: task.FileID},
		})

		if err != nil {
			panic(fmt.Sprintf("Failed to send message Error: %s", err))
		}
	}
}
