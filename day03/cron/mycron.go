package cron

import (
	"github.com/davecgh/go-spew/spew"
	"log"
	"sync"
	"time"
)

//MyCron 对象
type MyCron struct {
	entries   []*Entry       //工作集合
	running   bool           //是否在工作
	addEntry  chan *Entry    //运行中接受AddEntry的信号
	stop      chan bool      //接受停止信号
	runningMu sync.Mutex     //Cron 运行前需要抢占该锁，保证并发安全
	location  *time.Location //时区
	nextID    int            //下一个加入的Entry的ID
	wait      sync.WaitGroup
}

// Entry 每个任务
type Entry struct {
	EntryID int       //唯一标识符
	Next    time.Time //下次工作时间
	f       function  //用户传入的函数
	Cycle   Cycle     //函数工作周期
	remove  bool      //标记是否被删除
}

//Cycle 接口
type Cycle interface {
	//Next 获取当前Entry下次工作时间
	Next(time.Time) time.Time
}

// ParseOpt 解析器接口
type ParseOpt interface {
	// Parse 解析方法
	Parse(string) Cycle
}

//function 接口 用户自定义函数，即工作函数
type function interface {
	RUN()
}

//工作函数
type gotfunc func()

// RUN 实现function接口
func (fc gotfunc) RUN() { fc() }

// NewCron 新生成一个MyCron
func NewCron() *MyCron {
	c := &MyCron{
		entries:   nil,
		running:   false,
		addEntry:  nil,
		stop:      nil,
		runningMu: sync.Mutex{},
		location:  nil,
		nextID:    0,
	}
	return c
}

// AddEntry 新增任务
func (c *MyCron) AddEntry(cocy string, fc gotfunc) (int, error) {
	//解析
	cycle, err := c.Parse(cocy)
	if err != nil {
		log.Printf("%s\n解析错误", cocy)
		return -1, err
	}
	//获取锁
	c.runningMu.Lock()
	defer c.runningMu.Unlock()
	//设置ID
	c.nextID++
	//设置任务
	entry := &Entry{
		EntryID: c.nextID,
		f:       fc,
		Cycle:   cycle,
		remove:  false,
	}

	//检查MyCron运行状态
	//如果在运行,发给addEntry 否则追加到工作集合
	if !c.running {
		c.entries = append(c.entries, entry)
	} else {
		c.addEntry <- entry
	}
	//日志
	spew.Dump("add,success")
	return entry.EntryID, nil
}

// RemoveEntry 根据ID删除作业
func (c *MyCron) RemoveEntry(EntryId int) {
	//获取锁
	c.runningMu.Lock()
	defer c.runningMu.Unlock()
	//删除
	c.entries[EntryId].remove = true
	c.entries[EntryId].Next = time.Now().Add(99999 * time.Hour)
	//日志
	log.Printf("removed entyr :")
	spew.Dump(c.entries[EntryId])
}

// Stop 停止MyCron
func (c *MyCron) Stop() {
	//获取 锁
	c.runningMu.Lock()
	defer c.runningMu.Unlock()
	//如果在运行 发送停止信号
	if c.running {
		c.stop <- true
		c.running = false
	}
}

// Start 运行MyCron
func (c *MyCron) Start() {
	//获取锁
	c.runningMu.Lock()
	defer c.runningMu.Unlock()
	//在运行直接返回
	if c.running {
		return
	}
	c.running = true
	log.Printf("start cron")
	//
	c.wait.Add(1)
	go c.run()
	c.wait.Wait()
}

// SortEntry 把任务按时间排序
func (c *MyCron) SortEntry() {
	c.qsort(0, c.nextID-1)
}

//自定义排序
func (c *MyCron) qsort(l, r int) {
	if l >= r {
		return
	}
	var x, y int
	t := c.entries[(l+r+1)/2].Next
	x = l - 1
	y = r + 1

	for {
		if x >= y {
			break
		}
		//do-while
		y--
		for {
			if c.entries[y].Next.After(t) {
				y--
			} else {
				break
			}
		}
		x++
		for {
			if c.entries[x].Next.Before(t) {
				x++
			} else {
				break
			}
		}

		if x < y {
			tmp := c.entries[x].Next
			c.entries[x].Next = c.entries[y].Next
			c.entries[y].Next = tmp
		}
	}

	c.qsort(l, x-1)
	c.qsort(x, r)
}

//运行
func (c *MyCron) run() {
	now := time.Now()
	log.Println(now)
	//遍历c.entries 列表，通过 cycle.Next() 计算出这个任务下一次执行的时间，并赋值给entry.Next 字段
	for _, entry := range c.entries {
		entry.Next = entry.Cycle.Next(now)
	}

	for {
		//把作业按照时间按顺序排序 被删除的一定在最后
		c.SortEntry()
		//初始化定时器
		var timer *time.Timer
		if !c.entries[0].Next.IsZero() {
			timer = time.NewTimer(c.entries[0].Next.Sub(now))
		}
		//
		for {
			select {
			//定时器触发信号 触发后跑函数 再重新排序
			case now = <-timer.C:
				c.entries[0].f.RUN()
				c.entries[0].Next = c.entries[0].Cycle.Next(now)
			case newEntry := <-c.addEntry:
				//先停止计时
				timer.Stop()
				//追加
				c.entries = append(c.entries, newEntry)
			case <-c.stop:
				timer.Stop()
				log.Printf("stopped")
				c.wait.Done()
				return
			}
			break
		}
	}
}
