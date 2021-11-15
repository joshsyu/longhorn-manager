package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// BackingImageDownloadState is replaced by BackingImageState.
type BackingImageDownloadState string

type BackingImageState string

const (
	BackingImageStatePending          = BackingImageState("pending")
	BackingImageStateStarting         = BackingImageState("starting")
	BackingImageStateReadyForTransfer = BackingImageState("ready-for-transfer")
	BackingImageStateReady            = BackingImageState("ready")
	BackingImageStateInProgress       = BackingImageState("in-progress")
	BackingImageStateFailed           = BackingImageState("failed")
	BackingImageStateUnknown          = BackingImageState("unknown")
)

type BackingImageDiskFileStatus struct {
	// +optional
	State BackingImageState `json:"state"`
	// +optional
	Progress int `json:"progress"`
	// +optional
	Message string `json:"message"`
	// +optional
	LastStateTransitionTime string `json:"lastStateTransitionTime"`
}

type BackingImageSpec struct {
	// +optional
	Disks map[string]string `json:"disks"`
	// +optional
	Checksum string `json:"checksum"`
	// +optional
	SourceType BackingImageDataSourceType `json:"sourceType"`
	// +optional
	SourceParameters map[string]string `json:"sourceParameters"`
	// Deprecated: This kind of info will be included in the related BackingImageDataSource.
	// +optional
	ImageURL string `json:"imageURL"`
}

type BackingImageStatus struct {
	// +optional
	OwnerID string `json:"ownerID"`
	// +optional
	UUID string `json:"uuid"`
	// +optional
	Size int64 `json:"size"`
	// +optional
	Checksum string `json:"checksum"`
	// +optional
	DiskFileStatusMap map[string]*BackingImageDiskFileStatus `json:"diskFileStatusMap"`
	// +optional
	DiskLastRefAtMap map[string]string `json:"diskLastRefAtMap"`
	// Deprecated: Replaced by field `State` in `DiskFileStatusMap`.
	// +optional
	DiskDownloadStateMap map[string]BackingImageDownloadState `json:"diskDownloadStateMap"`
	// Deprecated: Replaced by field `Progress` in `DiskFileStatusMap`.
	// +optional
	DiskDownloadProgressMap map[string]int `json:"diskDownloadProgressMap"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName=lhbi
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Image",type=string,JSONPath=`.spec.image`,description="The backing image name"
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
type BackingImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackingImageSpec   `json:"spec,omitempty"`
	Status BackingImageStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BackingImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackingImage `json:"items"`
}