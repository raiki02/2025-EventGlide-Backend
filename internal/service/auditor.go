package service

import (
	"github.com/cqhasy/2025-Muxi-Team-auditor-Backend/api/request"
	"github.com/cqhasy/2025-Muxi-Team-auditor-Backend/api/response"
	"github.com/cqhasy/2025-Muxi-Team-auditor-Backend/sdk/client"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/model"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

var _ dao.AuditorRepository = (*dao.AuditorRepo)(nil)

const (
	SubjectActivity = "activities"
	SubjectPost     = "posts"
)

type AuditorService interface {
	UploadForm(c *gin.Context, creq any, FormId uint) error
	CreateAuditorForm(c *gin.Context, ActId, FormUrl, Sub string) (*model.AuditorForm, error)
}

type auditorService struct {
	ApiKey      string
	HookUrl     string
	MuxiCli     *client.MuxiAuditClient
	AuditorRepo dao.AuditorRepository

	l *zap.Logger
}

func NewAuditorService(repo dao.AuditorRepository, l *zap.Logger) AuditorService {
	httpCli := &http.Client{}
	c := &auditorService{
		ApiKey:      viper.GetString("auditor.apiKey"),
		HookUrl:     viper.GetString("auditor.hookUrl"),
		MuxiCli:     client.NewMuxiAuditClient(httpCli, viper.GetString("auditor.auditUrl")),
		AuditorRepo: repo,
		l:           l.Named("auditor/service"),
	}
	return c
}

func (a *auditorService) UploadForm(c *gin.Context, creq any, id uint) error {
	uploadReq := a.toUploadReq(creq, id)
	_, err := a.MuxiCli.UploadItem(a.ApiKey, uploadReq)
	if err != nil {
		a.l.Error("Upload to auditor failed", zap.Error(err))
		return err
	}
	return nil
}

func (a *auditorService) CreateAuditorForm(c *gin.Context, ActId, FormUrl string, sub string) (*model.AuditorForm, error) {
	return a.AuditorRepo.Insert(c, ActId, FormUrl, sub)
}

func (a *auditorService) toUploadReq(creq any, id uint) request.UploadReq {
	res := request.UploadReq{
		HookUrl:    a.HookUrl,
		Id:         id,
		Tags:       []string{"校灵通"},
		PublicTime: time.Now().Unix(),
	}

	switch creq.(type) {
	case req.CreateActReq:
		actReq := creq.(req.CreateActReq)

		res.Author = extractAuthors(actReq.LabelForm.Signer)
		res.Tags = append(res.Tags, "活动")
		res.Content = response.Contents{
			Topic: response.Topics{
				Title:    actReq.Title,
				Content:  actReq.Introduce,
				Pictures: actReq.ShowImg,
			},
		}
		if actReq.LabelForm.IfRegister == "是" {
			res.Tags = append(res.Tags, "含报名表需要审核")
			res.Content.Topic.Pictures = append(res.Content.Topic.Pictures, actReq.LabelForm.ActiveForm)
		}

	case req.CreatePostReq:
		postReq := creq.(req.CreatePostReq)

		res.Author = postReq.StudentID
		res.Tags = append(res.Tags, "帖子")
		res.Content = response.Contents{
			Topic: response.Topics{
				Title:    postReq.Title,
				Content:  postReq.Introduce,
				Pictures: postReq.ShowImg,
			},
		}
	}
	return res
}

func extractAuthors(signers []struct {
	StudentID string `json:"studentid"`
	Name      string `json:"name"`
}) string {
	builder := strings.Builder{}
	for _, s := range signers {
		builder.WriteString(s.Name + "-")
	}
	return builder.String()
}
