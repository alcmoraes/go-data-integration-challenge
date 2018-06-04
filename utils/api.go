package utils

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func MustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
}

func Upload(url string, values map[string]io.Reader) (h *http.Request, mp *multipart.Writer, err error) {
	var b bytes.Buffer
	m := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		if x, ok := r.(*os.File); ok {
			if fw, err = m.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		} else {
			if fw, err = m.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return nil, nil, err
		}
	}
	m.Close()
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return nil, nil, err
	}
	return req, m, nil
}
