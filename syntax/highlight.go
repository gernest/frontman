// Package syntaxhighlight provides syntax highlighting for code. It currently
// uses a language-independent lexer and performs decently on JavaScript, Java,
// Ruby, Python, Go, and C.
package syntax

import (
	"bytes"
	"io"
	"text/scanner"
	"text/template"
	"unicode"
	"unicode/utf8"

	"github.com/sourcegraph/annotate"
)

// Kind represents a syntax highlighting kind (class) which will be assigned to tokens.
// A syntax highlighting scheme (style) maps text style properties to each token kind.
type Kind uint8

const (
	Whitespace Kind = iota
	String
	Keyword
	Comment
	Type
	Literal
	Punctuation
	Plaintext
	Tag
	HTMLTag
	HTMLAttrName
	HTMLAttrValue
	Decimal
)

//go:generate gostringer -type=Kind

// HTMLConfig holds the HTML class configuration to be used by annotators when
// highlighting code.
type HTMLConfig struct {
	String        string
	Keyword       string
	Comment       string
	Type          string
	Literal       string
	Punctuation   string
	Plaintext     string
	Tag           string
	HTMLTag       string
	HTMLAttrName  string
	HTMLAttrValue string
	Decimal       string
	Whitespace    string
}

type Classifier interface {
	Print(w io.Writer, kind Kind, tokString string) error
	Class(Kind) string
	TokenKind(tok rune, tokString string) Kind
}

// Class returns the set class for a given token Kind.
func (c HTMLConfig) Class(kind Kind) string {
	switch kind {
	case String:
		return c.String
	case Keyword:
		return c.Keyword
	case Comment:
		return c.Comment
	case Type:
		return c.Type
	case Literal:
		return c.Literal
	case Punctuation:
		return c.Punctuation
	case Plaintext:
		return c.Plaintext
	case Tag:
		return c.Tag
	case HTMLTag:
		return c.HTMLTag
	case HTMLAttrName:
		return c.HTMLAttrName
	case HTMLAttrValue:
		return c.HTMLAttrValue
	case Decimal:
		return c.Decimal
	}
	return ""
}

type Annotator interface {
	Annotate(start int, kind Kind, tokText string) (*annotate.Annotation, error)
}

type HTMLAnnotator HTMLConfig

func (a HTMLAnnotator) Annotate(start int, kind Kind, tokText string) (*annotate.Annotation, error) {
	class := ((HTMLConfig)(a)).Class(kind)
	if class != "" {
		left := []byte(`<span class="`)
		left = append(left, []byte(class)...)
		left = append(left, []byte(`">`)...)
		return &annotate.Annotation{
			Start: start, End: start + len(tokText),
			Left: left, Right: []byte("</span>"),
		}, nil
	}
	return nil, nil
}

type HTMLClassifier HTMLConfig

func (h HTMLClassifier) TokenKind(tok rune, tokText string) Kind {
	switch tok {
	case scanner.Ident:
		if IsKeyword(tokText) {
			return Keyword
		}
		if r, _ := utf8.DecodeRuneInString(tokText); unicode.IsUpper(r) {
			return Type
		}
		return Plaintext
	case scanner.Float, scanner.Int:
		return Decimal
	case scanner.Char, scanner.String, scanner.RawString:
		return String
	case scanner.Comment:
		return Comment
	}
	if unicode.IsSpace(tok) {
		return Whitespace
	}
	return Punctuation
}

func (h HTMLClassifier) Class(kind Kind) string {
	return ((HTMLConfig)(h)).Class(kind)
}

func (h HTMLClassifier) Print(w io.Writer, kind Kind, tokText string) error {
	class := ((HTMLConfig)(h)).Class(kind)
	if class != "" {
		_, err := w.Write([]byte(`<span class="`))
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, class)
		if err != nil {
			return err
		}
		_, err = w.Write([]byte(`">`))
		if err != nil {
			return err
		}
	}
	template.HTMLEscape(w, []byte(tokText))
	if class != "" {
		_, err := w.Write([]byte(`</span>`))
		if err != nil {
			return err
		}
	}
	return nil
}

// DefaultHTMLConfig's class names match those of google-code-prettify
// (https://code.google.com/p/google-code-prettify/).
var DefaultHTMLConfig = HTMLConfig{
	String:        "str",
	Keyword:       "kwd",
	Comment:       "com",
	Type:          "typ",
	Literal:       "lit",
	Punctuation:   "pun",
	Plaintext:     "pln",
	Tag:           "tag",
	HTMLTag:       "htm",
	HTMLAttrName:  "atn",
	HTMLAttrValue: "atv",
	Decimal:       "dec",
	Whitespace:    "",
}

func Print(s *scanner.Scanner, w io.Writer, c Classifier) error {
	tok := s.Scan()
	for tok != scanner.EOF {
		tokText := s.TokenText()
		err := c.Print(w, c.TokenKind(tok, tokText), tokText)
		if err != nil {
			return err
		}

		tok = s.Scan()
	}

	return nil
}

func Annotate(src []byte, a Annotator) (annotate.Annotations, error) {
	s := NewScanner(src)

	var anns annotate.Annotations
	read := 0

	tok := s.Scan()
	for tok != scanner.EOF {
		tokText := s.TokenText()

		ann, err := a.Annotate(read, tokenKind(tok, tokText), tokText)
		if err != nil {
			return nil, err
		}
		read += len(tokText)
		if ann != nil {
			anns = append(anns, ann)
		}

		tok = s.Scan()
	}

	return anns, nil
}

func AsHTML(src []byte) ([]byte, error) {
	var buf bytes.Buffer
	err := Print(NewScanner(src), &buf, HTMLClassifier(DefaultHTMLConfig))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Highlight(src []byte, lang string, klasifier ...Classifier) ([]byte, error) {
	var c Classifier
	c = HTMLClassifier(DefaultHTMLConfig)
	if len(klasifier) > 0 {
		c = klasifier[0]
	}
	var buf bytes.Buffer
	err := Print(NewScanner(src), &buf, c)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// NewScanner is a helper that takes a []byte src, wraps it in a reader and creates a Scanner.
func NewScanner(src []byte) *scanner.Scanner {
	return NewScannerReader(bytes.NewReader(src))
}

// NewScannerReader takes a reader src and creates a Scanner.
func NewScannerReader(src io.Reader) *scanner.Scanner {
	var s scanner.Scanner
	s.Init(src)
	s.Error = func(_ *scanner.Scanner, _ string) {}
	s.Whitespace = 0
	s.Mode = s.Mode ^ scanner.SkipComments
	return &s
}

func tokenKind(tok rune, tokText string) Kind {
	switch tok {
	case scanner.Ident:
		if IsKeyword(tokText) {
			return Keyword
		}
		if r, _ := utf8.DecodeRuneInString(tokText); unicode.IsUpper(r) {
			return Type
		}
		return Plaintext
	case scanner.Float, scanner.Int:
		return Decimal
	case scanner.Char, scanner.String, scanner.RawString:
		return String
	case scanner.Comment:
		return Comment
	}
	if unicode.IsSpace(tok) {
		return Whitespace
	}
	return Punctuation
}
