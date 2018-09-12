package templates

import (
	"bytes"
	"errors"
	"go/format"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/nicklanng/modelgen/model"
)

type TemplateWriter struct {
	templateRoot string
	outputPath   string
	packageName  string
}

func NewTemplateWriter(templateRoot, outputPath, packageName string) *TemplateWriter {
	return &TemplateWriter{
		templateRoot: templateRoot,
		outputPath:   outputPath,
		packageName:  packageName,
	}
}

func (w *TemplateWriter) WriteModels(models []model.EntityDescriptor) error {
	var (
		modelTemplate []byte
		buffer        *bytes.Buffer
		formatted     []byte
		file          *os.File
		err           error
	)

	modelTemplate, err = box.MustBytes(filepath.Join(w.templateRoot, "model.go.template"))
	if err != nil {
		return errors.New("cannot load model template")
	}
	t := template.Must(template.New("model").Funcs(funcMap).Parse(string(modelTemplate)))

	for _, m := range models {
		templateData := model.TemplateData{
			Model:       m,
			Receiver:    strings.ToLower(string(m.Name[0])),
			PackageName: w.packageName,
		}

		buffer = new(bytes.Buffer)
		if err := t.Execute(buffer, templateData); err != nil {
			return err
		}

		if formatted, err = format.Source(buffer.Bytes()); err != nil {
			return err
		}

		buffer = bytes.NewBuffer(formatted)

		os.Mkdir(w.outputPath, 0777)

		p := filepath.Join(w.outputPath, m.TableName)
		if file, err = os.Create(p + ".go"); err != nil {
			return err
		}
		buffer.WriteTo(file)
		file.Close()
	}

	return nil
}

func (w *TemplateWriter) WriteHelpers() error {
	var (
		helperTemplate []byte
		formatted      []byte
		tmpl           *template.Template
		buffer         *bytes.Buffer
		file           *os.File
		err            error
	)

	templatePath := filepath.Join(w.templateRoot, "x_helpers.go.template")
	if helperTemplate, err = box.MustBytes(templatePath); err != nil {
		return errors.New("cannot retrieve helper template file")
	}

	tmpl = template.Must(template.New("helpers").Parse(string(helperTemplate)))
	buffer = new(bytes.Buffer)
	if err = tmpl.Execute(buffer, map[string]string{"PackageName": w.packageName}); err != nil {
		return err
	}

	if formatted, err = format.Source(buffer.Bytes()); err != nil {
		return err
	}

	buffer = bytes.NewBuffer(formatted)

	out := filepath.Join(w.outputPath, "x_helpers.go")
	if file, err = os.Create(out); err != nil {
		return err
	}
	defer file.Close()

	if _, err = io.Copy(file, buffer); err != nil {
		return err
	}

	return nil
}

func (w *TemplateWriter) WriteHelperTests() error {
	var (
		helperTemplate []byte
		formatted      []byte
		tmpl           *template.Template
		buffer         *bytes.Buffer
		file           *os.File
		err            error
	)

	templatePath := filepath.Join(w.templateRoot, "x_helpers_test.go.template")
	if helperTemplate, err = box.MustBytes(templatePath); err != nil {
		return errors.New("cannot retrieve helper template file")
	}

	tmpl = template.Must(template.New("helpers").Parse(string(helperTemplate)))
	buffer = new(bytes.Buffer)
	if err = tmpl.Execute(buffer, map[string]string{"PackageName": w.packageName}); err != nil {
		return err
	}

	if formatted, err = format.Source(buffer.Bytes()); err != nil {
		return err
	}

	buffer = bytes.NewBuffer(formatted)

	out := filepath.Join(w.outputPath, "x_helpers_test.go")
	if file, err = os.Create(out); err != nil {
		return err
	}
	defer file.Close()

	if _, err = io.Copy(file, buffer); err != nil {
		return err
	}

	return nil
}
