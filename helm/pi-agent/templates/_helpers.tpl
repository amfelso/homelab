# templates/_helpers.tpl

# Standard labels applied to all resources
{{- define "pi-agent.labels" -}}
helm.sh/chart: {{ .Chart.Name }}-{{ .Chart.Version }}
app.kubernetes.io/name: {{ .Chart.Name }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

# Metric collector selector labels
{{- define "pi-agent.collector.selectorLabels" -}}
app.kubernetes.io/component: {{ .Values.collector.name }}
{{- end }}