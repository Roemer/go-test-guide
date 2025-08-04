package main

import (
	"archive/zip"
	"fmt"
	"go/build"
	"io"
	"os"
	"path/filepath"

	"github.com/roemer/go-test-guide/internal"
	"github.com/roemer/gotaskr"
	"github.com/roemer/gotaskr/execr"
)

////////////////////////////////////////////////////////////
// Variables
////////////////////////////////////////////////////////////

var outputDirectory = ".build-output"
var reportsDirectory = ".test-reports"

////////////////////////////////////////////////////////////
// Main
////////////////////////////////////////////////////////////

func main() {
	os.Exit(gotaskr.Execute())
}

////////////////////////////////////////////////////////////
// Tasks
////////////////////////////////////////////////////////////

func init() {
	gotaskr.Task("Compile:All", func() error {
		return nil
	}).
		DependsOn("Compile:Windows").
		DependsOn("Compile:Linux").
		DependsOn("Compile:Mac").
		DependsOn("Compile:MacArm")

	gotaskr.Task("Compile:Windows", func() error {
		os.Setenv("GOOS", "windows")
		os.Setenv("GOARCH", "amd64")

		path, err := compile(".exe")
		if err != nil {
			return err
		}
		return zipRelease(path)
	})

	gotaskr.Task("Compile:Linux", func() error {
		os.Setenv("GOOS", "linux")
		os.Setenv("GOARCH", "amd64")
		os.Setenv("CGO_ENABLED", "0")

		path, err := compile("")
		if err != nil {
			return err
		}
		return zipRelease(path)
	})

	gotaskr.Task("Compile:Mac", func() error {
		os.Setenv("GOOS", "darwin")
		os.Setenv("GOARCH", "amd64")

		path, err := compile(".dmg")
		if err != nil {
			return err
		}
		return zipRelease(path)
	})

	gotaskr.Task("Compile:MacArm", func() error {
		os.Setenv("GOOS", "darwin")
		os.Setenv("GOARCH", "arm64")

		path, err := compile(".dmg")
		if err != nil {
			return err
		}
		return zipRelease(path)
	})

	gotaskr.Task("Test", func() error {
		if err := os.MkdirAll(reportsDirectory, os.ModePerm); err != nil {
			return err
		}
		goTestReport := filepath.Join(reportsDirectory, "go-test-report.txt")
		stdout, _, execErr := execr.RunGetOutput(false, "go", execr.SplitArgumentString("test -v ./...")...)
		if err := os.WriteFile(goTestReport, []byte(stdout), os.ModePerm); err != nil {
			return err
		}

		junitTestReport := filepath.Join(reportsDirectory, "junit-test-report.xml")
		if err := execr.Run(true, "go", "install", "github.com/jstemmer/go-junit-report/v2@v2.1.0"); err != nil {
			return err
		}
		if err := execr.Run(true, filepath.Join(build.Default.GOPATH, "bin/go-junit-report"), "-in", goTestReport, "-set-exit-code", "-out", junitTestReport); err != nil {
			return err
		}
		return execErr
	})
}

////////////////////////////////////////////////////////////
// Internal helper functions
////////////////////////////////////////////////////////////

func compile(ext string) (string, error) {
	outputFile := filepath.Join(outputDirectory, "go-test-guide"+ext)
	return outputFile, execr.Run(true, "go", "build", "-o", outputFile, "./cmd/go-test-guide")
}

func zipRelease(file string) error {
	zipFilePath := filepath.Join(outputDirectory, fmt.Sprintf("go-test-guide-%s-%s-%s.zip", os.Getenv("GOOS"), internal.Version, os.Getenv("GOARCH")))

	a, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer a.Close()

	if err := createFlatZip(a, file); err != nil {
		return err
	}
	return os.Remove(file)
}

func createFlatZip(w io.Writer, files ...string) error {
	z := zip.NewWriter(w)
	for _, file := range files {
		src, err := os.Open(file)
		if err != nil {
			return err
		}
		info, err := src.Stat()
		if err != nil {
			return err
		}
		hdr, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		hdr.Name = filepath.Base(file) // Write only the base name in the header
		dst, err := z.CreateHeader(hdr)
		if err != nil {
			return err
		}
		_, err = io.Copy(dst, src)
		if err != nil {
			return err
		}
		src.Close()
	}
	return z.Close()
}
