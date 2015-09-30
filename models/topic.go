package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"time"
)

func AddTopic(title, category, content string, labels []string) error {
	label := ""
	if len(labels) > 0 {
		label = "$" + strings.Join(labels, "#$") + "#"
	}
	o := orm.NewOrm()
	topic := &Topic{
		Title:    title,
		Category: category,
		Lables:   label,
		Content:  content,
		Created:  time.Now().Format("2006-01-02 15:04:05"),
		Updated:  time.Now().Format("2006-01-02 15:04:05"),
	}

	_, err := o.Insert(topic)
	if err != nil {
		return err
	}
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("id", category).One(cate)
	if err == nil {
		cate.TopicCount++
		_, err = o.Update(cate)
	}
	for i, n := 0, len(labels); i < n; i++ {
		labelOne := new(Label)
		qs := o.QueryTable("label")
		err = qs.Filter("id", labels[i]).One(labelOne)
		if err == nil {
			labelOne.TopicCount++
			_, err = o.Update(labelOne)
		}
	}
	return err
}

func GetAllTopics(category, lable string, isDesc bool) (topics []*Topic, err error) {
	o := orm.NewOrm()
	topics = make([]*Topic, 0)
	qs := o.QueryTable("topic")
	if isDesc {
		if len(category) > 0 {
			qs = qs.Filter("category", category)
		}
		if len(lable) > 0 {
			qs = qs.Filter("lables__contains", "$"+lable+"#")
		}
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

func DelTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	var oldCate string
	var oldLabel string
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		oldCate = topic.Category
		oldLabel = topic.Lables
		_, err = o.Delete(topic)
		if err != nil {
			return err
		}
	}
	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("id", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}

	if len(oldLabel) > 0 {
		label := strings.Replace(strings.Replace(oldLabel, "#", " ", -1), "$", "", -1)
		labels := strings.Split(label, " ")
		for i, n := 0, len(labels); i < n-1; i++ {
			labelOne := new(Label)
			qs := o.QueryTable("label")
			err = qs.Filter("id", labels[i]).One(labelOne)
			if err == nil {
				labelOne.TopicCount--
				_, err = o.Update(labelOne)
			}
		}
	}
	return err
}

func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	topic.Lables = strings.Replace(strings.Replace(
		topic.Lables, "#", " ", -1), "$", "", -1)
	n := len(topic.Lables)
	topic.Lables = topic.Lables[0 : n-1]
	return topic, nil
}

/*func GetTopicLabels(labelsId []string) ([]*Label, error) {
	o := orm.NewOrm()
	Labels := make([]*Label, 0)
	qs := o.QueryTable("label")
	var err error
	for i, n := 0, len(labelsId); i < n; i++ {
		err = qs.Filter("id", labelsId[i]).One(Labels[i])
		if err != nil {
			return nil, err
		}
	}
	return Labels, err
}*/

func EditTopic(tid, title, category, content string, labels []string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	label := ""
	if len(labels) > 0 {
		label = "$" + strings.Join(labels, "#$") + "#"
	}
	var oldCate, oldLabel string
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		oldCate = topic.Category
		oldLabel = topic.Lables
		topic.Title = title
		topic.Category = category
		topic.Lables = label
		topic.Content = content
		topic.Updated = time.Now().Format("2006-01-02 15:04:05")
		_, err = o.Update(topic)
		if err != nil {
			return err
		}
	}
	if len(oldCate) > 0 {
		cate := new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("id", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}
	if len(oldLabel) > 0 {
		label := strings.Replace(strings.Replace(oldLabel, "#", " ", -1), "$", "", -1)
		labels := strings.Split(label, " ")
		for i, n := 0, len(labels); i < n-1; i++ {
			labelOne := new(Label)
			qs := o.QueryTable("label")
			err = qs.Filter("id", labels[i]).One(labelOne)
			if err == nil {
				labelOne.TopicCount--
				_, err = o.Update(labelOne)
			}
		}
	}
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("id", category).One(cate)
	if err == nil {
		cate.TopicCount++
		_, err = o.Update(cate)
	}
	for i, n := 0, len(labels); i < n; i++ {
		labelOne := new(Label)
		qs := o.QueryTable("label")
		err = qs.Filter("id", labels[i]).One(labelOne)
		if err == nil {
			labelOne.TopicCount++
			_, err = o.Update(labelOne)
		}
	}
	return err
}
