package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimelineIssueEvents composed type wrapper for classes AddedToProjectIssueEventable, ConvertedNoteToIssueIssueEventable, DemilestonedIssueEventable, LabeledIssueEventable, LockedIssueEventable, MilestonedIssueEventable, MovedColumnInProjectIssueEventable, RemovedFromProjectIssueEventable, RenamedIssueEventable, ReviewDismissedIssueEventable, ReviewRequestedIssueEventable, ReviewRequestRemovedIssueEventable, StateChangeIssueEventable, TimelineAssignedIssueEventable, TimelineCommentEventable, TimelineCommitCommentedEventable, TimelineCommittedEventable, TimelineCrossReferencedEventable, TimelineLineCommentedEventable, TimelineReviewedEventable, TimelineUnassignedIssueEventable, UnlabeledIssueEventable
type TimelineIssueEvents struct {
    // Composed type representation for type AddedToProjectIssueEventable
    addedToProjectIssueEvent AddedToProjectIssueEventable
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
    // Composed type representation for type StateChangeIssueEventable
    stateChangeIssueEvent StateChangeIssueEventable
    // Composed type representation for type TimelineAssignedIssueEventable
    timelineAssignedIssueEvent TimelineAssignedIssueEventable
    // Composed type representation for type TimelineCommentEventable
    timelineCommentEvent TimelineCommentEventable
    // Composed type representation for type TimelineCommitCommentedEventable
    timelineCommitCommentedEvent TimelineCommitCommentedEventable
    // Composed type representation for type TimelineCommittedEventable
    timelineCommittedEvent TimelineCommittedEventable
    // Composed type representation for type TimelineCrossReferencedEventable
    timelineCrossReferencedEvent TimelineCrossReferencedEventable
    // Composed type representation for type TimelineLineCommentedEventable
    timelineLineCommentedEvent TimelineLineCommentedEventable
    // Composed type representation for type TimelineReviewedEventable
    timelineReviewedEvent TimelineReviewedEventable
    // Composed type representation for type TimelineUnassignedIssueEventable
    timelineUnassignedIssueEvent TimelineUnassignedIssueEventable
    // Composed type representation for type UnlabeledIssueEventable
    unlabeledIssueEvent UnlabeledIssueEventable
}
// NewTimelineIssueEvents instantiates a new TimelineIssueEvents and sets the default values.
func NewTimelineIssueEvents()(*TimelineIssueEvents) {
    m := &TimelineIssueEvents{
    }
    return m
}
// CreateTimelineIssueEventsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTimelineIssueEventsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewTimelineIssueEvents()
    if parseNode != nil {
        if val, err := parseNode.GetObjectValue(CreateAddedToProjectIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(AddedToProjectIssueEventable); ok {
                result.SetAddedToProjectIssueEvent(cast)
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
        } else if val, err := parseNode.GetObjectValue(CreateStateChangeIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(StateChangeIssueEventable); ok {
                result.SetStateChangeIssueEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateTimelineAssignedIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(TimelineAssignedIssueEventable); ok {
                result.SetTimelineAssignedIssueEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateTimelineCommentEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(TimelineCommentEventable); ok {
                result.SetTimelineCommentEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateTimelineCommitCommentedEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(TimelineCommitCommentedEventable); ok {
                result.SetTimelineCommitCommentedEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateTimelineCommittedEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(TimelineCommittedEventable); ok {
                result.SetTimelineCommittedEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateTimelineCrossReferencedEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(TimelineCrossReferencedEventable); ok {
                result.SetTimelineCrossReferencedEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateTimelineLineCommentedEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(TimelineLineCommentedEventable); ok {
                result.SetTimelineLineCommentedEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateTimelineReviewedEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(TimelineReviewedEventable); ok {
                result.SetTimelineReviewedEvent(cast)
            }
        } else if val, err := parseNode.GetObjectValue(CreateTimelineUnassignedIssueEventFromDiscriminatorValue); val != nil {
            if err != nil {
                return nil, err
            }
            if cast, ok := val.(TimelineUnassignedIssueEventable); ok {
                result.SetTimelineUnassignedIssueEvent(cast)
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
func (m *TimelineIssueEvents) GetAddedToProjectIssueEvent()(AddedToProjectIssueEventable) {
    return m.addedToProjectIssueEvent
}
// GetConvertedNoteToIssueIssueEvent gets the convertedNoteToIssueIssueEvent property value. Composed type representation for type ConvertedNoteToIssueIssueEventable
// returns a ConvertedNoteToIssueIssueEventable when successful
func (m *TimelineIssueEvents) GetConvertedNoteToIssueIssueEvent()(ConvertedNoteToIssueIssueEventable) {
    return m.convertedNoteToIssueIssueEvent
}
// GetDemilestonedIssueEvent gets the demilestonedIssueEvent property value. Composed type representation for type DemilestonedIssueEventable
// returns a DemilestonedIssueEventable when successful
func (m *TimelineIssueEvents) GetDemilestonedIssueEvent()(DemilestonedIssueEventable) {
    return m.demilestonedIssueEvent
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TimelineIssueEvents) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *TimelineIssueEvents) GetIsComposedType()(bool) {
    return true
}
// GetLabeledIssueEvent gets the labeledIssueEvent property value. Composed type representation for type LabeledIssueEventable
// returns a LabeledIssueEventable when successful
func (m *TimelineIssueEvents) GetLabeledIssueEvent()(LabeledIssueEventable) {
    return m.labeledIssueEvent
}
// GetLockedIssueEvent gets the lockedIssueEvent property value. Composed type representation for type LockedIssueEventable
// returns a LockedIssueEventable when successful
func (m *TimelineIssueEvents) GetLockedIssueEvent()(LockedIssueEventable) {
    return m.lockedIssueEvent
}
// GetMilestonedIssueEvent gets the milestonedIssueEvent property value. Composed type representation for type MilestonedIssueEventable
// returns a MilestonedIssueEventable when successful
func (m *TimelineIssueEvents) GetMilestonedIssueEvent()(MilestonedIssueEventable) {
    return m.milestonedIssueEvent
}
// GetMovedColumnInProjectIssueEvent gets the movedColumnInProjectIssueEvent property value. Composed type representation for type MovedColumnInProjectIssueEventable
// returns a MovedColumnInProjectIssueEventable when successful
func (m *TimelineIssueEvents) GetMovedColumnInProjectIssueEvent()(MovedColumnInProjectIssueEventable) {
    return m.movedColumnInProjectIssueEvent
}
// GetRemovedFromProjectIssueEvent gets the removedFromProjectIssueEvent property value. Composed type representation for type RemovedFromProjectIssueEventable
// returns a RemovedFromProjectIssueEventable when successful
func (m *TimelineIssueEvents) GetRemovedFromProjectIssueEvent()(RemovedFromProjectIssueEventable) {
    return m.removedFromProjectIssueEvent
}
// GetRenamedIssueEvent gets the renamedIssueEvent property value. Composed type representation for type RenamedIssueEventable
// returns a RenamedIssueEventable when successful
func (m *TimelineIssueEvents) GetRenamedIssueEvent()(RenamedIssueEventable) {
    return m.renamedIssueEvent
}
// GetReviewDismissedIssueEvent gets the reviewDismissedIssueEvent property value. Composed type representation for type ReviewDismissedIssueEventable
// returns a ReviewDismissedIssueEventable when successful
func (m *TimelineIssueEvents) GetReviewDismissedIssueEvent()(ReviewDismissedIssueEventable) {
    return m.reviewDismissedIssueEvent
}
// GetReviewRequestedIssueEvent gets the reviewRequestedIssueEvent property value. Composed type representation for type ReviewRequestedIssueEventable
// returns a ReviewRequestedIssueEventable when successful
func (m *TimelineIssueEvents) GetReviewRequestedIssueEvent()(ReviewRequestedIssueEventable) {
    return m.reviewRequestedIssueEvent
}
// GetReviewRequestRemovedIssueEvent gets the reviewRequestRemovedIssueEvent property value. Composed type representation for type ReviewRequestRemovedIssueEventable
// returns a ReviewRequestRemovedIssueEventable when successful
func (m *TimelineIssueEvents) GetReviewRequestRemovedIssueEvent()(ReviewRequestRemovedIssueEventable) {
    return m.reviewRequestRemovedIssueEvent
}
// GetStateChangeIssueEvent gets the stateChangeIssueEvent property value. Composed type representation for type StateChangeIssueEventable
// returns a StateChangeIssueEventable when successful
func (m *TimelineIssueEvents) GetStateChangeIssueEvent()(StateChangeIssueEventable) {
    return m.stateChangeIssueEvent
}
// GetTimelineAssignedIssueEvent gets the timelineAssignedIssueEvent property value. Composed type representation for type TimelineAssignedIssueEventable
// returns a TimelineAssignedIssueEventable when successful
func (m *TimelineIssueEvents) GetTimelineAssignedIssueEvent()(TimelineAssignedIssueEventable) {
    return m.timelineAssignedIssueEvent
}
// GetTimelineCommentEvent gets the timelineCommentEvent property value. Composed type representation for type TimelineCommentEventable
// returns a TimelineCommentEventable when successful
func (m *TimelineIssueEvents) GetTimelineCommentEvent()(TimelineCommentEventable) {
    return m.timelineCommentEvent
}
// GetTimelineCommitCommentedEvent gets the timelineCommitCommentedEvent property value. Composed type representation for type TimelineCommitCommentedEventable
// returns a TimelineCommitCommentedEventable when successful
func (m *TimelineIssueEvents) GetTimelineCommitCommentedEvent()(TimelineCommitCommentedEventable) {
    return m.timelineCommitCommentedEvent
}
// GetTimelineCommittedEvent gets the timelineCommittedEvent property value. Composed type representation for type TimelineCommittedEventable
// returns a TimelineCommittedEventable when successful
func (m *TimelineIssueEvents) GetTimelineCommittedEvent()(TimelineCommittedEventable) {
    return m.timelineCommittedEvent
}
// GetTimelineCrossReferencedEvent gets the timelineCrossReferencedEvent property value. Composed type representation for type TimelineCrossReferencedEventable
// returns a TimelineCrossReferencedEventable when successful
func (m *TimelineIssueEvents) GetTimelineCrossReferencedEvent()(TimelineCrossReferencedEventable) {
    return m.timelineCrossReferencedEvent
}
// GetTimelineLineCommentedEvent gets the timelineLineCommentedEvent property value. Composed type representation for type TimelineLineCommentedEventable
// returns a TimelineLineCommentedEventable when successful
func (m *TimelineIssueEvents) GetTimelineLineCommentedEvent()(TimelineLineCommentedEventable) {
    return m.timelineLineCommentedEvent
}
// GetTimelineReviewedEvent gets the timelineReviewedEvent property value. Composed type representation for type TimelineReviewedEventable
// returns a TimelineReviewedEventable when successful
func (m *TimelineIssueEvents) GetTimelineReviewedEvent()(TimelineReviewedEventable) {
    return m.timelineReviewedEvent
}
// GetTimelineUnassignedIssueEvent gets the timelineUnassignedIssueEvent property value. Composed type representation for type TimelineUnassignedIssueEventable
// returns a TimelineUnassignedIssueEventable when successful
func (m *TimelineIssueEvents) GetTimelineUnassignedIssueEvent()(TimelineUnassignedIssueEventable) {
    return m.timelineUnassignedIssueEvent
}
// GetUnlabeledIssueEvent gets the unlabeledIssueEvent property value. Composed type representation for type UnlabeledIssueEventable
// returns a UnlabeledIssueEventable when successful
func (m *TimelineIssueEvents) GetUnlabeledIssueEvent()(UnlabeledIssueEventable) {
    return m.unlabeledIssueEvent
}
// Serialize serializes information the current object
func (m *TimelineIssueEvents) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAddedToProjectIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetAddedToProjectIssueEvent())
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
    } else if m.GetStateChangeIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetStateChangeIssueEvent())
        if err != nil {
            return err
        }
    } else if m.GetTimelineAssignedIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetTimelineAssignedIssueEvent())
        if err != nil {
            return err
        }
    } else if m.GetTimelineCommentEvent() != nil {
        err := writer.WriteObjectValue("", m.GetTimelineCommentEvent())
        if err != nil {
            return err
        }
    } else if m.GetTimelineCommitCommentedEvent() != nil {
        err := writer.WriteObjectValue("", m.GetTimelineCommitCommentedEvent())
        if err != nil {
            return err
        }
    } else if m.GetTimelineCommittedEvent() != nil {
        err := writer.WriteObjectValue("", m.GetTimelineCommittedEvent())
        if err != nil {
            return err
        }
    } else if m.GetTimelineCrossReferencedEvent() != nil {
        err := writer.WriteObjectValue("", m.GetTimelineCrossReferencedEvent())
        if err != nil {
            return err
        }
    } else if m.GetTimelineLineCommentedEvent() != nil {
        err := writer.WriteObjectValue("", m.GetTimelineLineCommentedEvent())
        if err != nil {
            return err
        }
    } else if m.GetTimelineReviewedEvent() != nil {
        err := writer.WriteObjectValue("", m.GetTimelineReviewedEvent())
        if err != nil {
            return err
        }
    } else if m.GetTimelineUnassignedIssueEvent() != nil {
        err := writer.WriteObjectValue("", m.GetTimelineUnassignedIssueEvent())
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
func (m *TimelineIssueEvents) SetAddedToProjectIssueEvent(value AddedToProjectIssueEventable)() {
    m.addedToProjectIssueEvent = value
}
// SetConvertedNoteToIssueIssueEvent sets the convertedNoteToIssueIssueEvent property value. Composed type representation for type ConvertedNoteToIssueIssueEventable
func (m *TimelineIssueEvents) SetConvertedNoteToIssueIssueEvent(value ConvertedNoteToIssueIssueEventable)() {
    m.convertedNoteToIssueIssueEvent = value
}
// SetDemilestonedIssueEvent sets the demilestonedIssueEvent property value. Composed type representation for type DemilestonedIssueEventable
func (m *TimelineIssueEvents) SetDemilestonedIssueEvent(value DemilestonedIssueEventable)() {
    m.demilestonedIssueEvent = value
}
// SetLabeledIssueEvent sets the labeledIssueEvent property value. Composed type representation for type LabeledIssueEventable
func (m *TimelineIssueEvents) SetLabeledIssueEvent(value LabeledIssueEventable)() {
    m.labeledIssueEvent = value
}
// SetLockedIssueEvent sets the lockedIssueEvent property value. Composed type representation for type LockedIssueEventable
func (m *TimelineIssueEvents) SetLockedIssueEvent(value LockedIssueEventable)() {
    m.lockedIssueEvent = value
}
// SetMilestonedIssueEvent sets the milestonedIssueEvent property value. Composed type representation for type MilestonedIssueEventable
func (m *TimelineIssueEvents) SetMilestonedIssueEvent(value MilestonedIssueEventable)() {
    m.milestonedIssueEvent = value
}
// SetMovedColumnInProjectIssueEvent sets the movedColumnInProjectIssueEvent property value. Composed type representation for type MovedColumnInProjectIssueEventable
func (m *TimelineIssueEvents) SetMovedColumnInProjectIssueEvent(value MovedColumnInProjectIssueEventable)() {
    m.movedColumnInProjectIssueEvent = value
}
// SetRemovedFromProjectIssueEvent sets the removedFromProjectIssueEvent property value. Composed type representation for type RemovedFromProjectIssueEventable
func (m *TimelineIssueEvents) SetRemovedFromProjectIssueEvent(value RemovedFromProjectIssueEventable)() {
    m.removedFromProjectIssueEvent = value
}
// SetRenamedIssueEvent sets the renamedIssueEvent property value. Composed type representation for type RenamedIssueEventable
func (m *TimelineIssueEvents) SetRenamedIssueEvent(value RenamedIssueEventable)() {
    m.renamedIssueEvent = value
}
// SetReviewDismissedIssueEvent sets the reviewDismissedIssueEvent property value. Composed type representation for type ReviewDismissedIssueEventable
func (m *TimelineIssueEvents) SetReviewDismissedIssueEvent(value ReviewDismissedIssueEventable)() {
    m.reviewDismissedIssueEvent = value
}
// SetReviewRequestedIssueEvent sets the reviewRequestedIssueEvent property value. Composed type representation for type ReviewRequestedIssueEventable
func (m *TimelineIssueEvents) SetReviewRequestedIssueEvent(value ReviewRequestedIssueEventable)() {
    m.reviewRequestedIssueEvent = value
}
// SetReviewRequestRemovedIssueEvent sets the reviewRequestRemovedIssueEvent property value. Composed type representation for type ReviewRequestRemovedIssueEventable
func (m *TimelineIssueEvents) SetReviewRequestRemovedIssueEvent(value ReviewRequestRemovedIssueEventable)() {
    m.reviewRequestRemovedIssueEvent = value
}
// SetStateChangeIssueEvent sets the stateChangeIssueEvent property value. Composed type representation for type StateChangeIssueEventable
func (m *TimelineIssueEvents) SetStateChangeIssueEvent(value StateChangeIssueEventable)() {
    m.stateChangeIssueEvent = value
}
// SetTimelineAssignedIssueEvent sets the timelineAssignedIssueEvent property value. Composed type representation for type TimelineAssignedIssueEventable
func (m *TimelineIssueEvents) SetTimelineAssignedIssueEvent(value TimelineAssignedIssueEventable)() {
    m.timelineAssignedIssueEvent = value
}
// SetTimelineCommentEvent sets the timelineCommentEvent property value. Composed type representation for type TimelineCommentEventable
func (m *TimelineIssueEvents) SetTimelineCommentEvent(value TimelineCommentEventable)() {
    m.timelineCommentEvent = value
}
// SetTimelineCommitCommentedEvent sets the timelineCommitCommentedEvent property value. Composed type representation for type TimelineCommitCommentedEventable
func (m *TimelineIssueEvents) SetTimelineCommitCommentedEvent(value TimelineCommitCommentedEventable)() {
    m.timelineCommitCommentedEvent = value
}
// SetTimelineCommittedEvent sets the timelineCommittedEvent property value. Composed type representation for type TimelineCommittedEventable
func (m *TimelineIssueEvents) SetTimelineCommittedEvent(value TimelineCommittedEventable)() {
    m.timelineCommittedEvent = value
}
// SetTimelineCrossReferencedEvent sets the timelineCrossReferencedEvent property value. Composed type representation for type TimelineCrossReferencedEventable
func (m *TimelineIssueEvents) SetTimelineCrossReferencedEvent(value TimelineCrossReferencedEventable)() {
    m.timelineCrossReferencedEvent = value
}
// SetTimelineLineCommentedEvent sets the timelineLineCommentedEvent property value. Composed type representation for type TimelineLineCommentedEventable
func (m *TimelineIssueEvents) SetTimelineLineCommentedEvent(value TimelineLineCommentedEventable)() {
    m.timelineLineCommentedEvent = value
}
// SetTimelineReviewedEvent sets the timelineReviewedEvent property value. Composed type representation for type TimelineReviewedEventable
func (m *TimelineIssueEvents) SetTimelineReviewedEvent(value TimelineReviewedEventable)() {
    m.timelineReviewedEvent = value
}
// SetTimelineUnassignedIssueEvent sets the timelineUnassignedIssueEvent property value. Composed type representation for type TimelineUnassignedIssueEventable
func (m *TimelineIssueEvents) SetTimelineUnassignedIssueEvent(value TimelineUnassignedIssueEventable)() {
    m.timelineUnassignedIssueEvent = value
}
// SetUnlabeledIssueEvent sets the unlabeledIssueEvent property value. Composed type representation for type UnlabeledIssueEventable
func (m *TimelineIssueEvents) SetUnlabeledIssueEvent(value UnlabeledIssueEventable)() {
    m.unlabeledIssueEvent = value
}
type TimelineIssueEventsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddedToProjectIssueEvent()(AddedToProjectIssueEventable)
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
    GetStateChangeIssueEvent()(StateChangeIssueEventable)
    GetTimelineAssignedIssueEvent()(TimelineAssignedIssueEventable)
    GetTimelineCommentEvent()(TimelineCommentEventable)
    GetTimelineCommitCommentedEvent()(TimelineCommitCommentedEventable)
    GetTimelineCommittedEvent()(TimelineCommittedEventable)
    GetTimelineCrossReferencedEvent()(TimelineCrossReferencedEventable)
    GetTimelineLineCommentedEvent()(TimelineLineCommentedEventable)
    GetTimelineReviewedEvent()(TimelineReviewedEventable)
    GetTimelineUnassignedIssueEvent()(TimelineUnassignedIssueEventable)
    GetUnlabeledIssueEvent()(UnlabeledIssueEventable)
    SetAddedToProjectIssueEvent(value AddedToProjectIssueEventable)()
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
    SetStateChangeIssueEvent(value StateChangeIssueEventable)()
    SetTimelineAssignedIssueEvent(value TimelineAssignedIssueEventable)()
    SetTimelineCommentEvent(value TimelineCommentEventable)()
    SetTimelineCommitCommentedEvent(value TimelineCommitCommentedEventable)()
    SetTimelineCommittedEvent(value TimelineCommittedEventable)()
    SetTimelineCrossReferencedEvent(value TimelineCrossReferencedEventable)()
    SetTimelineLineCommentedEvent(value TimelineLineCommentedEventable)()
    SetTimelineReviewedEvent(value TimelineReviewedEventable)()
    SetTimelineUnassignedIssueEvent(value TimelineUnassignedIssueEventable)()
    SetUnlabeledIssueEvent(value UnlabeledIssueEventable)()
}
