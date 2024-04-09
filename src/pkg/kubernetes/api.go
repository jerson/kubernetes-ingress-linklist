package kubernetes

import (
	"context"
	"sort"

	v1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ByHost []networkingv1.Ingress

func (a ByHost) Len() int           { return len(a) }
func (a ByHost) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByHost) Less(i, j int) bool { return a[i].Name < a[j].Name }

type ByName []v1.Namespace

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

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
	sort.Sort(ByName(output.Items))
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
	sort.Sort(ByHost(output.Items))
	return output, nil
}
