package common

import (
	"bytes"
)

type StringBuffer struct {
	buffer bytes.Buffer
}

func (str *StringBuffer) Join(strs ...string) *StringBuffer {

	//var buffer bytes.Buffer //Buffer是一个实现了读写方法的可变大小的字节缓冲
	//buffer.WriteString(str)
	for i := 0; i < len(strs); i++ {
		str.buffer.WriteString(strs[i])
	}
	return str

}
func (str *StringBuffer) String() string {
	return str.buffer.String()

}

func GetStringBuffer() StringBuffer {
	return StringBuffer{}
}
