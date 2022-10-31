// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"storj.io/common/sync2"
)

func main() {
	parallel := flag.Int("parallel", runtime.GOMAXPROCS(0), "number of compiles to run in parallel")
	flag.Parse()

	ctx := context.Background()

	type testcase struct {
		compiler string
		os       string
		arch     string
	}

	test := func(compiler, os, arch string) testcase {
		return testcase{
			compiler: compiler,
			os:       os,
			arch:     arch,
		}
	}

	tests := []testcase{
		test("go", "linux", "amd64"),
		test("go", "linux", "386"),
		test("go", "linux", "arm64"),
		test("go", "linux", "arm"),
		test("go", "windows", "amd64"),
		test("go", "windows", "386"),
		test("go", "windows", "arm64"),
		test("go", "darwin", "amd64"),
		test("go", "darwin", "arm64"),
	}

	if _, err := exec.LookPath("go1.17.13"); err == nil {
		tests = append(tests,
			test("go1.17.13", "linux", "amd64"),
			test("go1.17.13", "linux", "386"),
			test("go1.17.13", "linux", "arm64"),
			test("go1.17.13", "linux", "arm"),
			test("go1.17.13", "windows", "amd64"),
			test("go1.17.13", "windows", "386"),
			test("go1.17.13", "darwin", "amd64"),
			test("go1.17.13", "darwin", "arm64"),
		)
	} else {
		fmt.Fprintf(os.Stderr, "skipping go1.17.13 tests: %v\n", err)
	}

	type result struct {
		title string
		out   string
		err   error
	}
	results := make([]result, len(tests))

	lim := sync2.NewLimiter(*parallel)
	for i, test := range tests {
		i, test := i, test
		lim.Go(ctx, func() {
			cmd := exec.Command(test.compiler, "vet", "storj.io/common/...")
			cmd.Env = append(os.Environ(),
				"GOOS="+test.os,
				"GOARCH="+test.arch,
			)

			data, err := cmd.CombinedOutput()
			results[i] = result{
				title: test.compiler + "/" + test.os + "/" + test.arch,
				out:   strings.TrimSpace(string(data)),
				err:   err,
			}
		})
	}
	lim.Wait()

	exit := 0
	for _, r := range results {
		if r.err == nil {
			fmt.Println("#", r.title, "SUCCESS")
		} else {
			fmt.Println("#", r.title, "FAILED", r.err)
			fmt.Println(r.out)
			fmt.Println()
			exit = 1
		}
	}
	os.Exit(exit)
}
