package strings

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

// ReplaceAll replace string
func ReplaceAll(oldStr, pattern, newStr string) string {
	p, _ := regexp.Compile(pattern)
	s := p.ReplaceAll([]byte(oldStr), []byte(newStr))
	return string(s)
}

// 自带方法补全html
func fixHTML(result string) string {
	// 取出所有标签
	tempResult := ReplaceAll(result, "(>)[^<>]*(<?)", "$1$2") // 把标签中间的所有内容都去掉了

	// 2. 去掉不需要结素标记的HTML标记
	tempResult = ReplaceAll(tempResult, "</?(embed|AREA|BASE|BASEFONT|BR|COL|COLGROUP|DD|DT|FRAME|HEAD|HR|IMG|INPUT|ISINDEX|LI|LINK|META|OPTION|PARAM|area|wbr|br|col|colgroup|dd|dt|frame|hr|img|input|isindex|link|meta|param)[^<>]*/?>", "")

	// 把<div class=xxx的class=xxx去掉
	tempResult = ReplaceAll(tempResult, "<(/?[a-zA-Z]+)[^<>]*>", "<$1>")

	// 3 只能用正则,+stack来去有结束的
	// golang的正则暂不支持back reference, 以后可以用它来去掉重复的标签
	p, _ := regexp.Compile("<(/?[a-zA-Z]+)[^<>]*>") // 得到所有的<div>, </div>...
	strs := p.FindAllString(tempResult, -1)

	//	fmt.Println(strs)
	stack := make([]string, len(strs))
	stackP := -1
	for _, each := range strs {
		if stackP >= 0 {
			// 匹配
			if stack[stackP][1:] == each[2:] {
				stackP--
				continue
			}
		}
		stackP++
		stack[stackP] = each
	}
	// 补全tag
	if stackP != -1 {
		fmt.Println(stack[0 : stackP+1])

		for _, each := range stack[0 : stackP+1] {
			if each[1] != '/' {
				result += "</" + each[1:]
			}
		}
	}

	return result
}

// SubStringHTML 获取摘要, HTML
func SubStringHTML(param string, length int, end string) string {
	if param == "" {
		return param
	}
	result := ""

	rStr := []rune(param)
	lenStr := len(rStr)

	if lenStr <= length {
		result = param
	} else {
		// 1
		n := 0
		var temp rune   // 中文问题, 用rune来解决
		isCode := false //是不是HTML代码
		isHTML := false //是不是HTML特殊字符,如&nbsp;
		var i = 0
		for ; i < lenStr; i++ {
			temp = rStr[i]
			if temp == '<' {
				isCode = true
			} else if temp == '&' {
				isHTML = true
			} else if temp == '>' && isCode {
				// n = n - 1
				isCode = false
			} else if temp == ';' && isHTML {
				isHTML = false
			}
			if !isCode && !isHTML {
				n = n + 1
			}
			// 每一次都相加, 速度非常慢!, 重新分配内存, 7倍的差距
			// result += string(temp)
			if n >= length {
				break
			}
		}

		result = string(rStr[0:i])

		if end != "" {
			result += end
		}
	}

	// 使用goquery来取出html, 为了补全html
	htmlReader := bytes.NewBufferString(result)
	dom, err1 := goquery.NewDocumentFromReader(htmlReader)
	if err1 == nil {
		html, _ := dom.Html()
		html = strings.Replace(html, "<html><head></head><body>", "", 1)
		html = strings.Replace(html, "</body></html>", "", 1)

		// TODO 把style="float: left"去掉
		return html

		// 如果有错误, 则使用自己的方法补全, 有风险
	}
	return fixHTML(result)
}
