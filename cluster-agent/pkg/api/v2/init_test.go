package v2_test

import "github.com/canonical/microk8s/cluster-agent/pkg/util"

func init() {
	util.SnapData = "testdata"
	util.Snap = "testdata"
}