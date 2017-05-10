// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

package main

import (
	"bitbucket.org/mexisme/build-dotenv/cmd"
	log "github.com/sirupsen/logrus"
)

/*
  $0 discover
    - Figure out how to get secrets, and present them to the app
      - Read config (i.e. Config file, as well as Env Vars, via Viper) to determine where the secrets are stored
      - Use the correct library to get them
      - Validate / convert to an internal representation -- use Viper + GoDotEnv?
      - Create a .env file in /secrets, as well as export them as Env Vars?
  $0 s3 --s3path ${S3-Path}
  $0 kms --kms-lookup ${KMS Path(s)?}
  $0 parameters --parameters-lookup ${KMS Path(s)?}
  $0 vault --secrets-lookup ${???}
     - Viper is used to get settings from a local config file, flags, env-vars, etc. The above CLI flags are representations of this.
     - All the above support [--dotenv-file ${Path}] [--export-env] for the export options
*/
func main() {
	log.Debug("Starting...")
	cmd.Execute()
}
