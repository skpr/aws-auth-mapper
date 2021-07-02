package maprole

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	iamauthenticatorv1beta1 "github.com/skpr/aws-auth-mapper/apis/iamauthenticator/v1beta1"
	"github.com/skpr/aws-auth-mapper/internal/configmap"
)

const (
	// FieldName which these mappings will be applied.
	FieldName = "mapRoles"
)

// Reconciler reconciles a MapRole object
type Reconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	// ConfigMap which will be updated by this controller.
	ConfigMap types.NamespacedName
}

// +kubebuilder:rbac:groups=iamauthenticator.skpr.io,resources=maproles,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=iamauthenticator.skpr.io,resources=maproles/status,verbs=get;update;patch

func (r *Reconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()

	log := r.Log.WithValues("maprole", req.NamespacedName)

	log.Info("Starting reconcile loop")

	list := &iamauthenticatorv1beta1.MapRoleList{}

	err := r.List(ctx, list)
	if err != nil {
		return reconcile.Result{}, errors.Wrap(err, "failed to list MapRoles")
	}

	log.Info("Updating ConfigMap")

	var specs []iamauthenticatorv1beta1.MapRoleSpec

	for _, item := range list.Items {
		specs = append(specs, item.Spec)
	}

	data, err := yaml.Marshal(specs)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to marshal ConfigMap data")
	}

	if err := configmap.UpdateDataWithKey(ctx, r, r.ConfigMap, FieldName, data); err != nil {
		return ctrl.Result{}, err
	}

	log.Info("Finished reconcile loop")

	return ctrl.Result{}, nil
}

func (r *Reconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&iamauthenticatorv1beta1.MapRole{}).
		Complete(r)
}
