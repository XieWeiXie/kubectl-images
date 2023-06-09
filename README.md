<h3 align="center"> Kubectl-images </h3>

<h4 align="center"> Show container images in K8s cluster </h4>


<p align="left"> Usages: </p>

```shell
>> go run kubectlimages.go --n namespace --o table

```

<p align="left"> Output: </p>

```shell
>> go run kubectlimages.go --n kube-system --o table

+-------------+-------------------------+-----------------------------------------------------------+-------------------------------------------------------------------------+-----------------+
|  Namespace  |        Container        |                            Pod                            |                                 Images                                  | ImagePullPolicy |
+-------------+-------------------------+-----------------------------------------------------------+-------------------------------------------------------------------------+-----------------+
| kube-system | coredns                 | coredns-59d64cd4d4-js4hq                                  | registry.aliyuncs.com/google_containers/coredns:v1.8.0                  | IfNotPresent    |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | coredns-59d64cd4d4-kwkhg                                  |                                                                         |                 |
+             +-------------------------+-----------------------------------------------------------+-------------------------------------------------------------------------+                 +
|             | etcd                    | etcd-l-gz-dxc-lyy-local-k8s-master-001                    | registry.aliyuncs.com/google_containers/etcd:3.4.13-0                   |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | etcd-l-gz-dxc-lyy-local-k8s-master-002                    |                                                                         |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | etcd-l-gz-dxc-lyy-local-k8s-master-003                    |                                                                         |                 |
+             +-------------------------+-----------------------------------------------------------+-------------------------------------------------------------------------+                 +
|             | kube-apiserver          | kube-apiserver-l-gz-dxc-lyy-local-k8s-master-001          | registry.aliyuncs.com/google_containers/kube-apiserver:v1.21.6          |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-apiserver-l-gz-dxc-lyy-local-k8s-master-002          |                                                                         |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-apiserver-l-gz-dxc-lyy-local-k8s-master-003          |                                                                         |                 |
+             +-------------------------+-----------------------------------------------------------+-------------------------------------------------------------------------+                 +
|             | kube-controller-manager | kube-controller-manager-l-gz-dxc-lyy-local-k8s-master-001 | registry.aliyuncs.com/google_containers/kube-controller-manager:v1.21.6 |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-controller-manager-l-gz-dxc-lyy-local-k8s-master-002 |                                                                         |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-controller-manager-l-gz-dxc-lyy-local-k8s-master-003 |                                                                         |                 |
+             +-------------------------+-----------------------------------------------------------+-------------------------------------------------------------------------+                 +
|             | kube-proxy              | kube-proxy-2msdv                                          | registry.aliyuncs.com/google_containers/kube-proxy:v1.21.6              |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-proxy-5shjs                                          |                                                                         |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-proxy-6cz7m                                          |                                                                         |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-proxy-6qtq2                                          |                                                                         |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-proxy-8tgdd                                          |                                                                         |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-proxy-9t6zt                                          |                                                                         |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-proxy-9zhvf                                          |                                                                         |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-proxy-dcvkp                                          |                                                                         |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-proxy-fztzc                                          |                                                                         |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-proxy-hjcmp                                          |                                                                         |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-proxy-j5d9s                                          |                                                                         |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-proxy-qgk5f                                          |                                                                         |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-proxy-rrdpb                                          |                                                                         |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-proxy-tkwqc                                          |                                                                         |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-proxy-vfms9                                          |                                                                         |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-proxy-z4xbq                                          |                                                                         |                 |
+             +-------------------------+-----------------------------------------------------------+-------------------------------------------------------------------------+                 +
|             | kube-scheduler          | kube-scheduler-l-gz-dxc-lyy-local-k8s-master-001          | registry.aliyuncs.com/google_containers/kube-scheduler:v1.21.6          |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-scheduler-l-gz-dxc-lyy-local-k8s-master-002          |                                                                         |                 |
+             +                         +-----------------------------------------------------------+                                                                         +                 +
|             |                         | kube-scheduler-l-gz-dxc-lyy-local-k8s-master-003          |                                                                         |                 |
+             +-------------------------+-----------------------------------------------------------+-------------------------------------------------------------------------+                 +
|             | metrics-server          | metrics-server-9dc4489d9-lh99s                            | uhub.service.ucloud.cn/jackhe/metrics-server:v0.5.1                     |                 |
+-------------+-------------------------+-----------------------------------------------------------+-------------------------------------------------------------------------+-----------------+

>> go run kubectlimages.go --n kube-system --o j

[
 {
  "namespace": "kube-system",
  "pod": "coredns-59d64cd4d4-js4hq",
  "container": "coredns",
  "image": "registry.aliyuncs.com/google_containers/coredns:v1.8.0",
  "imagePullPolicy": "IfNotPresent"
 },
 {
  "namespace": "kube-system",
  "pod": "coredns-59d64cd4d4-kwkhg",
  "container": "coredns",
  "image": "registry.aliyuncs.com/google_containers/coredns:v1.8.0",
  "imagePullPolicy": "IfNotPresent"
 }
 ...

]

>> go run kubectlimages.go --n kube-system --o y

- namespace: kube-system
  pod: coredns-59d64cd4d4-js4hq
  container: coredns
  image: registry.aliyuncs.com/google_containers/coredns:v1.8.0
  imagepullpolicy: IfNotPresent
- namespace: kube-system
  pod: coredns-59d64cd4d4-kwkhg
  container: coredns
  image: registry.aliyuncs.com/google_containers/coredns:v1.8.0
  imagepullpolicy: IfNotPresent
...
```


<p align="left"> Author </p>
<p align="left"> <a href="https://github.com/XieWeiXie"> @XieWei </a></p>