// Copyright 2018 @boseji <salearj@hotmail.com> All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is specially dedicated to Other utilities packaged with mserver

package mserver

import (
	"bytes"
	"fmt"
)

// BufToHexString converts the bytes.Buffer into a Hexadecimal string
func BufToHexString(data *bytes.Buffer) string {
	return fmt.Sprintf("%x", data.Bytes())
}
