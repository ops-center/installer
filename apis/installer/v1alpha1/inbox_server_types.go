/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindInboxServer = "InboxServer"
	ResourceInboxServer     = "inboxserver"
	ResourceInboxServers    = "inboxservers"
)

// InboxServer defines the schama for InboxServer operator installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
type InboxServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InboxServerSpec `json:"spec,omitempty"`
}

// InboxServerSpec is the schema for Inbox server values file
type InboxServerSpec struct {
	ReplicaCount       int                       `json:"replicaCount"`
	NameOverride       string                    `json:"nameOverride"`
	FullnameOverride   string                    `json:"fullnameOverride"`
	Dns                DnsSpec                   `json:"dns"`
	James              JamesSpec                 `json:"james"`
	ServiceMonitor     ServiceMonitorSpec        `json:"serviceMonitor"`
	Ingress            IngressSpec               `json:"ingress"`
	AdminIngress       IngressSpec               `json:"adminIngress"`
	Cassandra          CassandraSpec             `json:"cassandra"`
	Opensearch         OpensearchSpec            `json:"opensearch"`
	Rabbitmq           RabbitmqSpec              `json:"rabbitmq"`
	S3                 S3Spec                    `json:"s3"`
	Tika               TikaSpec                  `json:"tika"`
	PodAnnotations     map[string]string         `json:"podAnnotations"`
	PodLabels          map[string]string         `json:"podLabels"`
	PodSecurityContext *core.PodSecurityContext  `json:"podSecurityContext"`
	SecurityContext    *core.SecurityContext     `json:"securityContext"`
	Service            ServiceSpec               `json:"service"`
	Resources          core.ResourceRequirements `json:"resources"`
	LivenessProbe      *core.Probe               `json:"livenessProbe"`
	ReadinessProbe     *core.Probe               `json:"readinessProbe"`
	Volumes            []core.VolumeSource       `json:"volumes"`
	VolumeMounts       []core.VolumeMount        `json:"volumeMounts"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity"`
}

type DnsSpec struct {
	Domain      string   `json:"domain"`
	EmailDomain string   `json:"emailDomain"`
	TargetIPs   []string `json:"targetIPs"`
}

type JamesSpec struct {
	ReplicaJmapInstanceCount     int          `json:"replicaJmapInstanceCount"`
	ReplicaImapSmtpInstanceCount int          `json:"replicaImapSmtpInstanceCount"`
	Image                        string       `json:"image"`
	Env                          JamesEnv     `json:"env"`
	TLS                          TLSSpec      `json:"tls"`
	Secret                       *JamesSecret `json:"secret,omitempty"`
}

type JamesEnv struct {
	JamesCassandraKeyspace       string                    `json:"jamesCassandraKeyspace"`
	JamesCassandraCacheKeyspace  string                    `json:"jamesCassandraCacheKeyspace"`
	JamesEsMailboxIndex          string                    `json:"jamesEsMailboxIndex"`
	JamesEsClusterName           string                    `json:"jamesEsClusterName"`
	JamesEsHostScheme            string                    `json:"jamesEsHostScheme"`
	JamesEsSslValidationStrategy string                    `json:"jamesEsSslValidationStrategy"`
	JamesEsHostNameVerifier      string                    `json:"jamesEsHostNameVerifier"`
	JamesEsNbShards              int                       `json:"jamesEsNbShards"`
	JamesEsNbReplica             int                       `json:"jamesEsNbReplica"`
	JamesEsMailboxReadAlias      string                    `json:"jamesEsMailboxReadAlias"`
	JamesEsMailboxWriteAlias     string                    `json:"jamesEsMailboxWriteAlias"`
	JamesRabbitHost              string                    `json:"jamesRabbitHost"`
	JamesMessageSize             string                    `json:"jamesMessageSize"`
	JamesDkimSignSmtp            string                    `json:"jamesDkimSignSmtp"`
	JamesDkimSignDomain          string                    `json:"jamesDkimSignDomain"`
	CassandraReplicationFactor   int                       `json:"cassandraReplicationFactor"`
	JamesHELOMessage             string                    `json:"jamesHELOMessage"`
	JvmOpts                      string                    `json:"jvmOpts"`
	Glowroot                     GlowrootSpec              `json:"glowroot"`
	JamesResources               core.ResourceRequirements `json:"jamesResources"`
}

type JamesSecret struct {
	AdminJWTPublicKey string `json:"adminJWTPublicKey"`
	JwtPublicKey      string `json:"jwtPublicKey"`
	JwtPrivateKey     string `json:"jwtPrivateKey"`
	DkimPrivateKey    string `json:"dkimPrivateKey"`
}

type GlowrootSpec struct {
	Enabled bool `json:"enabled"`
}

type ServiceMonitorSpec struct {
	Enabled          bool              `json:"enabled"`
	AdditionalLabels map[string]string `json:"additionalLabels"`
	Interval         string            `json:"interval"`
}

type IngressSpec struct {
	Enabled     bool              `json:"enabled"`
	ClassName   string            `json:"className"`
	Annotations map[string]string `json:"annotations"`
	TLS         TLSSpec           `json:"tls"`
}

type TLSSpec struct {
	SecretName string `json:"secretName"`
}

type CassandraSpec struct {
	Image         ImageReference `json:"image"`
	LivenessProbe core.Probe     `json:"livenessProbe"`
}

type OpensearchSpec struct {
	Enabled bool           `json:"enabled"`
	Image   ImageReference `json:"image"`
}

type RabbitmqSpec struct {
	Image ImageReference `json:"image"`
}

type S3Spec struct {
	Enabled bool           `json:"enabled"`
	Image   ImageReference `json:"image"`
}

type TikaSpec struct {
	Enabled bool           `json:"enabled"`
	Image   ImageReference `json:"image"`
}

type ServiceSpec struct {
	Type string `json:"type"`
}

type ImageReference struct {
	Repository string `json:"repository"`
	PullPolicy string `json:"pullPolicy"`
	Tag        string `json:"tag"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InboxServerList is a list of InboxServers
type InboxServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of InboxServer CRD objects
	Items []InboxServer `json:"items,omitempty"`
}
