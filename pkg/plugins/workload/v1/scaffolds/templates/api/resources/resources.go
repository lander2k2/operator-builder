package resources

import (
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"

	workloadv1 "gitlab.eng.vmware.com/landerr/operator-builder/pkg/workload/v1"
)

var _ machinery.Template = &Resources{}

// Types scaffolds the main package for the companion CLI
type Resources struct {
	machinery.TemplateMixin
	machinery.BoilerplateMixin
	machinery.RepositoryMixin
	machinery.ResourceMixin

	ClusterScoped bool
	SourceFile    workloadv1.SourceFile
	PackageName   string
}

// SetTemplateDefaults implements file.Template
func (f *Resources) SetTemplateDefaults() error {

	f.Path = filepath.Join(
		"apis",
		f.Resource.Group,
		f.Resource.Version,
		f.PackageName,
		f.SourceFile.Filename,
	)

	f.TemplateBody = resourcesTemplate
	f.IfExistsAction = machinery.OverwriteFile

	return nil
}

const resourcesTemplate = `{{ .Boilerplate }}

package {{ .PackageName }}

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	{{ .Resource.ImportAlias }} "{{ .Resource.Path }}"
)

{{ range .SourceFile.Children }}
// Create{{ .UniqueName }} creates the {{ .Name }} {{ .Kind }} resource
func Create{{ .UniqueName }} (parent *{{ $.Resource.ImportAlias }}.{{ $.Resource.Kind }}) (metav1.Object, error) {

	{{ .SourceCode }}

	{{ if not $.ClusterScoped }}
	resourceObj.SetNamespace(parent.Namespace)
	{{ end }}

	return resourceObj, nil
}
{{ end }}
`
