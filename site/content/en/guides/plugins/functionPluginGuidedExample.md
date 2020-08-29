---
title: "Function example"
linkTitle: "Function example"
type: docs
description: >
    Function example
---

# Function Guided Example for Linux

[kustomize/functions/examples/template-go-nginx]: https://github.com/kubernetes-sigs/kustomize/tree/master/functions/examples/template-go-nginx

This is a (no reading allowed!) 60 second copy/paste guided
example.

Full plugin docs [here](README.md).

This demo uses a function, `template-go-nginx`,
that lives in the [kustomize/functions/examples/template-go-nginx].
This function is a container image. When triggered, it
generates a Deployment and a Service for nginx.

This is a guide to try it without damaging your
current setup.

#### requirements

* linux, docker

## Make a place to work

<!-- @setupDemoDir @testAgainstLatestRelease -->
```shell
# Keeping these separate to avoid cluttering the DEMO dir.
DEMO=$(mktemp -d)
tmpGoPath=$(mktemp -d)
```

## Install kustomize

Need v3.7.0 or later version for what follows. If you want to install it
from the head.

<!-- @installKustomize @testAgainstLatestRelease -->
```shell
GOPATH=$tmpGoPath go install sigs.k8s.io/kustomize/kustomize
```

## Create a kustomization

Make a kustomization directory to
hold all your config:

<!-- @createAppDir @testAgainstLatestRelease -->
```shell
MYAPP=$DEMO/myapp
mkdir -p $MYAPP
```

Make a function config file.

<!-- @fnConfig @testAgainstLatestRelease -->
```shell
cat <<EOF >$MYAPP/nginx.yaml
apiVersion: examples.config.kubernetes.io/v1beta1
kind: Nginx
metadata:
  name: demo
  annotations:
    config.kubernetes.io/function: |
      container:
        image: gcr.io/kustomize-functions/example-nginx:v0.2.0
spec:
  replicas: 4
EOF
```

This function will generate a Deployment with 4 replicas; we'll get to that shortly.

Make a kustomization file referencing the function
config:

<!-- @installKustomize @testAgainstLatestRelease -->
```shell
cat <<EOF >$MYAPP/kustomization.yaml
generators:
- nignx.yaml

namespace: kustomize-test
EOF
```

## Build your app

<!-- @build @testAgainstLatestRelease -->
```shell
$tmpGoPath/bin/kustomize build --enable_alpha_plugins $MYAPP
```

This should emit a Deployment and a Service for Nginx.
The Deployment has 4 replicas.

```yaml
apiVersion: v1
kind: Service
metadata:
  annotations:
    config.kubernetes.io/path: service_demo.yaml
  labels:
    app: nginx
    instance: demo
  name: demo
  namespace: kustomize-test
spec:
  ports:
  - name: http
    port: 80
    targetPort: 80
  selector:
    app: nginx
    instance: demo
---
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    config.kubernetes.io/path: deployment_demo.yaml
  labels:
    app: nginx
    instance: demo
  name: demo
  namespace: kustomize-test
spec:
  replicas: 4
  selector:
    matchLabels:
      app: nginx
      instance: demo
  template:
    metadata:
      labels:
        app: nginx
        instance: demo
    spec:
      containers:
      - image: nginx:1.7.9
        name: nginx
        ports:
        - containerPort: 80
```