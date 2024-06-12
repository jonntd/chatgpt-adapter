package coze

import (
	"github.com/bincooo/chatgpt-adapter/internal/common"
	"github.com/bincooo/chatgpt-adapter/internal/gin.handler/response"
	"github.com/bincooo/chatgpt-adapter/internal/plugin"
	"github.com/bincooo/chatgpt-adapter/logger"
	"github.com/bincooo/chatgpt-adapter/pkg"
	"github.com/bincooo/coze-api"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func completeToolCalls(ctx *gin.Context, cookie, proxies string, completion pkg.ChatCompletion) bool {
	logger.Info("completeTools ...")
	exec, err := plugin.CompleteToolCalls(ctx, completion, func(message string) (string, error) {
		message = strings.TrimSpace(message)
		system := ""
		if strings.HasPrefix(message, "<|system|>") {
			index := strings.Index(message, "<|end|>")
			system = message[:index+7]
			message = strings.TrimSpace(message[index+7:])
		}

		var pMessages []coze.Message
		if system != "" {
			pMessages = append(pMessages, coze.Message{
				Role:    "system",
				Content: system,
			})
		}

		pMessages = append(pMessages, coze.Message{
			Role:    "user",
			Content: message,
		})

		co, msToken := extCookie(cookie)
		options, mode, err := newOptions(proxies, completion.Model, pMessages)
		if err != nil {
			return "", logger.WarpError(err)
		}

		chat := coze.New(co, msToken, options)
		var lock *common.ExpireLock
		if mode == 'o' {
			l, e := draftBot(ctx, pMessages, chat, completion)
			if e != nil {
				return "", logger.WarpError(e.Err)
			}
			lock = l
		}

		query := ""
		if mode == 'w' {
			query = pMessages[len(pMessages)-1].Content
			chat.WebSdk(chat.TransferMessages(pMessages[:len(pMessages)-1]))
		} else {
			query = coze.MergeMessages(pMessages)
		}

		chatResponse, err := chat.Reply(ctx.Request.Context(), coze.Text, query)
		// 构建完请求即可解锁
		if lock != nil {
			lock.Unlock()
			botId := customBotId(completion.Model)
			rmLock(botId)
			logger.Infof("构建完成解锁：%s", botId)
		}

		if err != nil {
			return "", logger.WarpError(err)
		}

		return waitMessage(chatResponse, plugin.ToolCallCancel)
	})

	if err != nil {
		errMessage := err.Error()
		if strings.Contains(errMessage, "Login verification is invalid") {
			logger.Error(err)
			response.Error(ctx, http.StatusUnauthorized, errMessage)
			return true
		}

		logger.Error(err)
		response.Error(ctx, -1, errMessage)
		return true
	}

	return exec
}
