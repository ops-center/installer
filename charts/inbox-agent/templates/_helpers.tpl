{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "inbox-agent.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "inbox-agent.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- if contains $name .Release.Name -}}
{{- .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "inbox-agent.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "inbox-agent.labels" -}}
helm.sh/chart: {{ include "inbox-agent.chart" . }}
{{ include "inbox-agent.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{/*
Selector labels
*/}}
{{- define "inbox-agent.selectorLabels" -}}
app.kubernetes.io/name: {{ include "inbox-agent.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- with .Values.podLabels }}
{{- toYaml . | nindent 0 }}
{{- end -}}
{{- end -}}

{{/*
Create the name of the service account to use
*/}}
{{- define "inbox-agent.serviceAccountName" -}}
{{- if .Values.serviceAccount.create -}}
    {{ default (include "inbox-agent.fullname" .) .Values.serviceAccount.name }}
{{- else -}}
    {{ default "default" .Values.serviceAccount.name }}
{{- end -}}
{{- end -}}

{{- define "inbox-agent.webhookServiceName" -}}
{{- printf "%s-webhook" (include "inbox-agent.fullname" . ) | trunc 63 | trimPrefix "-" -}}
{{- end -}}

{{/*
Returns the registry used for operator docker image
*/}}
{{- define "operator.registry" -}}
{{- list .Values.registryFQDN .Values.operator.registry | compact | join "/" }}
{{- end }}

{{/*
Returns the registry used for kube-rbac-proxy docker image
*/}}
{{- define "rbacproxy.registry" -}}
{{- list .Values.registryFQDN .Values.rbacproxy.registry | compact | join "/" }}
{{- end }}

{{/*
Returns the registry used for cleaner docker image
*/}}
{{- define "cleaner.registry" -}}
{{- list .Values.registryFQDN .Values.cleaner.registry | compact | join "/" }}
{{- end }}

{{/*
Returns whether the cleaner job YAML will be generated or not
*/}}
{{- define "cleaner.generate" -}}
{{- ternary "false" "true" .Values.cleaner.skip -}}
{{- end }}

{{- define "appscode.imagePullSecrets" -}}
{{- with .Values.imagePullSecrets -}}
imagePullSecrets:
{{- toYaml . | nindent 2 }}
{{- end }}
{{- end }}

{{- define "image-pull-secrets" -}}
{{- with .Values.imagePullSecrets -}}
imagePullSecrets:
{{- toYaml . | nindent 2 }}
{{- end -}}
{{- end -}}
