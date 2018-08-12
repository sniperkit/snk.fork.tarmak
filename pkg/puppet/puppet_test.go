/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright Jetstack Ltd. See LICENSE for details.
package puppet

import (
	"strings"
	"testing"

	clusterv1alpha1 "github.com/sniperkit/snk.fork.tarmak/pkg/apis/cluster/v1alpha1"
)

func TestOIDCFields(t *testing.T) {
	c := clusterv1alpha1.ClusterKubernetes{
		APIServer: &clusterv1alpha1.ClusterKubernetesAPIServer{
			OIDC: &clusterv1alpha1.ClusterKubernetesAPIServerOIDC{
				IssuerURL:      "http://123",
				ClientID:       "client_id",
				SigningAlgs:    []string{"alg1", "alg2"},
				GroupsPrefix:   "groups-prefix",
				GroupsClaim:    "groups-claim",
				UsernamePrefix: "username-prefix",
				UsernameClaim:  "username-claim",
			},
		},
	}

	d := hieraData{}

	kubernetesClusterConfig(&c, &d)

	count := 0
	for _, v := range d.variables {
		if strings.Contains(v, "oidc") {
			count += 1
		}
	}

	if act, exp := count, 7; act != exp {
		t.Fatalf("unexpected number of variables: exp:%d act:%d", exp, act)
	}

}
