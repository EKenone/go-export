package export

import "fmt"

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
			v := fmt.Sprintf("%v", row[h.Field])
			one = append(one, v)
		}
		list = append(list, one)
	}
	return
}
