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
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func PDF(db, outDir, texCompiler string) {
	paths, err := filepath.Glob(filepath.Join(db, "*"))
	if err != nil {
		log.Fatal(err)
	}
	if len(paths) == 0 {
		log.WithFields(log.Fields{
			"db":     db,
			"outDir": outDir,
		}).Fatal("no source directories found!")
	}
	for _, p := range paths {
		s, err := readSrcObj(p)
		if err != nil {
			log.WithFields(log.Fields{
				"path": p,
			}).Warn("failed to read source object")
			log.Fatal(err)
		}
		if err := s.toPDF(outDir, texCompiler); err != nil {
			log.WithFields(log.Fields{
				"path":          p,
				"source object": s,
			}).Warn("failed to compile source object to PDF")
			log.Fatal(err)
		}
	}
}
