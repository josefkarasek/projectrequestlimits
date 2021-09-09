package webhook

import (
	"context"
	"fmt"
	"net/http"

	projectv1 "github.com/openshift/api/project/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	"sigs.k8s.io/yaml"
)

// +kubebuilder:webhook:path=/validate-v1-project,mutating=false,failurePolicy=fail,groups="project.openshift.io",resources=projects,verbs=create,versions=v1,name=vproject.kb.io,sideEffects=none,admissionReviewVersions=v1

// ProjectRequestLimitsValidator
type ProjectRequestLimitsValidator struct {
	Client  client.Client
	decoder *admission.Decoder
}

// ProjectRequestLimitsValidator
func (a *ProjectRequestLimitsValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	project := &projectv1.Project{}

	err := a.decoder.Decode(req, project)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	d, _ := yaml.Marshal(project)
	fmt.Println(string(d))

	return admission.Denied("Everyone is denied")
}

// InjectDecoder injects the decoder.
func (a *ProjectRequestLimitsValidator) InjectDecoder(d *admission.Decoder) error {
	a.decoder = d
	return nil
}
