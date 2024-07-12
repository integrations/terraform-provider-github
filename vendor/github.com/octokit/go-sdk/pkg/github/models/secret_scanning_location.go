package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SecretScanningLocation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The details property
    details SecretScanningLocation_SecretScanningLocation_detailsable
    // The location type. Because secrets may be found in different types of resources (ie. code, comments, issues, pull requests, discussions), this field identifies the type of resource where the secret was found.
    typeEscaped *SecretScanningLocation_type
}
// SecretScanningLocation_SecretScanningLocation_details composed type wrapper for classes SecretScanningLocationCommitable, SecretScanningLocationDiscussionBodyable, SecretScanningLocationDiscussionCommentable, SecretScanningLocationDiscussionTitleable, SecretScanningLocationIssueBodyable, SecretScanningLocationIssueCommentable, SecretScanningLocationIssueTitleable, SecretScanningLocationPullRequestBodyable, SecretScanningLocationPullRequestCommentable, SecretScanningLocationPullRequestReviewable, SecretScanningLocationPullRequestReviewCommentable, SecretScanningLocationPullRequestTitleable, SecretScanningLocationWikiCommitable
type SecretScanningLocation_SecretScanningLocation_details struct {
    // Composed type representation for type SecretScanningLocationCommitable
    secretScanningLocationCommit SecretScanningLocationCommitable
    // Composed type representation for type SecretScanningLocationDiscussionBodyable
    secretScanningLocationDiscussionBody SecretScanningLocationDiscussionBodyable
    // Composed type representation for type SecretScanningLocationDiscussionCommentable
    secretScanningLocationDiscussionComment SecretScanningLocationDiscussionCommentable
    // Composed type representation for type SecretScanningLocationDiscussionTitleable
    secretScanningLocationDiscussionTitle SecretScanningLocationDiscussionTitleable
    // Composed type representation for type SecretScanningLocationIssueBodyable
    secretScanningLocationIssueBody SecretScanningLocationIssueBodyable
    // Composed type representation for type SecretScanningLocationIssueCommentable
    secretScanningLocationIssueComment SecretScanningLocationIssueCommentable
    // Composed type representation for type SecretScanningLocationIssueTitleable
    secretScanningLocationIssueTitle SecretScanningLocationIssueTitleable
    // Composed type representation for type SecretScanningLocationPullRequestBodyable
    secretScanningLocationPullRequestBody SecretScanningLocationPullRequestBodyable
    // Composed type representation for type SecretScanningLocationPullRequestCommentable
    secretScanningLocationPullRequestComment SecretScanningLocationPullRequestCommentable
    // Composed type representation for type SecretScanningLocationPullRequestReviewable
    secretScanningLocationPullRequestReview SecretScanningLocationPullRequestReviewable
    // Composed type representation for type SecretScanningLocationPullRequestReviewCommentable
    secretScanningLocationPullRequestReviewComment SecretScanningLocationPullRequestReviewCommentable
    // Composed type representation for type SecretScanningLocationPullRequestTitleable
    secretScanningLocationPullRequestTitle SecretScanningLocationPullRequestTitleable
    // Composed type representation for type SecretScanningLocationWikiCommitable
    secretScanningLocationWikiCommit SecretScanningLocationWikiCommitable
}
// NewSecretScanningLocation_SecretScanningLocation_details instantiates a new SecretScanningLocation_SecretScanningLocation_details and sets the default values.
func NewSecretScanningLocation_SecretScanningLocation_details()(*SecretScanningLocation_SecretScanningLocation_details) {
    m := &SecretScanningLocation_SecretScanningLocation_details{
    }
    return m
}
// CreateSecretScanningLocation_SecretScanningLocation_detailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningLocation_SecretScanningLocation_detailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewSecretScanningLocation_SecretScanningLocation_details()
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
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningLocation_SecretScanningLocation_details) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *SecretScanningLocation_SecretScanningLocation_details) GetIsComposedType()(bool) {
    return true
}
// GetSecretScanningLocationCommit gets the secretScanningLocationCommit property value. Composed type representation for type SecretScanningLocationCommitable
// returns a SecretScanningLocationCommitable when successful
func (m *SecretScanningLocation_SecretScanningLocation_details) GetSecretScanningLocationCommit()(SecretScanningLocationCommitable) {
    return m.secretScanningLocationCommit
}
// GetSecretScanningLocationDiscussionBody gets the secretScanningLocationDiscussionBody property value. Composed type representation for type SecretScanningLocationDiscussionBodyable
// returns a SecretScanningLocationDiscussionBodyable when successful
func (m *SecretScanningLocation_SecretScanningLocation_details) GetSecretScanningLocationDiscussionBody()(SecretScanningLocationDiscussionBodyable) {
    return m.secretScanningLocationDiscussionBody
}
// GetSecretScanningLocationDiscussionComment gets the secretScanningLocationDiscussionComment property value. Composed type representation for type SecretScanningLocationDiscussionCommentable
// returns a SecretScanningLocationDiscussionCommentable when successful
func (m *SecretScanningLocation_SecretScanningLocation_details) GetSecretScanningLocationDiscussionComment()(SecretScanningLocationDiscussionCommentable) {
    return m.secretScanningLocationDiscussionComment
}
// GetSecretScanningLocationDiscussionTitle gets the secretScanningLocationDiscussionTitle property value. Composed type representation for type SecretScanningLocationDiscussionTitleable
// returns a SecretScanningLocationDiscussionTitleable when successful
func (m *SecretScanningLocation_SecretScanningLocation_details) GetSecretScanningLocationDiscussionTitle()(SecretScanningLocationDiscussionTitleable) {
    return m.secretScanningLocationDiscussionTitle
}
// GetSecretScanningLocationIssueBody gets the secretScanningLocationIssueBody property value. Composed type representation for type SecretScanningLocationIssueBodyable
// returns a SecretScanningLocationIssueBodyable when successful
func (m *SecretScanningLocation_SecretScanningLocation_details) GetSecretScanningLocationIssueBody()(SecretScanningLocationIssueBodyable) {
    return m.secretScanningLocationIssueBody
}
// GetSecretScanningLocationIssueComment gets the secretScanningLocationIssueComment property value. Composed type representation for type SecretScanningLocationIssueCommentable
// returns a SecretScanningLocationIssueCommentable when successful
func (m *SecretScanningLocation_SecretScanningLocation_details) GetSecretScanningLocationIssueComment()(SecretScanningLocationIssueCommentable) {
    return m.secretScanningLocationIssueComment
}
// GetSecretScanningLocationIssueTitle gets the secretScanningLocationIssueTitle property value. Composed type representation for type SecretScanningLocationIssueTitleable
// returns a SecretScanningLocationIssueTitleable when successful
func (m *SecretScanningLocation_SecretScanningLocation_details) GetSecretScanningLocationIssueTitle()(SecretScanningLocationIssueTitleable) {
    return m.secretScanningLocationIssueTitle
}
// GetSecretScanningLocationPullRequestBody gets the secretScanningLocationPullRequestBody property value. Composed type representation for type SecretScanningLocationPullRequestBodyable
// returns a SecretScanningLocationPullRequestBodyable when successful
func (m *SecretScanningLocation_SecretScanningLocation_details) GetSecretScanningLocationPullRequestBody()(SecretScanningLocationPullRequestBodyable) {
    return m.secretScanningLocationPullRequestBody
}
// GetSecretScanningLocationPullRequestComment gets the secretScanningLocationPullRequestComment property value. Composed type representation for type SecretScanningLocationPullRequestCommentable
// returns a SecretScanningLocationPullRequestCommentable when successful
func (m *SecretScanningLocation_SecretScanningLocation_details) GetSecretScanningLocationPullRequestComment()(SecretScanningLocationPullRequestCommentable) {
    return m.secretScanningLocationPullRequestComment
}
// GetSecretScanningLocationPullRequestReview gets the secretScanningLocationPullRequestReview property value. Composed type representation for type SecretScanningLocationPullRequestReviewable
// returns a SecretScanningLocationPullRequestReviewable when successful
func (m *SecretScanningLocation_SecretScanningLocation_details) GetSecretScanningLocationPullRequestReview()(SecretScanningLocationPullRequestReviewable) {
    return m.secretScanningLocationPullRequestReview
}
// GetSecretScanningLocationPullRequestReviewComment gets the secretScanningLocationPullRequestReviewComment property value. Composed type representation for type SecretScanningLocationPullRequestReviewCommentable
// returns a SecretScanningLocationPullRequestReviewCommentable when successful
func (m *SecretScanningLocation_SecretScanningLocation_details) GetSecretScanningLocationPullRequestReviewComment()(SecretScanningLocationPullRequestReviewCommentable) {
    return m.secretScanningLocationPullRequestReviewComment
}
// GetSecretScanningLocationPullRequestTitle gets the secretScanningLocationPullRequestTitle property value. Composed type representation for type SecretScanningLocationPullRequestTitleable
// returns a SecretScanningLocationPullRequestTitleable when successful
func (m *SecretScanningLocation_SecretScanningLocation_details) GetSecretScanningLocationPullRequestTitle()(SecretScanningLocationPullRequestTitleable) {
    return m.secretScanningLocationPullRequestTitle
}
// GetSecretScanningLocationWikiCommit gets the secretScanningLocationWikiCommit property value. Composed type representation for type SecretScanningLocationWikiCommitable
// returns a SecretScanningLocationWikiCommitable when successful
func (m *SecretScanningLocation_SecretScanningLocation_details) GetSecretScanningLocationWikiCommit()(SecretScanningLocationWikiCommitable) {
    return m.secretScanningLocationWikiCommit
}
// Serialize serializes information the current object
func (m *SecretScanningLocation_SecretScanningLocation_details) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetSecretScanningLocationCommit() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationCommit())
        if err != nil {
            return err
        }
    } else if m.GetSecretScanningLocationDiscussionBody() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationDiscussionBody())
        if err != nil {
            return err
        }
    } else if m.GetSecretScanningLocationDiscussionComment() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationDiscussionComment())
        if err != nil {
            return err
        }
    } else if m.GetSecretScanningLocationDiscussionTitle() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationDiscussionTitle())
        if err != nil {
            return err
        }
    } else if m.GetSecretScanningLocationIssueBody() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationIssueBody())
        if err != nil {
            return err
        }
    } else if m.GetSecretScanningLocationIssueComment() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationIssueComment())
        if err != nil {
            return err
        }
    } else if m.GetSecretScanningLocationIssueTitle() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationIssueTitle())
        if err != nil {
            return err
        }
    } else if m.GetSecretScanningLocationPullRequestBody() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationPullRequestBody())
        if err != nil {
            return err
        }
    } else if m.GetSecretScanningLocationPullRequestComment() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationPullRequestComment())
        if err != nil {
            return err
        }
    } else if m.GetSecretScanningLocationPullRequestReview() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationPullRequestReview())
        if err != nil {
            return err
        }
    } else if m.GetSecretScanningLocationPullRequestReviewComment() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationPullRequestReviewComment())
        if err != nil {
            return err
        }
    } else if m.GetSecretScanningLocationPullRequestTitle() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationPullRequestTitle())
        if err != nil {
            return err
        }
    } else if m.GetSecretScanningLocationWikiCommit() != nil {
        err := writer.WriteObjectValue("", m.GetSecretScanningLocationWikiCommit())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSecretScanningLocationCommit sets the secretScanningLocationCommit property value. Composed type representation for type SecretScanningLocationCommitable
func (m *SecretScanningLocation_SecretScanningLocation_details) SetSecretScanningLocationCommit(value SecretScanningLocationCommitable)() {
    m.secretScanningLocationCommit = value
}
// SetSecretScanningLocationDiscussionBody sets the secretScanningLocationDiscussionBody property value. Composed type representation for type SecretScanningLocationDiscussionBodyable
func (m *SecretScanningLocation_SecretScanningLocation_details) SetSecretScanningLocationDiscussionBody(value SecretScanningLocationDiscussionBodyable)() {
    m.secretScanningLocationDiscussionBody = value
}
// SetSecretScanningLocationDiscussionComment sets the secretScanningLocationDiscussionComment property value. Composed type representation for type SecretScanningLocationDiscussionCommentable
func (m *SecretScanningLocation_SecretScanningLocation_details) SetSecretScanningLocationDiscussionComment(value SecretScanningLocationDiscussionCommentable)() {
    m.secretScanningLocationDiscussionComment = value
}
// SetSecretScanningLocationDiscussionTitle sets the secretScanningLocationDiscussionTitle property value. Composed type representation for type SecretScanningLocationDiscussionTitleable
func (m *SecretScanningLocation_SecretScanningLocation_details) SetSecretScanningLocationDiscussionTitle(value SecretScanningLocationDiscussionTitleable)() {
    m.secretScanningLocationDiscussionTitle = value
}
// SetSecretScanningLocationIssueBody sets the secretScanningLocationIssueBody property value. Composed type representation for type SecretScanningLocationIssueBodyable
func (m *SecretScanningLocation_SecretScanningLocation_details) SetSecretScanningLocationIssueBody(value SecretScanningLocationIssueBodyable)() {
    m.secretScanningLocationIssueBody = value
}
// SetSecretScanningLocationIssueComment sets the secretScanningLocationIssueComment property value. Composed type representation for type SecretScanningLocationIssueCommentable
func (m *SecretScanningLocation_SecretScanningLocation_details) SetSecretScanningLocationIssueComment(value SecretScanningLocationIssueCommentable)() {
    m.secretScanningLocationIssueComment = value
}
// SetSecretScanningLocationIssueTitle sets the secretScanningLocationIssueTitle property value. Composed type representation for type SecretScanningLocationIssueTitleable
func (m *SecretScanningLocation_SecretScanningLocation_details) SetSecretScanningLocationIssueTitle(value SecretScanningLocationIssueTitleable)() {
    m.secretScanningLocationIssueTitle = value
}
// SetSecretScanningLocationPullRequestBody sets the secretScanningLocationPullRequestBody property value. Composed type representation for type SecretScanningLocationPullRequestBodyable
func (m *SecretScanningLocation_SecretScanningLocation_details) SetSecretScanningLocationPullRequestBody(value SecretScanningLocationPullRequestBodyable)() {
    m.secretScanningLocationPullRequestBody = value
}
// SetSecretScanningLocationPullRequestComment sets the secretScanningLocationPullRequestComment property value. Composed type representation for type SecretScanningLocationPullRequestCommentable
func (m *SecretScanningLocation_SecretScanningLocation_details) SetSecretScanningLocationPullRequestComment(value SecretScanningLocationPullRequestCommentable)() {
    m.secretScanningLocationPullRequestComment = value
}
// SetSecretScanningLocationPullRequestReview sets the secretScanningLocationPullRequestReview property value. Composed type representation for type SecretScanningLocationPullRequestReviewable
func (m *SecretScanningLocation_SecretScanningLocation_details) SetSecretScanningLocationPullRequestReview(value SecretScanningLocationPullRequestReviewable)() {
    m.secretScanningLocationPullRequestReview = value
}
// SetSecretScanningLocationPullRequestReviewComment sets the secretScanningLocationPullRequestReviewComment property value. Composed type representation for type SecretScanningLocationPullRequestReviewCommentable
func (m *SecretScanningLocation_SecretScanningLocation_details) SetSecretScanningLocationPullRequestReviewComment(value SecretScanningLocationPullRequestReviewCommentable)() {
    m.secretScanningLocationPullRequestReviewComment = value
}
// SetSecretScanningLocationPullRequestTitle sets the secretScanningLocationPullRequestTitle property value. Composed type representation for type SecretScanningLocationPullRequestTitleable
func (m *SecretScanningLocation_SecretScanningLocation_details) SetSecretScanningLocationPullRequestTitle(value SecretScanningLocationPullRequestTitleable)() {
    m.secretScanningLocationPullRequestTitle = value
}
// SetSecretScanningLocationWikiCommit sets the secretScanningLocationWikiCommit property value. Composed type representation for type SecretScanningLocationWikiCommitable
func (m *SecretScanningLocation_SecretScanningLocation_details) SetSecretScanningLocationWikiCommit(value SecretScanningLocationWikiCommitable)() {
    m.secretScanningLocationWikiCommit = value
}
type SecretScanningLocation_SecretScanningLocation_detailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSecretScanningLocationCommit()(SecretScanningLocationCommitable)
    GetSecretScanningLocationDiscussionBody()(SecretScanningLocationDiscussionBodyable)
    GetSecretScanningLocationDiscussionComment()(SecretScanningLocationDiscussionCommentable)
    GetSecretScanningLocationDiscussionTitle()(SecretScanningLocationDiscussionTitleable)
    GetSecretScanningLocationIssueBody()(SecretScanningLocationIssueBodyable)
    GetSecretScanningLocationIssueComment()(SecretScanningLocationIssueCommentable)
    GetSecretScanningLocationIssueTitle()(SecretScanningLocationIssueTitleable)
    GetSecretScanningLocationPullRequestBody()(SecretScanningLocationPullRequestBodyable)
    GetSecretScanningLocationPullRequestComment()(SecretScanningLocationPullRequestCommentable)
    GetSecretScanningLocationPullRequestReview()(SecretScanningLocationPullRequestReviewable)
    GetSecretScanningLocationPullRequestReviewComment()(SecretScanningLocationPullRequestReviewCommentable)
    GetSecretScanningLocationPullRequestTitle()(SecretScanningLocationPullRequestTitleable)
    GetSecretScanningLocationWikiCommit()(SecretScanningLocationWikiCommitable)
    SetSecretScanningLocationCommit(value SecretScanningLocationCommitable)()
    SetSecretScanningLocationDiscussionBody(value SecretScanningLocationDiscussionBodyable)()
    SetSecretScanningLocationDiscussionComment(value SecretScanningLocationDiscussionCommentable)()
    SetSecretScanningLocationDiscussionTitle(value SecretScanningLocationDiscussionTitleable)()
    SetSecretScanningLocationIssueBody(value SecretScanningLocationIssueBodyable)()
    SetSecretScanningLocationIssueComment(value SecretScanningLocationIssueCommentable)()
    SetSecretScanningLocationIssueTitle(value SecretScanningLocationIssueTitleable)()
    SetSecretScanningLocationPullRequestBody(value SecretScanningLocationPullRequestBodyable)()
    SetSecretScanningLocationPullRequestComment(value SecretScanningLocationPullRequestCommentable)()
    SetSecretScanningLocationPullRequestReview(value SecretScanningLocationPullRequestReviewable)()
    SetSecretScanningLocationPullRequestReviewComment(value SecretScanningLocationPullRequestReviewCommentable)()
    SetSecretScanningLocationPullRequestTitle(value SecretScanningLocationPullRequestTitleable)()
    SetSecretScanningLocationWikiCommit(value SecretScanningLocationWikiCommitable)()
}
// NewSecretScanningLocation instantiates a new SecretScanningLocation and sets the default values.
func NewSecretScanningLocation()(*SecretScanningLocation) {
    m := &SecretScanningLocation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSecretScanningLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecretScanningLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecretScanningLocation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SecretScanningLocation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDetails gets the details property value. The details property
// returns a SecretScanningLocation_SecretScanningLocation_detailsable when successful
func (m *SecretScanningLocation) GetDetails()(SecretScanningLocation_SecretScanningLocation_detailsable) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecretScanningLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecretScanningLocation_SecretScanningLocation_detailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val.(SecretScanningLocation_SecretScanningLocation_detailsable))
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecretScanningLocation_type)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val.(*SecretScanningLocation_type))
        }
        return nil
    }
    return res
}
// GetTypeEscaped gets the type property value. The location type. Because secrets may be found in different types of resources (ie. code, comments, issues, pull requests, discussions), this field identifies the type of resource where the secret was found.
// returns a *SecretScanningLocation_type when successful
func (m *SecretScanningLocation) GetTypeEscaped()(*SecretScanningLocation_type) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *SecretScanningLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    if m.GetTypeEscaped() != nil {
        cast := (*m.GetTypeEscaped()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *SecretScanningLocation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDetails sets the details property value. The details property
func (m *SecretScanningLocation) SetDetails(value SecretScanningLocation_SecretScanningLocation_detailsable)() {
    m.details = value
}
// SetTypeEscaped sets the type property value. The location type. Because secrets may be found in different types of resources (ie. code, comments, issues, pull requests, discussions), this field identifies the type of resource where the secret was found.
func (m *SecretScanningLocation) SetTypeEscaped(value *SecretScanningLocation_type)() {
    m.typeEscaped = value
}
type SecretScanningLocationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetails()(SecretScanningLocation_SecretScanningLocation_detailsable)
    GetTypeEscaped()(*SecretScanningLocation_type)
    SetDetails(value SecretScanningLocation_SecretScanningLocation_detailsable)()
    SetTypeEscaped(value *SecretScanningLocation_type)()
}
