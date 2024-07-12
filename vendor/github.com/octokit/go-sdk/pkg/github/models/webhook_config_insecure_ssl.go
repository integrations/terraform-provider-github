package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebhookConfigInsecureSsl composed type wrapper for classes float64, string
type WebhookConfigInsecureSsl struct {
    // Composed type representation for type float64
    double *float64
    // Composed type representation for type string
    string *string
    // Composed type representation for type float64
    webhookConfigInsecureSslDouble *float64
    // Composed type representation for type float64
    webhookConfigInsecureSslDouble0 *float64
    // Composed type representation for type float64
    webhookConfigInsecureSslDouble1 *float64
    // Composed type representation for type float64
    webhookConfigInsecureSslDouble2 *float64
    // Composed type representation for type float64
    webhookConfigInsecureSslDouble3 *float64
    // Composed type representation for type float64
    webhookConfigInsecureSslDouble4 *float64
    // Composed type representation for type string
    webhookConfigInsecureSslString *string
    // Composed type representation for type string
    webhookConfigInsecureSslString0 *string
    // Composed type representation for type string
    webhookConfigInsecureSslString1 *string
    // Composed type representation for type string
    webhookConfigInsecureSslString2 *string
    // Composed type representation for type string
    webhookConfigInsecureSslString3 *string
    // Composed type representation for type string
    webhookConfigInsecureSslString4 *string
}
// NewWebhookConfigInsecureSsl instantiates a new WebhookConfigInsecureSsl and sets the default values.
func NewWebhookConfigInsecureSsl()(*WebhookConfigInsecureSsl) {
    m := &WebhookConfigInsecureSsl{
    }
    return m
}
// CreateWebhookConfigInsecureSslFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWebhookConfigInsecureSslFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewWebhookConfigInsecureSsl()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetWebhookConfigInsecureSslDouble(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetWebhookConfigInsecureSslDouble0(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetWebhookConfigInsecureSslDouble1(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetWebhookConfigInsecureSslDouble2(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetWebhookConfigInsecureSslDouble3(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetWebhookConfigInsecureSslDouble4(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetWebhookConfigInsecureSslString(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetWebhookConfigInsecureSslString0(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetWebhookConfigInsecureSslString1(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetWebhookConfigInsecureSslString2(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetWebhookConfigInsecureSslString3(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetWebhookConfigInsecureSslString4(val)
    }
    return result, nil
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *WebhookConfigInsecureSsl) GetDouble()(*float64) {
    return m.double
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WebhookConfigInsecureSsl) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *WebhookConfigInsecureSsl) GetIsComposedType()(bool) {
    return true
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *WebhookConfigInsecureSsl) GetString()(*string) {
    return m.string
}
// GetWebhookConfigInsecureSslDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *WebhookConfigInsecureSsl) GetWebhookConfigInsecureSslDouble()(*float64) {
    return m.webhookConfigInsecureSslDouble
}
// GetWebhookConfigInsecureSslDouble0 gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *WebhookConfigInsecureSsl) GetWebhookConfigInsecureSslDouble0()(*float64) {
    return m.webhookConfigInsecureSslDouble0
}
// GetWebhookConfigInsecureSslDouble1 gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *WebhookConfigInsecureSsl) GetWebhookConfigInsecureSslDouble1()(*float64) {
    return m.webhookConfigInsecureSslDouble1
}
// GetWebhookConfigInsecureSslDouble2 gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *WebhookConfigInsecureSsl) GetWebhookConfigInsecureSslDouble2()(*float64) {
    return m.webhookConfigInsecureSslDouble2
}
// GetWebhookConfigInsecureSslDouble3 gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *WebhookConfigInsecureSsl) GetWebhookConfigInsecureSslDouble3()(*float64) {
    return m.webhookConfigInsecureSslDouble3
}
// GetWebhookConfigInsecureSslDouble4 gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *WebhookConfigInsecureSsl) GetWebhookConfigInsecureSslDouble4()(*float64) {
    return m.webhookConfigInsecureSslDouble4
}
// GetWebhookConfigInsecureSslString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *WebhookConfigInsecureSsl) GetWebhookConfigInsecureSslString()(*string) {
    return m.webhookConfigInsecureSslString
}
// GetWebhookConfigInsecureSslString0 gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *WebhookConfigInsecureSsl) GetWebhookConfigInsecureSslString0()(*string) {
    return m.webhookConfigInsecureSslString0
}
// GetWebhookConfigInsecureSslString1 gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *WebhookConfigInsecureSsl) GetWebhookConfigInsecureSslString1()(*string) {
    return m.webhookConfigInsecureSslString1
}
// GetWebhookConfigInsecureSslString2 gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *WebhookConfigInsecureSsl) GetWebhookConfigInsecureSslString2()(*string) {
    return m.webhookConfigInsecureSslString2
}
// GetWebhookConfigInsecureSslString3 gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *WebhookConfigInsecureSsl) GetWebhookConfigInsecureSslString3()(*string) {
    return m.webhookConfigInsecureSslString3
}
// GetWebhookConfigInsecureSslString4 gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *WebhookConfigInsecureSsl) GetWebhookConfigInsecureSslString4()(*string) {
    return m.webhookConfigInsecureSslString4
}
// Serialize serializes information the current object
func (m *WebhookConfigInsecureSsl) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    } else if m.GetWebhookConfigInsecureSslDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetWebhookConfigInsecureSslDouble())
        if err != nil {
            return err
        }
    } else if m.GetWebhookConfigInsecureSslDouble0() != nil {
        err := writer.WriteFloat64Value("", m.GetWebhookConfigInsecureSslDouble0())
        if err != nil {
            return err
        }
    } else if m.GetWebhookConfigInsecureSslDouble1() != nil {
        err := writer.WriteFloat64Value("", m.GetWebhookConfigInsecureSslDouble1())
        if err != nil {
            return err
        }
    } else if m.GetWebhookConfigInsecureSslDouble2() != nil {
        err := writer.WriteFloat64Value("", m.GetWebhookConfigInsecureSslDouble2())
        if err != nil {
            return err
        }
    } else if m.GetWebhookConfigInsecureSslDouble3() != nil {
        err := writer.WriteFloat64Value("", m.GetWebhookConfigInsecureSslDouble3())
        if err != nil {
            return err
        }
    } else if m.GetWebhookConfigInsecureSslDouble4() != nil {
        err := writer.WriteFloat64Value("", m.GetWebhookConfigInsecureSslDouble4())
        if err != nil {
            return err
        }
    } else if m.GetWebhookConfigInsecureSslString() != nil {
        err := writer.WriteStringValue("", m.GetWebhookConfigInsecureSslString())
        if err != nil {
            return err
        }
    } else if m.GetWebhookConfigInsecureSslString0() != nil {
        err := writer.WriteStringValue("", m.GetWebhookConfigInsecureSslString0())
        if err != nil {
            return err
        }
    } else if m.GetWebhookConfigInsecureSslString1() != nil {
        err := writer.WriteStringValue("", m.GetWebhookConfigInsecureSslString1())
        if err != nil {
            return err
        }
    } else if m.GetWebhookConfigInsecureSslString2() != nil {
        err := writer.WriteStringValue("", m.GetWebhookConfigInsecureSslString2())
        if err != nil {
            return err
        }
    } else if m.GetWebhookConfigInsecureSslString3() != nil {
        err := writer.WriteStringValue("", m.GetWebhookConfigInsecureSslString3())
        if err != nil {
            return err
        }
    } else if m.GetWebhookConfigInsecureSslString4() != nil {
        err := writer.WriteStringValue("", m.GetWebhookConfigInsecureSslString4())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *WebhookConfigInsecureSsl) SetDouble(value *float64)() {
    m.double = value
}
// SetString sets the string property value. Composed type representation for type string
func (m *WebhookConfigInsecureSsl) SetString(value *string)() {
    m.string = value
}
// SetWebhookConfigInsecureSslDouble sets the double property value. Composed type representation for type float64
func (m *WebhookConfigInsecureSsl) SetWebhookConfigInsecureSslDouble(value *float64)() {
    m.webhookConfigInsecureSslDouble = value
}
// SetWebhookConfigInsecureSslDouble0 sets the double property value. Composed type representation for type float64
func (m *WebhookConfigInsecureSsl) SetWebhookConfigInsecureSslDouble0(value *float64)() {
    m.webhookConfigInsecureSslDouble0 = value
}
// SetWebhookConfigInsecureSslDouble1 sets the double property value. Composed type representation for type float64
func (m *WebhookConfigInsecureSsl) SetWebhookConfigInsecureSslDouble1(value *float64)() {
    m.webhookConfigInsecureSslDouble1 = value
}
// SetWebhookConfigInsecureSslDouble2 sets the double property value. Composed type representation for type float64
func (m *WebhookConfigInsecureSsl) SetWebhookConfigInsecureSslDouble2(value *float64)() {
    m.webhookConfigInsecureSslDouble2 = value
}
// SetWebhookConfigInsecureSslDouble3 sets the double property value. Composed type representation for type float64
func (m *WebhookConfigInsecureSsl) SetWebhookConfigInsecureSslDouble3(value *float64)() {
    m.webhookConfigInsecureSslDouble3 = value
}
// SetWebhookConfigInsecureSslDouble4 sets the double property value. Composed type representation for type float64
func (m *WebhookConfigInsecureSsl) SetWebhookConfigInsecureSslDouble4(value *float64)() {
    m.webhookConfigInsecureSslDouble4 = value
}
// SetWebhookConfigInsecureSslString sets the string property value. Composed type representation for type string
func (m *WebhookConfigInsecureSsl) SetWebhookConfigInsecureSslString(value *string)() {
    m.webhookConfigInsecureSslString = value
}
// SetWebhookConfigInsecureSslString0 sets the string property value. Composed type representation for type string
func (m *WebhookConfigInsecureSsl) SetWebhookConfigInsecureSslString0(value *string)() {
    m.webhookConfigInsecureSslString0 = value
}
// SetWebhookConfigInsecureSslString1 sets the string property value. Composed type representation for type string
func (m *WebhookConfigInsecureSsl) SetWebhookConfigInsecureSslString1(value *string)() {
    m.webhookConfigInsecureSslString1 = value
}
// SetWebhookConfigInsecureSslString2 sets the string property value. Composed type representation for type string
func (m *WebhookConfigInsecureSsl) SetWebhookConfigInsecureSslString2(value *string)() {
    m.webhookConfigInsecureSslString2 = value
}
// SetWebhookConfigInsecureSslString3 sets the string property value. Composed type representation for type string
func (m *WebhookConfigInsecureSsl) SetWebhookConfigInsecureSslString3(value *string)() {
    m.webhookConfigInsecureSslString3 = value
}
// SetWebhookConfigInsecureSslString4 sets the string property value. Composed type representation for type string
func (m *WebhookConfigInsecureSsl) SetWebhookConfigInsecureSslString4(value *string)() {
    m.webhookConfigInsecureSslString4 = value
}
type WebhookConfigInsecureSslable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDouble()(*float64)
    GetString()(*string)
    GetWebhookConfigInsecureSslDouble()(*float64)
    GetWebhookConfigInsecureSslDouble0()(*float64)
    GetWebhookConfigInsecureSslDouble1()(*float64)
    GetWebhookConfigInsecureSslDouble2()(*float64)
    GetWebhookConfigInsecureSslDouble3()(*float64)
    GetWebhookConfigInsecureSslDouble4()(*float64)
    GetWebhookConfigInsecureSslString()(*string)
    GetWebhookConfigInsecureSslString0()(*string)
    GetWebhookConfigInsecureSslString1()(*string)
    GetWebhookConfigInsecureSslString2()(*string)
    GetWebhookConfigInsecureSslString3()(*string)
    GetWebhookConfigInsecureSslString4()(*string)
    SetDouble(value *float64)()
    SetString(value *string)()
    SetWebhookConfigInsecureSslDouble(value *float64)()
    SetWebhookConfigInsecureSslDouble0(value *float64)()
    SetWebhookConfigInsecureSslDouble1(value *float64)()
    SetWebhookConfigInsecureSslDouble2(value *float64)()
    SetWebhookConfigInsecureSslDouble3(value *float64)()
    SetWebhookConfigInsecureSslDouble4(value *float64)()
    SetWebhookConfigInsecureSslString(value *string)()
    SetWebhookConfigInsecureSslString0(value *string)()
    SetWebhookConfigInsecureSslString1(value *string)()
    SetWebhookConfigInsecureSslString2(value *string)()
    SetWebhookConfigInsecureSslString3(value *string)()
    SetWebhookConfigInsecureSslString4(value *string)()
}
