// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

// Package experiment implements feature flag propagation.
package experiment

import (
	"context"
	"github.com/spacemonkeygo/monkit/v3"
)

const (
	// Experiment is the key used in context and DRPC metadata header.
	Experiment = "Experiment"
)

var mon = monkit.Package()

// WithExperiment registers the feature flag of an ongoing experiment.
func WithExperiment(ctx context.Context, experiment string) context.Context {
	if experiment != "" {
		return context.WithValue(ctx, Experiment, experiment)
	}
	return ctx
}
