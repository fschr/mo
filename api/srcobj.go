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

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var exts = [...]string{".txt", ".tex"}

type srcObj struct {
	path    string
	meta    *metadata
	subObjs [](*subObj)
}

type metadata struct {
	title string
	tags  []string
}

type subObj struct {
	problemSrc  string
	solutionSrc string
}

func readMetadata(path string) (m *metadata, err error) {
	t := struct {
		Title string
		Tags  []string
	}{}

	buf, err := ioutil.ReadFile(filepath.Join(path, "metadata.yml"))
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(buf, &t); err != nil {
		return nil, err
	}
	return &metadata{title: t.Title, tags: t.Tags}, nil
}

func readSubObj(path, ext string) (s *subObj, err error) {
	srcLoader := func(basename string) (src string, err error) {
		fullpath := filepath.Join(path, basename)
		buf, err := ioutil.ReadFile(fullpath)
		if err != nil {
			return "", err
		}
		return string(buf), nil
	}
	problemSrc, err := srcLoader(fmt.Sprintf("problem%v", ext))
	if err != nil {
		return nil, err
	}
	solutionSrc, err := srcLoader(fmt.Sprintf("solution%v", ext))
	if err != nil {
		return nil, err
	}
	return &subObj{problemSrc: problemSrc, solutionSrc: solutionSrc}, nil
}

func readSrcObj(path string) (s *srcObj, err error) {
	subObjs := make([](*subObj), 0)
	for _, ext := range exts {
		subObj, err := readSubObj(path, ext)
		if err != nil {
			log.WithFields(log.Fields{
				"path": path,
				"ext":  ext,
			}).Warn("failed to read sub obj")
			return nil, err
		}
		subObjs = append(subObjs, subObj)
	}
	meta, err := readMetadata(path)
	if err != nil {
		log.WithFields(log.Fields{"path": path}).Warn("failed to read metadata")
		return nil, err
	}
	return &srcObj{path: path, meta: meta, subObjs: subObjs}, nil
}

func (s *srcObj) String() string {
	if s.meta != nil {
		return s.meta.title
	}
	return "no metadata associated with this problem"
}

func (s *srcObj) toPDF(outDir, texCompiler string) error {
	fmt.Printf("%v: generated PDF in %v using %v\n", s, outDir, texCompiler)
	return nil
}
