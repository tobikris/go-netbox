// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"github.com/digitalocean/go-netbox/netbox/client"
)

func main() {
	c := defaultClient()

	rs, err := c.Dcim.DcimRacksList(nil, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v\n", *(rs.Payload.Count))
}

func defaultClient() *client.NetBox {
	// Passing nil to this method results in all default configuration for the client.
	// That is, a client that connects to http://localhost:8000, and using default
	// field formats. (e.g. remove '-' from json field names)
	return client.NewHTTPClient(nil)
}