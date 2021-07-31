package template

import "github.com/hoisie/mustache"

type DetailView struct {
	Display string
	Value   string
	Detail  string
}

type RenderView struct {
	Id       string
	QrSource string
	Issuer   string
	Infos    []DetailView
}

type Renderer interface {
	Render(templatePath string, view RenderView) (string, error)
}

type renderer struct{}

func New() Renderer {
	return &renderer{}
}

func (*renderer) Render(templatePath string, view RenderView) (string, error) {
	t, err := mustache.ParseFile(templatePath)
	if err != nil {
		return "", err
	}
	return t.Render(view), nil
}
