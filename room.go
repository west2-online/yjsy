package yjsy

import (
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/west2-online/yjsy/constants"
	"golang.org/x/net/html"
)

func (s *Student) GetExamRoom(req ExamRoomReq) ([]*ExamRoomInfo, error) {

	res, err := s.SetQueryParams(constants.ExamRoomQueryURL, map[string]string{
		"strwhere": req.Term,
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

		examInfo := &ExamRoomInfo{
			CourseName: strings.TrimSpace(htmlquery.InnerText(columns[1])),
			Credit:     "", // 页面没有学分信息，留空
			Teacher:    "", // 页面没有教师信息，留空
			Date:       "", // 页面没有明确的考试日期，留空
			Time:       strings.TrimSpace(htmlquery.InnerText(columns[4])),
			Location:   strings.TrimSpace(htmlquery.InnerText(columns[5])),
		}

		examInfos = append(examInfos, examInfo)
	}

	return examInfos, nil
}
