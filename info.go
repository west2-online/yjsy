// go
package yjsy

import (
    "fmt"
    "bytes"
    "strings"

    "github.com/antchfx/htmlquery"
    "github.com/west2-online/yjsy/errno"
    "golang.org/x/net/html"
)

// GetStudentInfo 请求学生信息页面并解析为 HTML 文档后提取学生信息
func (s *Student) GetStudentInfo() (map[string]string, error) {
    // 设置 Referer 为页面中的 xs_main.htm
    req := s.NewRequest().SetHeader("Referer", "https://yjsglxt.fzu.edu.cn/xs_main.htm")
    resp, err := req.Get("https://yjsglxt.fzu.edu.cn/xsgl/xsxx_show.aspx")
    if err != nil {
        return nil, err
    }

    // 如果响应中包含要求重新登录，则返回 Cookie 错误
    if strings.Contains(string(resp.Body()), "重新登录") {
        return nil, errno.CookieError
    }

    doc, err := htmlquery.Parse(bytes.NewReader(resp.Body()))
    if err != nil {
        return nil, err
    }
    
    // 解析学生信息
    details, err := ExtractStudentDetails(doc)
    if err != nil {
        fmt.Println("解析学生信息失败:", err)
    }

    return details, nil
}

// ExtractStudentDetails 通过 xpath 从 HTML 文档中提取学生信息
func ExtractStudentDetails(doc *html.Node) (map[string]string, error) {
    // 以下 xpath 表达式请根据实际响应 HTML 结构调整
    xpaths := map[string]string{
        "stu_id":  "//*[@id='xxTable']/table/tbody/tr[1]/td[2]",
        "name":  "//*[@id='xxTable']/table/tbody/tr[1]/td[4]",
        "birthday":  "//*[@id='xxTable']/table/tbody/tr[3]/td[4]",
        "sex":  "//*[@id='xxTable']/table/tbody/tr[2]/td[4]",
        "college":  "//*[@id='xxTable']/table/tbody/tr[15]/td[2]",
        "grade":  "//*[@id='xxTable']/table/tbody/tr[14]/td[4]",
        "major":  "//*[@id='xxTable']/table/tbody/tr[15]/td[4]",
    }

    result := make(map[string]string)
    for key, xp := range xpaths {
        node := htmlquery.FindOne(doc, xp)
        if node == nil {
            return nil, fmt.Errorf("无法找到 '%s' 对应的节点", key)
        }
        result[key] = htmlquery.InnerText(node)
		fmt.Printf("key: %s \t value:%s\n", key,result[key])
    }
    return result, nil
}