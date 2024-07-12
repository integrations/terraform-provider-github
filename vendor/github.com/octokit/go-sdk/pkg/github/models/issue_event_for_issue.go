package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IssueEventForIssue composed type wrapper for classes AddedToProjectIssueEventable, AssignedIssueEventable, ConvertedNoteToIssueIssueEventable, DemilestonedIssueEventable, LabeledIssueEventable, LockedIssueEventable, MilestonedIssueEventable, MovedColumnInProjectIssueEventable, RemovedFromProjectIssueEventable, RenamedIssueEventable, ReviewDismissedIssueEventable, ReviewRequestedIssueEventable, ReviewRequestRemovedIssueEventable, UnassignedIssueEventable, UnlabeledIssueEventable
type IssueEventForIssue struct {
    // Composed type representation for type AddedToProjectIssueEventable
    addedToProjectIssueEvent AddedToProjectIssueEventable
    // Composed type representation for type AssignedIssueEventable
    assignedIssueEvent AssignedIssueEventable
    // Composed type representation for type ConvertedNoteToIssueIssueEventable
    convertedNoteToIssueIssueEvent ConvertedNoteToIssueIssueEventable
    // Composed type representation for type DemilestonedIssueEventable
    demilestonedIssueEvent DemilestonedIssueEventable
    // Composed type representation for type LabeledIssueEventable
    labeledIssueEvent LabeledIssueEventable
    // Composed type representation for type LockedIssueEventable
    lockedIssueEvent LockedIssueEventable
    // Composed type representation for type MilestonedIssueEventable
    milestonedIssueEvent MilestonedIssueEventable
    // Composed type representation for type MovedColumnInProjectIssueEventable
    movedColumnInProjectIssueEvent MovedColumnInProjectIssueEventable
    // Composed type representation for type RemovedFromProjectIssueEventable
    removedFromProjectIssueEvent RemovedFromProjectIssueEventable
    // Composed type representation for type RenamedIssueEventable
    renamedIssueEvent RenamedIssueEventable
    // Composed type representation for type ReviewDismissedIssueEventable
    reviewDismissedIssueEvent ReviewDismissedIssueEventable
    // Composed type representation for type ReviewRequestedIssueEventable
    reviewRequestedIssueEvent ReviewRequestedIssueEventable
    // Composed type representation for type ReviewRequestRemovedIssueEventable
    reviewRequestRemovedIssueEvent ReviewRequestRemovedIssueEventable
    // Composed type representation for type UnassignedIssueEventable
    unassignedIssueEvent UnassignedIssueEventable
    // Composed type representation for type UnlabeledIssueEventable
    unlabeledIssueEvent UnlabeledIssueEventable
}
// NewIssueEventForIssue instantiates a new IssueEventForIssue and sets the default values.
func NewIssueEventForIssue()(*IssueEventForIssue) {
    m := &IssueEventForIssue{
    }
    return m
}
// CreateIssueEventForIssueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIssueEventForIssueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewIssueEventForIssue()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateAddedToProjectIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(AddedToProjectIssueEventable); ok {
                result.SetAddedToProjectIssueEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateAssignedIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(AssignedIssueEventable); ok {
                result.SetAssignedIssueEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateConvertedNoteToIssueIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ConvertedNoteToIssueIssueEventable); ok {
                result.SetConvertedNoteToIssueIssueEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateDemilestonedIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(DemilestonedIssueEventable); ok {
                result.SetDemilestonedIssueEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateLabeledIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(LabeledIssueEventable); ok {
                result.SetLabeledIssueEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateLockedIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(LockedIssueEventable); ok {
                result.SetLockedIssueEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateMilestonedIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(MilestonedIssueEventable); ok {
                result.SetMilestonedIssueEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateMovedColumnInProjectIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(MovedColumnInProjectIssueEventable); ok {
                result.SetMovedColumnInProjectIssueEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateRemovedFromProjectIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(RemovedFromProjectIssueEventable); ok {
                result.SetRemovedFromProjectIssueEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateRenamedIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(RenamedIssueEventable); ok {
                result.SetRenamedIssueEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateReviewDismissedIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ReviewDismissedIssueEventable); ok {
                result.SetReviewDismissedIssueEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateReviewRequestedIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ReviewRequestedIssueEventable); ok {
                result.SetReviewRequestedIssueEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateReviewRequestRemovedIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(ReviewRequestRemovedIssueEventable); ok {
                result.SetReviewRequestRemovedIssueEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateUnassignedIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(UnassignedIssueEventable); ok {
                result.SetUnassignedIssueEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateUnlabeledIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(UnlabeledIssueEventable); ok {
                result.SetUnlabeledIssueEvent(cast)
            }
        }
    }
    return result, nil
}
// GetAddedToProjectIssueEvent gets the addedToProjectIssueEvent property value. Composed type representation for type AddedToProjectIssueEventable
// returns a AddedToProjectIssueEventable when successful
func (m *IssueEventForIssue) GetAddedToProjectIssueEvent()(AddedToProjectIssueEventable) {
    return m.addedToProjectIssueEvent
}
// GetAssignedIssueEvent gets the assignedIssueEvent property value. Composed type representation for type AssignedIssueEventable
// returns a AssignedIssueEventable when successful
func (m *IssueEventForIssue) GetAssignedIssueEvent()(AssignedIssueEventable) {
    return m.assignedIssueEvent
}
// GetConvertedNoteToIssueIssueEvent gets the convertedNoteToIssueIssueEvent property value. Composed type representation for type ConvertedNoteToIssueIssueEventable
// returns a ConvertedNoteToIssueIssueEventable when successful
func (m *IssueEventForIssue) GetConvertedNoteToIssueIssueEvent()(ConvertedNoteToIssueIssueEventable) {
    return m.convertedNoteToIssueIssueEvent
}
// GetDemilestonedIssueEvent gets the demilestonedIssueEvent property value. Composed type representation for type DemilestonedIssueEventable
// returns a DemilestonedIssueEventable when successful
func (m *IssueEventForIssue) GetDemilestonedIssueEvent()(DemilestonedIssueEventable) {
    return m.demilestonedIssueEvent
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IssueEventForIssue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *IssueEventForIssue) GetIsComposedType()(bool) {
    return true
}
// GetLabeledIssueEvent gets the labeledIssueEvent property value. Composed type representation for type LabeledIssueEventable
// returns a LabeledIssueEventable when successful
func (m *IssueEventForIssue) GetLabeledIssueEvent()(LabeledIssueEventable) {
    return m.labeledIssueEvent
}
// GetLockedIssueEvent gets the lockedIssueEvent property value. Composed type representation for type LockedIssueEventable
// returns a LockedIssueEventable when successful
func (m *IssueEventForIssue) GetLockedIssueEvent()(LockedIssueEventable) {
    return m.lockedIssueEvent
}
// GetMilestonedIssueEvent gets the milestonedIssueEvent property value. Composed type representation for type MilestonedIssueEventable
// returns a MilestonedIssueEventable when successful
func (m *IssueEventForIssue) GetMilestonedIssueEvent()(MilestonedIssueEventable) {
    return m.milestonedIssueEvent
}
// GetMovedColumnInProjectIssueEvent gets the movedColumnInProjectIssueEvent property value. Composed type representation for type MovedColumnInProjectIssueEventable
// returns a MovedColumnInProjectIssueEventable when successful
func (m *IssueEventForIssue) GetMovedColumnInProjectIssueEvent()(MovedColumnInProjectIssueEventable) {
    return m.movedColumnInProjectIssueEvent
}
// GetRemovedFromProjectIssueEvent gets the removedFromProjectIssueEvent property value. Composed type representation for type RemovedFromProjectIssueEventable
// returns a RemovedFromProjectIssueEventable when successful
func (m *IssueEventForIssue) GetRemovedFromProjectIssueEvent()(RemovedFromProjectIssueEventable) {
    return m.removedFromProjectIssueEvent
}
// GetRenamedIssueEvent gets the renamedIssueEvent property value. Composed type representation for type RenamedIssueEventable
// returns a RenamedIssueEventable when successful
func (m *IssueEventForIssue) GetRenamedIssueEvent()(RenamedIssueEventable) {
    return m.renamedIssueEvent
}
// GetReviewDismissedIssueEvent gets the reviewDismissedIssueEvent property value. Composed type representation for type ReviewDismissedIssueEventable
// returns a ReviewDismissedIssueEventable when successful
func (m *IssueEventForIssue) GetReviewDismissedIssueEvent()(ReviewDismissedIssueEventable) {
    return m.reviewDismissedIssueEvent
}
// GetReviewRequestedIssueEvent gets the reviewRequestedIssueEvent property value. Composed type representation for type ReviewRequestedIssueEventable
// returns a ReviewRequestedIssueEventable when successful
func (m *IssueEventForIssue) GetReviewRequestedIssueEvent()(ReviewRequestedIssueEventable) {
    return m.reviewRequestedIssueEvent
}
// GetReviewRequestRemovedIssueEvent gets the reviewRequestRemovedIssueEvent property value. Composed type representation for type ReviewRequestRemovedIssueEventable
// returns a ReviewRequestRemovedIssueEventable when successful
func (m *IssueEventForIssue) GetReviewRequestRemovedIssueEvent()(ReviewRequestRemovedIssueEventable) {
    return m.reviewRequestRemovedIssueEvent
}
// GetUnassignedIssueEvent gets the unassignedIssueEvent property value. Composed type representation for type UnassignedIssueEventable
// returns a UnassignedIssueEventable when successful
func (m *IssueEventForIssue) GetUnassignedIssueEvent()(UnassignedIssueEventable) {
    return m.unassignedIssueEvent
}
// GetUnlabeledIssueEvent gets the unlabeledIssueEvent property value. Composed type representation for type UnlabeledIssueEventable
// returns a UnlabeledIssueEventable when successful
func (m *IssueEventForIssue) GetUnlabeledIssueEvent()(UnlabeledIssueEventable) {
    return m.unlabeledIssueEvent
}
// Serialize serializes information the current object
func (m *IssueEventForIssue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAddedToProjectIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetAddedToProjectIssueEvent())
        if err != nil {
            return err
        }
    } else if m.GetAssignedIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetAssignedIssueEvent())
        if err != nil {
            return err
        }
    } else if m.GetConvertedNoteToIssueIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetConvertedNoteToIssueIssueEvent())
        if err != nil {
            return err
        }
    } else if m.GetDemilestonedIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetDemilestonedIssueEvent())
        if err != nil {
            return err
        }
    } else if m.GetLabeledIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetLabeledIssueEvent())
        if err != nil {
            return err
        }
    } else if m.GetLockedIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetLockedIssueEvent())
        if err != nil {
            return err
        }
    } else if m.GetMilestonedIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetMilestonedIssueEvent())
        if err != nil {
            return err
        }
    } else if m.GetMovedColumnInProjectIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetMovedColumnInProjectIssueEvent())
        if err != nil {
            return err
        }
    } else if m.GetRemovedFromProjectIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetRemovedFromProjectIssueEvent())
        if err != nil {
            return err
        }
    } else if m.GetRenamedIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetRenamedIssueEvent())
        if err != nil {
            return err
        }
    } else if m.GetReviewDismissedIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetReviewDismissedIssueEvent())
        if err != nil {
            return err
        }
    } else if m.GetReviewRequestedIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetReviewRequestedIssueEvent())
        if err != nil {
            return err
        }
    } else if m.GetReviewRequestRemovedIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetReviewRequestRemovedIssueEvent())
        if err != nil {
            return err
        }
    } else if m.GetUnassignedIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetUnassignedIssueEvent())
        if err != nil {
            return err
        }
    } else if m.GetUnlabeledIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetUnlabeledIssueEvent())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddedToProjectIssueEvent sets the addedToProjectIssueEvent property value. Composed type representation for type AddedToProjectIssueEventable
func (m *IssueEventForIssue) SetAddedToProjectIssueEvent(value AddedToProjectIssueEventable)() {
    m.addedToProjectIssueEvent = value
}
// SetAssignedIssueEvent sets the assignedIssueEvent property value. Composed type representation for type AssignedIssueEventable
func (m *IssueEventForIssue) SetAssignedIssueEvent(value AssignedIssueEventable)() {
    m.assignedIssueEvent = value
}
// SetConvertedNoteToIssueIssueEvent sets the convertedNoteToIssueIssueEvent property value. Composed type representation for type ConvertedNoteToIssueIssueEventable
func (m *IssueEventForIssue) SetConvertedNoteToIssueIssueEvent(value ConvertedNoteToIssueIssueEventable)() {
    m.convertedNoteToIssueIssueEvent = value
}
// SetDemilestonedIssueEvent sets the demilestonedIssueEvent property value. Composed type representation for type DemilestonedIssueEventable
func (m *IssueEventForIssue) SetDemilestonedIssueEvent(value DemilestonedIssueEventable)() {
    m.demilestonedIssueEvent = value
}
// SetLabeledIssueEvent sets the labeledIssueEvent property value. Composed type representation for type LabeledIssueEventable
func (m *IssueEventForIssue) SetLabeledIssueEvent(value LabeledIssueEventable)() {
    m.labeledIssueEvent = value
}
// SetLockedIssueEvent sets the lockedIssueEvent property value. Composed type representation for type LockedIssueEventable
func (m *IssueEventForIssue) SetLockedIssueEvent(value LockedIssueEventable)() {
    m.lockedIssueEvent = value
}
// SetMilestonedIssueEvent sets the milestonedIssueEvent property value. Composed type representation for type MilestonedIssueEventable
func (m *IssueEventForIssue) SetMilestonedIssueEvent(value MilestonedIssueEventable)() {
    m.milestonedIssueEvent = value
}
// SetMovedColumnInProjectIssueEvent sets the movedColumnInProjectIssueEvent property value. Composed type representation for type MovedColumnInProjectIssueEventable
func (m *IssueEventForIssue) SetMovedColumnInProjectIssueEvent(value MovedColumnInProjectIssueEventable)() {
    m.movedColumnInProjectIssueEvent = value
}
// SetRemovedFromProjectIssueEvent sets the removedFromProjectIssueEvent property value. Composed type representation for type RemovedFromProjectIssueEventable
func (m *IssueEventForIssue) SetRemovedFromProjectIssueEvent(value RemovedFromProjectIssueEventable)() {
    m.removedFromProjectIssueEvent = value
}
// SetRenamedIssueEvent sets the renamedIssueEvent property value. Composed type representation for type RenamedIssueEventable
func (m *IssueEventForIssue) SetRenamedIssueEvent(value RenamedIssueEventable)() {
    m.renamedIssueEvent = value
}
// SetReviewDismissedIssueEvent sets the reviewDismissedIssueEvent property value. Composed type representation for type ReviewDismissedIssueEventable
func (m *IssueEventForIssue) SetReviewDismissedIssueEvent(value ReviewDismissedIssueEventable)() {
    m.reviewDismissedIssueEvent = value
}
// SetReviewRequestedIssueEvent sets the reviewRequestedIssueEvent property value. Composed type representation for type ReviewRequestedIssueEventable
func (m *IssueEventForIssue) SetReviewRequestedIssueEvent(value ReviewRequestedIssueEventable)() {
    m.reviewRequestedIssueEvent = value
}
// SetReviewRequestRemovedIssueEvent sets the reviewRequestRemovedIssueEvent property value. Composed type representation for type ReviewRequestRemovedIssueEventable
func (m *IssueEventForIssue) SetReviewRequestRemovedIssueEvent(value ReviewRequestRemovedIssueEventable)() {
    m.reviewRequestRemovedIssueEvent = value
}
// SetUnassignedIssueEvent sets the unassignedIssueEvent property value. Composed type representation for type UnassignedIssueEventable
func (m *IssueEventForIssue) SetUnassignedIssueEvent(value UnassignedIssueEventable)() {
    m.unassignedIssueEvent = value
}
// SetUnlabeledIssueEvent sets the unlabeledIssueEvent property value. Composed type representation for type UnlabeledIssueEventable
func (m *IssueEventForIssue) SetUnlabeledIssueEvent(value UnlabeledIssueEventable)() {
    m.unlabeledIssueEvent = value
}
type IssueEventForIssueable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddedToProjectIssueEvent()(AddedToProjectIssueEventable)
    GetAssignedIssueEvent()(AssignedIssueEventable)
    GetConvertedNoteToIssueIssueEvent()(ConvertedNoteToIssueIssueEventable)
    GetDemilestonedIssueEvent()(DemilestonedIssueEventable)
    GetLabeledIssueEvent()(LabeledIssueEventable)
    GetLockedIssueEvent()(LockedIssueEventable)
    GetMilestonedIssueEvent()(MilestonedIssueEventable)
    GetMovedColumnInProjectIssueEvent()(MovedColumnInProjectIssueEventable)
    GetRemovedFromProjectIssueEvent()(RemovedFromProjectIssueEventable)
    GetRenamedIssueEvent()(RenamedIssueEventable)
    GetReviewDismissedIssueEvent()(ReviewDismissedIssueEventable)
    GetReviewRequestedIssueEvent()(ReviewRequestedIssueEventable)
    GetReviewRequestRemovedIssueEvent()(ReviewRequestRemovedIssueEventable)
    GetUnassignedIssueEvent()(UnassignedIssueEventable)
    GetUnlabeledIssueEvent()(UnlabeledIssueEventable)
    SetAddedToProjectIssueEvent(value AddedToProjectIssueEventable)()
    SetAssignedIssueEvent(value AssignedIssueEventable)()
    SetConvertedNoteToIssueIssueEvent(value ConvertedNoteToIssueIssueEventable)()
    SetDemilestonedIssueEvent(value DemilestonedIssueEventable)()
    SetLabeledIssueEvent(value LabeledIssueEventable)()
    SetLockedIssueEvent(value LockedIssueEventable)()
    SetMilestonedIssueEvent(value MilestonedIssueEventable)()
    SetMovedColumnInProjectIssueEvent(value MovedColumnInProjectIssueEventable)()
    SetRemovedFromProjectIssueEvent(value RemovedFromProjectIssueEventable)()
    SetRenamedIssueEvent(value RenamedIssueEventable)()
    SetReviewDismissedIssueEvent(value ReviewDismissedIssueEventable)()
    SetReviewRequestedIssueEvent(value ReviewRequestedIssueEventable)()
    SetReviewRequestRemovedIssueEvent(value ReviewRequestRemovedIssueEventable)()
    SetUnassignedIssueEvent(value UnassignedIssueEventable)()
    SetUnlabeledIssueEvent(value UnlabeledIssueEventable)()
}
