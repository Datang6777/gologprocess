package main

import (
	"strings"
	"fmt"
	"time"
)
type Reader interface {
	Read(rc chan string)
}
type Write interface {
	Write(wc chan string)
} 
type LogProcess struct {
	rc chan string
	wc chan string
	read Reader
	write Write
}
type ReadFormFile struct {
	path string //读取文件的路径
}
type WriteToInfluxDB struct {
	influxDBDsn string
}
func (r * ReadFormFile) Read(rc chan string) {
	//读取模块
	line :="message"
	rc <- line
}
//   * 结构体很大，用到了引用，不需要拷贝对象，性能优势
//func (l *LogProcess) ReadFormFile(){
//
//}


func (w *WriteToInfluxDB) Write(wc chan string){
	//写入模块
	fmt.Println(<-wc)
}


func (l *LogProcess) Process(){
	//解析模块
	data := <-l.rc
	l.wc <- strings.ToUpper(data)
}

func (l *LogProcess) WriteToInfluxDB(){

}


func main(){
	r := &ReadFormFile{
		path:"",
	}
	w := &WriteToInfluxDB{
		influxDBDsn:"username&password..",
	}
	lp := &LogProcess{
		rc :make(chan string),
		wc :make(chan string),
		read :r,
		write:w,

	}
	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)
	time.Sleep(1*time.Second)
}
