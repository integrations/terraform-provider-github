package notifications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NotificationsResponse 
// Deprecated: This class is obsolete. Use notificationsPutResponse instead.
type NotificationsResponse struct {
    NotificationsPutResponse
}
// NewNotificationsResponse instantiates a new notificationsResponse and sets the default values.
func NewNotificationsResponse()(*NotificationsResponse) {
    m := &NotificationsResponse{
        NotificationsPutResponse: *NewNotificationsPutResponse(),
    }
    return m
}
// CreateNotificationsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNotificationsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotificationsResponse(), nil
}
// NotificationsResponseable 
// Deprecated: This class is obsolete. Use notificationsPutResponse instead.
type NotificationsResponseable interface {
    NotificationsPutResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
