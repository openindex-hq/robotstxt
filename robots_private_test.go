// Copyright 2024 Axel Etcheverry
// Copyright 2020 Jim Smart
// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// This file tests the robots.txt parsing and matching code found in robots.cc
// against the current Robots Exclusion Protocol (REP) internet draft (I-D).
// https://tools.ietf.org/html/draft-koster-rep

//

// Converted 2020-04-21, from https://github.com/google/robotstxt/blob/master/robots_test.cc

package robotstxt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPathParamsQuery(t *testing.T) {
	tests := []struct {
		uri      string
		expected string
	}{
		{"", "/"},
		{"http://www.example.com", "/"},
		{"http://www.example.com/", "/"},
		{"http://www.example.com/a", "/a"},
		{"http://www.example.com/a/", "/a/"},
		{"http://www.example.com/a/b?c=http://d.e/", "/a/b?c=http://d.e/"},
		{"http://www.example.com/a/b?c=d&e=f#fragment", "/a/b?c=d&e=f"},
		{"example.com", "/"},
		{"example.com/", "/"},
		{"example.com/a", "/a"},
		{"example.com/a/", "/a/"},
		{"example.com/a/b?c=d&e=f#fragment", "/a/b?c=d&e=f"},
		{"a", "/"},
		{"a/", "/"},
		{"/a", "/a"},
		{"a/b", "/b"},
		{"example.com?a", "/?a"},
		{"example.com/a;b#c", "/a;b"},
		{"//a/b/c", "/b/c"},
	}

	for _, test := range tests {
		t.Run(test.uri, func(t *testing.T) {
			assert.Equal(t, test.expected, getPathParamsQuery(test.uri))
		})
	}
}

func TestMaybeEscapePattern(t *testing.T) {
	tests := []struct {
		uri      string
		expected string
	}{
		{"http://www.example.com", "http://www.example.com"},
		{"/a/b/c", "/a/b/c"},
		{"á", "%C3%A1"},
		{"%aa", "%AA"},
	}

	for _, test := range tests {
		t.Run(test.uri, func(t *testing.T) {
			assert.Equal(t, test.expected, escapePattern(test.uri))
		})
	}
}
