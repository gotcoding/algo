package stack

import "testing"

func TestBrowser(t *testing.T) {
	b := NewBrowser()
	b.NewPage("www.baidu.com")
	b.NewPage("www.huawei.com")
	b.NewPage("www.zhihu.com")
	if b.CanBack() {
		b.Back()
	}
	if b.CanForward() {
		b.Forward()
	}
	if b.CanBack() {
		b.Back()
	}
	if b.CanBack() {
		b.Back()
	}
	if b.CanBack() {
		b.Back()
	}
	b.Open("www.qq.com")
	if b.CanForward() {
		b.Forward()
	}
}
