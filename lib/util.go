package lib

import (
	"context"
	"math/rand"
	"time"
)

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

//generates unique base62 UUID and be sure
func generateUniqueShortLinkString(ctx context.Context, n int) (ret string) {
	rand.Seed(time.Now().UnixMilli())

	for i := 0; i < n; i++ {
		ret += string(alphabet[rand.Intn(len(alphabet))])
	}

	for checkIfExist(ctx, ret) {
		ret = generateUniqueShortLinkString(ctx, 7)
	}
	return
}
