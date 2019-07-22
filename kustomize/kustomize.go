package kustomize

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/kustomize/v3/internal/loadertest"
	"sigs.k8s.io/kustomize/v3/k8sdeps/kunstruct"
	"sigs.k8s.io/kustomize/v3/k8sdeps/transformer"
	"sigs.k8s.io/kustomize/v3/pkg/plugins"
	"sigs.k8s.io/kustomize/v3/pkg/resmap"
	"sigs.k8s.io/kustomize/v3/pkg/resource"
	"sigs.k8s.io/kustomize/v3/pkg/target"
	"sigs.k8s.io/kustomize/v3/pkg/types"
	"sigs.k8s.io/yaml"
)

func Kustomize(k types.Kustomization, res []unstructured.Unstructured) ([]unstructured.Unstructured,
	error) {

	// convert to bytes
	data, err := yaml.Marshal(res)
	if err != nil {
		return nil, err
	}

	// Kustomize bytes
	output, err := KustomizeBytes(k, data)

	// convert back to unstructured
	var results []unstructured.Unstructured
	err = yaml.Unmarshal(output, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func KustomizeBytes(k types.Kustomization, data []byte) ([]byte, error) {

	// create kustomization.yaml and resource file with a fake loader
	ldr := loadertest.NewFakeLoader("base")
	err := ldr.AddFile("base/resource.yaml", data)
	if err != nil {
		return nil, err
	}
	k.Resources = []string{"resource.yaml"}

	kcontent, err := yaml.Marshal(k)
	if err != nil {
		return nil, err
	}
	err = ldr.AddFile("base/kustomization.yaml", kcontent)
	if err != nil {
		return nil, err
	}

	// prepare
	uf := kunstruct.NewKunstructuredFactoryImpl()
	pf := transformer.NewFactoryImpl()
	rf := resmap.NewFactory(resource.NewFactory(uf), pf)
	pluginConfig := plugins.DefaultPluginConfig()
	pl := plugins.NewLoader(pluginConfig, rf)
	kt, err := target.NewKustTarget(ldr, rf, pf, pl)
	if err != nil {
		return nil, err
	}

	// Run kustomize with the fakeloader
	rm, err := kt.MakeCustomizedResMap()
	if err != nil {
		return nil, err
	}
	return rm.AsYaml()
}