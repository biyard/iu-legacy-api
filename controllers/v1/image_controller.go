package v1

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"biyard.co/common/base"
	"biyard.co/common/logger"
	"biyard.co/iuniverse-api/config"
	"biyard.co/iuniverse-api/errors"
)

type ImageController struct{}

type ImageRequest struct {
	Image string `form:"image"`
}

type ImageResponse struct {
	ImageName        string `json:"name"`
	ImageDescription string `json:"description"`
	ImageUri         string `json:"image"`
}

func NewImageController(ctx context.Context, cfg config.AppConfig) *ImageController {
	return &ImageController{}
}

func (r *ImageController) OnInit(route base.PathRouter) {
	route.GET("/", r.imageData)
}

func (r *ImageController) imageData(c *base.Context, req *ImageRequest) (*ImageResponse, *base.Error) {
	if strings.Contains(req.Image, "json") {
		resp, err := http.Get(req.Image)
		if err != nil {
			logger.Debugs(c, err)
			return nil, errors.ErrMsg
		}

		defer resp.Body.Close()

		// 결과 출력
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Debugs(c, err)
			return nil, errors.ErrMsg
		}
		var myJSONparse map[string]interface{}
		json.Unmarshal(data, &myJSONparse)
		logger.Debugf(c, "%s\n", myJSONparse["description"])
		return &ImageResponse{
			ImageName:        myJSONparse["name"].(string),
			ImageDescription: myJSONparse["description"].(string),
			ImageUri:         myJSONparse["image"].(string),
		}, nil
	}
	return &ImageResponse{
		ImageName:        req.Image,
		ImageDescription: "Incheon Universe",
		ImageUri:         "incheon universe",
	}, nil
}
