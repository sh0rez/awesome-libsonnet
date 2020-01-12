





# Awesome libsonnet
A curated, Kubernetes focused list of Jsonnet resources. For a more general list
of Jsonnet related things, make sure to check out
[metalmatze/awesome-jsonnet](https://github.com/metalmatze/awesome-jsonnet).

A core feature of Jsonnet, `import`, allows easy sharing and reusing of code.
Currently available libraries can be grouped by different types:

* [**Modules**](#application-modules) provide Kubernetes API resources to install entire Applications.
They can be compared to *Helm Charts*.
* [**Mixins**](#configuration-mixins) usually provide only configuration of some kind. This can be for example *Grafana dashboards*, *Prometheus alerts*, etc
<!-- * **Helpers** are abstract libraries that don't produce anything on their own, but can ease your life. The most prominent example is [`k.libsonnet`](https://github.com/ksonnet/ksonnet-lib). -->

| Maintained by: | :crown: Project maintainers | :busts_in_silhouette: Community | :no_entry: Unmaintained |
|----------------|-----------------------------|---------------------------------|-------------------------|

## Application modules


 
#### [Bitly oauth2_proxy](https://github.com/bitly/oauth2_proxy)



* :busts_in_silhouette:  by Grafana Labs: [`github.com/grafana/jsonnet-libs/oauth2-proxy`](https://github.com/grafana/jsonnet-libs/tree/master/oauth2-proxy)


 

 
#### [CoreOS etcd-operator](https://github.com/coreos/etcd-operator)



* :busts_in_silhouette:  by Grafana Labs: [`github.com/grafana/jsonnet-libs/etcd-operator`](https://github.com/grafana/jsonnet-libs/tree/master/etcd-operator)


 

 
#### [CoreOS prometheus-operator](https://github.com/coreos/prometheus-operator)



* :crown: [`github.com/coreos/prometheus-operator/jsonnet/prometheus-operator`](https://github.com/coreos/prometheus-operator/tree/master/jsonnet/prometheus-operator)


 

 
#### [Grafana Loki](https://grafana.com/loki)



* :crown:  Loki: [`github.com/grafana/loki/production/ksonnet/loki`](https://github.com/grafana/loki/tree/master/production/ksonnet/loki)

* :crown:  Promtail: [`github.com/grafana/loki/production/ksonnet/promtail`](https://github.com/grafana/loki/tree/master/production/ksonnet/promtail)

* :crown:  Canary: [`github.com/grafana/loki/production/ksonnet/loki-canary`](https://github.com/grafana/loki/tree/master/production/ksonnet/loki-canary)


 

 
#### [HashiCorp Consul](https://consul.io)



* :busts_in_silhouette:  by Grafana Labs: [`github.com/grafana/jsonnet-libs/consul`](https://github.com/grafana/jsonnet-libs/tree/master/consul)


 





 
#### [Prometheus](https://prometheus.io)



* :busts_in_silhouette:  kube-prometheus by CoreOS: [`github.com/coreos/kube-prometheus/jsonnet/kube-prometheus`](https://github.com/coreos/kube-prometheus/tree/master/jsonnet/kube-prometheus)

* :busts_in_silhouette:  prometheus-ksonnet by Grafana Labs: [`github.com/grafana/jsonnet-libs/prometheus-ksonnet`](https://github.com/grafana/jsonnet-libs/tree/master/prometheus-ksonnet)


 

 
#### [memcached](https://memcached.org)



* :busts_in_silhouette:  by Grafana Labs: [`github.com/grafana/jsonnet-libs/memcached`](https://github.com/grafana/jsonnet-libs/tree/master/memcached)


 


## Configuration mixins










 
#### [HashiCorp Consul](https://consul.io)



* :busts_in_silhouette:  by Grafana Labs: [`github.com/grafana/jsonnet-libs/consul-mixin`](https://github.com/grafana/jsonnet-libs/tree/master/consul-mixin)


 

 
#### [Jaeger](https://jaegertracing.com)



* :crown: [`github.com/jaegertracing/jaeger/monitoring/jaeger-mixin`](https://github.com/jaegertracing/jaeger/tree/master/monitoring/jaeger-mixin)


 

 
#### [Kubernetes](https://kubernetes.iou)



* :busts_in_silhouette: [`github.com/kubernetes-monitoring/kubernetes-mixin`](github.com/kubernetes-monitoring/kubernetes-mixin)


 

 
#### [Prometheus](https://prometheus.io)



* :crown: [`github.com/prometheus/prometheus/documentation/prometheus-mixin`](https://github.com/prometheus/prometheus/tree/master/documentation/prometheus-mixin)


 

 
#### [memcached](https://memcached.org)



* :busts_in_silhouette:  by Grafana Labs: [`github.com/grafana/jsonnet-libs/memcached-mixin`](https://github.com/grafana/jsonnet-libs/tree/master/memcached-mixin)


 

