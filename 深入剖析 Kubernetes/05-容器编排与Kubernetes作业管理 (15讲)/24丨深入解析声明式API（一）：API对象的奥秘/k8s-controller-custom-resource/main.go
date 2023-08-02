package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	clientset "github.com/resouer/k8s-controller-custom-resource/pkg/client/clientset/versioned"
	informers "github.com/resouer/k8s-controller-custom-resource/pkg/client/informers/externalversions"
	"github.com/resouer/k8s-controller-custom-resource/pkg/signals"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	flag.Parse()

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	//1、根据提供的 Master 配置（APIServer 的地址端口和 kubeconfig 的路径），创建一个 Kubernetes 的 client（kubeClient）和
	//Network 对象的 client（networkClient）。
	//如果没有提供 Master 配置，则会直接使用一种名叫InClusterConfig的方式来创建client
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	networkClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %s", err.Error())
	}

	//Network Informer 正是使用这个 networkClient，跟 APIServer 建立了连接。
	//不过，真正负责维护这个连接的，则是 Informer 所使用的 Reflector 包。
	networkInformerFactory := informers.NewSharedInformerFactory(networkClient, time.Second*30)

	controller := NewController(kubeClient, networkClient,
		networkInformerFactory.Samplecrd().V1().Networks())

	go networkInformerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}
