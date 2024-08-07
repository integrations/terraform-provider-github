package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TopicSearchResultItem topic Search Result Item
type TopicSearchResultItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The aliases property
    aliases []TopicSearchResultItem_aliasesable
    // The created_at property
    created_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The created_by property
    created_by *string
    // The curated property
    curated *bool
    // The description property
    description *string
    // The display_name property
    display_name *string
    // The featured property
    featured *bool
    // The logo_url property
    logo_url *string
    // The name property
    name *string
    // The related property
    related []TopicSearchResultItem_relatedable
    // The released property
    released *string
    // The repository_count property
    repository_count *int32
    // The score property
    score *float64
    // The short_description property
    short_description *string
    // The text_matches property
    text_matches []Topicsable
    // The updated_at property
    updated_at *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewTopicSearchResultItem instantiates a new TopicSearchResultItem and sets the default values.
func NewTopicSearchResultItem()(*TopicSearchResultItem) {
    m := &TopicSearchResultItem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTopicSearchResultItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTopicSearchResultItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTopicSearchResultItem(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TopicSearchResultItem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAliases gets the aliases property value. The aliases property
// returns a []TopicSearchResultItem_aliasesable when successful
func (m *TopicSearchResultItem) GetAliases()([]TopicSearchResultItem_aliasesable) {
    return m.aliases
}
// GetCreatedAt gets the created_at property value. The created_at property
// returns a *Time when successful
func (m *TopicSearchResultItem) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.created_at
}
// GetCreatedBy gets the created_by property value. The created_by property
// returns a *string when successful
func (m *TopicSearchResultItem) GetCreatedBy()(*string) {
    return m.created_by
}
// GetCurated gets the curated property value. The curated property
// returns a *bool when successful
func (m *TopicSearchResultItem) GetCurated()(*bool) {
    return m.curated
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *TopicSearchResultItem) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the display_name property value. The display_name property
// returns a *string when successful
func (m *TopicSearchResultItem) GetDisplayName()(*string) {
    return m.display_name
}
// GetFeatured gets the featured property value. The featured property
// returns a *bool when successful
func (m *TopicSearchResultItem) GetFeatured()(*bool) {
    return m.featured
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TopicSearchResultItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["aliases"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTopicSearchResultItem_aliasesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TopicSearchResultItem_aliasesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TopicSearchResultItem_aliasesable)
                }
            }
            m.SetAliases(res)
        }
        return nil
    }
    res["created_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["created_by"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val)
        }
        return nil
    }
    res["curated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurated(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["display_name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["featured"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatured(val)
        }
        return nil
    }
    res["logo_url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogoUrl(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["related"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTopicSearchResultItem_relatedFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TopicSearchResultItem_relatedable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TopicSearchResultItem_relatedable)
                }
            }
            m.SetRelated(res)
        }
        return nil
    }
    res["released"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleased(val)
        }
        return nil
    }
    res["repository_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepositoryCount(val)
        }
        return nil
    }
    res["score"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScore(val)
        }
        return nil
    }
    res["short_description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShortDescription(val)
        }
        return nil
    }
    res["text_matches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTopicsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Topicsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Topicsable)
                }
            }
            m.SetTextMatches(res)
        }
        return nil
    }
    res["updated_at"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    return res
}
// GetLogoUrl gets the logo_url property value. The logo_url property
// returns a *string when successful
func (m *TopicSearchResultItem) GetLogoUrl()(*string) {
    return m.logo_url
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *TopicSearchResultItem) GetName()(*string) {
    return m.name
}
// GetRelated gets the related property value. The related property
// returns a []TopicSearchResultItem_relatedable when successful
func (m *TopicSearchResultItem) GetRelated()([]TopicSearchResultItem_relatedable) {
    return m.related
}
// GetReleased gets the released property value. The released property
// returns a *string when successful
func (m *TopicSearchResultItem) GetReleased()(*string) {
    return m.released
}
// GetRepositoryCount gets the repository_count property value. The repository_count property
// returns a *int32 when successful
func (m *TopicSearchResultItem) GetRepositoryCount()(*int32) {
    return m.repository_count
}
// GetScore gets the score property value. The score property
// returns a *float64 when successful
func (m *TopicSearchResultItem) GetScore()(*float64) {
    return m.score
}
// GetShortDescription gets the short_description property value. The short_description property
// returns a *string when successful
func (m *TopicSearchResultItem) GetShortDescription()(*string) {
    return m.short_description
}
// GetTextMatches gets the text_matches property value. The text_matches property
// returns a []Topicsable when successful
func (m *TopicSearchResultItem) GetTextMatches()([]Topicsable) {
    return m.text_matches
}
// GetUpdatedAt gets the updated_at property value. The updated_at property
// returns a *Time when successful
func (m *TopicSearchResultItem) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updated_at
}
// Serialize serializes information the current object
func (m *TopicSearchResultItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAliases() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAliases()))
        for i, v := range m.GetAliases() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("aliases", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("created_at", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("created_by", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("curated", m.GetCurated())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("display_name", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("featured", m.GetFeatured())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("logo_url", m.GetLogoUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetRelated() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRelated()))
        for i, v := range m.GetRelated() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("related", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("released", m.GetReleased())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("repository_count", m.GetRepositoryCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("score", m.GetScore())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("short_description", m.GetShortDescription())
        if err != nil {
            return err
        }
    }
    if m.GetTextMatches() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTextMatches()))
        for i, v := range m.GetTextMatches() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("text_matches", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updated_at", m.GetUpdatedAt())
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
func (m *TopicSearchResultItem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAliases sets the aliases property value. The aliases property
func (m *TopicSearchResultItem) SetAliases(value []TopicSearchResultItem_aliasesable)() {
    m.aliases = value
}
// SetCreatedAt sets the created_at property value. The created_at property
func (m *TopicSearchResultItem) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.created_at = value
}
// SetCreatedBy sets the created_by property value. The created_by property
func (m *TopicSearchResultItem) SetCreatedBy(value *string)() {
    m.created_by = value
}
// SetCurated sets the curated property value. The curated property
func (m *TopicSearchResultItem) SetCurated(value *bool)() {
    m.curated = value
}
// SetDescription sets the description property value. The description property
func (m *TopicSearchResultItem) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the display_name property value. The display_name property
func (m *TopicSearchResultItem) SetDisplayName(value *string)() {
    m.display_name = value
}
// SetFeatured sets the featured property value. The featured property
func (m *TopicSearchResultItem) SetFeatured(value *bool)() {
    m.featured = value
}
// SetLogoUrl sets the logo_url property value. The logo_url property
func (m *TopicSearchResultItem) SetLogoUrl(value *string)() {
    m.logo_url = value
}
// SetName sets the name property value. The name property
func (m *TopicSearchResultItem) SetName(value *string)() {
    m.name = value
}
// SetRelated sets the related property value. The related property
func (m *TopicSearchResultItem) SetRelated(value []TopicSearchResultItem_relatedable)() {
    m.related = value
}
// SetReleased sets the released property value. The released property
func (m *TopicSearchResultItem) SetReleased(value *string)() {
    m.released = value
}
// SetRepositoryCount sets the repository_count property value. The repository_count property
func (m *TopicSearchResultItem) SetRepositoryCount(value *int32)() {
    m.repository_count = value
}
// SetScore sets the score property value. The score property
func (m *TopicSearchResultItem) SetScore(value *float64)() {
    m.score = value
}
// SetShortDescription sets the short_description property value. The short_description property
func (m *TopicSearchResultItem) SetShortDescription(value *string)() {
    m.short_description = value
}
// SetTextMatches sets the text_matches property value. The text_matches property
func (m *TopicSearchResultItem) SetTextMatches(value []Topicsable)() {
    m.text_matches = value
}
// SetUpdatedAt sets the updated_at property value. The updated_at property
func (m *TopicSearchResultItem) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updated_at = value
}
type TopicSearchResultItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAliases()([]TopicSearchResultItem_aliasesable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedBy()(*string)
    GetCurated()(*bool)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetFeatured()(*bool)
    GetLogoUrl()(*string)
    GetName()(*string)
    GetRelated()([]TopicSearchResultItem_relatedable)
    GetReleased()(*string)
    GetRepositoryCount()(*int32)
    GetScore()(*float64)
    GetShortDescription()(*string)
    GetTextMatches()([]Topicsable)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAliases(value []TopicSearchResultItem_aliasesable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedBy(value *string)()
    SetCurated(value *bool)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetFeatured(value *bool)()
    SetLogoUrl(value *string)()
    SetName(value *string)()
    SetRelated(value []TopicSearchResultItem_relatedable)()
    SetReleased(value *string)()
    SetRepositoryCount(value *int32)()
    SetScore(value *float64)()
    SetShortDescription(value *string)()
    SetTextMatches(value []Topicsable)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
