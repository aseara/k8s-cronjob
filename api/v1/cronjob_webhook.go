/*
Copyright 2022 aseara.
*/

package v1

import (
	"github.com/robfig/cron"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	validationutils "k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var cronjoblog = logf.Log.WithName("cronjob-resource")

func (r *CronJob) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-batch-aseara-github-com-v1-cronjob,mutating=true,failurePolicy=fail,sideEffects=None,groups=batch.aseara.github.com,resources=cronjobs,verbs=create;update,versions=v1,name=mcronjob.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &CronJob{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *CronJob) Default() {
	cronjoblog.Info("default", "name", r.Name)

	if r.Spec.ConcurrencyPolicy == "" {
		r.Spec.ConcurrencyPolicy = AllowConcurrent
	}
	if r.Spec.Suspend == nil {
		r.Spec.Suspend = new(bool)
	}
	if r.Spec.SuccessfulJobsHistoryLimit == nil {
		r.Spec.SuccessfulJobsHistoryLimit = new(int32)
		*r.Spec.SuccessfulJobsHistoryLimit = 3
	}
	if r.Spec.FailedJobsHistoryLimit == nil {
		r.Spec.FailedJobsHistoryLimit = new(int32)
		*r.Spec.FailedJobsHistoryLimit = 1
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-batch-aseara-github-com-v1-cronjob,mutating=false,failurePolicy=fail,sideEffects=None,groups=batch.aseara.github.com,resources=cronjobs,verbs=create;update;delete,versions=v1,name=vcronjob.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &CronJob{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *CronJob) ValidateCreate() error {
	cronjoblog.Info("validate create", "name", r.Name)

	return r.validateCronJob()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *CronJob) ValidateUpdate(old runtime.Object) error {
	cronjoblog.Info("validate update", "name", r.Name)

	return r.validateCronJob()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *CronJob) ValidateDelete() error {
	cronjoblog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

func (r *CronJob) validateCronJob() error {
	var allErrs field.ErrorList
	if err := r.validateCronJobName(); err != nil {
		allErrs = append(allErrs, err)
	}
	if err := r.validateCronJobSpec(); err != nil {
		allErrs = append(allErrs, err)
	}
	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "", Kind: "CronJob"},
		r.Name, allErrs)
}

func (r *CronJob) validateCronJobName() *field.Error {
	if len(r.ObjectMeta.Name) > validationutils.DNS1035LabelMaxLength-11 {
		return field.Invalid(field.NewPath("metadata").Child("name"), r.Name, "must be no more than 52 characters")
	}
	return nil
}

func (r *CronJob) validateCronJobSpec() *field.Error {
	// The field helpers from the kubernetes API machinery help us return nicely
	// structed validation errors.
	return validateScheduleFormat(
		r.Spec.Schedule,
		field.NewPath("spec").Child("schedule"))
}

func validateScheduleFormat(schedule string, fldPath *field.Path) *field.Error {
	if _, err := cron.ParseStandard(schedule); err != nil {
		return field.Invalid(fldPath, schedule, err.Error())
	}
	return nil
}
