/*
Copyright 2022 aseara.
*/

package v2

import (
	ctrl "sigs.k8s.io/controller-runtime"
)

// log is for logging in this package.
// var cronjoblog = logf.Log.WithName("cronjob-resource")

func (r *CronJob) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
