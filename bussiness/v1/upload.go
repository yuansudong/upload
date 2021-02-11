package v1

import (
	"io"
	"net/http"
	"upload/constant"
	"upload/mempool"
)

// Upload 用于描述一个上传服务
func Upload(rsp http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(1024)
	mReader, mErr := req.MultipartReader()
	if mErr != nil {
		http.Error(rsp, mErr.Error(), http.StatusBadRequest)
		return
	}
	mBuffer := mempool.Get().GetBytes()
	for {
		mPart, mErr := mReader.NextPart()
		if mErr != nil {
			if mErr == io.EOF {
				break
			}
			http.Error(rsp, mErr.Error(), http.StatusNoContent)
			return
		}

		if mPart.FileName() == constant.EmptyString {
			http.Error(rsp, constant.ErrNotFile, http.StatusNoContent)
			return
		}
		for {
			nSize, nErr := mPart.Read(mBuffer)
			if nErr != nil {
				if nErr == io.EOF {
					break
				}
			}

		}

		// mPart.Read(d []byte)
	}
}
