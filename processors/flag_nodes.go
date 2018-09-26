package processors

import (
	"context"
	"errors"

	"github.com/dmathieu/dice/kubernetes"
	kube "k8s.io/client-go/kubernetes"
)

type FlagNodesProcessor struct {
	kubeClient kube.Interface
}

func (p *FlagNodesProcessor) Process(context.Context) error {
	nodes, err := kubernetes.GetNodes(p.kubeClient, kubernetes.NodeFlagged())
	if err != nil {
		return err
	}
	if len(nodes) > 0 {
		return errors.New("found already flagged nodes. Looks like a roll process is already running")
	}

	nodes, err = kubernetes.GetNodes(p.kubeClient)
	if err != nil {
		return err
	}

	for _, n := range nodes {
		err = kubernetes.FlagNode(p.kubeClient, n)
		if err != nil {
			return err
		}
	}

	return nil
}
