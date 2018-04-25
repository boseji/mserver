// Copyright 2018 @boseji <salearj@hotmail.com> All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is specially dedicated to Crypto Hashes

package mserver

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"io"
)

// Sha1 function get the SHA1 Hash from a given bytes.Buffer and
// returns the result also in bytes.Buffer
// As per the FIPS 180-4 :  When a message of any length less than 2^64 bits
// We need to use SHA-1, SHA-224 and SHA-256
// That would be 2048 Petabytes
func Sha1(data *bytes.Buffer) *bytes.Buffer {
	m := sha1.New()
	io.Copy(m, data)
	return bytes.NewBuffer(m.Sum(nil))
}

// Sha256 function get the SHA2-256 Hash from a given bytes.Buffer and
// returns the result also in bytes.Buffer
func Sha256(data *bytes.Buffer) *bytes.Buffer {
	m := sha256.New()
	io.Copy(m, data)
	return bytes.NewBuffer(m.Sum(nil))
}

// Sha224 function get the SHA2-224 Hash from a given bytes.Buffer and
// returns the result also in bytes.Buffer
func Sha224(data *bytes.Buffer) *bytes.Buffer {
	m := sha256.New224()
	io.Copy(m, data)
	return bytes.NewBuffer(m.Sum(nil))
}

// Sha384 function get the SHA2-384 Hash from a given bytes.Buffer and
// returns the result also in bytes.Buffer
func Sha384(data *bytes.Buffer) *bytes.Buffer {
	m := sha512.New384()
	io.Copy(m, data)
	return bytes.NewBuffer(m.Sum(nil))
}

// Sha512 function get the SHA2-512 Hash from a given bytes.Buffer and
// returns the result also in bytes.Buffer
func Sha512(data *bytes.Buffer) *bytes.Buffer {
	m := sha512.New()
	io.Copy(m, data)
	return bytes.NewBuffer(m.Sum(nil))
}

// Sha512_224 function get the SHA2-512/224 Hash from a given bytes.Buffer and
// returns the result also in bytes.Buffer
func Sha512_224(data *bytes.Buffer) *bytes.Buffer {
	m := sha512.New512_224()
	io.Copy(m, data)
	return bytes.NewBuffer(m.Sum(nil))
}

// Sha512_256 function get the SHA2-512/256 Hash from a given bytes.Buffer and
// returns the result also in bytes.Buffer
func Sha512_256(data *bytes.Buffer) *bytes.Buffer {
	m := sha512.New512_256()
	io.Copy(m, data)
	return bytes.NewBuffer(m.Sum(nil))
}

// Md5 function get the MD5 Hash from a given bytes.Buffer and
// returns the result also in bytes.Buffer
func Md5(data *bytes.Buffer) *bytes.Buffer {
	m := md5.New()
	io.Copy(m, data)
	return bytes.NewBuffer(m.Sum(nil))
}
