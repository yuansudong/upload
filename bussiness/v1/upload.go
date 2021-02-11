package v1

import (
	"net/http"
	"upload/constant"
)

// Upload 用于描述一个上传服务
func Upload(rsp http.ResponseWriter, req *http.Request) {
	err := req.ParseMultipartForm(constant.ParseMultipartFormMemoryLimit)

	if err != nil {

	}
}
