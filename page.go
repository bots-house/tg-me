package tgme

import (
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type page struct {
	doc *goquery.Document

	Title       string
	Extra       string
	Description string
	Avatar      string
	Button      string
	HasPreview  bool
}

func newPage(body io.Reader) (*page, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	desc, _ := doc.Find(`meta[property="og:description"]`).Attr("content")
	desc = strings.Replace(desc, "\t", "\n", -1)

	avatar, _ := doc.Find(".tgme_page_photo_image").Attr("src")
	hasPreview := doc.Find(".tgme_action_button_label").Length() > 0

	return &page{
		doc:         doc,
		Title:       strings.TrimSpace(doc.Find(".tgme_page_title").Children().Text()),
		Extra:       strings.TrimSpace(doc.Find(".tgme_page_extra").Text()),
		Description: desc,
		Avatar:      avatar,
		Button:      doc.Find(`.tgme_action_button_new`).First().Text(),
		HasPreview:  hasPreview,
	}, nil
}

var numberRegex = regexp.MustCompile(`([0-9 ]+)`)

func (p *page) parseExtraAsNumber(i int) int {
	matches := numberRegex.FindAllString(p.Extra, -1)

	if i < len(matches) {
		match := matches[i]
		match = strings.Replace(match, " ", "", -1)
		v, _ := strconv.Atoi(match)
		return v
	}

	return 0
}

func (p *page) Parse() (*Result, error) {
	label := strings.ToLower(p.Button)

	switch label {
	case "send message":
		// if button text is send message it's bot or user
		return &Result{
			User: p.toUser(),
		}, nil
	case "view in telegram":
		// if button text is "view in telegram" then it's public chat or group
		if p.HasPreview {
			return &Result{
				Channel: p.toChannel(),
			}, nil
		}

		return &Result{
			Chat: p.toChat(),
		}, nil
	case "join channel":
		return &Result{
			Channel: p.toChannel(),
		}, nil
	case "join group":
		return &Result{
			Chat: p.toChat(),
		}, nil
	}

	return nil, nil
}

func (p *page) toChat() *Chat {
	return &Chat{
		Title:       p.Title,
		Members:     p.parseExtraAsNumber(0),
		Online:      p.parseExtraAsNumber(1),
		Description: p.Description,
		Avatar:      p.Avatar,
	}
}

func (p *page) toChannel() *Channel {
	return &Channel{
		Title:       p.Title,
		Members:     p.parseExtraAsNumber(0),
		Description: p.Description,
		Avatar:      p.Avatar,
	}
}

var noBioRegex = regexp.MustCompile(`You can contact @[a-z0-9_]+ right away\.`)

func (p *page) toUser() *User {
	bio := p.Description

	if noBioRegex.MatchString(bio) {
		bio = ""
	}

	return &User{
		Name:     p.Title,
		Username: p.Extra,
		Bio:      bio,
		Avatar:   p.Avatar,
	}
}
