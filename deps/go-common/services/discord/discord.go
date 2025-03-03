package discord

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"biyard.co/common/base"
	"biyard.co/common/logger"
	"biyard.co/common/utils/httpclient"
	"github.com/bwmarrin/discordgo"
)

var once sync.Once

type BotController struct {
	botToken              string
	verificationChannelId string
	channelId             string
	guildId               string
	botId                 string
	missionChannelID      string
	missionChannelIDEn    string
	missionBotToken       string
	oauthURI              string
	cli                   *httpclient.HTTPClient
}

type SendChannelParams struct {
	Title         string
	TitleEn       string
	Description   string
	DescriptionEn string
	Label         string
	Deadline      int64
	ConfirmURL    string
	ImageURL      string
}

type EditChannelParams struct {
	MessageID     string
	Title         string
	TitleEn       string
	Description   string
	DescriptionEn string
	Deadline      int64
	ImageURL      string
	VoteResult    string
	VoteResultEn  string
	AcceptNum     int
	IsEnded       bool
}

type DiscordSendMessageRequest struct {
	Components []ComponentsData `json:"components"`
	Embeds     []EmbedsData     `json:"embeds"`
}

type EmbedsData struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       ImageData `json:"image"`
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

type DiscordSendMessageResponse struct {
	Type      int    `json:"type"`
	MessageID string `json:"id"`
}

func NewBotController(token string, guildId string, botId string, verificationChannelId string, id string, oauthURI string, missionChannelID, missionChannelIDEn string, missionBotToken string) *BotController {
	var botController *BotController

	once.Do(func() {
		_cli, err := httpclient.NewHttpClient(httpclient.ConstantBackoff(3, 1*time.Second), httpclient.Limiter(1000), httpclient.Headers(http.Header{
			"content-type":  []string{"application/json"},
			"authorization": []string{"Bot " + missionBotToken},
		}), httpclient.Endpoint("https://discord.com/api"))

		if err != nil {
			logger.Criticalf(nil, "NewHttpClient error: %v", err)
		}

		botController = &BotController{
			guildId:               guildId,
			botId:                 botId,
			botToken:              token,
			verificationChannelId: verificationChannelId,
			channelId:             id,
			missionChannelID:      missionChannelID,
			missionChannelIDEn:    missionChannelIDEn,
			missionBotToken:       missionBotToken,
			oauthURI:              oauthURI,
			cli:                   _cli,
		}
	})

	return botController
}

func (r *BotController) ChannelMessageEditComplex(context *base.Context, editChannelParams EditChannelParams) *base.Error {
	discord, err := discordgo.New("Bot " + r.missionBotToken)
	if err != nil {
		logger.Errorf(context, "botToken: %+v err: %+v", r.missionBotToken, err)
		return base.ErrUnknown
	}

	loc, err := time.LoadLocation("Asia/Seoul")
	unixTime := time.Unix(editChannelParams.Deadline, 0)
	deadlineTimestamp := unixTime.In(loc)
	logger.Debugf(context, "unixTime: %+v", deadlineTimestamp)

	var missionChannelID = r.missionChannelID

	if editChannelParams.IsEnded {
		var contents = fmt.Sprintf("%+v\n%+v\n찬성(Accept): %+v\n미션 기한(Mission Deadline): %+v (KST)", editChannelParams.Description, editChannelParams.DescriptionEn, editChannelParams.AcceptNum, deadlineTimestamp.Format("2006-01-02 15:04:05"))
		var titles = fmt.Sprintf("%+v(%+v)\n%+v(%+v)", editChannelParams.Title, editChannelParams.VoteResult, editChannelParams.TitleEn, editChannelParams.VoteResultEn)

		logger.Debugf(context, "discord: %+v %+v %+v %+v %+v", contents, titles, editChannelParams.ImageURL, editChannelParams.MessageID, r.missionChannelID)
		msgSend := &discordgo.MessageEdit{
			Embeds: []*discordgo.MessageEmbed{{
				Type:        discordgo.EmbedTypeRich,
				Title:       titles,
				Description: contents,
				Image: &discordgo.MessageEmbedImage{
					URL:      editChannelParams.ImageURL,
					ProxyURL: editChannelParams.ImageURL,
					Width:    80,
					Height:   80,
				},
			}},
			ID:         editChannelParams.MessageID,
			Channel:    missionChannelID,
			Components: []discordgo.MessageComponent{},
		}

		_, err := discord.ChannelMessageEditComplex(msgSend)
		if err != nil {
			logger.Debugf(context, "error is happening: %+v", err)
		}
	} else {
		var contents = fmt.Sprintf("%+v\n%+v\n찬성(Accept): %+v\n미션 기한(Mission Deadline): %+v (KST)", editChannelParams.Description, editChannelParams.DescriptionEn, editChannelParams.AcceptNum, deadlineTimestamp.Format("2006-01-02 15:04:05"))
		logger.Debugf(context, "contents: %+v messageID: %+v", contents, editChannelParams.MessageID)
		var titles = fmt.Sprintf("%s(%s)", editChannelParams.Title, editChannelParams.TitleEn)
		msgSend := &discordgo.MessageEdit{
			Embeds: []*discordgo.MessageEmbed{{
				Type:        discordgo.EmbedTypeRich,
				Title:       titles,
				Description: contents,
				Image: &discordgo.MessageEmbedImage{
					URL:      editChannelParams.ImageURL,
					ProxyURL: editChannelParams.ImageURL,
					Width:    80,
					Height:   80,
				},
			}},
			ID:      editChannelParams.MessageID,
			Channel: missionChannelID,
		}

		_, err := discord.ChannelMessageEditComplex(msgSend)
		if err != nil {
			logger.Debugf(context, "error is happening: %+v", err)
		}
	}

	return nil
}

func (r *BotController) ChannelHiddenMessageEditComplex(context *base.Context, editChannelParams EditChannelParams) *base.Error {
	discord, err := discordgo.New("Bot " + r.missionBotToken)
	if err != nil {
		logger.Errorf(context, "botToken: %+v err: %+v", r.missionBotToken, err)
		return base.ErrUnknown
	}

	loc, err := time.LoadLocation("Asia/Seoul")
	unixTime := time.Unix(editChannelParams.Deadline, 0)
	deadlineTimestamp := unixTime.In(loc)
	logger.Debugf(context, "unixTime: %+v", deadlineTimestamp)

	var missionChannelID = r.missionChannelID

	if editChannelParams.IsEnded {
		var contents = fmt.Sprintf("미션명: %+v\n(%+v)\n\n%+v\n%+v\n찬성(Accept): %+v\n미션 기한(Mission Deadline): %+v (KST)", editChannelParams.Title, editChannelParams.TitleEn, editChannelParams.Description, editChannelParams.DescriptionEn, editChannelParams.AcceptNum, deadlineTimestamp.Format("2006-01-02 15:04:05"))
		var titles = fmt.Sprintf("히든 미션(%+v)\nHidden Mission(%+v)", editChannelParams.VoteResult, editChannelParams.VoteResultEn)

		logger.Debugf(context, "discord: %+v %+v %+v %+v %+v", contents, titles, editChannelParams.ImageURL, editChannelParams.MessageID, r.missionChannelID)
		msgSend := &discordgo.MessageEdit{
			Embeds: []*discordgo.MessageEmbed{{
				Type:        discordgo.EmbedTypeRich,
				Title:       titles,
				Description: contents,
				Image: &discordgo.MessageEmbedImage{
					URL:      editChannelParams.ImageURL,
					ProxyURL: editChannelParams.ImageURL,
					Width:    80,
					Height:   80,
				},
			}},
			ID:         editChannelParams.MessageID,
			Channel:    missionChannelID,
			Components: []discordgo.MessageComponent{},
		}

		_, err := discord.ChannelMessageEditComplex(msgSend)
		if err != nil {
			logger.Debugf(context, "error is happening: %+v", err)
		}
	} else {
		var contents = fmt.Sprintf("미션명: %+v\n(%+v)\n\n%+v\n%+v\n찬성(Accept): %+v\n미션 기한(Mission Deadline): %+v (KST)", editChannelParams.Title, editChannelParams.TitleEn, editChannelParams.Description, editChannelParams.DescriptionEn, editChannelParams.AcceptNum, deadlineTimestamp.Format("2006-01-02 15:04:05"))
		logger.Debugf(context, "contents: %+v messageID: %+v", contents, editChannelParams.MessageID)
		var titles = fmt.Sprintf("히든 미션\nHidden Mission")
		msgSend := &discordgo.MessageEdit{
			Embeds: []*discordgo.MessageEmbed{{
				Type:        discordgo.EmbedTypeRich,
				Title:       titles,
				Description: contents,
				Image: &discordgo.MessageEmbedImage{
					URL:      editChannelParams.ImageURL,
					ProxyURL: editChannelParams.ImageURL,
					Width:    80,
					Height:   80,
				},
			}},
			ID:      editChannelParams.MessageID,
			Channel: missionChannelID,
		}

		_, err := discord.ChannelMessageEditComplex(msgSend)
		if err != nil {
			logger.Debugf(context, "error is happening: %+v", err)
		}
	}

	return nil
}

func (r *BotController) ChannelMessage(context *base.Context, messageID string) (*discordgo.Message, *base.Error) {
	discord, err := discordgo.New("Bot " + r.missionBotToken)
	if err != nil {
		logger.Errorf(context, "botToken: %+v err: %+v", r.missionBotToken, err)
		return nil, base.ErrUnknown
	}

	message, err := discord.ChannelMessage(r.missionChannelID, messageID)

	if err != nil {
		logger.Errorf(context, "channelID: %+v messageID: %+v err: %+v", r.missionChannelID, messageID, err)
		return nil, base.ErrUnknown
	}

	return message, nil
}

func (r *BotController) MessageReactions(context *base.Context, messageID string) ([]*discordgo.User, *base.Error) {
	discord, err := discordgo.New("Bot " + r.missionBotToken)
	if err != nil {
		logger.Errorf(context, "botToken: %+v err: %+v", r.missionBotToken, err)
		return nil, base.ErrUnknown
	}

	var accept []*discordgo.User
	var getDtErr error

	for i := 0; i < 15; i++ {
		acceptDt, e := discord.MessageReactions(r.missionChannelID, messageID, "🙆", 100, "", "")

		if e == nil {
			getDtErr = nil
			accept = acceptDt
			break
		}

		getDtErr = e
	}

	if getDtErr != nil {
		logger.Errorf(context, "botToken: %+v err: %+v", r.missionBotToken, getDtErr)
		return nil, base.ErrUnknown
	}

	return accept, nil
}

func (r *BotController) SendChannelMessageComplex(context *base.Context, sendChannelParams SendChannelParams) (string, *base.Error) {
	var res DiscordSendMessageResponse

	loc, _ := time.LoadLocation("Asia/Seoul")
	unixTime := time.Unix(sendChannelParams.Deadline, 0)
	deadlineTimestamp := unixTime.In(loc)
	logger.Debugf(context, "unixTime: %+v", deadlineTimestamp)

	var titles = sendChannelParams.Title + "\n" + sendChannelParams.TitleEn
	var contents = fmt.Sprintf("%+v\n%+v\n찬성(Accept): 0\n미션 기한(Mission Deadline): %+v (KST)", sendChannelParams.Description, sendChannelParams.DescriptionEn, deadlineTimestamp.Format("2006-01-02 15:04:05"))

	code, e := r.cli.Post(context, fmt.Sprintf("/channels/%+v/messages", r.missionChannelID), &DiscordSendMessageRequest{
		Components: []ComponentsData{
			ComponentsData{
				Type: 1,
				Components: []SubComponentsData{
					SubComponentsData{
						Type:     2,
						Label:    sendChannelParams.Label,
						Disabled: false,
						Style:    5,
						Emoji: EmojiData{
							Name: "🙆",
						},
						Url: sendChannelParams.ConfirmURL,
					},
				},
			},
		},
		Embeds: []EmbedsData{
			EmbedsData{
				Title:       titles,
				Description: contents,
				Image: ImageData{
					URL:      sendChannelParams.ImageURL,
					ProxyURL: sendChannelParams.ImageURL,
				},
			},
		},
	}, &res)

	if e != nil {
		logger.Errorf(context, "send discord message error(%v): %v", code, e.Error())
		return "", base.ErrUnknown
	}

	return res.MessageID, nil
}

func (r *BotController) SendHiddenMissionChannelComplex(context *base.Context, sendChannelParams SendChannelParams) (string, *base.Error) {
	var res DiscordSendMessageResponse

	loc, _ := time.LoadLocation("Asia/Seoul")
	unixTime := time.Unix(sendChannelParams.Deadline, 0)
	deadlineTimestamp := unixTime.In(loc)
	logger.Debugf(context, "unixTime: %+v", deadlineTimestamp)

	var titles = "히든 미션\nHidden Mission"
	var contents = fmt.Sprintf("미션명: %+v\n(%+v)\n\n%+v\n%+v\n찬성(Accept): 0\n미션 기한(Mission Deadline): %+v (KST)", sendChannelParams.Title, sendChannelParams.TitleEn, sendChannelParams.Description, sendChannelParams.DescriptionEn, deadlineTimestamp.Format("2006-01-02 15:04:05"))

	code, e := r.cli.Post(context, fmt.Sprintf("/channels/%+v/messages", r.missionChannelID), &DiscordSendMessageRequest{
		Components: []ComponentsData{
			ComponentsData{
				Type: 1,
				Components: []SubComponentsData{
					SubComponentsData{
						Type:     2,
						Label:    sendChannelParams.Label,
						Disabled: false,
						Style:    5,
						Emoji: EmojiData{
							Name: "🙆",
						},
						Url: sendChannelParams.ConfirmURL,
					},
				},
			},
		},
		Embeds: []EmbedsData{
			EmbedsData{
				Title:       titles,
				Description: contents,
				Image: ImageData{
					URL:      sendChannelParams.ImageURL,
					ProxyURL: sendChannelParams.ImageURL,
				},
			},
		},
	}, &res)

	if e != nil {
		logger.Errorf(context, "send discord message error(%v): %v", code, e.Error())
		return "", base.ErrUnknown
	}

	return res.MessageID, nil
}

func (r *BotController) MessageReactionAdd(context *base.Context, messageID, voteValue string) *base.Error {
	discord, err := discordgo.New("Bot " + r.botToken)
	if err != nil {
		logger.Errorf(context, "botToken: %+v err: %+v", r.botToken, err)
		return base.ErrUnknown
	}

	discord.MessageReactionAdd(r.verificationChannelId, messageID, "🙆")

	return nil
}

func (r *BotController) SendEmbedMessage(context *base.Context, url string, title string, description string) *base.Error {
	discord, err := discordgo.New("Bot " + r.botToken)
	if err != nil {
		logger.Errorf(context, "botToken: %+v err: %+v", r.botToken, err)
		return base.ErrUnknown
	}

	discord.ChannelMessageSendEmbed(r.verificationChannelId, &discordgo.MessageEmbed{URL: url, Title: title, Description: description}, discordgo.WithContext(context))
	return nil
}

func (r *BotController) GrantMemberRole(context *base.Context, userId string) *base.Error {
	discord, err := discordgo.New("Bot " + r.botToken)
	if err != nil {
		logger.Errorf(context, "botToken: %+v err: %+v", r.botToken, err)
		return base.ErrUnknown
	}

	err2 := discord.GuildMemberRoleAdd(r.guildId, userId, r.botId, discordgo.WithContext(context))

	if err2 != nil {
		return base.ErrUnknown
	}

	return nil
}

func (r *BotController) RevokeMemberRole(context *base.Context, userId string) *base.Error {
	discord, err := discordgo.New("Bot " + r.botToken)
	if err != nil {
		return base.ErrUnknown
	}

	err2 := discord.GuildMemberRoleRemove(r.guildId, userId, r.botId, discordgo.WithContext(context))

	if err2 != nil {
		return base.ErrUnknown
	}

	return nil
}
