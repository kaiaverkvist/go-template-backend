package view

import (
	"errors"
	"github.com/kaiaverkvist/stick"
	"github.com/kaiaverkvist/stick/twig"
	"github.com/labstack/echo/v4"
	"os"
	"path/filepath"
)

type ViewConfig struct {
	BaseURI   string
	Folder    string
	Extension string
	BaseFile  string
	Caching   bool
}

// View is the struct used to contain an individual page template.
type View struct {
	Name string

	Variables map[string]stick.Value
	Context   echo.Context

	// Store our ViewConfig instance so we can use different configs for different views.
	ViewConfig ViewConfig
}

// Creates a new View instance.
func New(ctx echo.Context, name string) *View {
	view := &View{}

	view.Name = name

	// Set the ViewConfig.
	view.ViewConfig = ViewConfig{
		BaseURI:   "/",
		Folder:    "www",
		Extension: "html",
		BaseFile:  "base",
		Caching:   false,
	}

	// Empty holding map for the View variables.
	view.Variables = make(map[string]stick.Value)
	view.Context = ctx

	return view
}

// Renders a view .
func (v *View) Render() error {

	// Pre-emptively catch one of the common errors that may happen.
	if v.Name == "" {
		return errors.New("unable to render a view without first setting template name")
	}

	// Set up proper pathing for the template.
	workingDirectory, _ := os.Getwd()
	templatePath := filepath.Join(workingDirectory, v.ViewConfig.Folder)
	documentName := v.Name + "." + v.ViewConfig.Extension

	// Create the template from the path
	stickTemplate := twig.New(stick.NewFilesystemLoader(templatePath))

	// Render the template, or alternatively the error returned:
	err := stickTemplate.Execute(documentName, v.Context.Response(), v.Variables)
	if err != nil {
		return err
	}

	// No error, return nil!
	return nil
}
