package utils

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
)

func MustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
}

func GenerateMultipartHeadersFromLocal(values map[string]io.Reader) (bb *bytes.Buffer, mp *multipart.Writer, err error) {
	var b bytes.Buffer
	m := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		if x, ok := r.(*os.File); ok {
			if fw, err = m.CreateFormFile(key, x.Name()); err != nil {
				return nil, nil, err
			}
		} else {
			if fw, err = m.CreateFormField(key); err != nil {
				return nil, nil, err
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return nil, nil, err
		}
	}
	m.Close()
	return &b, m, nil
}
