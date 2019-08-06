package dao

import (
	"fmt"
	"strconv"
	"web_demo/framework"
	"web_demo/model"
)

type TopicDao struct {
}

func (p *TopicDao) Insert(Topic *model.Topic) int64 {
	result, err := framework.DB.Exec("INSERT INTO  `topics` (`title`,`topics_preview`,`topics_info`,`create_time`,`modify_time`) value(?,?,?,?,?)", Topic.Title, Topic.TopicsPreview, Topic.TopicsInfo, Topic.CreateTime, Topic.ModifyTime)
	if err != nil {
		fmt.Println("Insert error")
		return 0
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Insert error")
		return 0
	}
	return id
}

func (p *TopicDao) SelectTopicById(TopId string) []model.Topic {
	rows, err := framework.DB.Query("SELECT * FROM topics WHERE id = ?", TopId)
	if err != nil {
		fmt.Println("SelectTopicById error")
		fmt.Println(err)

		return nil
	}
	var Topics []model.Topic
	for rows.Next() {
		var Topic model.Topic
		err := rows.Scan(&Topic.ID, &Topic.Title, &Topic.TopicsPreview, &Topic.TopicsInfo, &Topic.CreateTime, &Topic.ModifyTime)
		if err != nil {
			fmt.Println("SelectTopicById error")
			continue
		}
		Topics = append(Topics, Topic)
	}
	rows.Close()
	return Topics
}

func (p *TopicDao) SelectAllTopic() []model.Topic {
	rows, err := framework.DB.Query("SELECT * FROM topics")
	if err != nil {
		fmt.Println("SelectAllTopic error")
		return nil
	}
	var Topics []model.Topic
	for rows.Next() {
		var Topic model.Topic
		err := rows.Scan(&Topic.ID, &Topic.Title, &Topic.TopicsPreview, &Topic.TopicsInfo, &Topic.CreateTime, &Topic.ModifyTime)
		if err != nil {
			fmt.Println("SelectAllTopic error")
			continue
		}
		Topics = append(Topics, Topic)
	}
	rows.Close()
	return Topics
}
func (p *TopicDao) TopicDeleteById(id string) int64 {
	stmt, err := framework.DB.Prepare("DELETE FROM topics WHERE id=?")
	if err != nil {
		fmt.Println("TopicDeleteById err")
	}
	res, err := stmt.Exec(id)
	if err != nil {
		fmt.Println("TopicDeleteById err")
	}
	num, err := res.RowsAffected()
	if err != nil {
		fmt.Println("TopicDeleteById err")
	}

	fmt.Println(num)
	stmt.Close()
	return num
}

func (p *TopicDao) TopicModifyB(title, topics_preview, topics_info, id string) int64 {
	stmt, err := framework.DB.Prepare("UPDATE   topics SET title  =?, topics_preview  =?, topics_info  = ? WHERE  id  = ? ;")
	if err != nil {
		fmt.Println("TopicDeleteById err")
	}
	res, err := stmt.Exec(title, topics_preview, topics_info, id)
	if err != nil {
		fmt.Println("TopicDeleteById err")
	}
	num, err := res.RowsAffected()
	if err != nil {
		fmt.Println("TopicDeleteById err")
	}

	fmt.Println(num)
	stmt.Close()
	return num
}

func (p *TopicDao) LimitList(pagesize int, pageno int) []model.Topic {

	rows, err := framework.DB.Query("SELECT * FROM topics order by id desc  limit  " + strconv.Itoa((pageno-1)*pagesize) + " , " + strconv.Itoa(pagesize))
	if err != nil {
		fmt.Println("LimitList error")
		return nil
	}
	var Topics []model.Topic
	for rows.Next() {
		var Topic model.Topic
		err := rows.Scan(&Topic.ID, &Topic.Title, &Topic.TopicsPreview, &Topic.TopicsInfo, &Topic.CreateTime, &Topic.ModifyTime)
		if err != nil {
			fmt.Println("LimitList error")
			continue
		}
		Topics = append(Topics, Topic)
	}
	rows.Close()
	return Topics
}

func (p *TopicDao) GetDataNum() int64 {

	rows, err := framework.DB.Query("SELECT count(id) FROM topics")
	if err != nil {
		fmt.Println("GetDataNum error")
		//return nil
	}
	var num int64
	for rows.Next() {
		err := rows.Scan(&num)
		if err != nil {
			fmt.Println("GetDataNum error")
		}
	}
	//	fmt.Println(num)
	return num
}
