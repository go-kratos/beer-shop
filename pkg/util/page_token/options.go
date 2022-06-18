package page_token

import "time"

type TokenOption func(*token)

// SetMaxTokenPageSize .
func SetMaxTokenPageSize(pageSize int32) TokenOption {
	return func(t *token) {
		if pageSize <= 0 {
			return
		}
		t.maxIndex = pageSize
	}
}

// SetTimeLimitation .
func SetTimeLimitation(timeLimitation time.Duration) TokenOption {
	return func(t *token) {
		t.timeLimitation = timeLimitation
	}
}

// SetSalt .
func SetSalt(salt string) TokenOption {
	return func(t *token) {
		t.salt = salt
	}
}
