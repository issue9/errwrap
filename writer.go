// SPDX-FileCopyrightText: 2020-2024 caixw
//
// SPDX-License-Identifier: MIT

package errwrap

import (
	"fmt"
	"io"
)

// Writer 对 io.WrWriter 的简单封闭
type Writer struct {
	io.Writer
	Err error
}

// WString 写入字符串
func (w *Writer) WString(str string) *Writer {
	if w.Err == nil {
		_, w.Err = io.WriteString(w, str)
	}
	return w
}

// WByte 写入单个字节内容
func (w *Writer) WByte(b byte) *Writer {
	return w.WBytes([]byte{b})
}

// WBytes 写入字节内容
func (w *Writer) WBytes(data []byte) *Writer {
	if w.Err == nil {
		_, w.Err = w.Write(data)
	}
	return w
}

// WRune 写入单个字节内容
func (w *Writer) WRune(r rune) *Writer {
	return w.WBytes([]byte(string(r)))
}

// WRunes 写入字节内容
func (w *Writer) WRunes(rs []rune) *Writer {
	return w.WBytes([]byte(string(rs)))
}

// Print 相当于 fmt.Fprint(builder, v...)
func (w *Writer) Print(v ...interface{}) *Writer {
	if w.Err == nil {
		_, w.Err = fmt.Fprint(w, v...)
	}
	return w
}

// Println 相当于 fmt.Fprintln(builder, v...)
func (w *Writer) Println(v ...interface{}) *Writer {
	if w.Err == nil {
		_, w.Err = fmt.Fprintln(w, v...)
	}
	return w
}

// Printf 相当于 fmt.Fprintf(builder, format, v...)
func (w *Writer) Printf(format string, v ...interface{}) *Writer {
	if w.Err == nil {
		_, w.Err = fmt.Fprintf(w, format, v...)
	}
	return w
}
