AWS Auth Mapper for EKS
=======================

Terraform provider and Kubernetes controller for managing configration for [aws-iam-authenticator](https://github.com/kubernetes-sigs/aws-iam-authenticator)

## What problem are we solving?

AWS IAM Authenticator is configurable using a Kubernetes `ConfigMap`.

https://github.com/kubernetes-sigs/aws-iam-authenticator#eksconfigmap

This approach isn't flexible when trying to orchestrate cluster access since you will most likely have multiple tools controlling what
can access the cluster eg.

* Terraform
* Ansible
* Other Kubernets Controllers

To solve this issue was have AWS Auth Mapper for EKS. This tool allows us to declare our configuration as native Kubernetes objects
(`CustomResourceDefinitions`) and generate the `ConfigMap`.

While AWS IAM Authenticator does provide a mechanism for using CRDs, it has following issues:

* It's an Alpha feature
* You need to redeploy and manage the aws-iam-authenticator yourself.

## Usage

### Terraform

The following example can be used to create Kubernetes objects.

This is particularly handy when provisioning AWS IAM resources and mapping them directly to Kubernetes RBAC objects.

```hcl
resource "aam_iamauthenticator_v1beta1_maprole" "node" {
  name = "node"

  role_arn = "arn:aws:iam::xxxyyyzzz:role/node"
  username = "system:node:{{EC2PrivateDNSName}}"
  groups = [
    "system:bootstrappers",
    "system:nodes",
  ]
}
```

```hcl
resource "aam_iamauthenticator_v1beta1_mapuser" "admin" {
  name = "admin"

  user_arn = "arn:aws:iam::xxxyyyzzz:user/admin"
  username = "cluster-admin"
  groups = [
    "system:masters",
  ]
}
```

### Kubernetes

These Kubernetes objects...

```yaml
apiVersion: iamauthenticator.skpr.io/v1beta1
kind: MapRole
metadata:
  name: node
spec:
  roleARN: arn:aws:iam::xxxyyyzzz:role/node
  username: system:node:{{EC2PrivateDNSName}}
  groups:
    - system:bootstrappers
    - system:nodes
---
apiVersion: iamauthenticator.skpr.io/v1beta1
kind: MapUser
metadata:
  name: admin
spec:
  userARN: arn:aws:iam::xxxyyyzzz:user/admin
  username: cluster-admin
  groups:
    - system:masters
```

result in the this ConfigMap.

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: aws-auth
  namespace: kube-system
data:
  mapRoles: |
    - rolearn: arn:aws:iam::xxxyyyzzz:role/node
      username: system:node:{{EC2PrivateDNSName}}
      groups:
        - system:bootstrappers
        - system:nodes
  mapUsers: |
    - userarn: arn:aws:iam::xxxyyyzzz:user/admin
      username: cluster-admin
      groups:
      - system:masters
```
