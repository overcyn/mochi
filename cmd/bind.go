package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"go/build"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func Bind(flags *Flags, args []string, dev bool) error {
	// Make $WORK.
	tempdir, err := NewTmpDir(flags, "")
	if err != nil {
		return err
	}
	if !flags.BuildWork {
		defer RemoveAll(flags, tempdir)
	}

	// Make $WORK/matcha-ios
	workOutputDir := filepath.Join(tempdir, "matcha-ios")
	if err := Mkdir(flags, workOutputDir); err != nil {
		return err
	}

	// Get $GOPATH/pkg/gomobile.
	gomobilepath, err := GoMobilePath()
	if err != nil {
		return err
	}

	// Get toolchain version.
	installedVersion, err := ReadFile(flags, filepath.Join(gomobilepath, "version"))
	if err != nil {
		return errors.New("toolchain partially installed, run `gomobile init`")
	}

	// Get go version.
	goVersion, err := GoVersion(flags)
	if err != nil {
		return err
	}

	// Check toolchain matches go version.
	if !bytes.Equal(installedVersion, goVersion) {
		return errors.New("toolchain out of date, run `gomobile init`")
	}

	// Get current working directory.
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// Create a build context.
	ctx := build.Default
	ctx.GOARCH = "arm"
	ctx.GOOS = "darwin"
	ctx.BuildTags = append(ctx.BuildTags, "ios")

	// Get import paths to be built.
	importPaths := []string{}
	srcDir := ""
	if len(args) == 0 {
		importPaths = append(importPaths, ".")
		srcDir = cwd
	} else {
		for _, i := range args {
			i = path.Clean(i)
			importPaths = append(importPaths, i)
		}
		srcDir = cwd
	}

	// Get packages to be built
	pkgs, err := ImportAll(&ctx, importPaths, srcDir, build.ImportComment)
	if err != nil {
		return err
	}

	// Check if any of the package is main.
	for _, pkg := range pkgs {
		if pkg.Name == "main" {
			return fmt.Errorf("binding 'main' package (%s) is not supported", pkg.ImportComment)
		}
	}

	// Copy package's ios directory if it imports gomatcha.io/bridge.
	for _, pkg := range pkgs {
		importsBridge := false
		for _, i := range pkg.Imports {
			if i == "gomatcha.io/bridge" {
				importsBridge = true
				break
			}
		}

		if importsBridge {
			files, err := ioutil.ReadDir(pkg.Dir)
			if err != nil {
				continue
			}

			for _, i := range files {
				if i.IsDir() && i.Name() == "ios" {
					// Copy directory
					src := filepath.Join(pkg.Dir, "ios")
					dst := filepath.Join(workOutputDir)
					CopyDirContents(flags, dst, src)
				}
			}
		}
	}

	genDir := filepath.Join(tempdir, "gen")
	binaryPath := filepath.Join(workOutputDir, "MatchaBridge", "MatchaBridge", "MatchaBridge.a")

	// Build the "matcha/bridge" dir
	bridgeDir := filepath.Join(genDir, "src", "gomatcha.io", "bridge")
	if err := Mkdir(flags, bridgeDir); err != nil {
		return err
	}

	// Create the "main" go package, that references the other go packages
	mainPath := filepath.Join(tempdir, "src", "iosbin", "main.go")
	err = WriteFile(flags, mainPath, func(w io.Writer) error {
		format := fmt.Sprintf(BindFile, args[0]) // TODO(KD): Should this be args[0] or should it use the logic to generate pkgs
		_, err := w.Write([]byte(format))
		return err
	})
	if err != nil {
		return fmt.Errorf("failed to create the binding package for iOS: %v", err)
	}

	// Get the supporting files
	objcPkg, err := ctx.Import("gomatcha.io/matcha/cmd", "", build.FindOnly)
	if err != nil {
		return err
	}
	if err := CopyFile(flags, filepath.Join(bridgeDir, "matchaobjc.h"), filepath.Join(objcPkg.Dir, "matchaobjc.h.support")); err != nil {
		return err
	}
	if err := CopyFile(flags, filepath.Join(bridgeDir, "matchaobjc.m"), filepath.Join(objcPkg.Dir, "matchaobjc.m.support")); err != nil {
		return err
	}
	if err := CopyFile(flags, filepath.Join(bridgeDir, "matchaobjc.go"), filepath.Join(objcPkg.Dir, "matchaobjc.go.support")); err != nil {
		return err
	}
	if err := CopyFile(flags, filepath.Join(bridgeDir, "matchago.h"), filepath.Join(objcPkg.Dir, "matchago.h.support")); err != nil {
		return err
	}
	if err := CopyFile(flags, filepath.Join(bridgeDir, "matchago.m"), filepath.Join(objcPkg.Dir, "matchago.m.support")); err != nil {
		return err
	}
	if err := CopyFile(flags, filepath.Join(bridgeDir, "matchago.go"), filepath.Join(objcPkg.Dir, "matchago.go.support")); err != nil {
		return err
	}

	// title := "MatchaBridge"
	// frameworkDir := filepath.Join(workOutputDir, "Matcha", title+".framework")

	// Build framework directory structure.
	// headersDir := filepath.Join(frameworkDir, "Versions", "A", "Headers")
	// resourcesDir := filepath.Join(frameworkDir, "Versions", "A", "Resources")
	// modulesDir := filepath.Join(frameworkDir, "Versions", "A", "Modules")
	// binaryPath := filepath.Join(frameworkDir, "Versions", "A", title+".a")
	// if err := Mkdir(flags, headersDir); err != nil {
	// 	return err
	// }
	// if err := Mkdir(flags, resourcesDir); err != nil {
	// 	return err
	// }
	// if err := Mkdir(flags, modulesDir); err != nil {
	// 	return err
	// }
	// if err := Symlink(flags, "A", filepath.Join(frameworkDir, "Versions", "Current")); err != nil {
	// 	return err
	// }
	// if err := Symlink(flags, filepath.Join("Versions", "Current", "Headers"), filepath.Join(frameworkDir, "Headers")); err != nil {
	// 	return err
	// }
	// if err := Symlink(flags, filepath.Join("Versions", "Current", "Resources"), filepath.Join(frameworkDir, "Resources")); err != nil {
	// 	return err
	// }
	// if err := Symlink(flags, filepath.Join("Versions", "Current", "Modules"), filepath.Join(frameworkDir, "Modules")); err != nil {
	// 	return err
	// }
	// if err := Symlink(flags, filepath.Join("Versions", "Current", title), filepath.Join(frameworkDir, title)); err != nil {
	// 	return err
	// }

	// // Copy in headers.
	// if err = CopyFile(flags, filepath.Join(headersDir, "matchaobjc.h"), filepath.Join(bridgeDir, "matchaobjc.h")); err != nil {
	// 	return err
	// }
	// if err = CopyFile(flags, filepath.Join(headersDir, "matchago.h"), filepath.Join(bridgeDir, "matchago.h")); err != nil {
	// 	return err
	// }

	// // Copy in resources.
	// if err := ioutil.WriteFile(filepath.Join(resourcesDir, "Info.plist"), []byte(InfoPlist), 0666); err != nil {
	// 	return err
	// }

	// // Write modulemap.
	// err = WriteModuleMap(flags, filepath.Join(modulesDir, "module.modulemap"), title)
	// if err != nil {
	// 	return err
	// }

	// Build platform binaries concurrently.
	matchaDarwinArmEnv, err := DarwinArmEnv(flags)
	if err != nil {
		return err
	}
	matchaDarwinArm64Env, err := DarwinArm64Env(flags)
	if err != nil {
		return err
	}
	matchaDarwin386Env, err := Darwin386Env(flags)
	if err != nil {
		return err
	}
	matchaDarwinAmd64Env, err := DarwinAmd64Env(flags)
	if err != nil {
		return err
	}

	type archPath struct {
		arch string
		path string
		err  error
	}
	archChan := make(chan archPath)
	for _, i := range [][]string{matchaDarwinArmEnv, matchaDarwinArm64Env, matchaDarwinAmd64Env, matchaDarwin386Env} {
		go func(env []string) {
			arch := Getenv(env, "GOARCH")
			env = append(env, "GOPATH="+genDir+string(filepath.ListSeparator)+os.Getenv("GOPATH"))
			path := filepath.Join(tempdir, "matcha-"+arch+".a")
			err := GoBuild(flags, mainPath, env, ctx, tempdir, "-buildmode=c-archive", "-o", path)
			archChan <- archPath{arch, path, err}
		}(i)
	}
	archs := []archPath{}
	for i := 0; i < 4; i++ {
		arch := <-archChan
		if arch.err != nil {
			return arch.err
		}
		archs = append(archs, arch)
	}

	// Lipo to build fat binary.
	cmd := exec.Command("xcrun", "lipo", "-create")
	for _, i := range archs {
		cmd.Args = append(cmd.Args, "-arch", ArchClang(i.arch), i.path)
	}
	cmd.Args = append(cmd.Args, "-o", binaryPath)
	if err := RunCmd(flags, tempdir, cmd); err != nil {
		return err
	}

	// Copy in headers.
	if err = CopyFile(flags, filepath.Join(workOutputDir, "MatchaBridge", "MatchaBridge", "matchaobjc.h"), filepath.Join(bridgeDir, "matchaobjc.h")); err != nil {
		return err
	}
	if err = CopyFile(flags, filepath.Join(workOutputDir, "MatchaBridge", "MatchaBridge", "matchago.h"), filepath.Join(bridgeDir, "matchago.h")); err != nil {
		return err
	}

	// Create output dir
	outputDir := flags.BuildO
	if outputDir == "" {
		outputDir = "Matcha-iOS"
	}
	if err := RemoveAll(flags, outputDir); err != nil {
		return err
	}

	// Copy output directory into place.
	if err := CopyDir(flags, outputDir, workOutputDir); err != nil {
		return err
	}

	return nil
}

var InfoPlist = `<?xml version="1.0" encoding="UTF-8"?>
    <!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
    <plist version="1.0">
      <dict>
      </dict>
    </plist>
`

var BindFile = `
package main

import (
    _ "gomatcha.io/bridge"
    _ "%s"
)

import "C"

func main() {}
`

var ModuleMapTemplate = template.Must(template.New("iosmmap").Parse(`framework module "{{.Module}}" {
    // header "ref.h"
{{range .Headers}}    header "{{.}}"
{{end}}
    export *
}`))

func WriteModuleMap(flags *Flags, filename string, title string) error {
	// Write modulemap.
	var mmVals = struct {
		Module  string
		Headers []string
	}{
		Module:  title,
		Headers: []string{"matchaobjc.h", "matchago.h"},
	}
	err := WriteFile(flags, filename, func(w io.Writer) error {
		return ModuleMapTemplate.Execute(w, mmVals)
	})
	if err != nil {
		return err
	}
	return nil
}
