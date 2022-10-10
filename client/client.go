package client

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"go.uber.org/zap"
)

type application struct{}

func (app *application) OnCreate(sessionID quickfix.SessionID) {
	zap.L().Info("received OnCreate message", zap.Any("session_id", sessionID))
}

func (app *application) OnLogon(sessionID quickfix.SessionID) {
	zap.L().Info("received OnLogon message", zap.Any("session_id", sessionID))
}

func (app *application) OnLogout(sessionID quickfix.SessionID) {
	zap.L().Info("received OnLogout message", zap.Any("session_id", sessionID))
}

func (app *application) FromAdmin(
	msg *quickfix.Message,
	sessionID quickfix.SessionID,
) (reject quickfix.MessageRejectError) {
	zap.L().Info(
		"received FromAdmin message",
		zap.Any("msg", msg),
		zap.Any("session_id", sessionID),
	)
	return nil
}

func (app *application) ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) {
	zap.L().Info("received ToAdmin message", zap.Any("msg", msg))

	isLogonMessage := msg.IsMsgTypeOf(enum.MsgType_LOGON)
	if isLogonMessage {
		zap.L().Info("processing ToAdmin logon message")
		handleLogonMessage(msg)
		return
	}
}

func (app *application) ToApp(
	msg *quickfix.Message,
	sessionID quickfix.SessionID,
) (err error) {
	zap.L().Info(
		"received ToApp message",
		zap.Any("msg", msg),
		zap.Any("session_id", sessionID),
	)
	return nil
}

func (app *application) FromApp(
	msg *quickfix.Message,
	sessionID quickfix.SessionID,
) (reject quickfix.MessageRejectError) {
	zap.L().Info(
		"received FromApp message",
		zap.Any("msg", msg),
		zap.Any("session_id", sessionID),
	)
	return nil
}

func NewClient() quickfix.Application {
	return &application{}
}
