package export

import (
	"encoding/csv"
	"fmt"
	"go-export/internal/conf"
	"math"
	"os"
	"sync"
	"time"
)

const (
	MaxRow   = 1000000 // excel 最大行数
	StepPart = 50      // 把数据分成50份进行步进更新，保证每个任务最多50次步进更新
	FlushNum = 10      // 从缓冲区写入文件行数
)

// 导出的配置结构
type exportConf struct {
	w        *csv.Writer
	file     *os.File
	lw       sync.Mutex
	fr       int
	ar       int
	sp       int
	mk       string
	fullName string
	gc       chan bool
}

// 导出任务结构
type exportTask struct {
	task map[string]*exportConf
	l    sync.Mutex
}

// 导出任务队列
var task = exportTask{
	task: make(map[string]*exportConf),
	l:    sync.Mutex{},
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

		sp := math.Ceil(float64(f.Total) / float64(StepPart))

		ec := &exportConf{
			w:        w,
			file:     file,
			lw:       sync.Mutex{},
			ar:       f.Total,
			sp:       int(sp),
			mk:       f.HashMark,
			fullName: filename,
			gc:       make(chan bool, 1),
		}

		task.task[f.HashMark] = ec

		// 初始化进度条
		InitProgress(f.HashMark, f.Total)

		// 自动回收协程
		go ec.autoGCTask()
	}

	return task.task[f.HashMark]
}

// 写入表格
func (ec *exportConf) WriteRow(v []string) {
	ec.lw.Lock()
	defer ec.lw.Unlock()

	ec.fr++
	ec.w.Write(v)

	if ec.fr%FlushNum == 0 || ec.fr >= ec.ar { // 到达缓冲区最大实现条数或者写入了最后一条
		ec.w.Flush()
	}

	// 写入总数已经达到总条数，关闭文件和删除任务
	if ec.fr >= ec.ar || ec.fr >= MaxRow {
		ec.taskEnd()
	} else {
		if ec.fr%ec.sp == 0 {
			Stepping(ec.mk, ec.fr)
		}
	}
}

// 自动回收任务
func (ec *exportConf) autoGCTask() {
	for {
		select {
		case <-ec.gc:
			return
		case <-time.After(5 * time.Minute):
			ec.taskEnd()
		}
	}
}

// 任务结束
func (ec *exportConf) taskEnd() {
	ec.file.Close()
	delete(task.task, ec.mk)
	ec.gc <- true
	go func() {
		url := toYun(ec.mk, ec.fullName)
		Finish(ec.mk, url)
	}()
}
