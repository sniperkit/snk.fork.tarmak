/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright Jetstack Ltd. See LICENSE for details.

// this package contains generated assets from the repository for use during
// tarmak runtime

// This should be used when running in release mode (!devmode)
// +build !devmode

package assets

//go:generate go-bindata -prefix ../../../ -pkg $GOPACKAGE -o assets_bindata.go ../../../terraform/amazon/modules/... ../../../terraform/amazon/templates/... ../../../puppet/... ../../../packer/...
