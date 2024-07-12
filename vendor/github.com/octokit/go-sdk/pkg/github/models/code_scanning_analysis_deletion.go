package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CodeScanningAnalysisDeletion successful deletion of a code scanning analysis
type CodeScanningAnalysisDeletion struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Next deletable analysis in chain, with last analysis deletion confirmation
    confirm_delete_url *string
    // Next deletable analysis in chain, without last analysis deletion confirmation
    next_analysis_url *string
}
// NewCodeScanningAnalysisDeletion instantiates a new CodeScanningAnalysisDeletion and sets the default values.
func NewCodeScanningAnalysisDeletion()(*CodeScanningAnalysisDeletion) {
    m := &CodeScanningAnalysisDeletion{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCodeScanningAnalysisDeletionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCodeScanningAnalysisDeletionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCodeScanningAnalysisDeletion(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CodeScanningAnalysisDeletion) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConfirmDeleteUrl gets the confirm_delete_url property value. Next deletable analysis in chain, with last analysis deletion confirmation
// returns a *string when successful
func (m *CodeScanningAnalysisDeletion) GetConfirmDeleteUrl()(*string) {
    return m.confirm_delete_url
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CodeScanningAnalysisDeletion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["confirm_delete_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfirmDeleteUrl(val)
        }
        return nil
    }
    res["next_analysis_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextAnalysisUrl(val)
        }
        return nil
    }
    return res
}
// GetNextAnalysisUrl gets the next_analysis_url property value. Next deletable analysis in chain, without last analysis deletion confirmation
// returns a *string when successful
func (m *CodeScanningAnalysisDeletion) GetNextAnalysisUrl()(*string) {
    return m.next_analysis_url
}
// Serialize serializes information the current object
func (m *CodeScanningAnalysisDeletion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CodeScanningAnalysisDeletion) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConfirmDeleteUrl sets the confirm_delete_url property value. Next deletable analysis in chain, with last analysis deletion confirmation
func (m *CodeScanningAnalysisDeletion) SetConfirmDeleteUrl(value *string)() {
    m.confirm_delete_url = value
}
// SetNextAnalysisUrl sets the next_analysis_url property value. Next deletable analysis in chain, without last analysis deletion confirmation
func (m *CodeScanningAnalysisDeletion) SetNextAnalysisUrl(value *string)() {
    m.next_analysis_url = value
}
type CodeScanningAnalysisDeletionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfirmDeleteUrl()(*string)
    GetNextAnalysisUrl()(*string)
    SetConfirmDeleteUrl(value *string)()
    SetNextAnalysisUrl(value *string)()
}
