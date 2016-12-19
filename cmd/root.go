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

package cmd

import "github.com/spf13/cobra"

var (
	db          string
	outDir      string
	texCompiler string
	RootCmd     = &cobra.Command{
		Use:   "mo",
		Short: "a magnum opus of quantitative problems and solutions",
		Long: `mo is composed of two things:

	1. A database of quantitative problems and solutions
	2. A tool, namely 'mo', which compiles a given database into various formats`,
	}
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&db, "db", "", "db", "the database directory")
	RootCmd.PersistentFlags().StringVarP(&outDir, "output", "o", "output", "output directory")
	RootCmd.PersistentFlags().StringVarP(&texCompiler, "tex-compiler", "", "pdflatex", "TeX compiler script")
}
