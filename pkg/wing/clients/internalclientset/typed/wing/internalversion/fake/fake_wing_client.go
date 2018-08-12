/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright Jetstack Ltd. See LICENSE for details.
package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"

	internalversion "github.com/sniperkit/snk.fork.tarmak/pkg/wing/clients/internalclientset/typed/wing/internalversion"
)

type FakeWing struct {
	*testing.Fake
}

func (c *FakeWing) Instances(namespace string) internalversion.InstanceInterface {
	return &FakeInstances{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeWing) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
