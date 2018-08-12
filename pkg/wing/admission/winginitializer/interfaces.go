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

// WantsInternalWingInformerFactory defines a function which sets InformerFactory for admission plugins that need it
type WantsInternalWingInformerFactory interface {
	SetInternalWingInformerFactory(informers.SharedInformerFactory)
	admission.InitializationValidator
}
