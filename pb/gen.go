// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

//go:build ignore

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	mainpkg   = flag.String("pkg", "storj.io/common/pb", "main package name")
	protoc    = flag.String("protoc", "protoc", "protoc compiler")
	protolock = flag.String("protolock", "protolock", "protolock tool")
)

var ignoreProto = map[string]bool{
	"gogo.proto": true,
	"pico.proto": true,
}

func ignore(files []string) []string {
	xs := []string{}
	for _, file := range files {
		if !ignoreProto[file] {
			xs = append(xs, file)
		}
	}
	return xs
}

// Programs needed for code generation:
//
// github.com/ckaznocha/protoc-gen-lint
// storj.io/drpc/cmd/protoc-gen-go-drpc
// github.com/nilslice/protolock/cmd/protolock

func main() {
	flag.Parse()

	{
		// cleanup previous files
		localfiles, err := filepath.Glob("*.pb.go")
		check(err)

		all := []string{}
		all = append(all, localfiles...)
		for _, match := range all {
			_ = os.Remove(match)
		}
	}

	{
		protofiles, err := filepath.Glob("*.proto")
		check(err)

		protofiles = ignore(protofiles)

		overrideImports := ",Mgoogle/protobuf/timestamp.proto=storj.io/common/pb" +
			",Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types"
		args := []string{
			"--lint_out=.",
			"--gogo_out=paths=source_relative" + overrideImports + ":.",
			"--go-drpc_out=protolib=github.com/gogo/protobuf,paths=source_relative:.",
			"-I=.",
		}
		args = append(args, protofiles...)

		// generate new code
		cmd := exec.Command(*protoc, args...)
		fmt.Println(strings.Join(cmd.Args, " "))
		out, err := cmd.CombinedOutput()
		if len(out) > 0 {
			fmt.Println(string(out))
		}
		check(err)
	}

	{
		files, err := filepath.Glob("*.pb.go")
		check(err)
		for _, file := range files {
			process(file)
		}
	}

	{
		// format code to get rid of extra imports
		out, err := exec.Command("goimports", "-local", "storj.io", "-w", ".").CombinedOutput()
		if len(out) > 0 {
			fmt.Println(string(out))
		}
		check(err)
	}

	{
		// also generate grant, which depends on the protobuf files in this folder
		out, err := exec.Command("go", "generate", "storj.io/common/grant/...").CombinedOutput()
		if len(out) > 0 {
			fmt.Println(string(out))
		}
		check(err)
	}

	{ // regenerate proto.lock file
		cmd := exec.Command(*protolock, "commit")
		cmd.Dir = ".."
		fmt.Println(strings.Join(cmd.Args, " "))
		out, err := cmd.CombinedOutput()
		if len(out) > 0 {
			fmt.Println(string(out))
		}
		check(err)
	}
}

var rxInit = regexp.MustCompile(`func init\(\) \{[^\}]*\}`)
var rxDescriptor = regexp.MustCompile(`(?msU)func \(\*?[a-zA-Z_]+\) Descriptor\(\) .*^}`)
var rxEnumDescriptor = regexp.MustCompile(`(?msU)func \(\*?[a-zA-Z_]+\) EnumDescriptor\(\) .*^}`)
var rxFileDescriptor = regexp.MustCompile(`(?msU)^var fileDescriptor_.*^}`)

var rxReferenceImportsComment = regexp.MustCompile("(?m)^\\/\\/ Reference imports to suppress .*$")
var rxReferenceImports = regexp.MustCompile("(?m)^var _ = [a-zA-Z.]+$")

func process(file string) {
	data, err := os.ReadFile(file)
	check(err)

	source := string(data)

	source = rxInit.ReplaceAllString(source, "")
	source = rxDescriptor.ReplaceAllString(source, "")
	source = rxEnumDescriptor.ReplaceAllString(source, "")
	source = rxFileDescriptor.ReplaceAllString(source, "")

	// goimports takes care of removing the unnecessary imports
	source = rxReferenceImportsComment.ReplaceAllString(source, "")
	source = rxReferenceImports.ReplaceAllString(source, "")

	// When generating code to the same path as proto, it will
	// end up generating an `import _ "."`, the following replace removes it.
	source = strings.Replace(source, `_ "."`, "", -1)

	err = os.WriteFile(file, []byte(source), 0644)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
