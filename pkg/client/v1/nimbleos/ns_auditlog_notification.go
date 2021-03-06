// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAuditlogNotification - Represents an auditlog notification message.
// Export NsAuditlogNotificationFields for advance operations like search filter etc.
var NsAuditlogNotificationFields *NsAuditlogNotification

func init() {
	Activityfield := "activity"
	ObjectIdfield := "object_id"
	ObjectNamefield := "object_name"

	NsAuditlogNotificationFields = &NsAuditlogNotification{
		Activity:   &Activityfield,
		ObjectId:   &ObjectIdfield,
		ObjectName: &ObjectNamefield,
	}
}

type NsAuditlogNotification struct {
	// SequenceNumber - Notification Sequence Number.
	SequenceNumber *int64 `json:"sequence_number,omitempty"`
	// NotificationType - Represents the type of the notification.
	NotificationType *NsNotificationType `json:"notification_type,omitempty"`
	// Activity - Represents CUD message of auditlog notification.
	Activity *string `json:"activity,omitempty"`
	// ObjectType - Represents the object type of an auditlog based notification.
	ObjectType *NsObjectType `json:"object_type,omitempty"`
	// ObjectId - Represents the object of an auditlog based notification.
	ObjectId *string `json:"object_id,omitempty"`
	// ObjectName - Represents the object name of an auditlog based notification.
	ObjectName *string `json:"object_name,omitempty"`
	// Timestamp - The timestamp when the activity happened in seconds since last epoch.
	Timestamp *int64 `json:"timestamp,omitempty"`
}
