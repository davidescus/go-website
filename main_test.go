package main

import (
	"testing"
)

func TestParseTemplateFiles(t *testing.T) {
	parse := parseTemplateFiles(templatePath)
	t.Log(parse)

}
