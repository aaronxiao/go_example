package templateMethod

import "fmt"

//使用继承机制，把通用步骤和通用方法放到父类中，把具体实现延迟到子类中实现。使得实现符合开闭原则

type Downloader interface {
	Download(uri string)
}

type implement interface {
	download()
	save()
}

//
type template struct {
	implement
	uri string
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

var _ implement = new(template)

func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

func (t *template) save() { //实现了一个通过save方法
	fmt.Print("default save\n")
}

type HTTPDownloader struct {
	*template
}

func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	template1 := newTemplate(downloader)
	downloader.template = template1
	return downloader
}

func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

func (*HTTPDownloader) save() {
	fmt.Printf("http save\n")
}

type FTPDownloader struct {
	*template //指针类型也可以继承save方法
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}

//注意这里只实现的download方法  save用template继承过来的
