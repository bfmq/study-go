/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	ingressv1beta1 "crd/kubebuilder-demo/api/v1beta1"
	"crd/kubebuilder-demo/controllers/utils"
)

// AppReconciler reconciles a App object
type AppReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=ingress.bfmq.com,resources=apps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ingress.bfmq.com,resources=apps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ingress.bfmq.com,resources=apps/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the App object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.2/pkg/reconcile
func (r *AppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// TODO(user): your logic here
	app := &ingressv1beta1.App{}
	err := r.Get(ctx, req.NamespacedName, app)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Deployment的处理
	depoyment := utils.NewDeployment(app)
	err = controllerutil.SetControllerReference(app, depoyment, r.Scheme)
	if err != nil {
		return ctrl.Result{}, err
	}
	d := &appsv1.Deployment{}
	err = r.Get(ctx, req.NamespacedName, d)
	if errors.IsNotFound(err) {
		err = r.Create(ctx, depoyment)
		if err != nil {
			logger.Error(err, "create deploy failed")
			return ctrl.Result{}, err
		}
	} else {
		err = r.Update(ctx, depoyment)
		if err != nil {
			logger.Error(err, "update deploy failed")
			return ctrl.Result{}, err
		}
	}

	// Service的处理
	service := utils.NewService(app)
	err = controllerutil.SetControllerReference(app, service, r.Scheme)
	if err != nil {
		return ctrl.Result{}, err
	}
	s := &corev1.Service{}
	err = r.Get(ctx, req.NamespacedName, s)
	if errors.IsNotFound(err) {
		if app.Spec.EnabledService {
			err = r.Create(ctx, service)
			if err != nil {
				logger.Error(err, "create service failed")
				return ctrl.Result{}, err
			}
		}
	} else {
		if app.Spec.EnabledService {
			logger.Info("skip service update")
			// err = r.Update(ctx, service)
			// if err != nil {
			// 	logger.Error(err, "update service failed")
			// 	return ctrl.Result{}, err
			// }
		} else {
			err = r.Delete(ctx, service)
			if err != nil {
				logger.Error(err, "delete service failed")
				return ctrl.Result{}, err
			}
		}
	}

	// Ingress的处理
	ingress := utils.NewIngress(app)
	err = controllerutil.SetControllerReference(app, ingress, r.Scheme)
	if err != nil {
		return ctrl.Result{}, err
	}
	i := &netv1.Ingress{}
	err = r.Get(ctx, req.NamespacedName, i)
	if errors.IsNotFound(err) {
		if app.Spec.EnabledIngress {
			err = r.Create(ctx, ingress)
			if err != nil {
				logger.Error(err, "create ingress failed")
				return ctrl.Result{}, err
			}
		}
	} else {
		if app.Spec.EnabledIngress {
			logger.Info("skip ingress update")
			// err = r.Update(ctx, ingress)
			// if err != nil {
			// 	logger.Error(err, "update ingress failed")
			// 	return ctrl.Result{}, err
			// }
		} else {
			err = r.Delete(ctx, ingress)
			if err != nil {
				logger.Error(err, "delete ingress failed")
				return ctrl.Result{}, err
			}
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ingressv1beta1.App{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Owns(&netv1.Ingress{}).
		Complete(r)
}
