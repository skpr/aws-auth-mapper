package main

import (
	"flag"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	iamauthenticatorv1beta1 "github.com/skpr/aws-auth-mapper/apis/iamauthenticator/v1beta1"
	iamauthenticatorcmaprole "github.com/skpr/aws-auth-mapper/controllers/iamauthenticator/maprole"
	iamauthenticatorcmapuser "github.com/skpr/aws-auth-mapper/controllers/iamauthenticator/mapuser"
	// +kubebuilder:scaffold:imports
)

const (
	// ConfigMapNamespace which will be updated.
	ConfigMapNamespace = "kube-system"
	// ConfigMapName which will be updated.
	ConfigMapName = "aws-auth"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	if err := clientgoscheme.AddToScheme(scheme); err != nil {
		panic(err)
	}

	if err := iamauthenticatorv1beta1.AddToScheme(scheme); err != nil {
		panic(err)
	}
	// +kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		Port:               9443,
		LeaderElection:     enableLeaderElection,
		LeaderElectionID:   "8d440429.skpr.io",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if err = (&iamauthenticatorcmaprole.Reconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("MapUser"),
		Scheme: mgr.GetScheme(),
		ConfigMap: types.NamespacedName{
			Namespace: ConfigMapNamespace,
			Name:           ConfigMapName,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "MapUser")
		os.Exit(1)
	}

	if err = (&iamauthenticatorcmapuser.Reconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("MapRole"),
		Scheme: mgr.GetScheme(),
		ConfigMap: types.NamespacedName{
			Namespace: ConfigMapNamespace,
			Name:      ConfigMapName,
		},
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "MapRole")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
