// go
package yjsy

import (
	"github.com/west2-online/yjsy/constants"
)

// GetStudentInfo 请求学生信息页面并解析为 HTML 文档后提取学生信息
func (s *Student) GetStudentInfo() (map[string]string, error) {
	resp, err := s.GetWithIdentifier(constants.UserInfoURL, nil)
	if err != nil {
		return nil, err
	}

	// 解析学生信息
	// 以下 xpath 表达式请根据实际响应 HTML 结构调整
	xpaths := map[string]string{
		"stu_id":   "//*[@id='xxTable']/table/tbody/tr[1]/td[2]",
		"name":     "//*[@id='xxTable']/table/tbody/tr[1]/td[4]",
		"birthday": "//*[@id='xxTable']/table/tbody/tr[3]/td[4]",
		"sex":      "//*[@id='xxTable']/table/tbody/tr[2]/td[4]",
		"college":  "//*[@id='xxTable']/table/tbody/tr[15]/td[2]",
		"grade":    "//*[@id='xxTable']/table/tbody/tr[14]/td[4]",
		"major":    "//*[@id='xxTable']/table/tbody/tr[15]/td[4]",
	}

	result := make(map[string]string)
	for key, xp := range xpaths {
		result[key] = safeExtractHTMLFirst(resp, xp)
	}
	return result, nil
}
