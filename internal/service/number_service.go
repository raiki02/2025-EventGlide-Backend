package service

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/raiki02/EG/api/req"
	"github.com/raiki02/EG/internal/dao"
	"github.com/raiki02/EG/internal/ioc"
	"github.com/raiki02/EG/internal/model"
	"time"
)

type NumberServiceHdl interface {
	Send(*gin.Context, req.NumberSendReq) error
	Recv() error
	Delete(*gin.Context, req.NumberDelReq) error
	Search(*gin.Context, req.NumberSearchReq) error
}

type NumberService struct {
	nd *dao.NumberDao
	k  *ioc.Kafka
}

func NewNumberService(nd *dao.NumberDao, k *ioc.Kafka) *NumberService {
	ns := &NumberService{
		nd: nd,
		k:  k,
	}
	go ns.Recv()
	return ns
}

func (ns *NumberService) Send(c *gin.Context, req req.NumberSendReq) error {
	var msg = &sarama.ProducerMessage{
		Topic: "notification",
		Value: sarama.ByteEncoder(marshal(req)),
	}
	_, _, err := ns.k.P.SendMessage(msg)
	return err
}

func (ns *NumberService) Recv() error {
	pc, err := ns.k.C.ConsumePartition("notification", 0, sarama.OffsetNewest)
	if err != nil {
		return err
	}

	for {
		select {
		case msg := <-pc.Messages():
			number := unMarshalToPO(msg.Value)
			_ = ns.nd.Insert(&number)
		default:
			time.Sleep(time.Second * 5)
		}
	}
}

func (ns *NumberService) Delete(c *gin.Context, req req.NumberDelReq) error {
	return ns.nd.Delete(c, req.StudentID, req.Object)
}

func (ns *NumberService) Search(c *gin.Context, req req.NumberSearchReq) ([]*model.Number, int, error) {
	return ns.nd.Search(c, req.StudentID, req.Object, req.Action)
}

func marshal(req req.NumberSendReq) []byte {
	b, _ := json.Marshal(req)
	return b
}

func unMarshalToPO(b []byte) model.Number {
	var request req.NumberSendReq
	_ = json.Unmarshal(b, &request)
	return model.Number{
		FromSid:   request.FromSid,
		ToSid:     request.ToSid,
		Object:    request.Object,
		Action:    request.Action,
		CreatedAt: time.Now(),
		IsRead:    false,
		Content:   processContent(request),
	}
}

func processContent(req req.NumberSendReq) string {
	return req.FromSid + " " + req.Action + " your " + req.Object
}
