// Copyright 2018 The klaytn Authors
// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.
//
// This file is derived from internal/jsre/deps/deps.go (2018/06/04).
// Modified and improved for the klaytn development.

/*
Package deps contains the console's dependencies of Javascript libraries to be embedded in Go source code.

Klaytn console provides functions of bignumber.js and web3.js which are included in the source tree. deps is using go-bindata<https://github.com/jteeuwen/go-bindata> to make those javascript source files into assets which can be used by go environment.

Source Files

Each file provides following features
  - deps.go      	: Defines `go:generate` options to use go-bindata in making bignumber.js and web3.js as assets
  - bindata.go  	: Provides functions to load and to get information of assets. This file is generated by go-bindata
  - bignumber.js	: A JavaScript library for arbitrary-precision decimal and non-decimal arithmetic
  - web3.js      	: JavaScript API which connects to the Generic JSON RPC spec

*/
package deps
