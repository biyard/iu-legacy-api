package discord

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"biyard.co/common/base"
	"biyard.co/common/logger"
	"biyard.co/common/utils/httpclient"
)

type DiscordAPIConfig struct {
	BotToken  string
	ChannelID string
}

type DiscordAPI struct {
	cli       *httpclient.HTTPClient
	channelID string
}

func NewDiscordAPI(ctx context.Context, cfg DiscordAPIConfig) (*DiscordAPI, error) {
	cli, err := httpclient.NewHttpClient(httpclient.ConstantBackoff(3, 1*time.Second), httpclient.Limiter(1000), httpclient.Headers(http.Header{
		"content-type":  []string{"application/json"},
		"authorization": []string{"Bot " + cfg.BotToken},
	}), httpclient.Endpoint("https://discord.com/api"))
	if err != nil {
		return nil, err
	}

	return &DiscordAPI{
		cli:       cli,
		channelID: cfg.ChannelID,
	}, nil

}

type DiscordButton struct {
	Label string
	URL   string
	Emoji string
}

type DiscordEmbedRequest struct {
	Title       string
	Description string
	ImageURL    string
	Buttons     []DiscordButton
}

type DiscordSendMessageResponse struct {
	Type      int    `json:"type"`
	MessageID string `json:"id"`
}

func (r *DiscordAPI) SendMessage(ctx context.Context, req DiscordEmbedRequest) (string, error) {
	buttons := []SubComponentsData{}
	for _, button := range req.Buttons {
		buttons = append(buttons, SubComponentsData{
			Type:     2,
			Label:    button.Label,
			Disabled: false,
			Style:    5,
			Emoji:    EmojiData{Name: button.Emoji},
			Url:      button.URL,
		})
	}

	data := EmbedsData{
		Title:       req.Title,
		Description: req.Description,
	}

	if req.ImageURL != "" {
		data.Image = &ImageData{
			URL:      req.ImageURL,
			ProxyURL: req.ImageURL,
		}
	}

	var res DiscordSendMessageResponse

	code, e := r.cli.Post(ctx, fmt.Sprintf("/channels/%+v/messages", r.channelID), &DiscordSendMessageRequest{
		Components: []ComponentsData{
			ComponentsData{
				Type:       1,
				Components: buttons,
			},
		},
		Embeds: []EmbedsData{data},
	}, &res)

	if e != nil {
		logger.Errorf(ctx, "send discord message error(%v): %v", code, e.Error())
		return "", base.ErrUnknown
	}

	return res.MessageID, nil
}

type DiscordSendMessageRequest struct {
	Components []ComponentsData `json:"components"`
	Embeds     []EmbedsData     `json:"embeds"`
}

type EmbedsData struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Image       *ImageData `json:"image,omitempty"`
}

type ImageData struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url"`
}

type ComponentsData struct {
	Type       int                 `json:"type"`
	Components []SubComponentsData `json:"components"`
}

type SubComponentsData struct {
	Type     int       `json:"type"`
	Label    string    `json:"label"`
	Disabled bool      `json:"disabled"`
	Style    int       `json:"style"`
	Emoji    EmojiData `json:"emoji"`
	Url      string    `json:"url"`
}

type EmojiData struct {
	Name string `json:"name"`
}
