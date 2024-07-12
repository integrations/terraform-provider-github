package repos

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
)

type ItemItemActionsWorkflowsItemRunsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The total_count property
    total_count *int32
    // The workflow_runs property
    workflow_runs []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WorkflowRunable
}
// NewItemItemActionsWorkflowsItemRunsGetResponse instantiates a new ItemItemActionsWorkflowsItemRunsGetResponse and sets the default values.
func NewItemItemActionsWorkflowsItemRunsGetResponse()(*ItemItemActionsWorkflowsItemRunsGetResponse) {
    m := &ItemItemActionsWorkflowsItemRunsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemActionsWorkflowsItemRunsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemActionsWorkflowsItemRunsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemActionsWorkflowsItemRunsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemItemActionsWorkflowsItemRunsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemActionsWorkflowsItemRunsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["total_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalCount(val)
        }
        return nil
    }
    res["workflow_runs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateWorkflowRunFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WorkflowRunable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WorkflowRunable)
                }
            }
            m.SetWorkflowRuns(res)
        }
        return nil
    }
    return res
}
// GetTotalCount gets the total_count property value. The total_count property
// returns a *int32 when successful
func (m *ItemItemActionsWorkflowsItemRunsGetResponse) GetTotalCount()(*int32) {
    return m.total_count
}
// GetWorkflowRuns gets the workflow_runs property value. The workflow_runs property
// returns a []WorkflowRunable when successful
func (m *ItemItemActionsWorkflowsItemRunsGetResponse) GetWorkflowRuns()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WorkflowRunable) {
    return m.workflow_runs
}
// Serialize serializes information the current object
func (m *ItemItemActionsWorkflowsItemRunsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("total_count", m.GetTotalCount())
        if err != nil {
            return err
        }
    }
    if m.GetWorkflowRuns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWorkflowRuns()))
        for i, v := range m.GetWorkflowRuns() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("workflow_runs", cast)
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
func (m *ItemItemActionsWorkflowsItemRunsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetTotalCount sets the total_count property value. The total_count property
func (m *ItemItemActionsWorkflowsItemRunsGetResponse) SetTotalCount(value *int32)() {
    m.total_count = value
}
// SetWorkflowRuns sets the workflow_runs property value. The workflow_runs property
func (m *ItemItemActionsWorkflowsItemRunsGetResponse) SetWorkflowRuns(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WorkflowRunable)() {
    m.workflow_runs = value
}
type ItemItemActionsWorkflowsItemRunsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTotalCount()(*int32)
    GetWorkflowRuns()([]i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WorkflowRunable)
    SetTotalCount(value *int32)()
    SetWorkflowRuns(value []i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.WorkflowRunable)()
}
