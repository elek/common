// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package tlsopts

import (
	"crypto/tls"
	"crypto/x509"
	"strings"

	"storj.io/common/identity"
	"storj.io/common/peertls"
	"storj.io/common/storj"
)

// ServerTLSConfig returns a TSLConfig for use as a server in handshaking with a peer.
func (opts *Options) ServerTLSConfig() *tls.Config {
	return opts.tlsConfig(true)
}

// ClientTLSConfig returns a TSLConfig for use as a client in handshaking with a peer.
func (opts *Options) ClientTLSConfig(id storj.NodeID) *tls.Config {
	return opts.tlsConfig(false, verifyIdentity(id))
}

// ClientTLSConfigPrefix returns a TSLConfig for use as a client in handshaking with a peer.
// The peer node id is validated to match the given prefix
func (opts *Options) ClientTLSConfigPrefix(idPrefix string) *tls.Config {
	return opts.tlsConfig(false, verifyIdentityPrefix(idPrefix))
}

// UnverifiedClientTLSConfig returns a TLSConfig for use as a client in handshaking with
// an unknown peer.
func (opts *Options) UnverifiedClientTLSConfig() *tls.Config {
	return opts.tlsConfig(false)
}

func (opts *Options) tlsConfig(isServer bool, verificationFuncs ...peertls.PeerCertVerificationFunc) *tls.Config {
	verificationFuncs = append(
		[]peertls.PeerCertVerificationFunc{
			peertls.VerifyPeerCertChains,
		},
		verificationFuncs...,
	)

	switch isServer {
	case true:
		verificationFuncs = append(
			verificationFuncs,
			opts.VerificationFuncs.server...,
		)
	case false:
		verificationFuncs = append(
			verificationFuncs,
			opts.VerificationFuncs.client...,
		)
	}

	config := &tls.Config{
		Certificates:                []tls.Certificate{*opts.Cert},
		InsecureSkipVerify:          true,
		MinVersion:                  tls.VersionTLS12,
		DynamicRecordSizingDisabled: true, // always start with big records
		VerifyPeerCertificate: peertls.VerifyPeerFunc(
			verificationFuncs...,
		),
	}

	if isServer {
		config.ClientAuth = tls.RequireAnyClientCert
	}

	return config
}

func verifyIdentity(id storj.NodeID) peertls.PeerCertVerificationFunc {
	return func(_ [][]byte, parsedChains [][]*x509.Certificate) (err error) {
		defer mon.TaskNamed("verifyIdentity")(nil)(&err)
		peer, err := identity.PeerIdentityFromChain(parsedChains[0])
		if err != nil {
			return err
		}

		if peer.ID.String() != id.String() {
			return Error.New("peer ID did not match requested ID")
		}

		return nil
	}
}

func verifyIdentityPrefix(idPrefix string) peertls.PeerCertVerificationFunc {
	return func(_ [][]byte, parsedChains [][]*x509.Certificate) (err error) {
		defer mon.TaskNamed("verifyIdentityPrefix")(nil)(&err)
		peer, err := identity.PeerIdentityFromChain(parsedChains[0])
		if err != nil {
			return err
		}

		if !strings.HasPrefix(peer.ID.String(), idPrefix) {
			return Error.New("peer ID did not match requested ID prefix")
		}

		return nil
	}
}
