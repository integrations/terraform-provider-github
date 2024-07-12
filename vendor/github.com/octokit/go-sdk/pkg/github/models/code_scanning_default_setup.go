package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeScanningDefaultSetup configuration for code scanning default setup.
type CodeScanningDefaultSetup struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Languages to be analyzed.
    languages []CodeScanningDefaultSetup_languages
    // CodeQL query suite to be used.
    query_suite *CodeScanningDefaultSetup_query_suite
    // The frequency of the periodic analysis.
    schedule *CodeScanningDefaultSetup_schedule
    // Code scanning default setup has been configured or not.
    state *CodeScanningDefaultSetup_state
    // Timestamp of latest configuration update.
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewCodeScanningDefaultSetup instantiates a new CodeScanningDefaultSetup and sets the default values.
func NewCodeScanningDefaultSetup()(*CodeScanningDefaultSetup) {
    m := &CodeScanningDefaultSetup{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningDefaultSetupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningDefaultSetupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningDefaultSetup(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningDefaultSetup) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningDefaultSetup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["languages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseCodeScanningDefaultSetup_languages)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CodeScanningDefaultSetup_languages, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*CodeScanningDefaultSetup_languages))
                }
            }
            m.SetLanguages(res)
        }
        return nil
    }
    res["query_suite"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeScanningDefaultSetup_query_suite)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuerySuite(val.(*CodeScanningDefaultSetup_query_suite))
        }
        return nil
    }
    res["schedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeScanningDefaultSetup_schedule)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(*CodeScanningDefaultSetup_schedule))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeScanningDefaultSetup_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*CodeScanningDefaultSetup_state))
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    return res
}
// GetLanguages gets the languages property value. Languages to be analyzed.
// returns a []CodeScanningDefaultSetup_languages when successful
func (m *CodeScanningDefaultSetup) GetLanguages()([]CodeScanningDefaultSetup_languages) {
    return m.languages
}
// GetQuerySuite gets the query_suite property value. CodeQL query suite to be used.
// returns a *CodeScanningDefaultSetup_query_suite when successful
func (m *CodeScanningDefaultSetup) GetQuerySuite()(*CodeScanningDefaultSetup_query_suite) {
    return m.query_suite
}
// GetSchedule gets the schedule property value. The frequency of the periodic analysis.
// returns a *CodeScanningDefaultSetup_schedule when successful
func (m *CodeScanningDefaultSetup) GetSchedule()(*CodeScanningDefaultSetup_schedule) {
    return m.schedule
}
// GetState gets the state property value. Code scanning default setup has been configured or not.
// returns a *CodeScanningDefaultSetup_state when successful
func (m *CodeScanningDefaultSetup) GetState()(*CodeScanningDefaultSetup_state) {
    return m.state
}
// GetUpdatedAt gets the updated_at property value. Timestamp of latest configuration update.
// returns a *Time when successful
func (m *CodeScanningDefaultSetup) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *CodeScanningDefaultSetup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLanguages() != nil {
        err := writer.WriteCollectionOfStringValues("languages", SerializeCodeScanningDefaultSetup_languages(m.GetLanguages()))
        if err != nil {
            return err
        }
    }
    if m.GetQuerySuite() != nil {
        cast := (*m.GetQuerySuite()).String()
        err := writer.WriteStringValue("query_suite", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSchedule() != nil {
        cast := (*m.GetSchedule()).String()
        err := writer.WriteStringValue("schedule", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CodeScanningDefaultSetup) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLanguages sets the languages property value. Languages to be analyzed.
func (m *CodeScanningDefaultSetup) SetLanguages(value []CodeScanningDefaultSetup_languages)() {
    m.languages = value
}
// SetQuerySuite sets the query_suite property value. CodeQL query suite to be used.
func (m *CodeScanningDefaultSetup) SetQuerySuite(value *CodeScanningDefaultSetup_query_suite)() {
    m.query_suite = value
}
// SetSchedule sets the schedule property value. The frequency of the periodic analysis.
func (m *CodeScanningDefaultSetup) SetSchedule(value *CodeScanningDefaultSetup_schedule)() {
    m.schedule = value
}
// SetState sets the state property value. Code scanning default setup has been configured or not.
func (m *CodeScanningDefaultSetup) SetState(value *CodeScanningDefaultSetup_state)() {
    m.state = value
}
// SetUpdatedAt sets the updated_at property value. Timestamp of latest configuration update.
func (m *CodeScanningDefaultSetup) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
type CodeScanningDefaultSetupable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLanguages()([]CodeScanningDefaultSetup_languages)
    GetQuerySuite()(*CodeScanningDefaultSetup_query_suite)
    GetSchedule()(*CodeScanningDefaultSetup_schedule)
    GetState()(*CodeScanningDefaultSetup_state)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetLanguages(value []CodeScanningDefaultSetup_languages)()
    SetQuerySuite(value *CodeScanningDefaultSetup_query_suite)()
    SetSchedule(value *CodeScanningDefaultSetup_schedule)()
    SetState(value *CodeScanningDefaultSetup_state)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
