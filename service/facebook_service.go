package service

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
)

var (
	FacebookUri = "https://graph.facebook.com/v2.9/me?access_token=%s&fields=name,email,picture,first_name,last_name,link&method=get&pretty=0&sdk=joey&suppress_http_code=1"
)

func GetFacebookInfo(accessToken string) ([]byte, int) {
	uri := fmt.Sprintf(FacebookUri, accessToken)
	request := httplib.NewBeegoRequest(uri, "GET")
	return doRequest(request)
}

func doRequest(request *httplib.BeegoHTTPRequest) (body []byte, status int) {
	resp, err := request.DoRequest()
	if err == nil {
		status = resp.StatusCode
		if resp.Body != nil {
			body, _ = ioutil.ReadAll(resp.Body)
		}
	} else {
		status = http.StatusBadRequest
		glog.Error(err.Error())
	}
	return body, status
}
