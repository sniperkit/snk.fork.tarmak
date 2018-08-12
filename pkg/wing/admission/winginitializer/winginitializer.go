/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright Jetstack Ltd. See LICENSE for details.

package winginitializer

import (
	"k8s.io/apiserver/pkg/admission"

	informers "github.com/sniperkit/snk.fork.tarmak/pkg/wing/informers/internalversion"
)

type pluginInitializer struct {
	informers informers.SharedInformerFactory
}

var _ admission.PluginInitializer = pluginInitializer{}

// New creates an instance of wing admission plugins initializer.
func New(informers informers.SharedInformerFactory) (pluginInitializer, error) {
	return pluginInitializer{
		informers: informers,
	}, nil
}

// Initialize checks the initialization interfaces implemented by a plugin
// and provide the appropriate initialization data
func (i pluginInitializer) Initialize(plugin admission.Interface) {
	if wants, ok := plugin.(WantsInternalWingInformerFactory); ok {
		wants.SetInternalWingInformerFactory(i.informers)
	}
}
