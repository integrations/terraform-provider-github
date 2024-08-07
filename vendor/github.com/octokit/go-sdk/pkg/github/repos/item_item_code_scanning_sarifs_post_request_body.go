package repos

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemItemCodeScanningSarifsPostRequestBody struct {
    // The base directory used in the analysis, as it appears in the SARIF file.This property is used to convert file paths from absolute to relative, so that alerts can be mapped to their correct location in the repository.
    checkout_uri *string
    // The SHA of the commit to which the analysis you are uploading relates.
    commit_sha *string
    // The full Git reference, formatted as `refs/heads/<branch name>`,`refs/tags/<tag>`, `refs/pull/<number>/merge`, or `refs/pull/<number>/head`.
    ref *string
    // A Base64 string representing the SARIF file to upload. You must first compress your SARIF file using [`gzip`](http://www.gnu.org/software/gzip/manual/gzip.html) and then translate the contents of the file into a Base64 encoding string. For more information, see "[SARIF support for code scanning](https://docs.github.com/code-security/secure-coding/sarif-support-for-code-scanning)."
    sarif *string
    // The time that the analysis run began. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    started_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the tool used to generate the code scanning analysis. If this parameter is not used, the tool name defaults to "API". If the uploaded SARIF contains a tool GUID, this will be available for filtering using the `tool_guid` parameter of operations such as `GET /repos/{owner}/{repo}/code-scanning/alerts`.
    tool_name *string
    // Whether the SARIF file will be validated according to the code scanning specifications.This parameter is intended to help integrators ensure that the uploaded SARIF files are correctly rendered by code scanning.
    validate *bool
}
// NewItemItemCodeScanningSarifsPostRequestBody instantiates a new ItemItemCodeScanningSarifsPostRequestBody and sets the default values.
func NewItemItemCodeScanningSarifsPostRequestBody()(*ItemItemCodeScanningSarifsPostRequestBody) {
    m := &ItemItemCodeScanningSarifsPostRequestBody{
    }
    return m
}
// CreateItemItemCodeScanningSarifsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemCodeScanningSarifsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemCodeScanningSarifsPostRequestBody(), nil
}
// GetCheckoutUri gets the checkout_uri property value. The base directory used in the analysis, as it appears in the SARIF file.This property is used to convert file paths from absolute to relative, so that alerts can be mapped to their correct location in the repository.
// returns a *string when successful
func (m *ItemItemCodeScanningSarifsPostRequestBody) GetCheckoutUri()(*string) {
    return m.checkout_uri
}
// GetCommitSha gets the commit_sha property value. The SHA of the commit to which the analysis you are uploading relates.
// returns a *string when successful
func (m *ItemItemCodeScanningSarifsPostRequestBody) GetCommitSha()(*string) {
    return m.commit_sha
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemItemCodeScanningSarifsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["checkout_uri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckoutUri(val)
        }
        return nil
    }
    res["commit_sha"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitSha(val)
        }
        return nil
    }
    res["ref"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRef(val)
        }
        return nil
    }
    res["sarif"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSarif(val)
        }
        return nil
    }
    res["started_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartedAt(val)
        }
        return nil
    }
    res["tool_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToolName(val)
        }
        return nil
    }
    res["validate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidate(val)
        }
        return nil
    }
    return res
}
// GetRef gets the ref property value. The full Git reference, formatted as `refs/heads/<branch name>`,`refs/tags/<tag>`, `refs/pull/<number>/merge`, or `refs/pull/<number>/head`.
// returns a *string when successful
func (m *ItemItemCodeScanningSarifsPostRequestBody) GetRef()(*string) {
    return m.ref
}
// GetSarif gets the sarif property value. A Base64 string representing the SARIF file to upload. You must first compress your SARIF file using [`gzip`](http://www.gnu.org/software/gzip/manual/gzip.html) and then translate the contents of the file into a Base64 encoding string. For more information, see "[SARIF support for code scanning](https://docs.github.com/code-security/secure-coding/sarif-support-for-code-scanning)."
// returns a *string when successful
func (m *ItemItemCodeScanningSarifsPostRequestBody) GetSarif()(*string) {
    return m.sarif
}
// GetStartedAt gets the started_at property value. The time that the analysis run began. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
// returns a *Time when successful
func (m *ItemItemCodeScanningSarifsPostRequestBody) GetStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.started_at
}
// GetToolName gets the tool_name property value. The name of the tool used to generate the code scanning analysis. If this parameter is not used, the tool name defaults to "API". If the uploaded SARIF contains a tool GUID, this will be available for filtering using the `tool_guid` parameter of operations such as `GET /repos/{owner}/{repo}/code-scanning/alerts`.
// returns a *string when successful
func (m *ItemItemCodeScanningSarifsPostRequestBody) GetToolName()(*string) {
    return m.tool_name
}
// GetValidate gets the validate property value. Whether the SARIF file will be validated according to the code scanning specifications.This parameter is intended to help integrators ensure that the uploaded SARIF files are correctly rendered by code scanning.
// returns a *bool when successful
func (m *ItemItemCodeScanningSarifsPostRequestBody) GetValidate()(*bool) {
    return m.validate
}
// Serialize serializes information the current object
func (m *ItemItemCodeScanningSarifsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("checkout_uri", m.GetCheckoutUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commit_sha", m.GetCommitSha())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ref", m.GetRef())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sarif", m.GetSarif())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("started_at", m.GetStartedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tool_name", m.GetToolName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("validate", m.GetValidate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCheckoutUri sets the checkout_uri property value. The base directory used in the analysis, as it appears in the SARIF file.This property is used to convert file paths from absolute to relative, so that alerts can be mapped to their correct location in the repository.
func (m *ItemItemCodeScanningSarifsPostRequestBody) SetCheckoutUri(value *string)() {
    m.checkout_uri = value
}
// SetCommitSha sets the commit_sha property value. The SHA of the commit to which the analysis you are uploading relates.
func (m *ItemItemCodeScanningSarifsPostRequestBody) SetCommitSha(value *string)() {
    m.commit_sha = value
}
// SetRef sets the ref property value. The full Git reference, formatted as `refs/heads/<branch name>`,`refs/tags/<tag>`, `refs/pull/<number>/merge`, or `refs/pull/<number>/head`.
func (m *ItemItemCodeScanningSarifsPostRequestBody) SetRef(value *string)() {
    m.ref = value
}
// SetSarif sets the sarif property value. A Base64 string representing the SARIF file to upload. You must first compress your SARIF file using [`gzip`](http://www.gnu.org/software/gzip/manual/gzip.html) and then translate the contents of the file into a Base64 encoding string. For more information, see "[SARIF support for code scanning](https://docs.github.com/code-security/secure-coding/sarif-support-for-code-scanning)."
func (m *ItemItemCodeScanningSarifsPostRequestBody) SetSarif(value *string)() {
    m.sarif = value
}
// SetStartedAt sets the started_at property value. The time that the analysis run began. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
func (m *ItemItemCodeScanningSarifsPostRequestBody) SetStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.started_at = value
}
// SetToolName sets the tool_name property value. The name of the tool used to generate the code scanning analysis. If this parameter is not used, the tool name defaults to "API". If the uploaded SARIF contains a tool GUID, this will be available for filtering using the `tool_guid` parameter of operations such as `GET /repos/{owner}/{repo}/code-scanning/alerts`.
func (m *ItemItemCodeScanningSarifsPostRequestBody) SetToolName(value *string)() {
    m.tool_name = value
}
// SetValidate sets the validate property value. Whether the SARIF file will be validated according to the code scanning specifications.This parameter is intended to help integrators ensure that the uploaded SARIF files are correctly rendered by code scanning.
func (m *ItemItemCodeScanningSarifsPostRequestBody) SetValidate(value *bool)() {
    m.validate = value
}
type ItemItemCodeScanningSarifsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCheckoutUri()(*string)
    GetCommitSha()(*string)
    GetRef()(*string)
    GetSarif()(*string)
    GetStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetToolName()(*string)
    GetValidate()(*bool)
    SetCheckoutUri(value *string)()
    SetCommitSha(value *string)()
    SetRef(value *string)()
    SetSarif(value *string)()
    SetStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetToolName(value *string)()
    SetValidate(value *bool)()
}
