/*
Copyright 2021.

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
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"

	jobv1alpha1 "github.com/spartan-operators/job-operator/api/v1alpha1"
)

// VrTestJobReconciler reconciles a VrTestJob object
type VrTestJobReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=job.vr.fmwk.com,resources=vrtestjobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=job.vr.fmwk.com,resources=vrtestjobs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=job.vr.fmwk.com,resources=vrtestjobs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the VrTestJob object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.9.2/pkg/reconcile
func (r *VrTestJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Fetch the VrTestJob Resource
	vrTestJob := &jobv1alpha1.VrTestJob{}
	err := r.Get(ctx, req.NamespacedName, vrTestJob)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Info("VrTestJob resource not found. Ignoring since the object will be deleted")
			return ctrl.Result{}, nil
		}
	}

	// Check if the Job already exists, If not create a new one
	found := &batchv1.Job{}
	err = r.Get(ctx, types.NamespacedName{
		Namespace: vrTestJob.Namespace,
		Name:      vrTestJob.Name,
	}, found)

	if err != nil && errors.IsNotFound(err) {
		// create the K8s JobSpecs
	} else {
		log.Error(err, "Failed to get the Job Spec")
		return ctrl.Result{}, err
	}

	//reconcile size for Job replicas.
	//update CR status with the pod names

	return ctrl.Result{}, nil
}

func testJobDescriptor(vrTestJob *jobv1alpha1.VrTestJob) *batchv1.Job {
	metadata := &metav1.ObjectMeta{
		Name:        "batch-job",
		Namespace:   "default",
		Labels: 	 map[string]string{},
		Annotations: map[string]string{},
	}

	podSpec := &apiv1.PodSpec{
		Containers: 	[]apiv1.Container {
			{
				Name:                     "batch-job",
				Image:                    "luksa/batch-job:latest",
				Command:                  []string{},
				Args:                     []string{},
			},
		},
		RestartPolicy:	apiv1.RestartPolicyOnFailure,
	}

	podTemplate := &apiv1.PodTemplateSpec{
		ObjectMeta: *metadata,
		Spec:       *podSpec,
	}

	jobSpec := &batchv1.JobSpec{
		Template:  *podTemplate,
	}

	return &batchv1.Job{
		ObjectMeta: *metadata,
		Spec:       *jobSpec,
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *VrTestJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&jobv1alpha1.VrTestJob{}).
		Complete(r)
}
