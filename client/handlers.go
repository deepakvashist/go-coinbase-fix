package client

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix42/heartbeat"
	"go.uber.org/zap"
)

func onHeartBeat(
	msg heartbeat.Heartbeat,
	sessionID quickfix.SessionID,
) quickfix.MessageRejectError {
	zap.L().Info(
		"received onHeartBeat message",
		zap.Any("msg", msg),
		zap.Any("session_id", sessionID),
	)
	return nil
}
