package database

import (
	"fmt"
	"time"
)

func SetBlacklistToken(token string, userID string) error {
	return rdb.Set(ctx, fmt.Sprintf("blacklist_token:%s", token), userID, time.Hour*1).Err()
}

func IsBlacklistToken(token string) (bool, error) {
	val, err := rdb.Exists(ctx, fmt.Sprintf("blacklist_token:%s", token)).Result()
	if err != nil {
		return false, err
	}
	return val == 1, nil
}
