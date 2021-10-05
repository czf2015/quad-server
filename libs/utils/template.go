package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	textTemplate "text/template"
)

func ParseHTML(name, tmpl string, param interface{}) template.HTML {
	t := template.New(name)
	t, err := t.Parse(tmpl)
	if err != nil {
		fmt.Println("utils parse html error", err)
		return ""
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, param)
	if err != nil {
		fmt.Println("utils parse html error", err)
		return ""
	}
	return template.HTML(buf.String())
}

func ParseText(name, tmpl string, param interface{}) string {
	t := textTemplate.New(name)
	t, err := t.Parse(tmpl)
	if err != nil {
		fmt.Println("utils parse text error", err)
		return ""
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, param)
	if err != nil {
		fmt.Println("utils parse text error", err)
		return ""
	}
	return buf.String()
}

func CompressHTML(h *template.HTML) {
	st := strings.Split(string(*h), "\n")
	var ss []string
	for i := 0; i < len(st); i++ {
		st[i] = strings.TrimSpace(st[i])
		if st[i] != "" {
			ss = append(ss, st[i])
		}
	}
	*h = template.HTML(strings.Join(ss, "\n"))
}
