/*
 * SPDX-License-Identifier: AGPL-3.0-only
 * Copyright (c) 2023, daeuniverse Organization <team@v2raya.org>
 */

package main

import (
	"github.com/daeuniverse/dae-wing/cmd"
    "github.com/daeuniverse/dae-wing/pkg/prof"
	"github.com/json-iterator/go/extra"
	"os"
	"runtime"
)

import (
	_ "github.com/daeuniverse/dae/component/outbound"
)

func main() {
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)
    prof.Start()
	extra.RegisterFuzzyDecoders()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
