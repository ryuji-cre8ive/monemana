package ui

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"github.com/ryuji-cre8ive/monemana/internal/usecase"

	"golang.org/x/xerrors"
)

type (
	WebhookHandler interface {
		PostWebhook(c echo.Context) error
	}

	webhookHandler struct {
		usecase.WebhookUsecase
	}
)

func (h *webhookHandler) PostWebhook(c echo.Context) error {
	Secret := os.Getenv("LINE_BOT_CHANNEL_SECRET")
	cb, err := webhook.ParseRequest(Secret, c.Request())
	if err != nil {
		return xerrors.Errorf("webhook.ParseRequest error: %w", err)
	}

	relatedUserList := make([]string, 0, 0)
	targetUserList := make([]string, 0, 0)

	for _, event := range cb.Events {
		switch event := event.(type) {
		case webhook.MessageEvent:
			switch source := event.Source.(type) {
			case webhook.GroupSource:
				switch message := event.Message.(type) {
				case webhook.TextMessageContent:
					if strings.Contains(message.Text, "名前変更") {
						newName := strings.TrimSpace(strings.TrimPrefix(message.Text, "名前変更"))
						h.WebhookUsecase.UpdateUserName(c, source.UserId, newName)
						if err := replyMessage(event.ReplyToken, "名前変更完了👍"); err != nil {
							return xerrors.Errorf("failed to reply message: %w", err)
						}
					}
					if message.Text == "集計" {
						aggregateMessage, err := h.WebhookUsecase.AggregateTransaction(c, source.GroupId)
						if err != nil {
							return xerrors.Errorf("aggregate transaction err: %w", err)
						}
						if aggregateMessage == "" {
							if err := replyMessage(event.ReplyToken, "まだ何も登録されてないよ😢"); err != nil {
								return xerrors.Errorf("failed to reply message: %w", err)
							}
						}
						if err := replyMessage(event.ReplyToken, aggregateMessage); err != nil {
							return xerrors.Errorf("failed to reply message: %w", err)
						}
					}
					if message.Mention != nil {
						for _, mentionElement := range message.Mention.Mentionees {
							switch mention := mentionElement.(type) {
							case webhook.UserMentionee:
								targetUserList = append(targetUserList, mention.UserId)
							}
						}

						relatedUserList = append(targetUserList, source.UserId)
						groupID := source.GroupId
						userID := source.UserId

						for _, userID := range relatedUserList {
							// ユーザーが存在するかチェック
							exists, err := h.WebhookUsecase.CheckUserExists(c, userID)

							// ユーザーが存在しない場合のみ登録
							if !exists || err != nil {
								h.WebhookUsecase.CreateUser(c, userID, userID)
							}
						}
						pattern := regexp.MustCompile("[\\s　]+")
						re := regexp.MustCompile(`@.*?\n`)
						text := re.ReplaceAllString(message.Text, "")

						splitText := pattern.Split(text, -1)
						if len(splitText) < 2 {
							if err := replyMessage(event.ReplyToken, "フォーマットが正しくないかも😢"); err != nil {
								return xerrors.Errorf("failed to reply message: %w", err)
							}
						}
						title, priceStr := splitText[0], splitText[1]
						price, parseIntErr := strconv.ParseUint(priceStr, 10, 64)
						if parseIntErr != nil {
							return xerrors.Errorf("price parse err: %w", parseIntErr)
						}

						TransactionErr := h.WebhookUsecase.CreateTransaction(c, title, price, userID, targetUserList, groupID)
						if TransactionErr != nil {
							return xerrors.Errorf("failed to create transaction: %w", TransactionErr)
						}
						if err := replyMessage(event.ReplyToken, "登録完了👍"); err != nil {
							return xerrors.Errorf("failed to reply message: %w", err)
						}
					} else {
						if err := replyMessage(event.ReplyToken, "登録されてないコマンドかも😢\n```\n名前変更 <あなたの名前>\n```\nで名前変更できるよ！\n```@<友達の名前>\n<商品の名前> <値段>\n```\nで登録できるよ！"); err != nil {
							return xerrors.Errorf("failed to reply message: %w", err)
						}
					}
				}
			}
		case webhook.JoinEvent:
			joinMessage := "グループに招待してくれてありがとう🥺\n使い方を説明するね👍\nまずは全員が名前変更してね。やり方はこうだよ\n```\n名前変更 <あなたの名前>\n```\nそうすると名前が変更されてみやすくなるよ🙌\n次に登録方法だよ\n```\n@<友達の名前>\n<商品の名前> <値段>\n```\nで登録できるよ！こんな感じで送ってね！\n```\n@田中\n苺大福 380\n```\n\n最後に集計方法だよ！\n```\n集計\n```\nで集計できるよ！\nわからないことがあったらX（旧Twitter）の@ryuji_vlogにお問い合わせてね😢"
			if err := replyMessage(event.ReplyToken, joinMessage); err != nil {
				return xerrors.Errorf("failed to reply message: %w", err)
			}
		}
		if err != nil {
			return xerrors.Errorf("failed to post webhook: %w", err)
		}
	}
	return c.NoContent(200)
}
func replyMessage(token string, message string) error {
	bot, err := messaging_api.NewMessagingApiAPI(
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		return xerrors.Errorf("bot err: %w", err)
	}
	if _, err := bot.ReplyMessage(&messaging_api.ReplyMessageRequest{
		ReplyToken: token,
		Messages: []messaging_api.MessageInterface{
			messaging_api.TextMessage{
				Text: message,
			},
		},
	}); err != nil {
		xerrors.Errorf("reply message err: %w", err)
	}
	return nil
}
