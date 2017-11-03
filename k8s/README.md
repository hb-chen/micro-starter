### Kubernetes
[micro/kubernetes](http://github.com/micro/kubernetes)

### minikube安装gcr.io被墙镜像
```bash
# 查看哪些pod失败
$ kubectl get pods --all-namespaces

# 查看相应pod的image
$ kubectl describe pods --all-namespaces

# 进入虚拟机，拉取镜像并重新打回gcr.io镜像
$ minikube ssh
$ docker pull hbchen/kubernetes-dashboard-amd64:v1.7.0
$ docler tag {IMAGE_ID} gcr.io/google_containers/kubernetes-dashboard-amd64:v1.7.0

# 镜像列表
gcr.io/google_containers/kubernetes-dashboard-init-amd64:v1.0.1
  └ hbchen/kubernetes-dashboard-init-amd64:v1.0.1
gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.5
  └ hbchen/k8s-dns-sidecar-amd64:1.14.5
gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.5
  └ hbchen/k8s-dns-kube-dns-amd64:1.14.5
gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.5
  └ hbchen/k8s-dns-dnsmasq-nanny-amd64:1.14.5
gcr.io/google_containers/kubernetes-dashboard-amd64:v1.7.1
  └ hbchen/kubernetes-dashboard-amd64:v1.7.1
gcr.io/google_containers/kubernetes-dashboard-amd64:v1.7.0
  └ hbchen/kubernetes-dashboard-amd64:v1.7.0
gcr.io/google_containers/kubernetes-dashboard-amd64:v1.6.3
  └ hbchen/kubernetes-dashboard-amd64:v1.6.3
gcr.io/google-containers/kube-addon-manager:v6.4-beta.2
  └ hbchen/kube-addon-manager:v6.4-beta.2
gcr.io/google_containers/pause-amd64:3.0
  └ hbchen/pause-amd64:3.0
```

```bash
# 运行
$ ./run.sh start

$ minikube ip
192.168.64.4
```

### API
```bash
http://192.168.64.4:30001/account/account/login?nickname=hobo&pwd=123456
```

### Web
```bash
http://192.168.64.4:30003/
```