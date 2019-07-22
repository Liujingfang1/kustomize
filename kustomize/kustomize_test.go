package kustomize

import (
	"fmt"
	"testing"

	"sigs.k8s.io/kustomize/v3/pkg/types"
)

func TestKustomizeBytes(t *testing.T) {
	data := `
apiVersion: v1
kind: MyKind
metadata:
  name: nm
spec:
  foo: bar
`
	k := types.Kustomization{
		CommonLabels: map[string]string{
			"foo": "bar",
		},
		NamePrefix: "pre-",
		NameSuffix: "-suffix",
		Namespace:  "ns",
	}
	output, err := KustomizeBytes(k, []byte(data))
	if err != nil {
		t.Fatalf("unexpected error")
	}

	expected := `
apiVersion: v1
kind: MyKind
metadata:
  name: pre-nm-suffix
  namespace: ns
  labels:
    foo: bar
`
	if expected != string(output) {
		fmt.Printf("got\n%s\n", string(output))
	}
}

func TestKustomizeBytes2(t *testing.T) {
	data := `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: demo
  labels:
    app: demo
spec:
  selector:
    matchLabels:
      app: demo
  template:
    metadata:
      labels:
        app: demo
    spec:
      containers:
        - name: demo
          image: demo
          ports:
            - containerPort: 8080
          volumeMounts:
            - name: demo-config
              mountPath: /config
      volumes:
        - name: "demo-config"
          configMap:
            name: demo-configmap
`
	k := types.Kustomization{
		CommonLabels: map[string]string{
			"foo": "bar",
		},
		NamePrefix: "pre-",
		NameSuffix: "-suffix",
		Namespace:  "ns",
		ConfigMapGenerator: []types.ConfigMapArgs{
			{
				GeneratorArgs: types.GeneratorArgs{
					Name: "demo-configmap",
					DataSources: types.DataSources{
						LiteralSources: []string{
							"foo=bar",
						},
					},
				},
			},
		},
	}
	output, err := KustomizeBytes(k, []byte(data))
	if err != nil {
		t.Fatalf("unexpected error")
	}

	expected := `
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: demo
    foo: bar
  name: pre-demo-suffix
  namespace: ns
spec:
  selector:
    matchLabels:
      app: demo
      foo: bar
  template:
    metadata:
      labels:
        app: demo
        foo: bar
    spec:
      containers:
      - image: demo
        name: demo
        ports:
        - containerPort: 8080
        volumeMounts:
        - mountPath: /config
          name: demo-config
      volumes:
      - configMap:
          name: pre-demo-configmap-suffix-2mm7gfhtgt
        name: demo-config
---
apiVersion: v1
data:
  foo: bar
kind: ConfigMap
metadata:
  labels:
    foo: bar
  name: pre-demo-configmap-suffix-2mm7gfhtgt
  namespace: ns
`
	if expected != string(output) {
		fmt.Printf("got\n%s\n", string(output))
	}
}
