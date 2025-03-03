# Inbox Server Lite

[Inbox Server Lite by AppsCode](https://github.com/ops-center/james-project) - Inbox Server Lite by AppsCode

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/inbox-server --version=v2025.2.28
$ helm upgrade -i inbox-server appscode/inbox-server -n monitoring --create-namespace --version=v2025.2.28
```

## Introduction

This chart deploys Inbox Server Lite on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+

## Installing the Chart

To install/upgrade the chart with the release name `inbox-server`:

```bash
$ helm upgrade -i inbox-server appscode/inbox-server -n monitoring --create-namespace --version=v2025.2.28
```

The command deploys Inbox Server Lite on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `inbox-server`:

```bash
$ helm uninstall inbox-server -n monitoring
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `inbox-server` chart and their default values.

|                 Parameter                 |                                                               Description                                                               |                                                                                            Default                                                                                             |
|-------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| nameOverride                              | Overrides name template                                                                                                                 | <code>""</code>                                                                                                                                                                                |
| fullnameOverride                          | Overrides fullname template                                                                                                             | <code>""</code>                                                                                                                                                                                |
| dns.domain                                | # Default domain for all emailing components jmapUrl: "jmap.example.com" adminUrl: "admin.example.com" smtpHostname: "smtp.example.com" | <code>""</code>                                                                                                                                                                                |
| dns.emailDomain                           | in ip mode, the emailDomain is ace.internal, otherwise same as domain                                                                   | <code>""</code>                                                                                                                                                                                |
| dns.targetIPs                             |                                                                                                                                         | <code>[]</code>                                                                                                                                                                                |
| james.replicaJmapInstanceCount            |                                                                                                                                         | <code>1</code>                                                                                                                                                                                 |
| james.replicaImapSmtpInstanceCount        |                                                                                                                                         | <code>0</code>                                                                                                                                                                                 |
| james.image                               |                                                                                                                                         | <code>ghcr.io/appscode/inbox-server:postgres-latest</code>                                                                                                                                     |
| james.env.jamesMessageSize                |                                                                                                                                         | <code>25M</code>                                                                                                                                                                               |
| james.env.jamesDkimSignSmtp               |                                                                                                                                         | <code>"dkimselector"</code>                                                                                                                                                                    |
| james.env.jamesDkimSignDomain             |                                                                                                                                         | <code>"mail.example.com"</code>                                                                                                                                                                |
| james.env.jamesKeyStorePassword           |                                                                                                                                         | <code>james72laBalle</code>                                                                                                                                                                    |
| james.env.jamesHELOMessage                |                                                                                                                                         | <code>"change-me"</code>                                                                                                                                                                       |
| james.env.jvmOpts                         |                                                                                                                                         | <code>"-Xms3g -Xmx3g"</code>                                                                                                                                                                   |
| james.env.glowroot.enabled                |                                                                                                                                         | <code>false</code>                                                                                                                                                                             |
| james.env.jamesResources.limits.cpu       |                                                                                                                                         | <code>"2"</code>                                                                                                                                                                               |
| james.env.jamesResources.limits.memory    |                                                                                                                                         | <code>"4Gi"</code>                                                                                                                                                                             |
| james.env.jamesResources.requests.cpu     |                                                                                                                                         | <code>"150m"</code>                                                                                                                                                                            |
| james.env.jamesResources.requests.memory  |                                                                                                                                         | <code>"512Mi"</code>                                                                                                                                                                           |
| james.tls.secretName                      | secretName: the-name-of-a-secret                                                                                                        | <code>""</code>                                                                                                                                                                                |
| serviceMonitor.enabled                    |                                                                                                                                         | <code>false</code>                                                                                                                                                                             |
| serviceMonitor.additionalLabels.change-me |                                                                                                                                         | <code>change-me</code>                                                                                                                                                                         |
| serviceMonitor.interval                   | Scrape interval. Use Prometheus default value if not specified                                                                          | <code>30s</code>                                                                                                                                                                               |
| ingress.enabled                           |                                                                                                                                         | <code>false</code>                                                                                                                                                                             |
| ingress.className                         |                                                                                                                                         | <code>""</code>                                                                                                                                                                                |
| ingress.annotations                       |                                                                                                                                         | <code>{}</code>                                                                                                                                                                                |
| ingress.tls.secretName                    |                                                                                                                                         | <code>the-name-of-a-secret</code>                                                                                                                                                              |
| adminIngress.enabled                      |                                                                                                                                         | <code>false</code>                                                                                                                                                                             |
| adminIngress.className                    |                                                                                                                                         | <code>""</code>                                                                                                                                                                                |
| adminIngress.annotations                  |                                                                                                                                         | <code>{}</code>                                                                                                                                                                                |
| adminIngress.tls.secretName               |                                                                                                                                         | <code>the-name-of-a-secret</code>                                                                                                                                                              |
| postgresql.enabled                        | # @param postgresql.enabled enable the bitnami/postgresql subchart and deploy Postgres                                                  | <code>false</code>                                                                                                                                                                             |
| migrate.image.repository                  |                                                                                                                                         | <code>ghcr.io/appscode-images/postgres</code>                                                                                                                                                  |
| migrate.image.tag                         | Overrides the image tag whose default is the chart appVersion.                                                                          | <code>"16.4-alpine"</code>                                                                                                                                                                     |
| migrate.image.pullPolicy                  |                                                                                                                                         | <code>Always</code>                                                                                                                                                                            |
| migrate.podAnnotations                    |                                                                                                                                         | <code>{}</code>                                                                                                                                                                                |
| migrate.podSecurityContext                |                                                                                                                                         | <code>{}</code>                                                                                                                                                                                |
| migrate.securityContext                   |                                                                                                                                         | <code>{"allowPrivilegeEscalation":false,"capabilities":{"drop":["ALL"]},"readOnlyRootFilesystem":true,"runAsNonRoot":true,"runAsUser":65534,"seccompProfile":{"type":"RuntimeDefault"}}</code> |
| migrate.resources                         |                                                                                                                                         | <code>{}</code>                                                                                                                                                                                |
| migrate.nodeSelector                      |                                                                                                                                         | <code>{}</code>                                                                                                                                                                                |
| migrate.tolerations                       |                                                                                                                                         | <code>[]</code>                                                                                                                                                                                |
| migrate.affinity                          |                                                                                                                                         | <code>{}</code>                                                                                                                                                                                |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i inbox-server appscode/inbox-server -n monitoring --create-namespace --version=v2025.2.28 --set james.replicaJmapInstanceCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i inbox-server appscode/inbox-server -n monitoring --create-namespace --version=v2025.2.28 --values values.yaml
```
