package main

import (
	"github.com/pkg/errors"
	"github.com/project-planton/argocd-kubernetes-pulumi-module/pkg"
	argocdkubernetesv1 "github.com/project-planton/project-planton/apis/go/project/planton/provider/kubernetes/argocdkubernetes/v1"
	"github.com/project-planton/project-planton/pkg/pulmod/stackinput"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInput := &argocdkubernetesv1.ArgocdKubernetesStackInput{}

		if err := stackinput.LoadStackInput(ctx, stackInput); err != nil {
			return errors.Wrap(err, "failed to load stack-input")
		}

		return pkg.Resources(ctx, stackInput)
	})
}
