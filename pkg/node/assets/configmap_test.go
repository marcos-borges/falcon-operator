package assets

import (
	"testing"

	"github.com/crowdstrike/falcon-operator/apis/falcon/v1alpha1"
	"github.com/crowdstrike/falcon-operator/pkg/common"
	"github.com/crowdstrike/falcon-operator/pkg/node"
	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestDaemonsetConfigMap(t *testing.T) {
	falconNode := v1alpha1.FalconNodeSensor{}
	falconNode.Spec.FalconAPI = nil
	falconCID := "1234567890ABCDEF1234567890ABCDEF-12"
	falconImage := "testMyImage"
	cfgName := "test"

	want := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cfgName,
			Namespace: cfgName,
			Labels: map[string]string{
				common.FalconInstanceNameKey: "configmap",
				common.FalconInstanceKey:     cfgName,
				common.FalconComponentKey:    common.FalconKernelSensor,
				common.FalconManagedByKey:    common.FalconManagedByValue,
				common.FalconProviderKey:     common.FalconProviderValue,
				common.FalconPartOfKey:       common.FalconPartOfValue,
				common.FalconCreatedKey:      common.FalconCreatedValue,
			},
		},
		Data: map[string]string{
			"FALCONCTL_OPT_CID": "1234567890ABCDEF1234567890ABCDEF-12",
		},
	}

	config := node.ConfigCacheTest(falconCID, falconImage, &falconNode)

	got := DaemonsetConfigMap("test", "test", config)
	if diff := cmp.Diff(&want, &got); diff != "" {
		t.Errorf("getTermGracePeriod() mismatch (-want +got): %s", diff)
	}
}
