/*

 */

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	apisv1 "github.com/erikh/k8s-api/api/v1"
)

// UUIDReconciler reconciles a UUID object
type UUIDReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=apis.example.org,resources=uuids,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apis.example.org,resources=uuids/status,verbs=get;update;patch

func (r *UUIDReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("uuid", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *UUIDReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apisv1.UUID{}).
		Complete(r)
}
