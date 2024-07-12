package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApiOverview_domains struct {
    // The actions property
    actions []string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The codespaces property
    codespaces []string
    // The copilot property
    copilot []string
    // The packages property
    packages []string
    // The website property
    website []string
}
// NewApiOverview_domains instantiates a new ApiOverview_domains and sets the default values.
func NewApiOverview_domains()(*ApiOverview_domains) {
    m := &ApiOverview_domains{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApiOverview_domainsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApiOverview_domainsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApiOverview_domains(), nil
}
// GetActions gets the actions property value. The actions property
// returns a []string when successful
func (m *ApiOverview_domains) GetActions()([]string) {
    return m.actions
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ApiOverview_domains) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCodespaces gets the codespaces property value. The codespaces property
// returns a []string when successful
func (m *ApiOverview_domains) GetCodespaces()([]string) {
    return m.codespaces
}
// GetCopilot gets the copilot property value. The copilot property
// returns a []string when successful
func (m *ApiOverview_domains) GetCopilot()([]string) {
    return m.copilot
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApiOverview_domains) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetActions(res)
        }
        return nil
    }
    res["codespaces"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetCodespaces(res)
        }
        return nil
    }
    res["copilot"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetCopilot(res)
        }
        return nil
    }
    res["packages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetPackages(res)
        }
        return nil
    }
    res["website"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetWebsite(res)
        }
        return nil
    }
    return res
}
// GetPackages gets the packages property value. The packages property
// returns a []string when successful
func (m *ApiOverview_domains) GetPackages()([]string) {
    return m.packages
}
// GetWebsite gets the website property value. The website property
// returns a []string when successful
func (m *ApiOverview_domains) GetWebsite()([]string) {
    return m.website
}
// Serialize serializes information the current object
func (m *ApiOverview_domains) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActions() != nil {
        err := writer.WriteCollectionOfStringValues("actions", m.GetActions())
        if err != nil {
            return err
        }
    }
    if m.GetCodespaces() != nil {
        err := writer.WriteCollectionOfStringValues("codespaces", m.GetCodespaces())
        if err != nil {
            return err
        }
    }
    if m.GetCopilot() != nil {
        err := writer.WriteCollectionOfStringValues("copilot", m.GetCopilot())
        if err != nil {
            return err
        }
    }
    if m.GetPackages() != nil {
        err := writer.WriteCollectionOfStringValues("packages", m.GetPackages())
        if err != nil {
            return err
        }
    }
    if m.GetWebsite() != nil {
        err := writer.WriteCollectionOfStringValues("website", m.GetWebsite())
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
// SetActions sets the actions property value. The actions property
func (m *ApiOverview_domains) SetActions(value []string)() {
    m.actions = value
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApiOverview_domains) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCodespaces sets the codespaces property value. The codespaces property
func (m *ApiOverview_domains) SetCodespaces(value []string)() {
    m.codespaces = value
}
// SetCopilot sets the copilot property value. The copilot property
func (m *ApiOverview_domains) SetCopilot(value []string)() {
    m.copilot = value
}
// SetPackages sets the packages property value. The packages property
func (m *ApiOverview_domains) SetPackages(value []string)() {
    m.packages = value
}
// SetWebsite sets the website property value. The website property
func (m *ApiOverview_domains) SetWebsite(value []string)() {
    m.website = value
}
type ApiOverview_domainsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()([]string)
    GetCodespaces()([]string)
    GetCopilot()([]string)
    GetPackages()([]string)
    GetWebsite()([]string)
    SetActions(value []string)()
    SetCodespaces(value []string)()
    SetCopilot(value []string)()
    SetPackages(value []string)()
    SetWebsite(value []string)()
}
