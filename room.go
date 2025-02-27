package yjsy

import (
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/west2-online/yjsy/constants"
	"golang.org/x/net/html"
)

func (s *Student) GetExamRoom(req ExamRoomReq) ([]*ExamRoomInfo, error) {

	res, err := s.GetWithIdentifier(constants.ExamRoomQueryURL, map[string]string{
		"strwhere": "XQXN='" + req.Term + "'",
	})
	if err != nil {
		return nil, err
	}
	examInfos, err := parseExamRoom(res)
	if err != nil {
		return nil, err
	}
	return examInfos, nil
}
func parseExamRoom(doc *html.Node) ([]*ExamRoomInfo, error) {
	var examInfos []*ExamRoomInfo

	table := htmlquery.FindOne(doc, "//div[@id='divContent']/table")
	rows := htmlquery.Find(table, "//tr[position()>1]")

	for _, row := range rows {
		columns := htmlquery.Find(row, "./td")

		dateTime := strings.Fields(htmlquery.InnerText(columns[4]))
		if len(dateTime) == 0 {
			continue
		}
		// 如果为空，说明目前没有安排考试

		date := strings.Replace(dateTime[0], "/", "-", 2)
		time := dateTime[1]
		examInfo := &ExamRoomInfo{
			CourseName: strings.TrimSpace(htmlquery.InnerText(columns[1])),
			Credit:     "", // 页面没有学分信息，留空
			Teacher:    "", // 页面没有教师信息，留空
			Date:       date,
			Time:       time,
			Location:   strings.TrimSpace(htmlquery.InnerText(columns[5])),
		}

		examInfos = append(examInfos, examInfo)
	}

	return examInfos, nil
}
