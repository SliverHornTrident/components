package interfaces

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
)

type OssOptions func(*OssOption)

type OssOption struct {
	file           *os.File
	buffer         *bytes.Buffer
	fileHeader     *multipart.FileHeader
	fileHeaderFile multipart.File
}

func (o *OssOption) Defer() {
	if o.file != nil {
		_ = o.file.Close()
	}
	if o.fileHeaderFile != nil {
		_ = o.fileHeaderFile.Close()
	}
}

func (o *OssOption) GetReader() io.Reader {
	if o.file != nil {
		return o.file
	}
	if o.buffer != nil {
		return o.buffer
	}
	if o.fileHeaderFile == nil {
		o.fileHeaderFile, _ = o.fileHeader.Open()
		return o.fileHeaderFile
	}
	return nil
}

func (o *OssOption) GetHeader(key string) string {
	if o.fileHeaderFile != nil {
		return o.fileHeader.Header.Get(key)
	}
	return ""
}

func (o *OssOption) GetFileInfo() os.FileInfo {
	if o.file != nil {
		info, _ := o.file.Stat()
		return info
	}
	if o.fileHeader != nil && o.fileHeaderFile != nil {
		file, ok := o.fileHeaderFile.(*os.File)
		if ok {
			info, _ := file.Stat()
			return info
		}
	}
	return nil
}

func NewOssOption(options ...OssOptions) *OssOption {
	oss := &OssOption{}
	for i := 0; i < len(options); i++ {
		options[i](oss)
	}
	return oss
}

func WithFile(file *os.File) OssOptions {
	return func(o *OssOption) {
		o.file = file
	}
}

func WithBytes(body []byte) OssOptions {
	return func(o *OssOption) {
		o.buffer = bytes.NewBuffer(body)
	}
}

func WithFileHeader(header *multipart.FileHeader) OssOptions {
	return func(o *OssOption) {
		o.fileHeader = header
	}
}
