// Copyright © 2016 Thomas Fischer <tdf.tomfischer@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import "testing"

func TestReadSubObj(t *testing.T) {
	testcases := []struct {
		path        string
		ext         string
		expected    subObj
		expectedErr error
	}{
		{
			path: "testdb/0001",
			ext:  ".txt",
			expected: subObj{
				problemSrc:  "test problem.txt text\n",
				solutionSrc: "test solution.txt text\n",
			},
			expectedErr: nil,
		},
		{
			path: "testdb/0001",
			ext:  ".tex",
			expected: subObj{
				problemSrc:  "test problem.tex tex\n",
				solutionSrc: "test solution.tex tex\n",
			},
			expectedErr: nil,
		},
	}
	for i, tc := range testcases {
		got, err := readSubObj(tc.path, tc.ext)
		if err != tc.expectedErr {
			t.Errorf("%v: expected err %v, got err %v", i, tc.expectedErr, err)
		}
		if *got != tc.expected {
			t.Errorf("%v: expected %v, got %v", i, tc.expected, got)
		}
	}
}

func TestReadMetadata(t *testing.T) {
	testcases := []struct {
		path        string
		expected    metadata
		expectedErr error
	}{
		{
			path: "testdb/0001",
			expected: metadata{
				title: "Test Title",
				tags: []string{
					"test tag 1",
					"test tag 2",
					"test tag 3",
				},
			},
			expectedErr: nil,
		},
	}
	for i, tc := range testcases {
		got, err := readMetadata(tc.path)
		if err != tc.expectedErr {
			t.Errorf("%v: expected err %v, got err %v", i, tc.expectedErr, err)
		}
		if got.title != tc.expected.title {
			t.Errorf("%v: expected %v, got %v", i, tc.expected.title, got.title)
		}
		if len(got.tags) != len(tc.expected.tags) {
			t.Errorf("%v: expected %v, got %v", i, tc.expected.tags, got.tags)
		}
		for j, gotTag := range got.tags {
			if gotTag != tc.expected.tags[j] {
				t.Errorf("%v: expected %v, got %v", i, tc.expected.tags, got.tags)
			}
		}
	}
}

func TestReadSrcObj(t *testing.T) {
	testcases := []struct {
		path        string
		expected    srcObj
		expectedErr error
	}{
		{
			path: "testdb/0001",
			expected: srcObj{
				meta: &metadata{
					title: "Test Title",
					tags: []string{
						"test tag 1",
						"test tag 2",
						"test tag 3",
					},
				},
				subObjs: [](*subObj){
					&subObj{
						problemSrc:  "test problem.txt text\n",
						solutionSrc: "test solution.txt text\n",
					},
					&subObj{
						problemSrc:  "test problem.tex tex\n",
						solutionSrc: "test solution.tex tex\n",
					},
				},
			},
			expectedErr: nil,
		},
	}
	for i, tc := range testcases {
		got, err := readSrcObj(tc.path)
		if err != tc.expectedErr {
			t.Errorf("%v: expected err %v, got err %v", i, tc.expectedErr, err)
		}
		if got.meta.title != tc.expected.meta.title {
			t.Errorf("%v: expected %v, got %v", i, tc.expected.meta.title, got.meta.title)
		}
		if len(got.meta.tags) != len(tc.expected.meta.tags) {
			t.Errorf("%v: expected %v, got %v", i, tc.expected.meta.tags, got.meta.tags)
		}
		for j, gotTag := range got.meta.tags {
			if gotTag != tc.expected.meta.tags[j] {
				t.Errorf("%v: expected %v, got %v", i, tc.expected.meta.tags, got.meta.tags)
			}
		}
		if len(got.subObjs) != len(tc.expected.subObjs) {
			t.Errorf("%v: expected %v, got %v", i, tc.expected.subObjs, got.subObjs)
		}
		expectedSet := make(map[subObj]struct{})
		for _, expectedSubObj := range tc.expected.subObjs {
			expectedSet[*expectedSubObj] = struct{}{}
		}
		for _, gotSubObj := range got.subObjs {
			_, ok := expectedSet[*gotSubObj]
			if !ok {
				t.Errorf("%v: expected %v, got %v", i, tc.expected.subObjs, got.subObjs)
			}
		}
	}
}
