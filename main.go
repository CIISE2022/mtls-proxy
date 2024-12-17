// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/CIISE2022/mtls-proxy/internal/configuration"
	"github.com/CIISE2022/mtls-proxy/internal/httpproxy"
	"github.com/CIISE2022/mtls-proxy/internal/tcpproxy"
)

func main() {

	fmt.Println("Configure")
	cfg := configuration.NewConfiguration()

	time.Local = time.UTC

	tlsConfig := &tls.Config{
		ClientAuth:               tls.RequireAndVerifyClientCert,
		ClientCAs:                cfg.ClientCAPool,
		RootCAs:                  cfg.ClientCAPool,
		MinVersion:               tls.VersionTLS13,
		SessionTicketsDisabled:   true,
		PreferServerCipherSuites: true,
		Certificates:             cfg.ServerCertificates,
	}

	fmt.Println("Start")
	switch cfg.Mode {
	case "http":
		httpproxy.Start(cfg, tlsConfig)
	case "tcp":
		tcpproxy.Start(cfg, tlsConfig)
	}
}
