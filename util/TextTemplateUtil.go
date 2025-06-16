package util

import (
	"bytes"
	"fmt"
	"text/template"
)

func TextTemplateReader(data interface{}, template *template.Template) (string, error) {
	var buf bytes.Buffer
	if err := template.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("渲染模板失败: %w", err)
	}
	return buf.String(), nil
}
