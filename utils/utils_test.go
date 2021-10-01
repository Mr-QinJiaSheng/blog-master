package utils

import (
	"fmt"
	"testing"

	"github.com/astaxie/beego"
	"github.com/stretchr/testify/require"
)

func TestUtilsGetTopic(t *testing.T) {
	topic, content := GetTopic("#asd#000")
	require.EqualValues(t, topic, "#asd#")
	require.EqualValues(t, content, "000")

	topic, content = GetTopic("#你好#000")
	require.EqualValues(t, topic, "#你好#")
	require.EqualValues(t, content, "000")

	topic, content = GetTopic("你#好#000")
	require.EqualValues(t, topic, "")
	require.EqualValues(t, content, "你#好#000")

	topic, content = GetTopic("#你好#000#")
	require.EqualValues(t, topic, "#你好#")
	require.EqualValues(t, content, "000#")
}

func TestUtilsDownImage(t *testing.T) {

	fmt.Println(beego.WorkPath)
	path, err := DownImage("https://img9.doubanio.com/view/subject/s/public/s3745215.jpg")
	if err != nil {
		fmt.Println(err.Error())
	}
	require.EqualValues(t, path, "/var/local/go-blog/")
}
