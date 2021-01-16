package cmd

import "testing"

func TestRenderTpl(t *testing.T) {
	valFile = "../example/values.yaml"
	tplFile = "../example/template"

	val, err := parseVal()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", val)

	renderTpl(val)
}
