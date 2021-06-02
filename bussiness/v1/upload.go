package v1

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"upload/constant"
	"upload/mempool"
)

// Upload 用于描述一个上传服务
func Upload(rsp http.ResponseWriter, req *http.Request) {
	mReader, mErr := req.MultipartReader()
	if mErr != nil {
		http.Error(rsp, mErr.Error(), http.StatusBadRequest)
		return
	}
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
		log.Println(mPart.FileName(), mPart.FormName(), mPart.Header)
		mFileHandler, fErr := os.OpenFile(
			"./"+mPart.FileName(),
			os.O_RDWR|os.O_CREATE|os.O_TRUNC,
			0755,
		)
		if fErr != nil {
			http.Error(
				rsp,
				fErr.Error(),
				http.StatusServiceUnavailable,
			)
			return
		}
		tTotalSize, tErr := _Copy(mPart, mFileHandler)
		if tErr != nil {
			mFileHandler.Close()
			mPart.Close()
			http.Error(
				rsp,
				fErr.Error(),
				http.StatusServiceUnavailable,
			)
		}
		log.Println(tTotalSize)
		log.Println(mFileHandler.Close())
		log.Println(mPart.Close())
	}
}

// _Copy 用于拷贝
func _Copy(r io.Reader, w io.Writer) (oTotalSize int, oErr error) {
	mBuffer := mempool.Get().GetBytes()
	for {
		rSize, nErr := r.Read(mBuffer)
		if rSize > 0 {
			wSize, wErr := w.Write(mBuffer[0:rSize])
			if wErr != nil {
				oErr = wErr
				goto end
			}
			if wSize != rSize {
				oErr = errors.New(constant.ErrRWNotEQ)
				goto end
			}
			oTotalSize += wSize
		}
		if nErr != nil {
			if nErr != io.EOF {
				oErr = nErr
			}
			goto end
		}
	}
end:
	mempool.Get().PutBytes(mBuffer)
	return oTotalSize, oErr
}
