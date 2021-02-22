package export

import (
	"encoding/csv"
	"fmt"
	"go-export/internal/conf"
	"go-export/pkg/upload"
	"os"
	"sync"
)

const (
	MaxRow       = 1000000 // excel 最大行数
	SaveStepping = 10
)

// 导出的配置结构
type exportConf struct {
	w        *csv.Writer
	file     *os.File
	lw       *sync.Mutex
	fr       int
	ar       int
	mk       string
	fullName string
}

// 导出任务结构
type exportTask struct {
	task map[string]*exportConf
	l    *sync.Mutex
}

// 导出任务队列
var task = exportTask{
	task: make(map[string]*exportConf),
	l:    &sync.Mutex{},
}

// 初始化导出配置
func InitExportConf(f Form) *exportConf {
	task.l.Lock()
	defer task.l.Unlock()

	if _, ok := task.task[f.HashMark]; !ok {
		filename := conf.Conf.Ept.Path + f.HashMark + ".csv"
		_, err := os.Stat(filename)
		var wHeader = true
		if err == nil || os.IsExist(err) {
			wHeader = false
		}

		// 不存在则创建;存在则追加;读写模式;
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			fmt.Println("open file is failed, err: ", err)
		}

		// 写入UTF-8 BOM，防止中文乱码
		_, _ = file.WriteString("\xEF\xBB\xBF")

		w := csv.NewWriter(file)

		if wHeader {
			_ = w.Write(f.GetHeaderRow())
		}

		ec := &exportConf{
			w:        w,
			file:     file,
			lw:       &sync.Mutex{},
			ar:       f.Total,
			mk:       f.HashMark,
			fullName: filename,
		}

		task.task[f.HashMark] = ec
		InitProgress(f.HashMark, f.Total)
	}

	return task.task[f.HashMark]
}

// 写入表格
func (ec *exportConf) WriteRow(v []string) {
	ec.lw.Lock()
	defer ec.lw.Unlock()

	ec.w.Write(v)
	ec.w.Flush()
	ec.fr++

	// 写入总数已经达到总条数，关闭文件和删除任务
	if ec.fr >= ec.ar || ec.fr >= MaxRow {
		go func() {
			url := upload.ToYun(ec.mk, ec.fullName)
			Finish(ec.mk, url)
		}()
		ec.file.Close()
		delete(task.task, ec.mk)
	} else {
		if ec.fr%SaveStepping == 0 {
			Stepping(ec.mk, ec.fr)
		}
	}
}
