package main

import (
	"os"
	"path"

	. "github.com/dave/jennifer/jen"
	. "github.com/gagliardetto/utilz"
)

func ToPackageName(s string) string {
	return ToSnake(ToCamel(s))
}
func NewGoFile(programName string, includeBoilerplace bool) *File {
	file := NewFile(ToPackageName(programName))
	// Set a prefix to avoid collision between variable names and packages:
	file.PackagePrefix = "ag"
	// Add comment to file:
	// file.HeaderComment("Code generated by https://github.com/gagliardetto. DO NOT EDIT.")

	if includeBoilerplace {
		{
			// main function:
			// file.Func().Id("main").Params().Block()
		}
	}
	return file
}

// SaveGoFile encodes to a file the provided *jen.File.
func SaveGoFile(outDir string, assetFileName string, file *File) error {
	// Save Go assets:
	assetFilepath := path.Join(outDir, assetFileName)

	// Create file Golang file:
	goFile, err := os.Create(assetFilepath)
	if err != nil {
		panic(err)
	}
	defer goFile.Close()

	// Write generated Golang to file:
	Infof("Saving Golang assets to %q", MustAbs(assetFilepath))
	return file.Render(goFile)
}

func DoGroup(f func(*Group)) *Statement {
	g := &Group{}
	g.CustomFunc(Options{
		Multi: false,
	}, f)
	s := newStatement()
	*s = append(*s, g)
	return s
}

func DoGroupMultiline(f func(*Group)) *Statement {
	g := &Group{}
	g.CustomFunc(Options{
		Multi: true,
	}, f)
	s := newStatement()
	*s = append(*s, g)
	return s
}
func newStatement() *Statement {
	return &Statement{}
}
