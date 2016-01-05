package down

import (
	"io"
	"io/ioutil"

	"github.com/gernest/douceur/inliner"
	"github.com/gernest/douceur/parser"
	"github.com/gernest/mark"
)

const DeafultTheme = "prettity"

func Mark(out io.Writer, src io.Reader, highlight, embed bool, theme ...string) error {
	style := DeafultTheme
	if len(theme) > 0 {
		style = theme[0]
	}
	data, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}
	opts := &mark.Options{
		Gfm:       true,
		Highlight: highlight,
	}
	m := mark.New(string(data), opts)
	result := m.Render()
	if embed {
		result, err = Embed(result, style)
		if err != nil {
			return err
		}
	}
	_, err = out.Write([]byte(result))
	return err
}

func Embed(html string, style string) (string, error) {
	sheet, err := parser.Parse(GetStyle(style))
	if err != nil {
		return "", err
	}
	return inliner.NewInliner(html, sheet).Inline()
}
func GetStyle(name string) string {
	style, ok := themes[name]
	if ok {
		return style
	}
	return themes[DeafultTheme]
}
