package service

import (
	"time"
	"web_demo/dao"
	"web_demo/model"
)

type TopicService struct {
}

var TopicDao = new(dao.TopicDao)

func (p *TopicService) Insert(title, topicspreview, topicsinfo string) int64 {
	return TopicDao.Insert(&model.Topic{Title: title, TopicsPreview: topicspreview, TopicsInfo: topicsinfo, CreateTime: time.Now(), ModifyTime: time.Now()})
}

func (p *TopicService) SelectTopicById(TopicId string) []model.Topic {
	return TopicDao.SelectTopicById(TopicId)
}

func (p *TopicService) SelectAllTopic() []model.Topic {
	return TopicDao.SelectAllTopic()
}

func (p *TopicService) TopicDeleteById(id string) int64 {
	return TopicDao.TopicDeleteById(id)
}

func (p *TopicService) TopicModifyB(title, topics_preview, topics_info, id string) int64 {
	return TopicDao.TopicModifyB(title, topics_preview, topics_info, id)
}

func (p *TopicService) LimitList(pagesize int, pageno int) []model.Topic {
	return TopicDao.LimitList(pagesize, pageno)
}

func (p *TopicService) GetDataNum() int64 {
	return TopicDao.GetDataNum()
}
