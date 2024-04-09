package kubernetes

import (
	"context"

	v1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NamespacesOutput ...
type NamespacesOutput struct {
	Items []v1.Namespace `json:"items"`
}

// GetNamespaces ...
func (k8s *Kubernetes) GetNamespaces() (NamespacesOutput, error) {
	k8s.init()
	output := NamespacesOutput{}
	namespaces, err := k8s.clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return output, err
	}

	output.Items = namespaces.Items
	return output, nil
}

// IngressesOutput ...
type IngressesOutput struct {
	Items []networkingv1.Ingress `json:"items"`
}

// GetIngresses ...
func (k8s *Kubernetes) GetIngresses() (IngressesOutput, error) {
	k8s.init()
	output := IngressesOutput{}

	ingresses, err := k8s.clientset.NetworkingV1().Ingresses(k8s.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return output, err
	}

	output.Items = ingresses.Items
	return output, nil
}
