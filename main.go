package wordwrap

import (
	"bufio"
	"bytes"
	"errors"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	LF    = "\n" // LF  = byte(10) // 012 0A 00001010 LF 	&#010; Line Feed (\n)
	HT    = "\t" // HT  = byte(9)  // 011 09 00001001 HT 	&#009; Horizontal Tab (\t)
	SPACE = ` `  // SPACE = byte(32)  //	040	20 00100000   &#32; Space (&nbsp;)

	ColGuide = `---------------------------------------------------------|-----------|--|------|---------------------------------------|` // 58, 70, 72, 80, 120
)

func Wrap(s string, breakCol int) string {
	var (
		// n       int // counter
		cc      int // char count
		wordbuf bytes.Buffer
		outbuf  strings.Builder
		err     error
		indent  []rune
		isFirst = true
		// ok      bool
		// softLimit = breakCol - 8
	)
	in := bufio.NewScanner(strings.NewReader(s))
	in.Split(bufio.ScanRunes)
	for in.Scan() {

		err = in.Err()
		if !errors.Is(err, nil) {
			log.Fatal(err)
		}

		rs := []rune(in.Text())
		if len(rs) < 1 {
			continue
		}

		if isFirst && unicode.IsSpace(rs[0]) {
			// special case as one tab char is worth 4 space chars
			if string(rs[0]) == HT {
				rs = []rune("    ")
			}
			indent = append(indent, rs...)
			continue
		}

		if isFirst {
			outbuf.WriteString(string(indent))
			cc = len(indent)
			isFirst = false
		}

		// word count
		wordbuf.WriteString(string(rs))

		// soft break
		// if cc > softLimit && unicode.IsSpace(rs[0]) {
		// 	outbuf.WriteString(LF)
		// 	cc = 0

		// 	n, _ = outbuf.Write(wordbuf.Bytes())
		// 	cc += n
		// 	wordbuf.Reset()
		// 	continue
		// }

		if unicode.IsSpace(rs[0]) {
			wordLen := utf8.RuneCount(wordbuf.Bytes())
			if cc+wordLen > breakCol {
				outbuf.WriteString(LF)
				outbuf.WriteString(string(indent))
				cc = len(indent)
			}

			outbuf.Write(wordbuf.Bytes())
			cc += wordLen
			wordbuf.Reset()
		}
	}
	// @TODO deal with last/trailing/ending words ...
	if wordbuf.Len() > 0 {
		wordLen := utf8.RuneCount(wordbuf.Bytes())
		if cc+wordLen > breakCol {
			outbuf.WriteString(LF)
			outbuf.WriteString(string(indent))
			cc = len(indent)
		}

		outbuf.Write(wordbuf.Bytes())
	}

	return outbuf.String()
}

func Simple(s string, breakCol int) string {
	var (
		// n       int // counter
		cc      int // char count
		wordbuf bytes.Buffer
		outbuf  strings.Builder
		err     error
		indent  []rune
		isFirst = true
		// ok      bool
		softLimit = breakCol - 8
	)
	in := bufio.NewScanner(strings.NewReader(s))
	in.Split(bufio.ScanRunes)
	for in.Scan() {

		err = in.Err()
		if !errors.Is(err, nil) {
			log.Fatal(err)
		}

		rs := []rune(in.Text())
		if len(rs) < 1 {
			continue
		}

		if isFirst && unicode.IsSpace(rs[0]) {
			// special case as one tab char is worth 4 space chars
			if string(rs[0]) == HT {
				rs = []rune("    ")
			}
			indent = append(indent, rs...)
			continue
		}

		if isFirst {
			outbuf.WriteString(string(indent))
			cc = len(indent)
			isFirst = false
		}

		// word count
		wordbuf.WriteString(string(rs))

		// soft break
		if cc > softLimit && unicode.IsSpace(rs[0]) {
			wordLen := utf8.RuneCount(wordbuf.Bytes())
			outbuf.WriteString(LF)
			outbuf.WriteString(string(indent))
			cc = len(indent)

			outbuf.Write(wordbuf.Bytes())
			cc += wordLen
			wordbuf.Reset()
			continue
		}

		if unicode.IsSpace(rs[0]) {
			wordLen := utf8.RuneCount(wordbuf.Bytes())
			if cc+wordLen > breakCol {
				outbuf.WriteString(LF)
				outbuf.WriteString(string(indent))
				cc = len(indent)
			}

			outbuf.Write(wordbuf.Bytes())
			cc += wordLen
			wordbuf.Reset()
		}
	}
	// @TODO deal with last/trailing/ending words ...
	if wordbuf.Len() > 0 {
		wordLen := utf8.RuneCount(wordbuf.Bytes())
		if cc+wordLen > breakCol {
			outbuf.WriteString(LF)
			outbuf.WriteString(string(indent))
			cc = len(indent)
		}

		outbuf.Write(wordbuf.Bytes())
	}

	return outbuf.String()
}

func ReWrap(s string, breakCol int) string {
	var (
		// n       int // counter
		cc      int // char count
		wordbuf bytes.Buffer
		outbuf  strings.Builder
		err     error
		indent  []rune
		isFirst = true
		// ok      bool
		// softLimit = breakCol - 8
	)
	in := bufio.NewScanner(strings.NewReader(s))
	in.Split(bufio.ScanRunes)
	for in.Scan() {

		err = in.Err()
		if !errors.Is(err, nil) {
			log.Fatal(err)
		}

		rs := []rune(in.Text())
		if len(rs) < 1 {
			continue
		}

		if string(rs[0]) == LF {
			rs = []rune(SPACE)
			isFirst = false
		}

		if isFirst && unicode.IsSpace(rs[0]) {
			// special case as one tab char is worth 4 space chars
			if string(rs[0]) == HT {
				rs = []rune("    ")
			}
			indent = append(indent, rs...)
			continue
		}

		if isFirst {
			outbuf.WriteString(string(indent))
			cc = len(indent)
			isFirst = false
		}

		// word count
		wordbuf.WriteString(string(rs))

		// soft break
		// if cc > softLimit && unicode.IsSpace(rs[0]) {
		// 	outbuf.WriteString(LF)
		// 	cc = 0

		// 	n, _ = outbuf.Write(wordbuf.Bytes())
		// 	cc += n
		// 	wordbuf.Reset()
		// 	continue
		// }

		if unicode.IsSpace(rs[0]) {
			wordLen := utf8.RuneCount(wordbuf.Bytes())
			if cc+wordLen > breakCol {
				outbuf.WriteString(LF)
				outbuf.WriteString(string(indent))
				cc = len(indent)
			}

			outbuf.Write(wordbuf.Bytes())
			cc += wordLen
			wordbuf.Reset()
		}
	}
	// @TODO deal with last/trailing/ending words ...
	if wordbuf.Len() > 0 {
		wordLen := utf8.RuneCount(wordbuf.Bytes())
		if cc+wordLen > breakCol {
			outbuf.WriteString(LF)
			outbuf.WriteString(string(indent))
			cc = len(indent)
		}

		outbuf.Write(wordbuf.Bytes())
	}

	return outbuf.String()
}
