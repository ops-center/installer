# AppsCode OpenTelemetry Stack

[AppsCode OpenTelemetry Stack](https://github.com/ops-center) - AppsCode OpenTelemetry Stack

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/appscode-otel-stack --version=v2025.2.28
$ helm upgrade -i appscode-otel-stack appscode/appscode-otel-stack -n monitoring --create-namespace --version=v2025.2.28
```

## Introduction

This chart deploys AppsCode OpenTelemetry Stack on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.25+

## Installing the Chart

To install/upgrade the chart with the release name `appscode-otel-stack`:

```bash
$ helm upgrade -i appscode-otel-stack appscode/appscode-otel-stack -n monitoring --create-namespace --version=v2025.2.28
```

The command deploys AppsCode OpenTelemetry Stack on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `appscode-otel-stack`:

```bash
$ helm uninstall appscode-otel-stack -n monitoring
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `appscode-otel-stack` chart and their default values.

|                                                    Parameter                                                    |                  Description                  |                                     Default                                      |
|-----------------------------------------------------------------------------------------------------------------|-----------------------------------------------|----------------------------------------------------------------------------------|
| nameOverride                                                                                                    | This is to override the chart name.           | <code>""</code>                                                                  |
| fullnameOverride                                                                                                |                                               | <code>""</code>                                                                  |
| clickhouse.enabled                                                                                              |                                               | <code>true</code>                                                                |
| clickhouse.version                                                                                              | Global ClickHouse version and deletion policy | <code>24.4.1</code>                                                              |
| clickhouse.deletionPolicy                                                                                       |                                               | <code>WipeOut</code>                                                             |
| clickhouse.deploymentMode                                                                                       |                                               | <code>standalone</code>                                                          |
| clickhouse.storage.storageClassName                                                                             |                                               | <code>""</code>                                                                  |
| clickhouse.storage.resources.requests.storage                                                                   |                                               | <code>1Gi</code>                                                                 |
| clickhouse.clusterTopology.clickHouseKeeper.externallyManaged                                                   |                                               | <code>false</code>                                                               |
| clickhouse.clusterTopology.clickHouseKeeper.replicas                                                            |                                               | <code>1</code>                                                                   |
| clickhouse.clusterTopology.clickHouseKeeper.storage.storageClassName                                            |                                               | <code>""</code>                                                                  |
| clickhouse.clusterTopology.clickHouseKeeper.storage.resources.requests.storage                                  |                                               | <code>1Gi</code>                                                                 |
| clickhouse.clusterTopology.cluster.name                                                                         |                                               | <code>cluster</code>                                                             |
| clickhouse.clusterTopology.cluster.shards                                                                       |                                               | <code>2</code>                                                                   |
| clickhouse.clusterTopology.cluster.replicas                                                                     |                                               | <code>2</code>                                                                   |
| clickhouse.clusterTopology.cluster.storage.storageClassName                                                     |                                               | <code>""</code>                                                                  |
| clickhouse.clusterTopology.cluster.storage.resources.requests.storage                                           |                                               | <code>1Gi</code>                                                                 |
| opentelemetry-kube-stack.nameOverride                                                                           |                                               | <code>"appscode-otel-stack"</code>                                               |
| opentelemetry-kube-stack.enabled                                                                                |                                               | <code>true</code>                                                                |
| opentelemetry-kube-stack.clusterName                                                                            |                                               | <code>""</code>                                                                  |
| opentelemetry-kube-stack.collectors.cluster.config.exporters.clickhouse.database                                |                                               | <code>otel</code>                                                                |
| opentelemetry-kube-stack.collectors.cluster.config.exporters.clickhouse.endpoint                                |                                               | <code>tcp://appscode-otel-stack-clickhouse:9000</code>                           |
| opentelemetry-kube-stack.collectors.cluster.config.exporters.clickhouse.logs_table_name                         |                                               | <code>otel_logs</code>                                                           |
| opentelemetry-kube-stack.collectors.cluster.config.exporters.clickhouse.password                                |                                               | <code>${CLICKHOUSE_PASSWORD}</code>                                              |
| opentelemetry-kube-stack.collectors.cluster.config.exporters.clickhouse.retry_on_failure.enabled                |                                               | <code>true</code>                                                                |
| opentelemetry-kube-stack.collectors.cluster.config.exporters.clickhouse.retry_on_failure.initial_interval       |                                               | <code>5s</code>                                                                  |
| opentelemetry-kube-stack.collectors.cluster.config.exporters.clickhouse.retry_on_failure.max_elapsed_time       |                                               | <code>300s</code>                                                                |
| opentelemetry-kube-stack.collectors.cluster.config.exporters.clickhouse.retry_on_failure.max_interval           |                                               | <code>30s</code>                                                                 |
| opentelemetry-kube-stack.collectors.cluster.config.exporters.clickhouse.sending_queue.queue_size                |                                               | <code>100</code>                                                                 |
| opentelemetry-kube-stack.collectors.cluster.config.exporters.clickhouse.timeout                                 |                                               | <code>10s</code>                                                                 |
| opentelemetry-kube-stack.collectors.cluster.config.exporters.clickhouse.ttl                                     |                                               | <code>0</code>                                                                   |
| opentelemetry-kube-stack.collectors.cluster.config.exporters.clickhouse.username                                |                                               | <code>${CLICKHOUSE_USERNAME}</code>                                              |
| opentelemetry-kube-stack.collectors.cluster.config.exporters.otlphttp/prometheus.endpoint                       |                                               | <code>http://appscode-otel-stack-prometheus:9090/api/v1/otlp</code>              |
| opentelemetry-kube-stack.collectors.cluster.config.exporters.otlphttp/prometheus.tls.insecure                   |                                               | <code>true</code>                                                                |
| opentelemetry-kube-stack.collectors.cluster.config.receivers.otlp.protocols.grpc.endpoint                       |                                               | <code>0.0.0.0:4317</code>                                                        |
| opentelemetry-kube-stack.collectors.cluster.config.receivers.otlp.protocols.http.endpoint                       |                                               | <code>0.0.0.0:4318</code>                                                        |
| opentelemetry-kube-stack.collectors.cluster.enabled                                                             |                                               | <code>true</code>                                                                |
| opentelemetry-kube-stack.collectors.cluster.image.repository                                                    |                                               | <code>otel/opentelemetry-collector-contrib</code>                                |
| opentelemetry-kube-stack.collectors.cluster.image.tag                                                           |                                               | <code>0.95.0</code>                                                              |
| opentelemetry-kube-stack.collectors.cluster.labels.otel-collector-type                                          |                                               | <code>deployment</code>                                                          |
| opentelemetry-kube-stack.collectors.cluster.presets.clusterMetrics.enabled                                      |                                               | <code>false</code>                                                               |
| opentelemetry-kube-stack.collectors.cluster.suffix                                                              |                                               | <code>gateway</code>                                                             |
| opentelemetry-kube-stack.collectors.daemon.config.exporters.otlp.endpoint                                       |                                               | <code>appscode-otel-stack-gateway-collector:4317</code>                          |
| opentelemetry-kube-stack.collectors.daemon.config.exporters.otlp.tls.insecure                                   |                                               | <code>true</code>                                                                |
| opentelemetry-kube-stack.collectors.daemon.enabled                                                              |                                               | <code>true</code>                                                                |
| opentelemetry-kube-stack.collectors.daemon.labels.otel-collector-type                                           |                                               | <code>daemonset</code>                                                           |
| opentelemetry-kube-stack.collectors.daemon.presets.hostMetrics.enabled                                          |                                               | <code>false</code>                                                               |
| opentelemetry-kube-stack.collectors.daemon.presets.kubeletMetrics.enabled                                       |                                               | <code>false</code>                                                               |
| opentelemetry-kube-stack.collectors.daemon.presets.kubernetesAttributes.enabled                                 |                                               | <code>false</code>                                                               |
| opentelemetry-kube-stack.collectors.daemon.presets.logsCollection.enabled                                       |                                               | <code>true</code>                                                                |
| opentelemetry-kube-stack.collectors.daemon.scrape_configs_file                                                  |                                               | <code>config/kubelet_scrape_configs.yaml</code>                                  |
| opentelemetry-kube-stack.collectors.daemon.suffix                                                               |                                               | <code>agent</code>                                                               |
| opentelemetry-kube-stack.collectors.daemon.targetAllocator.allocationStrategy                                   |                                               | <code>per-node</code>                                                            |
| opentelemetry-kube-stack.collectors.daemon.targetAllocator.enabled                                              |                                               | <code>true</code>                                                                |
| opentelemetry-kube-stack.collectors.daemon.targetAllocator.image                                                |                                               | <code>ghcr.io/open-telemetry/opentelemetry-operator/target-allocator:main</code> |
| opentelemetry-kube-stack.collectors.daemon.targetAllocator.prometheusCR.enabled                                 |                                               | <code>true</code>                                                                |
| opentelemetry-kube-stack.collectors.daemon.targetAllocator.prometheusCR.podMonitorSelector                      |                                               | <code>{}</code>                                                                  |
| opentelemetry-kube-stack.collectors.daemon.targetAllocator.prometheusCR.scrapeInterval                          |                                               | <code>30s</code>                                                                 |
| opentelemetry-kube-stack.collectors.daemon.targetAllocator.prometheusCR.serviceMonitorSelector                  |                                               | <code>{}</code>                                                                  |
| opentelemetry-kube-stack.instrumentation.enabled                                                                |                                               | <code>false</code>                                                               |
| opentelemetry-kube-stack.kubeApiServer.enabled                                                                  |                                               | <code>false</code>                                                               |
| opentelemetry-kube-stack.kubeControllerManager.enabled                                                          |                                               | <code>false</code>                                                               |
| opentelemetry-kube-stack.kubeDns.enabled                                                                        |                                               | <code>false</code>                                                               |
| opentelemetry-kube-stack.kubeEtcd.enabled                                                                       |                                               | <code>false</code>                                                               |
| opentelemetry-kube-stack.kubeProxy.enabled                                                                      |                                               | <code>false</code>                                                               |
| opentelemetry-kube-stack.kubeScheduler.enabled                                                                  |                                               | <code>false</code>                                                               |
| opentelemetry-kube-stack.kubeStateMetrics.enabled                                                               |                                               | <code>true</code>                                                                |
| opentelemetry-kube-stack.kubelet.enabled                                                                        |                                               | <code>false</code>                                                               |
| opentelemetry-kube-stack.kubernetesServiceMonitors.enabled                                                      |                                               | <code>true</code>                                                                |
| opentelemetry-kube-stack.nodeExporter.enabled                                                                   |                                               | <code>true</code>                                                                |
| opentelemetry-kube-stack.opAMPBridge.enabled                                                                    |                                               | <code>false</code>                                                               |
| kube-prometheus-stack.nameOverride                                                                              |                                               | <code>"appscode-otel-stack"</code>                                               |
| kube-prometheus-stack.enabled                                                                                   |                                               | <code>true</code>                                                                |
| kube-prometheus-stack.grafana.enabled                                                                           |                                               | <code>false</code>                                                               |
| kube-prometheus-stack.prometheus.prometheusSpec.serviceMonitorSelector.matchLabels.scrape                       |                                               | <code>disabled</code>                                                            |
| kube-prometheus-stack.prometheus.prometheusSpec.serviceMonitorSelectorNilUsesHelmValues                         |                                               | <code>true</code>                                                                |
| kube-prometheus-stack.prometheus.prometheusSpec.storageSpec.volumeClaimTemplate.spec.resources.requests.storage |                                               | <code>50Gi</code>                                                                |
| kube-prometheus-stack.kubernetesServiceMonitors.enabled                                                         |                                               | <code>false</code>                                                               |
| kube-prometheus-stack.kubeApiServer.enabled                                                                     |                                               | <code>false</code>                                                               |
| kube-prometheus-stack.kubelet.enabled                                                                           |                                               | <code>false</code>                                                               |
| kube-prometheus-stack.kubeControllerManager.enabled                                                             |                                               | <code>false</code>                                                               |
| kube-prometheus-stack.kubeDns.enabled                                                                           |                                               | <code>false</code>                                                               |
| kube-prometheus-stack.kubeEtcd.enabled                                                                          |                                               | <code>false</code>                                                               |
| kube-prometheus-stack.kubeScheduler.enabled                                                                     |                                               | <code>false</code>                                                               |
| kube-prometheus-stack.kubeProxy.enabled                                                                         |                                               | <code>false</code>                                                               |
| kube-prometheus-stack.kubeStateMetrics.enabled                                                                  |                                               | <code>false</code>                                                               |
| kube-prometheus-stack.nodeExporter.enabled                                                                      |                                               | <code>false</code>                                                               |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i appscode-otel-stack appscode/appscode-otel-stack -n monitoring --create-namespace --version=v2025.2.28 --set clickhouse.version=24.4.1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i appscode-otel-stack appscode/appscode-otel-stack -n monitoring --create-namespace --version=v2025.2.28 --values values.yaml
```
