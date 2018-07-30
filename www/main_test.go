package main

import (
	"testing"
)

func TestIsUrl(t *testing.T) {
	t.Log(IsUrl("ifth"))
	t.Log(IsUrl("ifth.net"))
	t.Log(IsUrl("http://ifth.net"))
	t.Log(IsUrl("http://ifth.zz"))
	t.Log(IsUrl("http://www.ifth.net"))
	t.Log(IsUrl("http://www.test.ifth.net"))
	t.Log(IsUrl("http://www.ifth.net/test=3&test2=4&test-3[]=h_h"))
	t.Log(IsUrl("https://www.zhizhu-inc.com"))
}
