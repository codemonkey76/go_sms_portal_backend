package commands

import (
	"math/rand"
	"sms_portal/env"
	"sms_portal/ui"
)

func GenerateKeyCommand(args []string) error {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 32)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	env.Env(".env").Set("APP_KEY", string(b))
	ui.Info("Key generated successfully.")
	return nil
}
