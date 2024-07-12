package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GitTag metadata for a Git tag
type GitTag struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Message describing the purpose of the tag
    message *string
    // The node_id property
    node_id *string
    // The object property
    object GitTag_objectable
    // The sha property
    sha *string
    // Name of the tag
    tag *string
    // The tagger property
    tagger GitTag_taggerable
    // URL for the tag
    url *string
    // The verification property
    verification Verificationable
}
// NewGitTag instantiates a new GitTag and sets the default values.
func NewGitTag()(*GitTag) {
    m := &GitTag{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGitTagFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGitTagFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGitTag(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *GitTag) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GitTag) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
        }
        return nil
    }
    res["node_id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeId(val)
        }
        return nil
    }
    res["object"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGitTag_objectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObject(val.(GitTag_objectable))
        }
        return nil
    }
    res["sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha(val)
        }
        return nil
    }
    res["tag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTag(val)
        }
        return nil
    }
    res["tagger"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGitTag_taggerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTagger(val.(GitTag_taggerable))
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    res["verification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVerificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerification(val.(Verificationable))
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. Message describing the purpose of the tag
// returns a *string when successful
func (m *GitTag) GetMessage()(*string) {
    return m.message
}
// GetNodeId gets the node_id property value. The node_id property
// returns a *string when successful
func (m *GitTag) GetNodeId()(*string) {
    return m.node_id
}
// GetObject gets the object property value. The object property
// returns a GitTag_objectable when successful
func (m *GitTag) GetObject()(GitTag_objectable) {
    return m.object
}
// GetSha gets the sha property value. The sha property
// returns a *string when successful
func (m *GitTag) GetSha()(*string) {
    return m.sha
}
// GetTag gets the tag property value. Name of the tag
// returns a *string when successful
func (m *GitTag) GetTag()(*string) {
    return m.tag
}
// GetTagger gets the tagger property value. The tagger property
// returns a GitTag_taggerable when successful
func (m *GitTag) GetTagger()(GitTag_taggerable) {
    return m.tagger
}
// GetUrl gets the url property value. URL for the tag
// returns a *string when successful
func (m *GitTag) GetUrl()(*string) {
    return m.url
}
// GetVerification gets the verification property value. The verification property
// returns a Verificationable when successful
func (m *GitTag) GetVerification()(Verificationable) {
    return m.verification
}
// Serialize serializes information the current object
func (m *GitTag) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("node_id", m.GetNodeId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("object", m.GetObject())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha", m.GetSha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tag", m.GetTag())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("tagger", m.GetTagger())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("verification", m.GetVerification())
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
func (m *GitTag) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMessage sets the message property value. Message describing the purpose of the tag
func (m *GitTag) SetMessage(value *string)() {
    m.message = value
}
// SetNodeId sets the node_id property value. The node_id property
func (m *GitTag) SetNodeId(value *string)() {
    m.node_id = value
}
// SetObject sets the object property value. The object property
func (m *GitTag) SetObject(value GitTag_objectable)() {
    m.object = value
}
// SetSha sets the sha property value. The sha property
func (m *GitTag) SetSha(value *string)() {
    m.sha = value
}
// SetTag sets the tag property value. Name of the tag
func (m *GitTag) SetTag(value *string)() {
    m.tag = value
}
// SetTagger sets the tagger property value. The tagger property
func (m *GitTag) SetTagger(value GitTag_taggerable)() {
    m.tagger = value
}
// SetUrl sets the url property value. URL for the tag
func (m *GitTag) SetUrl(value *string)() {
    m.url = value
}
// SetVerification sets the verification property value. The verification property
func (m *GitTag) SetVerification(value Verificationable)() {
    m.verification = value
}
type GitTagable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMessage()(*string)
    GetNodeId()(*string)
    GetObject()(GitTag_objectable)
    GetSha()(*string)
    GetTag()(*string)
    GetTagger()(GitTag_taggerable)
    GetUrl()(*string)
    GetVerification()(Verificationable)
    SetMessage(value *string)()
    SetNodeId(value *string)()
    SetObject(value GitTag_objectable)()
    SetSha(value *string)()
    SetTag(value *string)()
    SetTagger(value GitTag_taggerable)()
    SetUrl(value *string)()
    SetVerification(value Verificationable)()
}
