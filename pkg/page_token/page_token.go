package page_token

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	ErrInvalidToken         = errors.New(400, "Token failed", "The field `page_token` is invalid")
	ErrOverdueToken         = errors.New(400, "Token overdue", "The field `page_token` is overdue")
	ErrOverMaxPageSizeToken = errors.New(400, "Token over max page size", "The field `page_token` is over max page size")
	ErrInvalidPageSize      = errors.New(400, "PageSize failed", "The page size provided must not be negative.")
)

const (
	defaultMaxIndex    = 0
	defaultMaxElements = 0
	defaultSalt        = "beer-shop"
	layout             = "2006-01-02 15-04-05"
)

type PageToken interface {
	TokenGenerator
	ProcessPageTokens
}

type TokenGenerator interface {
	ForIndex(int) string
	GetIndex(string) (int, error)
}

type ProcessPageTokens interface {
	ProcessPageTokens(numElements int, pageSize int32, pageToken string) (start, end int32, nextToken string, err error)
}

type token struct {
	maxIndex       int32         // Maximum index
	maxElements    int           // Maximum number of elements
	timeLimitation time.Duration // token Time limitation
	salt           string        // Special identification
}

func (t *token) ForIndex(i int) string {
	return base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s%s:%d", t.salt, time.Now().Format(layout), i)))
}

func (t *token) GetIndex(s string) (int, error) {
	if s == "" {
		return 0, nil
	}
	bs, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return -1, ErrInvalidToken
	}
	if !strings.HasPrefix(string(bs), t.salt) {
		return -1, ErrInvalidToken
	}
	parseToken := strings.Split(strings.TrimPrefix(string(bs), t.salt), ":")
	if len(parseToken) != 2 {
		return -1, ErrInvalidToken
	}
	if t.timeLimitation != 0 {
		generateTime, err := time.Parse(layout, parseToken[0])
		if err != nil {
			return -1, ErrInvalidToken
		}
		if generateTime.Add(t.timeLimitation).After(time.Now()) {
			return -1, ErrOverdueToken
		}
	}
	i, err := strconv.Atoi(parseToken[1])
	if err != nil {
		return -1, ErrInvalidToken
	}
	if t.maxIndex != defaultMaxIndex && int32(i) > t.maxIndex {
		return -1, ErrOverMaxPageSizeToken
	}
	return i, nil
}

func NewTokenGenerate(options ...TokenOption) PageToken {
	t := &token{
		maxIndex:       defaultMaxIndex,
		maxElements:    defaultMaxElements,
		timeLimitation: 0,
		salt:           defaultSalt,
	}
	for _, option := range options {
		option(t)
	}
	return t
}

func (t *token) ProcessPageTokens(numElements int, pageSize int32, pageToken string) (start, end int32, nextToken string, err error) {
	if pageSize < 0 {
		return 0, 0, "", ErrInvalidPageSize
	}

	if t.maxElements != defaultMaxElements && numElements > t.maxElements {
		numElements = t.maxElements
	}

	if pageToken != "" {
		index, err := t.GetIndex(pageToken)
		if err != nil {
			return 0, 0, "", err
		}

		token32 := int32(index)
		if token32 < 0 || token32 >= int32(numElements) {
			return 0, 0, "", ErrInvalidToken
		}
		start = token32
	}

	if pageSize == 0 {
		pageSize = int32(numElements)
	}
	end = min(start+pageSize, int32(numElements))

	if end < int32(numElements) {
		nextToken = t.ForIndex(int(end))
	}

	return start, end, nextToken, nil
}

func min(a, b int32) int32 {
	if a > b {
		return b
	}
	return a
}
