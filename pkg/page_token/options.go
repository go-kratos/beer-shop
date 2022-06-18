package page_token

import "time"

type TokenOption func(*token)

// SetTokenMaxIndex .
func SetTokenMaxIndex(index int32) TokenOption {
	return func(t *token) {
		if index <= 0 {
			return
		}
		t.maxIndex = index
	}
}

// SetTokenMaxElements .
func SetTokenMaxElements(elements int) TokenOption {
	return func(t *token) {
		if elements <= 0 {
			return
		}
		t.maxElements = elements
	}
}

// SetTokenTimeLimitation .
func SetTokenTimeLimitation(timeLimitation time.Duration) TokenOption {
	return func(t *token) {
		t.timeLimitation = timeLimitation
	}
}

// SetTokenSalt .
func SetTokenSalt(salt string) TokenOption {
	return func(t *token) {
		t.salt = salt
	}
}
