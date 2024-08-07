package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Snapshot create a new snapshot of a repository's dependencies.
type Snapshot struct {
    // A description of the detector used.
    detector Snapshot_detectorable
    // The job property
    job Snapshot_jobable
    // A collection of package manifests, which are a collection of related dependencies declared in a file or representing a logical group of dependencies.
    manifests Snapshot_manifestsable
    // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
    metadata Metadataable
    // The repository branch that triggered this snapshot.
    ref *string
    // The time at which the snapshot was scanned.
    scanned *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The commit SHA associated with this dependency snapshot. Maximum length: 40 characters.
    sha *string
    // The version of the repository snapshot submission.
    version *int32
}
// NewSnapshot instantiates a new Snapshot and sets the default values.
func NewSnapshot()(*Snapshot) {
    m := &Snapshot{
    }
    return m
}
// CreateSnapshotFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSnapshotFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSnapshot(), nil
}
// GetDetector gets the detector property value. A description of the detector used.
// returns a Snapshot_detectorable when successful
func (m *Snapshot) GetDetector()(Snapshot_detectorable) {
    return m.detector
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Snapshot) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["detector"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSnapshot_detectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetector(val.(Snapshot_detectorable))
        }
        return nil
    }
    res["job"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSnapshot_jobFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJob(val.(Snapshot_jobable))
        }
        return nil
    }
    res["manifests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSnapshot_manifestsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManifests(val.(Snapshot_manifestsable))
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadata(val.(Metadataable))
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
    res["scanned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScanned(val)
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
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetJob gets the job property value. The job property
// returns a Snapshot_jobable when successful
func (m *Snapshot) GetJob()(Snapshot_jobable) {
    return m.job
}
// GetManifests gets the manifests property value. A collection of package manifests, which are a collection of related dependencies declared in a file or representing a logical group of dependencies.
// returns a Snapshot_manifestsable when successful
func (m *Snapshot) GetManifests()(Snapshot_manifestsable) {
    return m.manifests
}
// GetMetadata gets the metadata property value. User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
// returns a Metadataable when successful
func (m *Snapshot) GetMetadata()(Metadataable) {
    return m.metadata
}
// GetRef gets the ref property value. The repository branch that triggered this snapshot.
// returns a *string when successful
func (m *Snapshot) GetRef()(*string) {
    return m.ref
}
// GetScanned gets the scanned property value. The time at which the snapshot was scanned.
// returns a *Time when successful
func (m *Snapshot) GetScanned()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.scanned
}
// GetSha gets the sha property value. The commit SHA associated with this dependency snapshot. Maximum length: 40 characters.
// returns a *string when successful
func (m *Snapshot) GetSha()(*string) {
    return m.sha
}
// GetVersion gets the version property value. The version of the repository snapshot submission.
// returns a *int32 when successful
func (m *Snapshot) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *Snapshot) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("detector", m.GetDetector())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("job", m.GetJob())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("manifests", m.GetManifests())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("metadata", m.GetMetadata())
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
        err := writer.WriteTimeValue("scanned", m.GetScanned())
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
        err := writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDetector sets the detector property value. A description of the detector used.
func (m *Snapshot) SetDetector(value Snapshot_detectorable)() {
    m.detector = value
}
// SetJob sets the job property value. The job property
func (m *Snapshot) SetJob(value Snapshot_jobable)() {
    m.job = value
}
// SetManifests sets the manifests property value. A collection of package manifests, which are a collection of related dependencies declared in a file or representing a logical group of dependencies.
func (m *Snapshot) SetManifests(value Snapshot_manifestsable)() {
    m.manifests = value
}
// SetMetadata sets the metadata property value. User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
func (m *Snapshot) SetMetadata(value Metadataable)() {
    m.metadata = value
}
// SetRef sets the ref property value. The repository branch that triggered this snapshot.
func (m *Snapshot) SetRef(value *string)() {
    m.ref = value
}
// SetScanned sets the scanned property value. The time at which the snapshot was scanned.
func (m *Snapshot) SetScanned(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.scanned = value
}
// SetSha sets the sha property value. The commit SHA associated with this dependency snapshot. Maximum length: 40 characters.
func (m *Snapshot) SetSha(value *string)() {
    m.sha = value
}
// SetVersion sets the version property value. The version of the repository snapshot submission.
func (m *Snapshot) SetVersion(value *int32)() {
    m.version = value
}
type Snapshotable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDetector()(Snapshot_detectorable)
    GetJob()(Snapshot_jobable)
    GetManifests()(Snapshot_manifestsable)
    GetMetadata()(Metadataable)
    GetRef()(*string)
    GetScanned()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSha()(*string)
    GetVersion()(*int32)
    SetDetector(value Snapshot_detectorable)()
    SetJob(value Snapshot_jobable)()
    SetManifests(value Snapshot_manifestsable)()
    SetMetadata(value Metadataable)()
    SetRef(value *string)()
    SetScanned(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSha(value *string)()
    SetVersion(value *int32)()
}
