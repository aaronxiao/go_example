// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package terrors

import (
	"errors"
	"testing"
)

func TestNewEqual(t *testing.T) {
	// Different allocations should not be equal.
	if errors.New("abc") == errors.New("abc") {		//返回的是地址  这里是不同的两个对象的地址  所以会不相等
		t.Errorf(`New("abc") == New("abc")`)
	}
	if errors.New("abc") == errors.New("xyz") {
		t.Errorf(`New("abc") == New("xyz")`)
	}

	// Same allocation should be equal to itself (not crash).
	err := errors.New("jkl")
	if err != err {					//自己等于自己
		t.Errorf(`err != err`)
	}
}

func TestErrorMethod(t *testing.T) {
	err := errors.New("abc")
	if err.Error() != "abc" {			//相等字符串 肯定相等
		t.Errorf(`New("abc").Error() = %q, want %q`, err.Error(), "abc")
	}
}


