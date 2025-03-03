package metadata

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"biyard.co/common/logger"
	"biyard.co/common/utils/httpclient"
)

type MetadataAPI struct {
	ctx            context.Context
	metadataCli    *httpclient.HTTPClient
	token          string
	owner          string
	repo           string
	basePath       string
	committerName  string
	committerEmail string
}

type MetadataRequest struct {
	Message   string        `json:"message"`
	Committer CommitterInfo `json:"committer"`
	Content   string        `json:"content"`
}

type CommitterInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MetadataResponse struct {
	Code        int
	MetadataURI string
}

func NewMetadataAPI(ctx context.Context, token string, owner, repo, basePath, committerName, committerEmail string) *MetadataAPI {
	t := fmt.Sprintf("Bearer %+v", token)
	uri := fmt.Sprintf("https://api.github.com/repos/%+v/%+v/contents/%+v", owner, repo, basePath)
	metadataCli, err := httpclient.NewHttpClient(httpclient.ConstantBackoff(3, 1*time.Second), httpclient.Limiter(1000), httpclient.Headers(http.Header{
		"Authorization":        []string{t},
		"Content-Type":         []string{"application/json"},
		"Accept":               []string{"application/vnd.github+json"},
		"X-GitHub-Api-Version": []string{"2022-11-28"},
	}), httpclient.Endpoint(uri))

	if err != nil {
		logger.Criticalf(nil, "NewHttpClient error(get metadata): %+v", err)
	}

	return &MetadataAPI{
		ctx:            ctx,
		metadataCli:    metadataCli,
		token:          t,
		owner:          owner,
		repo:           repo,
		basePath:       basePath,
		committerName:  committerName,
		committerEmail: committerEmail,
	}
}

func (r *MetadataAPI) UploadMetadata(ctx context.Context, fileName string, content string, path string) *MetadataResponse {
	var jsonData any
	code, e := r.metadataCli.Put(ctx, path, &MetadataRequest{
		Message: "image upload",
		Committer: CommitterInfo{
			Name:  r.committerName,
			Email: r.committerEmail,
		},
		Content: content,
	}, &jsonData)

	if e != nil {
		logger.Errorf(ctx, "request data: %+v, err: %+v", jsonData, e)
		return &MetadataResponse{
			Code:        code,
			MetadataURI: "",
		}
	}

	metadataURI := fmt.Sprintf("https://%+v/%+v/%+v", r.repo, r.basePath, fileName)
	logger.Debugf(ctx, "code: %+v", code)

	return &MetadataResponse{
		Code:        code,
		MetadataURI: metadataURI,
	}
}
