package watcher

import (
	"context"
	"flag"
	"path/filepath"

	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s "k8s.io/client-go/kubernetes"
)

var (
	k8sClient *k8s.Clientset
)

type ConfigMapWatcher struct {
	configMapName string
	namespace     string
	ConfigMap     *v1.ConfigMap
	Watcher       <-chan watch.Event
}

func getClient() *k8s.Clientset {
	if k8sClient != nil {
		return k8sClient
	}

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	k8sClient, err = k8s.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return k8sClient
}

func NewConfigMapWatcher(configMapName, namespace string) (*ConfigMapWatcher, error) {
	k8sClient := getClient()
	watcher, err := k8sClient.CoreV1().ConfigMaps(namespace).Watch(context.Background(),
		metav1.SingleObject(metav1.ObjectMeta{Name: configMapName, Namespace: namespace}))
	if err != nil {
		return nil, err
	}

	cm, err := k8sClient.CoreV1().ConfigMaps(namespace).Get(context.Background(), configMapName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return &ConfigMapWatcher{
		ConfigMap:     cm,
		configMapName: configMapName,
		namespace:     namespace,
		Watcher:       watcher.ResultChan(),
	}, nil
}
