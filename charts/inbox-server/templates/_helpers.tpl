{{/*
Expand the name of the chart.
*/}}
{{- define "inbox-server.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "inbox-server.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "inbox-server.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "inbox-server.labels" -}}
helm.sh/chart: {{ include "inbox-server.chart" . }}
{{ include "inbox-server.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "inbox-server.selectorLabels" -}}
app.kubernetes.io/name: {{ include "inbox-server.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "inbox-server.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "inbox-server.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Builds up the JAVA_TOOL_OPTIONS env variable for java tunning James
*/}}
{{- define "inbox-server.jvmOpts" -}}
{{- if .Values.james.env.glowroot.enabled }}
{{- printf "%s -javaagent:/root/glowroot/glowroot.jar" .Values.james.env.jvmOpts }}
{{- else }}
{{- .Values.james.env.jvmOpts }}
{{- end }}
{{- end }}