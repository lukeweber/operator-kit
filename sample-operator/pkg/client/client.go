package client

import (
	"k8s.io/client-go/rest"
	"github.com/rook/operator-kit/sample-operator/pkg/generated/clientset/versioned"
)

// TODO: Copied from etcd operator //

func MustNewInCluster() versioned.Interface {
	cfg, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}
	return MustNew(cfg)
}

func MustNew(cfg *rest.Config) versioned.Interface {
	cli, err := versioned.NewForConfig(cfg)
	if err != nil {
		panic(err)
	}
	return cli
}

// TODO: end Copied from etcd operator //
