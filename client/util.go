package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"os"
	"strings"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/tag"
)

func handleLogonMessage(msg *quickfix.Message) {
	msg.Body.SetInt(98, 0)
	msg.Body.SetInt(108, 30)
	msg.Body.SetString(554, os.Getenv("COINBASE_API_KEY_PASSPHRASE"))
	msg.Body.SetString(96, getLogonRawData(msg))
	msg.Body.SetString(8013, "S")
	msg.Body.SetString(9406, "Y")
}

// getLogonRawData returns the logon message raw data.
// See: https://docs.cloud.coinbase.com/exchange/docs/messages#logon-a
func getLogonRawData(msg *quickfix.Message) string {
	sendingtime, _ := msg.Header.GetString(tag.SendingTime)
	msgType := "A"
	msgSeqNum, _ := msg.Header.GetString(tag.MsgSeqNum)
	senderCompID := os.Getenv("COINBASE_API_KEY")
	targetCompID := "Coinbase"
	password := os.Getenv("COINBASE_API_KEY_PASSPHRASE")

	presign := strings.Join(
		[]string{
			sendingtime,
			msgType,
			msgSeqNum,
			senderCompID,
			targetCompID,
			password,
		},
		string("\x01"),
	)

	apiKeySecret := os.Getenv("COINBASE_API_KEY_SECRET")
	key, _ := base64.StdEncoding.DecodeString(apiKeySecret)
	hmacHash := hmac.New(sha256.New, key)
	hmacHash.Write([]byte(presign))
	encoded := base64.StdEncoding.EncodeToString(hmacHash.Sum(nil))

	return encoded
}
