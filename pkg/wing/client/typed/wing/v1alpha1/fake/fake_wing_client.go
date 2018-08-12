/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright Jetstack Ltd. See LICENSE for details.
package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/sniperkit/snk.fork.tarmak/pkg/wing/client/typed/wing/v1alpha1"
)

type FakeWingV1alpha1 struct {
	*testing.Fake
}

func (c *FakeWingV1alpha1) Instances(namespace string) v1alpha1.InstanceInterface {
	return &FakeInstances{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeWingV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
