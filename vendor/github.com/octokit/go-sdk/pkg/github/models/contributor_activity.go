package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContributorActivity contributor Activity
type ContributorActivity struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A GitHub user.
    author NullableSimpleUserable
    // The total property
    total *int32
    // The weeks property
    weeks []ContributorActivity_weeksable
}
// NewContributorActivity instantiates a new ContributorActivity and sets the default values.
func NewContributorActivity()(*ContributorActivity) {
    m := &ContributorActivity{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateContributorActivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContributorActivityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContributorActivity(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContributorActivity) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthor gets the author property value. A GitHub user.
// returns a NullableSimpleUserable when successful
func (m *ContributorActivity) GetAuthor()(NullableSimpleUserable) {
    return m.author
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContributorActivity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["author"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNullableSimpleUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthor(val.(NullableSimpleUserable))
        }
        return nil
    }
    res["total"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
        }
        return nil
    }
    res["weeks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateContributorActivity_weeksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ContributorActivity_weeksable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ContributorActivity_weeksable)
                }
            }
            m.SetWeeks(res)
        }
        return nil
    }
    return res
}
// GetTotal gets the total property value. The total property
// returns a *int32 when successful
func (m *ContributorActivity) GetTotal()(*int32) {
    return m.total
}
// GetWeeks gets the weeks property value. The weeks property
// returns a []ContributorActivity_weeksable when successful
func (m *ContributorActivity) GetWeeks()([]ContributorActivity_weeksable) {
    return m.weeks
}
// Serialize serializes information the current object
func (m *ContributorActivity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("author", m.GetAuthor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("total", m.GetTotal())
        if err != nil {
            return err
        }
    }
    if m.GetWeeks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWeeks()))
        for i, v := range m.GetWeeks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("weeks", cast)
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
func (m *ContributorActivity) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthor sets the author property value. A GitHub user.
func (m *ContributorActivity) SetAuthor(value NullableSimpleUserable)() {
    m.author = value
}
// SetTotal sets the total property value. The total property
func (m *ContributorActivity) SetTotal(value *int32)() {
    m.total = value
}
// SetWeeks sets the weeks property value. The weeks property
func (m *ContributorActivity) SetWeeks(value []ContributorActivity_weeksable)() {
    m.weeks = value
}
type ContributorActivityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthor()(NullableSimpleUserable)
    GetTotal()(*int32)
    GetWeeks()([]ContributorActivity_weeksable)
    SetAuthor(value NullableSimpleUserable)()
    SetTotal(value *int32)()
    SetWeeks(value []ContributorActivity_weeksable)()
}
