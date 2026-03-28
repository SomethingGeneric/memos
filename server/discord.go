package server

import (
	"context"
	"fmt"

	"github.com/usememos/memos/plugin/discord"
	"github.com/usememos/memos/store"
)

type discordHandler struct {
	store *store.Store
}

func newDiscordHandler(store *store.Store) *discordHandler {
	return &discordHandler{store: store}
}

func (d *discordHandler) BotToken(ctx context.Context) string {
	return d.store.GetSystemSettingValueWithDefault(&ctx, apiv1.SystemSettingDiscordBotTokenName.String(), "")
}

func (d *discordHandler) MessageHandle(ctx context.Context, bot *discord.Bot, message discord.Message, attachments []discord.Attachment) error {
	// Handle Discord message logic here
	return nil
}

func (d *discordHandler) CallbackQueryHandle(ctx context.Context, bot *discord.Bot, callbackQuery discord.CallbackQuery) error {
	// Handle Discord callback query logic here
	return nil
}

func generateKeyboardForMemoID(id int32) [][]discord.InlineKeyboardButton {
	// Generate Discord keyboard for memo ID here
	return nil
}

func convertToMarkdown(text string, messageEntities []discord.MessageEntity) string {
	// Convert Discord text to Markdown format here
	return ""
}
