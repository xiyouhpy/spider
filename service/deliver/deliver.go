package deliver

import (
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// downloadTimeOut 下载超时设置
const downloadTimeOut = time.Second * 60

// Download 下载url的图片，返回下载文件名
func Download(strURL string, dstPath string) error {
	client := new(http.Client)
	client.Timeout = downloadTimeOut
	rsp, rspErr := client.Get(strURL)
	if rspErr != nil {
		logrus.Warnf("client.Get err, url:%s, err:%s", strURL, rspErr.Error())
		return rspErr
	}
	defer rsp.Body.Close()

	if !strings.Contains(dstPath, ".") {
		dstPath += ".jpg"
	}
	file, fileErr := os.Create(dstPath)
	if fileErr != nil {
		logrus.Warnf("os.Create err, file:%s, err:%s", dstPath, fileErr.Error())
		return fileErr
	}
	defer file.Close()

	_, copyErr := io.Copy(file, rsp.Body)
	if copyErr != nil {
		logrus.Warnf("os.Copy err, err:%s", copyErr.Error())
		return copyErr
	}

	return nil
}
