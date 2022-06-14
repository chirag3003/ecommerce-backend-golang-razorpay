package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gofiber/fiber/v2"
	"os"
	"strings"
)

func ValidateWebhook(ctx *fiber.Ctx) bool {
	webhookSignature := ctx.GetReqHeaders()["X-Razorpay-Signature"]
	if strings.TrimSpace(webhookSignature) == "" {
		return false
	}
	h := hmac.New(sha256.New, []byte(os.Getenv("RAZORPAY_WEBHOOK_SECRET")))
	h.Write(ctx.Body())
	sha := hex.EncodeToString(h.Sum(nil))
	return sha == webhookSignature
}
