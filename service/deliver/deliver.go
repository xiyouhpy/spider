package deliver

import (
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/xiyouhpy/spider/base"
)

// httpReq http 请求参数
type httpReq struct {
	strMethod string
	strURL    string
	strParams string
	arrHeader map[string]string
	timeOut   time.Duration
}

// download 下载数据
func (r *httpReq) download(dstPath string) error {
	var tmpData io.Reader
	if r.strParams != "" {
		tmpData = strings.NewReader(r.strParams)
	}

	var httpReq *http.Request
	httpReq, _ = http.NewRequest(r.strMethod, r.strURL, tmpData)
	for key, value := range r.arrHeader {
		httpReq.Header.Add(key, value)
	}
	strCookie := &http.Cookie{}
	httpReq.AddCookie(strCookie)

	client := &http.Client{Timeout: r.timeOut}
	rsp, rspErr := client.Do(httpReq)
	if rspErr != nil {
		logrus.Warnf("client.Get err, url:%s, err:%s", r.strURL, rspErr.Error())
		return rspErr
	}
	defer rsp.Body.Close()

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

// DownloadImage 下载图片，返回下载文件名
func DownloadImage(strURL string, dstPath string) error {
	if strURL == "" || dstPath == "" {
		logrus.Warnf("params is empty, url:%s, path:%s", strURL, dstPath)
		return base.ErrParamsError
	}
	if !strings.Contains(dstPath, ".") {
		dstPath += ".jpg"
	}

	arrHeader := map[string]string{}
	req := &httpReq{"GET", strURL, "", arrHeader, time.Second * 60}
	return req.download(dstPath)
}

// DownloadHtml 下载网页，返回下载文件名
func DownloadHtml(strURL string, dstPath string) error {
	if strURL == "" || dstPath == "" {
		logrus.Warnf("params is empty, url:%s, path:%s", strURL, dstPath)
		return base.ErrParamsError
	}
	if !strings.Contains(dstPath, ".") {
		dstPath += ".html"
	}

	arrHeader := map[string]string{}
	req := &httpReq{"GET", strURL, "", arrHeader, time.Second * 60}
	return req.download(dstPath)
}
