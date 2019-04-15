/*
 * Copyright 2019 gotp
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package name_resolver

import (
	"google.golang.org/grpc/resolver"
	"github.com/gotp/template_server/config"
)

func init() {
	resolver.Register(NewLocalResolver())
}

// NewResolver creates a new resolver builder
func NewLocalResolver() *LocalResolver {
	return &LocalResolver{
		scheme: "local",
	}
}

// Resolver is also a resolver builder.
// It's build() function always returns itself.
type LocalResolver struct {
	scheme string
	// Fields actually belong to the resolver.
	clientConn	resolver.ClientConn
}

// Build returns itself for Resolver, because it's both a builder and a resolver.
func (this *LocalResolver) Build(target resolver.Target, clientConn resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	this.clientConn = clientConn
	addrs, found := config.GetRouterTable().FindAddressByName(target.Endpoint)
	if (found) {
		this.clientConn.NewAddress(buildAddress(addrs))
	} else {
		return nil, nil
	}
	return this, nil
}

// Scheme returns the test scheme.
func (this *LocalResolver) Scheme() string {
	return this.scheme
}

// ResolveNow is a noop for Resolver.
func (*LocalResolver) ResolveNow(o resolver.ResolveNowOption) {}

// Close is a noop for Resolver.
func (*LocalResolver) Close() {}

func buildAddress(strAddrs []string) []resolver.Address {
	var addrs []resolver.Address
	for _, strAddr := range strAddrs {
		addrs = append(addrs, resolver.Address{Addr: strAddr})
	}
	return addrs
}
