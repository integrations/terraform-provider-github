package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemNotificationsResponse 
// Deprecated: This class is obsolete. Use notificationsPutResponse instead.
type ItemItemNotificationsResponse struct {
    ItemItemNotificationsPutResponse
}
// NewItemItemNotificationsResponse instantiates a new ItemItemNotificationsResponse and sets the default values.
func NewItemItemNotificationsResponse()(*ItemItemNotificationsResponse) {
    m := &ItemItemNotificationsResponse{
        ItemItemNotificationsPutResponse: *NewItemItemNotificationsPutResponse(),
    }
    return m
}
// CreateItemItemNotificationsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemNotificationsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemNotificationsResponse(), nil
}
// ItemItemNotificationsResponseable 
// Deprecated: This class is obsolete. Use notificationsPutResponse instead.
type ItemItemNotificationsResponseable interface {
    ItemItemNotificationsPutResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
