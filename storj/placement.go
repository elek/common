// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package storj

// PlacementConstraint is the ID of the placement/geofencing rule.
type PlacementConstraint int

const (

	// EveryCountry includes all countries.
	EveryCountry PlacementConstraint = iota

	// EU includes only the 27 members of European Union.
	EU

	// EEA defines the European Economic Area (EU + 3 countries), the area where GDPR is valid.
	EEA

	// US filters nodes only from the United States.
	US

	// DE placement uses nodes only from Germany.
	DE
)
