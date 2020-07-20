package tgme

import (
	"context"
	"net/http"
	"strings"
)

type Parser struct {
	Client *http.Client
}

// User it's user or bot
type User struct {
	// Name of user or bot
	Name string `json:"name"`

	// Username of user or bot
	Username string `json:"username"`

	// Bio of user or bot
	Bio string `json:"bio,omitempty"`

	// Link to avatar
	Avatar string `json:"avatar,omitempty"`
}

// Channel it's Telegram channel.
type Channel struct {
	// Title of channel
	Title string `json:"title"`

	// Members count
	Members int `json:"members"`

	// Description of channel
	Description string `json:"description"`

	// URL of channel avatar
	Avatar string `json:"avatar"`
}

// Chat it's Telegram Chat
type Chat struct {
	// Title of channel
	Title string `json:"title"`

	// Members count
	Members int `json:"members"`

	// Online members count
	Online int `json:"online"`

	// Description of channel
	Description string `json:"description,omitempty"`

	// URL of channel avatar
	Avatar string `json:"avatar,omitempty"`
}

type Result struct {
	User    *User    `json:"user,omitempty"`
	Channel *Channel `json:"channel,omitempty"`
	Chat    *Chat    `json:"chat,omitempty"`
}

// Parse returns info parsed from t.me. Link can be something like:
// - User: t.me/MrLinch
// - Bot: t.me/crosser_bot
// - Channel: t.me/crosser_live
// - Chat: t.me/crosser_chat
// - Private Channel: t.me/joinchat/AAAAAENM1m0f_WHVNXjP4w
// - Private Chat: https://t.me/joinchat/BEqxykAVc_3aNu906D2v7A
func (psr *Parser) Parse(ctx context.Context, link string) (*Result, error) {
	client := psr.Client

	if client == nil {
		client = http.DefaultClient
	}

	if !strings.Contains(link, "http://") && !strings.Contains(link, "https://") {
		link = "https://" + link
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	page, err := newPage(res.Body)
	if err != nil {
		return nil, err
	}

	return page.Parse()
}

var defaultParser = &Parser{Client: http.DefaultClient}

func Parse(ctx context.Context, link string) (*Result, error) {
	return defaultParser.Parse(ctx, link)
}
