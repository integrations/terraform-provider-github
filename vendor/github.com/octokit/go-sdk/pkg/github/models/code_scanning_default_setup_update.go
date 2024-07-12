package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeScanningDefaultSetupUpdate configuration for code scanning default setup.
type CodeScanningDefaultSetupUpdate struct {
    // CodeQL languages to be analyzed.
    languages []CodeScanningDefaultSetupUpdate_languages
    // CodeQL query suite to be used.
    query_suite *CodeScanningDefaultSetupUpdate_query_suite
    // The desired state of code scanning default setup.
    state *CodeScanningDefaultSetupUpdate_state
}
// NewCodeScanningDefaultSetupUpdate instantiates a new CodeScanningDefaultSetupUpdate and sets the default values.
func NewCodeScanningDefaultSetupUpdate()(*CodeScanningDefaultSetupUpdate) {
    m := &CodeScanningDefaultSetupUpdate{
    }
    return m
}
// CreateCodeScanningDefaultSetupUpdateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningDefaultSetupUpdateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningDefaultSetupUpdate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningDefaultSetupUpdate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["languages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseCodeScanningDefaultSetupUpdate_languages)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CodeScanningDefaultSetupUpdate_languages, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*CodeScanningDefaultSetupUpdate_languages))
                }
            }
            m.SetLanguages(res)
        }
        return nil
    }
    res["query_suite"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeScanningDefaultSetupUpdate_query_suite)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuerySuite(val.(*CodeScanningDefaultSetupUpdate_query_suite))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCodeScanningDefaultSetupUpdate_state)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*CodeScanningDefaultSetupUpdate_state))
        }
        return nil
    }
    return res
}
// GetLanguages gets the languages property value. CodeQL languages to be analyzed.
// returns a []CodeScanningDefaultSetupUpdate_languages when successful
func (m *CodeScanningDefaultSetupUpdate) GetLanguages()([]CodeScanningDefaultSetupUpdate_languages) {
    return m.languages
}
// GetQuerySuite gets the query_suite property value. CodeQL query suite to be used.
// returns a *CodeScanningDefaultSetupUpdate_query_suite when successful
func (m *CodeScanningDefaultSetupUpdate) GetQuerySuite()(*CodeScanningDefaultSetupUpdate_query_suite) {
    return m.query_suite
}
// GetState gets the state property value. The desired state of code scanning default setup.
// returns a *CodeScanningDefaultSetupUpdate_state when successful
func (m *CodeScanningDefaultSetupUpdate) GetState()(*CodeScanningDefaultSetupUpdate_state) {
    return m.state
}
// Serialize serializes information the current object
func (m *CodeScanningDefaultSetupUpdate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetLanguages() != nil {
        err := writer.WriteCollectionOfStringValues("languages", SerializeCodeScanningDefaultSetupUpdate_languages(m.GetLanguages()))
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
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLanguages sets the languages property value. CodeQL languages to be analyzed.
func (m *CodeScanningDefaultSetupUpdate) SetLanguages(value []CodeScanningDefaultSetupUpdate_languages)() {
    m.languages = value
}
// SetQuerySuite sets the query_suite property value. CodeQL query suite to be used.
func (m *CodeScanningDefaultSetupUpdate) SetQuerySuite(value *CodeScanningDefaultSetupUpdate_query_suite)() {
    m.query_suite = value
}
// SetState sets the state property value. The desired state of code scanning default setup.
func (m *CodeScanningDefaultSetupUpdate) SetState(value *CodeScanningDefaultSetupUpdate_state)() {
    m.state = value
}
type CodeScanningDefaultSetupUpdateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLanguages()([]CodeScanningDefaultSetupUpdate_languages)
    GetQuerySuite()(*CodeScanningDefaultSetupUpdate_query_suite)
    GetState()(*CodeScanningDefaultSetupUpdate_state)
    SetLanguages(value []CodeScanningDefaultSetupUpdate_languages)()
    SetQuerySuite(value *CodeScanningDefaultSetupUpdate_query_suite)()
    SetState(value *CodeScanningDefaultSetupUpdate_state)()
}
