// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package apachepulsar

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "apachepulsar", asset.ModuleFieldsPri, AssetApachepulsar); err != nil {
		panic(err)
	}
}

// AssetApachepulsar returns asset data.
// This is the base64 encoded gzipped contents of module/apachepulsar.
func AssetApachepulsar() string {
	return "eJx8j01uwyAUhPecYuS9L8Ciux7k1UxbZDAIsBLfPrJJIvyjfMsBfTOvx8hFQ6IM/4yzy5IUUGxx1OjauFNAoqNkavywiAIM85BsLDZMGl8KwM4EH8zsqIBfS2ey3n70mMTz1LlSlkiNvxTm+EwuGva21uhZkh3yO78SrhzPeHFZVjmrjyPaIbyLj9vhLXXMyOUWkjm8fahe+a7CWqoeAQAA//+DEnfJ"
}
