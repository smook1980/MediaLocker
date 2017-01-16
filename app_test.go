package medialocker

import (
	"testing"
	"bytes"
	"io"
	"os"
)

type TestingApp struct {
	App
	in  bytes.Buffer
	out bytes.Buffer
	err bytes.Buffer
	logs bytes.Buffer
}

func (ab *AppBuilder) TestBuild() (*TestingApp, *App, []error) {
	testApp := &TestingApp{App: *ab.app}
	ab.app.Fs = GetTestFileSystem()
	ab.app.In = io.Reader(&testApp.in)
	ab.app.Out = io.Writer(&testApp.out)
	ab.app.Err = io.Writer(&testApp.err)
	ab.app.Log = NewTestLogger(os.Stderr)

	ab.app.Config.MemDB = true
	ab.app.Config.LogSQL = true
	ab.app.Registry = *NewRegistry(ab.app.Log, ab.app.Config)

	return testApp, ab.app, ab.errors
}

func TestAppConfig(t *testing.T) {
	EnableTestFileSystem()
	config, errors := BuildConfig(DefaultConfiguration())

	for _, err := range(errors) {
		t.Errorf("Unexpected error returned from BuildConfig! %s", err)
	}

	if config == EmptyConfig {
		t.Errorf("Expected not EmptyConfig! %+v", EmptyConfig)
	}
}

func TestAppBuilder(t *testing.T) {
	testApp, app, errs := NewAppBuilder().TestBuild()
	for _, err := range(errs) {
		t.Errorf("Unexpected error for AppBuilder: %s", err)
	}

	if testApp == nil {
		t.Errorf("Expected TestApp to not be nil, got nil!")
	}

	if app == nil {
		t.Errorf("Expected app to not be nil, got nil!")
	}

	app.Shutdown()
}

func TestAppRegistry(t *testing.T) {
	_, app, errs := NewAppBuilder().TestBuild()
	for _, err := range(errs) {
		t.Errorf("Unexpected error for AppBuilder: %s", err)
	}

	if app.Registry.dbFactory == nil {
		t.Error("Got nil back for Registry.dbFactory, expected it to be present.")
	}

	db, err := app.Registry.dbFactory()

	if err != nil {
		t.Errorf("Expected error to be nil, got: %s", err)
	}

	if db == nil {
		t.Error("Expected db to be present, instead got nil!")
	}

	if err := db.Ping(); err != nil {
		t.Errorf("Unable to open DB, got err: %s", err)
	}

	app.Shutdown()
}