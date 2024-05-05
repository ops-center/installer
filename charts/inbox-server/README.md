# Inbox Server

[Inbox Server by AppsCode](https://github.com/ops-center/james-project) - Inbox Server by AppsCode

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/inbox-server --version=v2024.5.3
$ helm upgrade -i inbox-server appscode/inbox-server -n monitoring --create-namespace --version=v2024.5.3
```

## Introduction

This chart deploys Inbox Server on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `inbox-server`:

```bash
$ helm upgrade -i inbox-server appscode/inbox-server -n monitoring --create-namespace --version=v2024.5.3
```

The command deploys Inbox Server on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `inbox-server`:

```bash
$ helm uninstall inbox-server -n monitoring
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `inbox-server` chart and their default values.

|                 Parameter                 |                                                               Description                                                               |                   Default                   |
|-------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------|
| dns.emailDomain                           | # Default domain for all emailing components jmapUrl: "jmap.example.com" adminUrl: "admin.example.com" smtpHostname: "smtp.example.com" | <code>"example.com"</code>                  |
| dns.targetIPs                             |                                                                                                                                         | <code>[]</code>                             |
| james.replicaJmapInstanceCount            |                                                                                                                                         | <code>1</code>                              |
| james.replicaImapSmtpInstanceCount        |                                                                                                                                         | <code>1</code>                              |
| james.image                               |                                                                                                                                         | <code>apache/james:distributed-3.8.0</code> |
| james.env.jamesCassandraKeyspace          |                                                                                                                                         | <code>sandbox_james</code>                  |
| james.env.jamesCassandraCacheKeyspace     |                                                                                                                                         | <code>sandbox_james_cache</code>            |
| james.env.jamesEsMailboxIndex             |                                                                                                                                         | <code>"mailbox_v1"</code>                   |
| james.env.jamesEsClusterName              |                                                                                                                                         | <code>"change-me"</code>                    |
| james.env.jamesEsHostScheme               |                                                                                                                                         | <code>"http"</code>                         |
| james.env.jamesEsSslValidationStrategy    |                                                                                                                                         | <code>"default"</code>                      |
| james.env.jamesEsHostNameVerifier         |                                                                                                                                         | <code>"default"</code>                      |
| james.env.jamesEsNbShards                 |                                                                                                                                         | <code>5</code>                              |
| james.env.jamesEsNbReplica                |                                                                                                                                         | <code>1</code>                              |
| james.env.jamesEsMailboxReadAlias         |                                                                                                                                         | <code>"read-mailbox"</code>                 |
| james.env.jamesEsMailboxWriteAlias        |                                                                                                                                         | <code>"write-mailbox"</code>                |
| james.env.jamesRabbitHost                 |                                                                                                                                         | <code>"change-me"</code>                    |
| james.env.jamesMessageSize                |                                                                                                                                         | <code>25M</code>                            |
| james.env.jamesDkimSignSmtp               |                                                                                                                                         | <code>"dkimselector"</code>                 |
| james.env.jamesDkimSignDomain             |                                                                                                                                         | <code>"mail.example.com"</code>             |
| james.env.cassandraReplicationFactor      |                                                                                                                                         | <code>3</code>                              |
| james.env.jamesHELOMessage                |                                                                                                                                         | <code>"change-me"</code>                    |
| james.env.jvmOpts                         |                                                                                                                                         | <code>"-Xms3g -Xmx3g"</code>                |
| james.env.glowroot.enabled                |                                                                                                                                         | <code>false</code>                          |
| james.env.jamesResources.limits.cpu       |                                                                                                                                         | <code>"2000m"</code>                        |
| james.env.jamesResources.limits.memory    |                                                                                                                                         | <code>"4Gi"</code>                          |
| james.env.jamesResources.requests.cpu     |                                                                                                                                         | <code>"1000m"</code>                        |
| james.env.jamesResources.requests.memory  |                                                                                                                                         | <code>"4Gi"</code>                          |
| james.tls.secretName                      | secretName: the-name-of-a-secret                                                                                                        | <code>""</code>                             |
| serviceMonitor.enabled                    |                                                                                                                                         | <code>false</code>                          |
| serviceMonitor.additionalLabels.change-me |                                                                                                                                         | <code>change-me</code>                      |
| serviceMonitor.interval                   | Scrape interval. Use Prometheus default value if not specified                                                                          | <code>30s</code>                            |
| ingress.enabled                           |                                                                                                                                         | <code>false</code>                          |
| ingress.className                         |                                                                                                                                         | <code>""</code>                             |
| ingress.annotations                       |                                                                                                                                         | <code>{}</code>                             |
| ingress.tls.secretName                    |                                                                                                                                         | <code>the-name-of-a-secret</code>           |
| adminIngress.enabled                      |                                                                                                                                         | <code>false</code>                          |
| adminIngress.className                    |                                                                                                                                         | <code>""</code>                             |
| adminIngress.annotations                  |                                                                                                                                         | <code>{}</code>                             |
| adminIngress.tls.secretName               |                                                                                                                                         | <code>the-name-of-a-secret</code>           |
| cassandra.image.repository                |                                                                                                                                         | <code>cassandra</code>                      |
| cassandra.image.pullPolicy                |                                                                                                                                         | <code>IfNotPresent</code>                   |
| cassandra.image.tag                       | Overrides the image tag whose default is the chart appVersion.                                                                          | <code>4.1.3</code>                          |
| cassandra.livenessProbe.failureThreshold  |                                                                                                                                         | <code>5</code>                              |
| cassandra.livenessProbe.periodSeconds     |                                                                                                                                         | <code>3</code>                              |
| cassandra.livenessProbe.timeoutSeconds    |                                                                                                                                         | <code>20</code>                             |
| opensearch.enabled                        |                                                                                                                                         | <code>true</code>                           |
| opensearch.image.repository               |                                                                                                                                         | <code>opensearchproject/opensearch</code>   |
| opensearch.image.pullPolicy               |                                                                                                                                         | <code>IfNotPresent</code>                   |
| opensearch.image.tag                      | Overrides the image tag whose default is the chart appVersion.                                                                          | <code>2.1.0</code>                          |
| rabbitmq.image.repository                 |                                                                                                                                         | <code>rabbitmq</code>                       |
| rabbitmq.image.pullPolicy                 |                                                                                                                                         | <code>IfNotPresent</code>                   |
| rabbitmq.image.tag                        | Overrides the image tag whose default is the chart appVersion.                                                                          | <code>3.12.1-management</code>              |
| s3.enabled                                |                                                                                                                                         | <code>true</code>                           |
| s3.image.repository                       |                                                                                                                                         | <code>ghcr.io/appscode/cloudserver</code>   |
| s3.image.pullPolicy                       |                                                                                                                                         | <code>IfNotPresent</code>                   |
| s3.image.tag                              | Overrides the image tag whose default is the chart appVersion.                                                                          | <code>8.8.20</code>                         |
| tika.enabled                              |                                                                                                                                         | <code>true</code>                           |
| tika.image.repository                     |                                                                                                                                         | <code>apache/tika</code>                    |
| tika.image.pullPolicy                     |                                                                                                                                         | <code>IfNotPresent</code>                   |
| tika.image.tag                            | Overrides the image tag whose default is the chart appVersion.                                                                          | <code>2.8.0.0</code>                        |
| replicaCount                              | -------------------                                                                                                                     | <code>1</code>                              |
| nameOverride                              | imagePullSecrets: []                                                                                                                    | <code>""</code>                             |
| fullnameOverride                          |                                                                                                                                         | <code>""</code>                             |
| podAnnotations                            |                                                                                                                                         | <code>{}</code>                             |
| podLabels                                 |                                                                                                                                         | <code>{}</code>                             |
| podSecurityContext                        |                                                                                                                                         | <code>{}</code>                             |
| securityContext                           |                                                                                                                                         | <code>{}</code>                             |
| service.type                              |                                                                                                                                         | <code>ClusterIP</code>                      |
| resources                                 |                                                                                                                                         | <code>{}</code>                             |
| livenessProbe.httpGet.path                |                                                                                                                                         | <code>/</code>                              |
| livenessProbe.httpGet.port                |                                                                                                                                         | <code>http</code>                           |
| readinessProbe.httpGet.path               |                                                                                                                                         | <code>/</code>                              |
| readinessProbe.httpGet.port               |                                                                                                                                         | <code>http</code>                           |
| volumes                                   | Additional volumes on the output Deployment definition.                                                                                 | <code>[]</code>                             |
| volumeMounts                              | Additional volumeMounts on the output Deployment definition.                                                                            | <code>[]</code>                             |
| nodeSelector                              |                                                                                                                                         | <code>{}</code>                             |
| tolerations                               |                                                                                                                                         | <code>[]</code>                             |
| affinity                                  |                                                                                                                                         | <code>{}</code>                             |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i inbox-server appscode/inbox-server -n monitoring --create-namespace --version=v2024.5.3 --set dns.emailDomain="example.com"
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i inbox-server appscode/inbox-server -n monitoring --create-namespace --version=v2024.5.3 --values values.yaml
```
