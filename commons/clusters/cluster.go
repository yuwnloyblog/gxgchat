package clusters

import "github.com/yuwnloyblog/gmicro"

var ImCluster *gmicro.Cluster

func init() {
	ImCluster = gmicro.NewSingleCluster("imcluster", "testNode")
}
