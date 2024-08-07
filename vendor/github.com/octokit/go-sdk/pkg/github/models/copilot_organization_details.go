package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CopilotOrganizationDetails information about the seat breakdown and policies set for an organization with a Copilot Business subscription.
type CopilotOrganizationDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The organization policy for allowing or disallowing organization members to use Copilot within their CLI.
    cli *CopilotOrganizationDetails_cli
    // The organization policy for allowing or disallowing organization members to use Copilot Chat within their editor.
    ide_chat *CopilotOrganizationDetails_ide_chat
    // The organization policy for allowing or disallowing organization members to use Copilot features within github.com.
    platform_chat *CopilotOrganizationDetails_platform_chat
    // The organization policy for allowing or disallowing Copilot to make suggestions that match public code.
    public_code_suggestions *CopilotOrganizationDetails_public_code_suggestions
    // The breakdown of Copilot Business seats for the organization.
    seat_breakdown CopilotSeatBreakdownable
    // The mode of assigning new seats.
    seat_management_setting *CopilotOrganizationDetails_seat_management_setting
}
// NewCopilotOrganizationDetails instantiates a new CopilotOrganizationDetails and sets the default values.
func NewCopilotOrganizationDetails()(*CopilotOrganizationDetails) {
    m := &CopilotOrganizationDetails{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCopilotOrganizationDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCopilotOrganizationDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopilotOrganizationDetails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CopilotOrganizationDetails) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCli gets the cli property value. The organization policy for allowing or disallowing organization members to use Copilot within their CLI.
// returns a *CopilotOrganizationDetails_cli when successful
func (m *CopilotOrganizationDetails) GetCli()(*CopilotOrganizationDetails_cli) {
    return m.cli
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CopilotOrganizationDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cli"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCopilotOrganizationDetails_cli)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCli(val.(*CopilotOrganizationDetails_cli))
        }
        return nil
    }
    res["ide_chat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCopilotOrganizationDetails_ide_chat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdeChat(val.(*CopilotOrganizationDetails_ide_chat))
        }
        return nil
    }
    res["platform_chat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCopilotOrganizationDetails_platform_chat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformChat(val.(*CopilotOrganizationDetails_platform_chat))
        }
        return nil
    }
    res["public_code_suggestions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCopilotOrganizationDetails_public_code_suggestions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicCodeSuggestions(val.(*CopilotOrganizationDetails_public_code_suggestions))
        }
        return nil
    }
    res["seat_breakdown"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCopilotSeatBreakdownFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeatBreakdown(val.(CopilotSeatBreakdownable))
        }
        return nil
    }
    res["seat_management_setting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCopilotOrganizationDetails_seat_management_setting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeatManagementSetting(val.(*CopilotOrganizationDetails_seat_management_setting))
        }
        return nil
    }
    return res
}
// GetIdeChat gets the ide_chat property value. The organization policy for allowing or disallowing organization members to use Copilot Chat within their editor.
// returns a *CopilotOrganizationDetails_ide_chat when successful
func (m *CopilotOrganizationDetails) GetIdeChat()(*CopilotOrganizationDetails_ide_chat) {
    return m.ide_chat
}
// GetPlatformChat gets the platform_chat property value. The organization policy for allowing or disallowing organization members to use Copilot features within github.com.
// returns a *CopilotOrganizationDetails_platform_chat when successful
func (m *CopilotOrganizationDetails) GetPlatformChat()(*CopilotOrganizationDetails_platform_chat) {
    return m.platform_chat
}
// GetPublicCodeSuggestions gets the public_code_suggestions property value. The organization policy for allowing or disallowing Copilot to make suggestions that match public code.
// returns a *CopilotOrganizationDetails_public_code_suggestions when successful
func (m *CopilotOrganizationDetails) GetPublicCodeSuggestions()(*CopilotOrganizationDetails_public_code_suggestions) {
    return m.public_code_suggestions
}
// GetSeatBreakdown gets the seat_breakdown property value. The breakdown of Copilot Business seats for the organization.
// returns a CopilotSeatBreakdownable when successful
func (m *CopilotOrganizationDetails) GetSeatBreakdown()(CopilotSeatBreakdownable) {
    return m.seat_breakdown
}
// GetSeatManagementSetting gets the seat_management_setting property value. The mode of assigning new seats.
// returns a *CopilotOrganizationDetails_seat_management_setting when successful
func (m *CopilotOrganizationDetails) GetSeatManagementSetting()(*CopilotOrganizationDetails_seat_management_setting) {
    return m.seat_management_setting
}
// Serialize serializes information the current object
func (m *CopilotOrganizationDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCli() != nil {
        cast := (*m.GetCli()).String()
        err := writer.WriteStringValue("cli", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIdeChat() != nil {
        cast := (*m.GetIdeChat()).String()
        err := writer.WriteStringValue("ide_chat", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlatformChat() != nil {
        cast := (*m.GetPlatformChat()).String()
        err := writer.WriteStringValue("platform_chat", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPublicCodeSuggestions() != nil {
        cast := (*m.GetPublicCodeSuggestions()).String()
        err := writer.WriteStringValue("public_code_suggestions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("seat_breakdown", m.GetSeatBreakdown())
        if err != nil {
            return err
        }
    }
    if m.GetSeatManagementSetting() != nil {
        cast := (*m.GetSeatManagementSetting()).String()
        err := writer.WriteStringValue("seat_management_setting", &cast)
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
func (m *CopilotOrganizationDetails) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCli sets the cli property value. The organization policy for allowing or disallowing organization members to use Copilot within their CLI.
func (m *CopilotOrganizationDetails) SetCli(value *CopilotOrganizationDetails_cli)() {
    m.cli = value
}
// SetIdeChat sets the ide_chat property value. The organization policy for allowing or disallowing organization members to use Copilot Chat within their editor.
func (m *CopilotOrganizationDetails) SetIdeChat(value *CopilotOrganizationDetails_ide_chat)() {
    m.ide_chat = value
}
// SetPlatformChat sets the platform_chat property value. The organization policy for allowing or disallowing organization members to use Copilot features within github.com.
func (m *CopilotOrganizationDetails) SetPlatformChat(value *CopilotOrganizationDetails_platform_chat)() {
    m.platform_chat = value
}
// SetPublicCodeSuggestions sets the public_code_suggestions property value. The organization policy for allowing or disallowing Copilot to make suggestions that match public code.
func (m *CopilotOrganizationDetails) SetPublicCodeSuggestions(value *CopilotOrganizationDetails_public_code_suggestions)() {
    m.public_code_suggestions = value
}
// SetSeatBreakdown sets the seat_breakdown property value. The breakdown of Copilot Business seats for the organization.
func (m *CopilotOrganizationDetails) SetSeatBreakdown(value CopilotSeatBreakdownable)() {
    m.seat_breakdown = value
}
// SetSeatManagementSetting sets the seat_management_setting property value. The mode of assigning new seats.
func (m *CopilotOrganizationDetails) SetSeatManagementSetting(value *CopilotOrganizationDetails_seat_management_setting)() {
    m.seat_management_setting = value
}
type CopilotOrganizationDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCli()(*CopilotOrganizationDetails_cli)
    GetIdeChat()(*CopilotOrganizationDetails_ide_chat)
    GetPlatformChat()(*CopilotOrganizationDetails_platform_chat)
    GetPublicCodeSuggestions()(*CopilotOrganizationDetails_public_code_suggestions)
    GetSeatBreakdown()(CopilotSeatBreakdownable)
    GetSeatManagementSetting()(*CopilotOrganizationDetails_seat_management_setting)
    SetCli(value *CopilotOrganizationDetails_cli)()
    SetIdeChat(value *CopilotOrganizationDetails_ide_chat)()
    SetPlatformChat(value *CopilotOrganizationDetails_platform_chat)()
    SetPublicCodeSuggestions(value *CopilotOrganizationDetails_public_code_suggestions)()
    SetSeatBreakdown(value CopilotSeatBreakdownable)()
    SetSeatManagementSetting(value *CopilotOrganizationDetails_seat_management_setting)()
}
