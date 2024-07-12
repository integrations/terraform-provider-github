package github

import (
    "context"
    i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488 "github.com/microsoft/kiota-serialization-json-go"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347 "github.com/microsoft/kiota-serialization-form-go"
    i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426 "github.com/microsoft/kiota-serialization-multipart-go"
    i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83 "github.com/microsoft/kiota-serialization-text-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i0418eb5bc6d6e5c5596c278b5f327e24b1594d0a2338da0da064365a642eceff "github.com/octokit/go-sdk/pkg/github/app"
    i04dfd3ae6494a1aec921b0f0ae334653c1b3e6091b7b31a50e7977b2b4e54f66 "github.com/octokit/go-sdk/pkg/github/users"
    i0737a59f72bd0b3677f51f977eb9a1a5e217db56b30eda4683f072dd51df524a "github.com/octokit/go-sdk/pkg/github/user"
    i08b5ce560895c2e4c4bd1673fc739fcd3d31e314c5b53adc6a89aba3d8c6244e "github.com/octokit/go-sdk/pkg/github/appmanifests"
    i0ee0b8f649e15b17c2eabdd46a0f8c70c497e9c6cdef6ff35be832f2eb719ac3 "github.com/octokit/go-sdk/pkg/github/classrooms"
    i12a9b43201d9a86bb60f5b677ba8d0e7e5da47bc806080ad3218e4c65628b957 "github.com/octokit/go-sdk/pkg/github/rate_limit"
    i24ae3cb00d5734ee78f4378d5739457021b8117fbda874d208f0e79ba3767e84 "github.com/octokit/go-sdk/pkg/github/enterprises"
    i331f27c1d84b94cd34b6f5276a80b9b495557537e64366b12a82d4f6256cc315 "github.com/octokit/go-sdk/pkg/github/markdown"
    i3d7d4dc978cfc98dfee1e9704d82dee19f232386607bb0ec32b0288e7ad55112 "github.com/octokit/go-sdk/pkg/github/notifications"
    i4000a3f52da0039065e50768fd1044a9efa2154163159903dfe0efad17ee681d "github.com/octokit/go-sdk/pkg/github/repos"
    i47b66686b0190707417a08d06857b07dec037d9b5534bb91752a10f16de0e9db "github.com/octokit/go-sdk/pkg/github/search"
    i4d1eee1704a96ea8ab53e424182986725fde88abc66df8b117e9a2f20429d93a "github.com/octokit/go-sdk/pkg/github/meta"
    i54045f2cdc31d6122b70e132f86739f81aea3c9aaaa4bbebeafc3e09be67700a "github.com/octokit/go-sdk/pkg/github/organizations"
    i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6 "github.com/octokit/go-sdk/pkg/github/models"
    i5c383b7fca29547223d6e4abf6eaff83f09b72e11524fd25346fd31be394ace5 "github.com/octokit/go-sdk/pkg/github/issues"
    i5cf6ed0734619011c475ed18bdb3bf27f375f600864b00ae52feb4473ffe1081 "github.com/octokit/go-sdk/pkg/github/emojis"
    i6265cecf37e9b8b6daa13472abde6c43c88081bb9d68ee62d9c56c001ae179d4 "github.com/octokit/go-sdk/pkg/github/feeds"
    i68b78648401e10ac283b41c6c62b9e34e6b84e36aed7bbf9a408157262621541 "github.com/octokit/go-sdk/pkg/github/gitignore"
    i6fbd16012f9d384db4ed793592fd1e4960e1e71a355ae0e5a6a457dec525b9fa "github.com/octokit/go-sdk/pkg/github/orgs"
    i77422b4816b2e4034320d80ba8150f21a4209c1d31d547c1f9105e20e5bd24ac "github.com/octokit/go-sdk/pkg/github/applications"
    i7c2f4266a964aa71671d1f93b3dedd60ab0fb58fc1c48d67ec5323e06f987d13 "github.com/octokit/go-sdk/pkg/github/installation"
    i81237a20a3d3d568f85359ecb683b25d592945354af01b3bb96015225d355d0b "github.com/octokit/go-sdk/pkg/github/advisories"
    i9058dc6558645b18c3dabb68ce06f623305fd2f87c0952167619d0d8b4249490 "github.com/octokit/go-sdk/pkg/github/zen"
    i92a99a8b5d9294c6b4600047e7f16515a40994715a7ed07d4afa2393ed306c04 "github.com/octokit/go-sdk/pkg/github/octocat"
    i986f660025ed63efc77aa38f829d1f8dbd3f63e6ad3c8df735c5335e37b607be "github.com/octokit/go-sdk/pkg/github/licenses"
    i9a3e59ac75b1a35f0b41cb5bba910e278ebbbf0527883428330d78ca41ad925b "github.com/octokit/go-sdk/pkg/github/networks"
    i9c8518268801f75b016b6b68b07c0ac4c95512b503cc56ba270c337a1b265384 "github.com/octokit/go-sdk/pkg/github/repositories"
    ia5478699530639ac7ac52c373afc47e03f523b842a72b212b381ca49aae9255d "github.com/octokit/go-sdk/pkg/github/codes_of_conduct"
    ia7fd50a86ade62d3f84a4d485b2c3bb76080fbbce4f4ab65fdcbd7789bd82705 "github.com/octokit/go-sdk/pkg/github/gists"
    ib62b62162b22a0c33176b41c32ba6062508c8c70e490bfdd926b74d9500eb171 "github.com/octokit/go-sdk/pkg/github/events"
    icccce443123713e2f41a15534fb2982842bc02bef783c80ed5f0f20406addd52 "github.com/octokit/go-sdk/pkg/github/teams"
    id2671b72dcd915381cbe90f24655e8c8c52db963e200cc8f20b1738da7beb68f "github.com/octokit/go-sdk/pkg/github/assignments"
    ie7ab5250c8bf9bd78939f5773ba8b0e78d558cfedd3cf8858458622c0c8ee879 "github.com/octokit/go-sdk/pkg/github/marketplace_listing"
    if0a678a7358b449c2d0570a0c4327fbd4d0b69bdec4884ec0cd79748e725fc19 "github.com/octokit/go-sdk/pkg/github/projects"
    if6a0b82222a4a600b84ed4730a3e16ae4912e617120d6fd543a9e35c9c81e38e "github.com/octokit/go-sdk/pkg/github/apps"
    if9262d059a5df6366c8597430a86d0c7247500ee47d3cb11ea28c4da6de74620 "github.com/octokit/go-sdk/pkg/github/versions"
)

// ApiClient the main entry point of the SDK, exposes the configuration and the fluent API.
type ApiClient struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Advisories the advisories property
// returns a *AdvisoriesRequestBuilder when successful
func (m *ApiClient) Advisories()(*i81237a20a3d3d568f85359ecb683b25d592945354af01b3bb96015225d355d0b.AdvisoriesRequestBuilder) {
    return i81237a20a3d3d568f85359ecb683b25d592945354af01b3bb96015225d355d0b.NewAdvisoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// App the app property
// returns a *AppRequestBuilder when successful
func (m *ApiClient) App()(*i0418eb5bc6d6e5c5596c278b5f327e24b1594d0a2338da0da064365a642eceff.AppRequestBuilder) {
    return i0418eb5bc6d6e5c5596c278b5f327e24b1594d0a2338da0da064365a642eceff.NewAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Applications the applications property
// returns a *ApplicationsRequestBuilder when successful
func (m *ApiClient) Applications()(*i77422b4816b2e4034320d80ba8150f21a4209c1d31d547c1f9105e20e5bd24ac.ApplicationsRequestBuilder) {
    return i77422b4816b2e4034320d80ba8150f21a4209c1d31d547c1f9105e20e5bd24ac.NewApplicationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppManifests the appManifests property
// returns a *AppManifestsRequestBuilder when successful
func (m *ApiClient) AppManifests()(*i08b5ce560895c2e4c4bd1673fc739fcd3d31e314c5b53adc6a89aba3d8c6244e.AppManifestsRequestBuilder) {
    return i08b5ce560895c2e4c4bd1673fc739fcd3d31e314c5b53adc6a89aba3d8c6244e.NewAppManifestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Apps the apps property
// returns a *AppsRequestBuilder when successful
func (m *ApiClient) Apps()(*if6a0b82222a4a600b84ed4730a3e16ae4912e617120d6fd543a9e35c9c81e38e.AppsRequestBuilder) {
    return if6a0b82222a4a600b84ed4730a3e16ae4912e617120d6fd543a9e35c9c81e38e.NewAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments the assignments property
// returns a *AssignmentsRequestBuilder when successful
func (m *ApiClient) Assignments()(*id2671b72dcd915381cbe90f24655e8c8c52db963e200cc8f20b1738da7beb68f.AssignmentsRequestBuilder) {
    return id2671b72dcd915381cbe90f24655e8c8c52db963e200cc8f20b1738da7beb68f.NewAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Classrooms the classrooms property
// returns a *ClassroomsRequestBuilder when successful
func (m *ApiClient) Classrooms()(*i0ee0b8f649e15b17c2eabdd46a0f8c70c497e9c6cdef6ff35be832f2eb719ac3.ClassroomsRequestBuilder) {
    return i0ee0b8f649e15b17c2eabdd46a0f8c70c497e9c6cdef6ff35be832f2eb719ac3.NewClassroomsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Codes_of_conduct the codes_of_conduct property
// returns a *Codes_of_conductRequestBuilder when successful
func (m *ApiClient) Codes_of_conduct()(*ia5478699530639ac7ac52c373afc47e03f523b842a72b212b381ca49aae9255d.Codes_of_conductRequestBuilder) {
    return ia5478699530639ac7ac52c373afc47e03f523b842a72b212b381ca49aae9255d.NewCodes_of_conductRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewApiClient instantiates a new ApiClient and sets the default values.
func NewApiClient(requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApiClient) {
    m := &ApiClient{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}", map[string]string{}),
    }
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultSerializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriterFactory { return i56887720f41ac882814261620b1c8459c4a992a0207af547c4453dd39fabc426.NewMultipartSerializationWriterFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i25911dc319edd61cbac496af7eab5ef20b6069a42515e22ec6a9bc97bf598488.NewJsonParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i7294a22093d408fdca300f11b81a887d89c47b764af06c8b803e2323973fdb83.NewTextParseNodeFactory() })
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RegisterDefaultDeserializer(func() i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNodeFactory { return i4bcdc892e61ac17e2afc10b5e2b536b29f4fd6c1ad30f4a5a68df47495db3347.NewFormParseNodeFactory() })
    if m.BaseRequestBuilder.RequestAdapter.GetBaseUrl() == "" {
        m.BaseRequestBuilder.RequestAdapter.SetBaseUrl("https://api.github.com")
    }
    m.BaseRequestBuilder.PathParameters["baseurl"] = m.BaseRequestBuilder.RequestAdapter.GetBaseUrl()
    return m
}
// Emojis the emojis property
// returns a *EmojisRequestBuilder when successful
func (m *ApiClient) Emojis()(*i5cf6ed0734619011c475ed18bdb3bf27f375f600864b00ae52feb4473ffe1081.EmojisRequestBuilder) {
    return i5cf6ed0734619011c475ed18bdb3bf27f375f600864b00ae52feb4473ffe1081.NewEmojisRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Enterprises the enterprises property
// returns a *EnterprisesRequestBuilder when successful
func (m *ApiClient) Enterprises()(*i24ae3cb00d5734ee78f4378d5739457021b8117fbda874d208f0e79ba3767e84.EnterprisesRequestBuilder) {
    return i24ae3cb00d5734ee78f4378d5739457021b8117fbda874d208f0e79ba3767e84.NewEnterprisesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Events the events property
// returns a *EventsRequestBuilder when successful
func (m *ApiClient) Events()(*ib62b62162b22a0c33176b41c32ba6062508c8c70e490bfdd926b74d9500eb171.EventsRequestBuilder) {
    return ib62b62162b22a0c33176b41c32ba6062508c8c70e490bfdd926b74d9500eb171.NewEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Feeds the feeds property
// returns a *FeedsRequestBuilder when successful
func (m *ApiClient) Feeds()(*i6265cecf37e9b8b6daa13472abde6c43c88081bb9d68ee62d9c56c001ae179d4.FeedsRequestBuilder) {
    return i6265cecf37e9b8b6daa13472abde6c43c88081bb9d68ee62d9c56c001ae179d4.NewFeedsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get Hypermedia links to resources accessible in GitHub's REST API
// returns a Rootable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/rest/meta/meta#github-api-root
func (m *ApiClient) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Rootable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.CreateRootFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i59ea7d99994c6a4bb9ef742ed717844297d055c7fd3742131406eea67a6404b6.Rootable), nil
}
// Gists the gists property
// returns a *GistsRequestBuilder when successful
func (m *ApiClient) Gists()(*ia7fd50a86ade62d3f84a4d485b2c3bb76080fbbce4f4ab65fdcbd7789bd82705.GistsRequestBuilder) {
    return ia7fd50a86ade62d3f84a4d485b2c3bb76080fbbce4f4ab65fdcbd7789bd82705.NewGistsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Gitignore the gitignore property
// returns a *GitignoreRequestBuilder when successful
func (m *ApiClient) Gitignore()(*i68b78648401e10ac283b41c6c62b9e34e6b84e36aed7bbf9a408157262621541.GitignoreRequestBuilder) {
    return i68b78648401e10ac283b41c6c62b9e34e6b84e36aed7bbf9a408157262621541.NewGitignoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Installation the installation property
// returns a *InstallationRequestBuilder when successful
func (m *ApiClient) Installation()(*i7c2f4266a964aa71671d1f93b3dedd60ab0fb58fc1c48d67ec5323e06f987d13.InstallationRequestBuilder) {
    return i7c2f4266a964aa71671d1f93b3dedd60ab0fb58fc1c48d67ec5323e06f987d13.NewInstallationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Issues the issues property
// returns a *IssuesRequestBuilder when successful
func (m *ApiClient) Issues()(*i5c383b7fca29547223d6e4abf6eaff83f09b72e11524fd25346fd31be394ace5.IssuesRequestBuilder) {
    return i5c383b7fca29547223d6e4abf6eaff83f09b72e11524fd25346fd31be394ace5.NewIssuesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Licenses the licenses property
// returns a *LicensesRequestBuilder when successful
func (m *ApiClient) Licenses()(*i986f660025ed63efc77aa38f829d1f8dbd3f63e6ad3c8df735c5335e37b607be.LicensesRequestBuilder) {
    return i986f660025ed63efc77aa38f829d1f8dbd3f63e6ad3c8df735c5335e37b607be.NewLicensesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Markdown the markdown property
// returns a *MarkdownRequestBuilder when successful
func (m *ApiClient) Markdown()(*i331f27c1d84b94cd34b6f5276a80b9b495557537e64366b12a82d4f6256cc315.MarkdownRequestBuilder) {
    return i331f27c1d84b94cd34b6f5276a80b9b495557537e64366b12a82d4f6256cc315.NewMarkdownRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Marketplace_listing the marketplace_listing property
// returns a *Marketplace_listingRequestBuilder when successful
func (m *ApiClient) Marketplace_listing()(*ie7ab5250c8bf9bd78939f5773ba8b0e78d558cfedd3cf8858458622c0c8ee879.Marketplace_listingRequestBuilder) {
    return ie7ab5250c8bf9bd78939f5773ba8b0e78d558cfedd3cf8858458622c0c8ee879.NewMarketplace_listingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Meta the meta property
// returns a *MetaRequestBuilder when successful
func (m *ApiClient) Meta()(*i4d1eee1704a96ea8ab53e424182986725fde88abc66df8b117e9a2f20429d93a.MetaRequestBuilder) {
    return i4d1eee1704a96ea8ab53e424182986725fde88abc66df8b117e9a2f20429d93a.NewMetaRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Networks the networks property
// returns a *NetworksRequestBuilder when successful
func (m *ApiClient) Networks()(*i9a3e59ac75b1a35f0b41cb5bba910e278ebbbf0527883428330d78ca41ad925b.NetworksRequestBuilder) {
    return i9a3e59ac75b1a35f0b41cb5bba910e278ebbbf0527883428330d78ca41ad925b.NewNetworksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Notifications the notifications property
// returns a *NotificationsRequestBuilder when successful
func (m *ApiClient) Notifications()(*i3d7d4dc978cfc98dfee1e9704d82dee19f232386607bb0ec32b0288e7ad55112.NotificationsRequestBuilder) {
    return i3d7d4dc978cfc98dfee1e9704d82dee19f232386607bb0ec32b0288e7ad55112.NewNotificationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Octocat the octocat property
// returns a *OctocatRequestBuilder when successful
func (m *ApiClient) Octocat()(*i92a99a8b5d9294c6b4600047e7f16515a40994715a7ed07d4afa2393ed306c04.OctocatRequestBuilder) {
    return i92a99a8b5d9294c6b4600047e7f16515a40994715a7ed07d4afa2393ed306c04.NewOctocatRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Organizations the organizations property
// returns a *OrganizationsRequestBuilder when successful
func (m *ApiClient) Organizations()(*i54045f2cdc31d6122b70e132f86739f81aea3c9aaaa4bbebeafc3e09be67700a.OrganizationsRequestBuilder) {
    return i54045f2cdc31d6122b70e132f86739f81aea3c9aaaa4bbebeafc3e09be67700a.NewOrganizationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Orgs the orgs property
// returns a *OrgsRequestBuilder when successful
func (m *ApiClient) Orgs()(*i6fbd16012f9d384db4ed793592fd1e4960e1e71a355ae0e5a6a457dec525b9fa.OrgsRequestBuilder) {
    return i6fbd16012f9d384db4ed793592fd1e4960e1e71a355ae0e5a6a457dec525b9fa.NewOrgsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Projects the projects property
// returns a *ProjectsRequestBuilder when successful
func (m *ApiClient) Projects()(*if0a678a7358b449c2d0570a0c4327fbd4d0b69bdec4884ec0cd79748e725fc19.ProjectsRequestBuilder) {
    return if0a678a7358b449c2d0570a0c4327fbd4d0b69bdec4884ec0cd79748e725fc19.NewProjectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rate_limit the rate_limit property
// returns a *Rate_limitRequestBuilder when successful
func (m *ApiClient) Rate_limit()(*i12a9b43201d9a86bb60f5b677ba8d0e7e5da47bc806080ad3218e4c65628b957.Rate_limitRequestBuilder) {
    return i12a9b43201d9a86bb60f5b677ba8d0e7e5da47bc806080ad3218e4c65628b957.NewRate_limitRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repos the repos property
// returns a *ReposRequestBuilder when successful
func (m *ApiClient) Repos()(*i4000a3f52da0039065e50768fd1044a9efa2154163159903dfe0efad17ee681d.ReposRequestBuilder) {
    return i4000a3f52da0039065e50768fd1044a9efa2154163159903dfe0efad17ee681d.NewReposRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Repositories the repositories property
// returns a *RepositoriesRequestBuilder when successful
func (m *ApiClient) Repositories()(*i9c8518268801f75b016b6b68b07c0ac4c95512b503cc56ba270c337a1b265384.RepositoriesRequestBuilder) {
    return i9c8518268801f75b016b6b68b07c0ac4c95512b503cc56ba270c337a1b265384.NewRepositoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Search the search property
// returns a *SearchRequestBuilder when successful
func (m *ApiClient) Search()(*i47b66686b0190707417a08d06857b07dec037d9b5534bb91752a10f16de0e9db.SearchRequestBuilder) {
    return i47b66686b0190707417a08d06857b07dec037d9b5534bb91752a10f16de0e9db.NewSearchRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Teams the teams property
// returns a *TeamsRequestBuilder when successful
func (m *ApiClient) Teams()(*icccce443123713e2f41a15534fb2982842bc02bef783c80ed5f0f20406addd52.TeamsRequestBuilder) {
    return icccce443123713e2f41a15534fb2982842bc02bef783c80ed5f0f20406addd52.NewTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get Hypermedia links to resources accessible in GitHub's REST API
// returns a *RequestInformation when successful
func (m *ApiClient) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// User the user property
// returns a *UserRequestBuilder when successful
func (m *ApiClient) User()(*i0737a59f72bd0b3677f51f977eb9a1a5e217db56b30eda4683f072dd51df524a.UserRequestBuilder) {
    return i0737a59f72bd0b3677f51f977eb9a1a5e217db56b30eda4683f072dd51df524a.NewUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Users the users property
// returns a *UsersRequestBuilder when successful
func (m *ApiClient) Users()(*i04dfd3ae6494a1aec921b0f0ae334653c1b3e6091b7b31a50e7977b2b4e54f66.UsersRequestBuilder) {
    return i04dfd3ae6494a1aec921b0f0ae334653c1b3e6091b7b31a50e7977b2b4e54f66.NewUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Versions the versions property
// returns a *VersionsRequestBuilder when successful
func (m *ApiClient) Versions()(*if9262d059a5df6366c8597430a86d0c7247500ee47d3cb11ea28c4da6de74620.VersionsRequestBuilder) {
    return if9262d059a5df6366c8597430a86d0c7247500ee47d3cb11ea28c4da6de74620.NewVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Zen the zen property
// returns a *ZenRequestBuilder when successful
func (m *ApiClient) Zen()(*i9058dc6558645b18c3dabb68ce06f623305fd2f87c0952167619d0d8b4249490.ZenRequestBuilder) {
    return i9058dc6558645b18c3dabb68ce06f623305fd2f87c0952167619d0d8b4249490.NewZenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
