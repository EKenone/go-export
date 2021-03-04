package export

import (
	"fmt"
	"strings"
)

type HeaderField struct {
	Field string
	Title string
}

type Form struct {
	HashMark string                   `json:"hash_mark"`
	Header   []HeaderField            `json:"header"`
	Data     []map[string]interface{} `json:"data"`
	Total    int                      `json:"total"`
}

// 获取表头
func (f Form) GetHeaderRow() (header []string) {
	for _, v := range f.Header {
		header = append(header, v.Title)
	}
	return
}

// 获取行数据
func (f Form) GetExportList() (list [][]string) {
	for _, row := range f.Data {
		var one []string
		for _, h := range f.Header {

			// 看看字段有多深
			deepLen := strings.Count(h.Field, ".")
			v, dr := "", row

			if deepLen > 0 {
				deep := strings.Split(h.Field, ".")
				for k, df := range deep {

					// 如果当前深度的行没有想要的值，直接返回空数据
					if dr == nil {
						v = ""
						break
					}

					// 没到最深的那层数据
					if k != deepLen {
						switch dr[df].(type) {
						case map[string]interface{}:
							dr = dr[df].(map[string]interface{})
						default:
							dr = nil
						}
						continue
					}

					// 最深数据的时候直接取值
					v = fmt.Sprintf("%v", dr[df])

				}
			} else {
				v = fmt.Sprintf("%v", dr[h.Field])
			}
			one = append(one, v)

		}
		list = append(list, one)
	}
	return
}
