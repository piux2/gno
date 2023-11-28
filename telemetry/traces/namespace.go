package traces

import (
	"context"
)

type namespace string

const (
	NamespaceVM     namespace = "vm"
	NamespaceVMInit namespace = "vmInit"
)

var namespaces = make(map[namespace]context.Context)

func InitNamespace(ctx context.Context, ns namespace) {
	if ctx == nil {
		ctx = context.Background()
	}

	namespaces[ns] = ctx
}
