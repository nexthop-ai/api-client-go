# Search
(*Client.Search*)

## Overview

### Available Operations

* [QueryAsAdmin](#queryasadmin) - Search the index (admin)
* [Autocomplete](#autocomplete) - Autocomplete
* [RetrieveFeed](#retrievefeed) - Feed of documents and events
* [Recommendations](#recommendations) - Recommend documents
* [Query](#query) - Search

## QueryAsAdmin

Retrieves results for search query without respect for permissions. This is available only to privileged users.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/gleanwork/api-client-go/models/components"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(components.Security{
            APIToken: apiclientgo.String(os.Getenv("GLEAN_API_TOKEN")),
        }),
    )

    res, err := s.Client.Search.QueryAsAdmin(ctx, &components.SearchRequest{
        TrackingToken: apiclientgo.String("trackingToken"),
        SourceDocument: &components.Document{
            Metadata: &components.DocumentMetadata{
                Datasource: apiclientgo.String("datasource"),
                ObjectType: apiclientgo.String("Feature Request"),
                Container: apiclientgo.String("container"),
                ParentID: apiclientgo.String("JIRA_EN-1337"),
                MimeType: apiclientgo.String("mimeType"),
                DocumentID: apiclientgo.String("documentId"),
                CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                Author: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    RelatedDocuments: []components.RelatedDocuments{
                        components.RelatedDocuments{
                            QuerySuggestion: &components.QuerySuggestion{
                                Query: "app:github type:pull author:mortimer",
                                SearchProviderInfo: &components.SearchProviderInfo{
                                    Name: apiclientgo.String("Google"),
                                    SearchLinkURLTemplate: apiclientgo.String("https://www.google.com/search?q={query}&hl=en"),
                                },
                                Label: apiclientgo.String("Mortimer's PRs"),
                                Datasource: apiclientgo.String("github"),
                                RequestOptions: &components.SearchRequestOptions{
                                    DatasourceFilter: apiclientgo.String("JIRA"),
                                    DatasourcesFilter: []string{
                                        "JIRA",
                                    },
                                    QueryOverridesFacetFilters: apiclientgo.Bool(true),
                                    FacetFilters: []components.FacetFilter{
                                        components.FacetFilter{
                                            FieldName: apiclientgo.String("type"),
                                            Values: []components.FacetFilterValue{
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.String("Spreadsheet"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.String("Presentation"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                            },
                                        },
                                    },
                                    FacetFilterSets: []components.FacetFilterSet{
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.String("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.String("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.String("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    FacetBucketSize: 629241,
                                    AuthTokens: []components.AuthToken{
                                        components.AuthToken{
                                            AccessToken: "123abc",
                                            Datasource: "gmail",
                                            Scope: apiclientgo.String("email profile https://www.googleapis.com/auth/gmail.readonly"),
                                            TokenType: apiclientgo.String("Bearer"),
                                            AuthUser: apiclientgo.String("1"),
                                        },
                                    },
                                },
                                Ranges: []components.TextRange{
                                    components.TextRange{
                                        StartIndex: 927545,
                                    },
                                },
                                InputDetails: &components.SearchRequestInputDetails{
                                    HasCopyPaste: apiclientgo.Bool(true),
                                },
                            },
                            Results: []components.SearchResult{
                                components.SearchResult{
                                    Title: apiclientgo.String("title"),
                                    URL: "https://example.com/foo/bar",
                                    NativeAppURL: apiclientgo.String("slack://foo/bar"),
                                    Snippets: []components.SearchResultSnippet{
                                        components.SearchResultSnippet{
                                            Snippet: "snippet",
                                            MimeType: apiclientgo.String("mimeType"),
                                        },
                                    },
                                },
                            },
                        },
                        components.RelatedDocuments{
                            QuerySuggestion: &components.QuerySuggestion{
                                Query: "app:github type:pull author:mortimer",
                                SearchProviderInfo: &components.SearchProviderInfo{
                                    Name: apiclientgo.String("Google"),
                                    SearchLinkURLTemplate: apiclientgo.String("https://www.google.com/search?q={query}&hl=en"),
                                },
                                Label: apiclientgo.String("Mortimer's PRs"),
                                Datasource: apiclientgo.String("github"),
                                RequestOptions: &components.SearchRequestOptions{
                                    DatasourceFilter: apiclientgo.String("JIRA"),
                                    DatasourcesFilter: []string{
                                        "JIRA",
                                    },
                                    QueryOverridesFacetFilters: apiclientgo.Bool(true),
                                    FacetFilters: []components.FacetFilter{
                                        components.FacetFilter{
                                            FieldName: apiclientgo.String("type"),
                                            Values: []components.FacetFilterValue{
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.String("Spreadsheet"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.String("Presentation"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                            },
                                        },
                                    },
                                    FacetFilterSets: []components.FacetFilterSet{
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.String("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.String("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.String("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    FacetBucketSize: 629241,
                                    AuthTokens: []components.AuthToken{
                                        components.AuthToken{
                                            AccessToken: "123abc",
                                            Datasource: "gmail",
                                            Scope: apiclientgo.String("email profile https://www.googleapis.com/auth/gmail.readonly"),
                                            TokenType: apiclientgo.String("Bearer"),
                                            AuthUser: apiclientgo.String("1"),
                                        },
                                    },
                                },
                                Ranges: []components.TextRange{
                                    components.TextRange{
                                        StartIndex: 927545,
                                    },
                                },
                                InputDetails: &components.SearchRequestInputDetails{
                                    HasCopyPaste: apiclientgo.Bool(true),
                                },
                            },
                            Results: []components.SearchResult{
                                components.SearchResult{
                                    Title: apiclientgo.String("title"),
                                    URL: "https://example.com/foo/bar",
                                    NativeAppURL: apiclientgo.String("slack://foo/bar"),
                                    Snippets: []components.SearchResultSnippet{
                                        components.SearchResultSnippet{
                                            Snippet: "snippet",
                                            MimeType: apiclientgo.String("mimeType"),
                                        },
                                    },
                                },
                            },
                        },
                        components.RelatedDocuments{
                            QuerySuggestion: &components.QuerySuggestion{
                                Query: "app:github type:pull author:mortimer",
                                SearchProviderInfo: &components.SearchProviderInfo{
                                    Name: apiclientgo.String("Google"),
                                    SearchLinkURLTemplate: apiclientgo.String("https://www.google.com/search?q={query}&hl=en"),
                                },
                                Label: apiclientgo.String("Mortimer's PRs"),
                                Datasource: apiclientgo.String("github"),
                                RequestOptions: &components.SearchRequestOptions{
                                    DatasourceFilter: apiclientgo.String("JIRA"),
                                    DatasourcesFilter: []string{
                                        "JIRA",
                                    },
                                    QueryOverridesFacetFilters: apiclientgo.Bool(true),
                                    FacetFilters: []components.FacetFilter{
                                        components.FacetFilter{
                                            FieldName: apiclientgo.String("type"),
                                            Values: []components.FacetFilterValue{
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.String("Spreadsheet"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.String("Presentation"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                            },
                                        },
                                    },
                                    FacetFilterSets: []components.FacetFilterSet{
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.String("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.String("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.String("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    FacetBucketSize: 629241,
                                    AuthTokens: []components.AuthToken{
                                        components.AuthToken{
                                            AccessToken: "123abc",
                                            Datasource: "gmail",
                                            Scope: apiclientgo.String("email profile https://www.googleapis.com/auth/gmail.readonly"),
                                            TokenType: apiclientgo.String("Bearer"),
                                            AuthUser: apiclientgo.String("1"),
                                        },
                                    },
                                },
                                Ranges: []components.TextRange{
                                    components.TextRange{
                                        StartIndex: 927545,
                                    },
                                },
                                InputDetails: &components.SearchRequestInputDetails{
                                    HasCopyPaste: apiclientgo.Bool(true),
                                },
                            },
                            Results: []components.SearchResult{
                                components.SearchResult{
                                    Title: apiclientgo.String("title"),
                                    URL: "https://example.com/foo/bar",
                                    NativeAppURL: apiclientgo.String("slack://foo/bar"),
                                    Snippets: []components.SearchResultSnippet{
                                        components.SearchResultSnippet{
                                            Snippet: "snippet",
                                            MimeType: apiclientgo.String("mimeType"),
                                        },
                                    },
                                },
                            },
                        },
                    },
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.String("Actor"),
                        Department: apiclientgo.String("Movies"),
                        Email: apiclientgo.String("george@example.com"),
                        Location: apiclientgo.String("Hollywood, CA"),
                        Phone: apiclientgo.String("6505551234"),
                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                        StartDate: types.MustNewDateFromString("2000-01-23"),
                        DatasourceProfile: []components.DatasourceProfile{
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                        },
                        QuerySuggestions: &components.QuerySuggestionList{
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.String("Mortimer's PRs"),
                                    Datasource: apiclientgo.String("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                            },
                        },
                        CustomFields: []components.CustomFieldData{
                            components.CustomFieldData{
                                Label: "<value>",
                                Values: []components.CustomFieldValue{},
                            },
                            components.CustomFieldData{
                                Label: "<value>",
                                Values: []components.CustomFieldValue{},
                            },
                            components.CustomFieldData{
                                Label: "<value>",
                                Values: []components.CustomFieldValue{},
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.String("deployment_name_new_hire"),
                                DisplayName: apiclientgo.String("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.String("#343CED"),
                                    Key: apiclientgo.String("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.String("user"),
                                },
                            },
                        },
                    },
                },
                Owner: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.String("Actor"),
                        Department: apiclientgo.String("Movies"),
                        Email: apiclientgo.String("george@example.com"),
                        Location: apiclientgo.String("Hollywood, CA"),
                        Phone: apiclientgo.String("6505551234"),
                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                        StartDate: types.MustNewDateFromString("2000-01-23"),
                        DatasourceProfile: []components.DatasourceProfile{
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                        },
                        QuerySuggestions: &components.QuerySuggestionList{
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.String("Mortimer's PRs"),
                                    Datasource: apiclientgo.String("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.String("deployment_name_new_hire"),
                                DisplayName: apiclientgo.String("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.String("#343CED"),
                                    Key: apiclientgo.String("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.String("user"),
                                },
                            },
                        },
                    },
                },
                MentionedPeople: []components.Person{
                    components.Person{
                        Name: "George Clooney",
                        ObfuscatedID: "abc123",
                    },
                },
                Components: []string{
                    "Backend",
                    "Networking",
                },
                Status: apiclientgo.String("[\"Done\"]"),
                Pins: []components.PinDocument{
                    components.PinDocument{
                        AudienceFilters: []components.FacetFilter{
                            components.FacetFilter{
                                FieldName: apiclientgo.String("type"),
                                Values: []components.FacetFilterValue{
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Spreadsheet"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Presentation"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                },
                            },
                        },
                        DocumentID: "<id>",
                        Attribution: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
                AssignedTo: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.String("Actor"),
                        Department: apiclientgo.String("Movies"),
                        Email: apiclientgo.String("george@example.com"),
                        Location: apiclientgo.String("Hollywood, CA"),
                        Phone: apiclientgo.String("6505551234"),
                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                        StartDate: types.MustNewDateFromString("2000-01-23"),
                        DatasourceProfile: []components.DatasourceProfile{
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                        },
                        QuerySuggestions: &components.QuerySuggestionList{
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.String("Mortimer's PRs"),
                                    Datasource: apiclientgo.String("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.String("deployment_name_new_hire"),
                                DisplayName: apiclientgo.String("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.String("#343CED"),
                                    Key: apiclientgo.String("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.String("user"),
                                },
                            },
                        },
                    },
                },
                UpdatedBy: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.String("Actor"),
                        Department: apiclientgo.String("Movies"),
                        Email: apiclientgo.String("george@example.com"),
                        Location: apiclientgo.String("Hollywood, CA"),
                        Phone: apiclientgo.String("6505551234"),
                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                        StartDate: types.MustNewDateFromString("2000-01-23"),
                        DatasourceProfile: []components.DatasourceProfile{
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                        },
                        QuerySuggestions: &components.QuerySuggestionList{
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.String("Mortimer's PRs"),
                                    Datasource: apiclientgo.String("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.String("deployment_name_new_hire"),
                                DisplayName: apiclientgo.String("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.String("#343CED"),
                                    Key: apiclientgo.String("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.String("user"),
                                },
                            },
                        },
                    },
                },
                Collections: []components.Collection{
                    components.Collection{
                        Name: "<value>",
                        Description: "gadzooks aside turret although as before exalted hospitalization option whether",
                        AddedRoles: []components.UserRoleSpecification{
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleVerifier,
                            },
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleVerifier,
                            },
                        },
                        RemovedRoles: []components.UserRoleSpecification{
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleViewer,
                            },
                        },
                        AudienceFilters: []components.FacetFilter{
                            components.FacetFilter{
                                FieldName: apiclientgo.String("type"),
                                Values: []components.FacetFilterValue{
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Spreadsheet"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Presentation"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                },
                            },
                        },
                        ID: 740835,
                        Creator: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        Items: []components.CollectionItem{
                            components.CollectionItem{
                                CollectionID: 177661,
                                CreatedBy: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Shortcut: &components.Shortcut{
                                    InputAlias: "<value>",
                                    CreatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    UpdatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    Roles: []components.UserRoleSpecification{
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleVerifier,
                                        },
                                    },
                                },
                                ItemType: components.CollectionItemItemTypeText,
                            },
                            components.CollectionItem{
                                CollectionID: 177661,
                                CreatedBy: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Shortcut: &components.Shortcut{
                                    InputAlias: "<value>",
                                    CreatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    UpdatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    Roles: []components.UserRoleSpecification{
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleVerifier,
                                        },
                                    },
                                },
                                ItemType: components.CollectionItemItemTypeText,
                            },
                            components.CollectionItem{
                                CollectionID: 177661,
                                CreatedBy: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Shortcut: &components.Shortcut{
                                    InputAlias: "<value>",
                                    CreatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    UpdatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    Roles: []components.UserRoleSpecification{
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleVerifier,
                                        },
                                    },
                                },
                                ItemType: components.CollectionItemItemTypeText,
                            },
                        },
                    },
                },
                Interactions: &components.DocumentInteractions{
                    Reacts: []components.Reaction{
                        components.Reaction{
                            Reactors: []components.Person{
                                components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                },
                            },
                        },
                        components.Reaction{
                            Reactors: []components.Person{
                                components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                },
                            },
                        },
                    },
                    Shares: []components.Share{
                        components.Share{
                            NumDaysAgo: 867476,
                            Sharer: &components.Person{
                                Name: "George Clooney",
                                ObfuscatedID: "abc123",
                                Metadata: &components.PersonMetadata{
                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                    Title: apiclientgo.String("Actor"),
                                    Department: apiclientgo.String("Movies"),
                                    Email: apiclientgo.String("george@example.com"),
                                    Location: apiclientgo.String("Hollywood, CA"),
                                    Phone: apiclientgo.String("6505551234"),
                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                    DatasourceProfile: []components.DatasourceProfile{
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                    },
                                    QuerySuggestions: &components.QuerySuggestionList{
                                        Suggestions: []components.QuerySuggestion{
                                            components.QuerySuggestion{
                                                Query: "app:github type:pull author:mortimer",
                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                Datasource: apiclientgo.String("github"),
                                            },
                                        },
                                    },
                                    InviteInfo: &components.InviteInfo{
                                        Invites: []components.ChannelInviteInfo{
                                            components.ChannelInviteInfo{},
                                        },
                                    },
                                    Badges: []components.Badge{
                                        components.Badge{
                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                            DisplayName: apiclientgo.String("New hire"),
                                            IconConfig: &components.IconConfig{
                                                Color: apiclientgo.String("#343CED"),
                                                Key: apiclientgo.String("person_icon"),
                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                Name: apiclientgo.String("user"),
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        components.Share{
                            NumDaysAgo: 867476,
                            Sharer: &components.Person{
                                Name: "George Clooney",
                                ObfuscatedID: "abc123",
                                Metadata: &components.PersonMetadata{
                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                    Title: apiclientgo.String("Actor"),
                                    Department: apiclientgo.String("Movies"),
                                    Email: apiclientgo.String("george@example.com"),
                                    Location: apiclientgo.String("Hollywood, CA"),
                                    Phone: apiclientgo.String("6505551234"),
                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                    DatasourceProfile: []components.DatasourceProfile{
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                    },
                                    QuerySuggestions: &components.QuerySuggestionList{
                                        Suggestions: []components.QuerySuggestion{
                                            components.QuerySuggestion{
                                                Query: "app:github type:pull author:mortimer",
                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                Datasource: apiclientgo.String("github"),
                                            },
                                        },
                                    },
                                    InviteInfo: &components.InviteInfo{
                                        Invites: []components.ChannelInviteInfo{
                                            components.ChannelInviteInfo{},
                                        },
                                    },
                                    Badges: []components.Badge{
                                        components.Badge{
                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                            DisplayName: apiclientgo.String("New hire"),
                                            IconConfig: &components.IconConfig{
                                                Color: apiclientgo.String("#343CED"),
                                                Key: apiclientgo.String("person_icon"),
                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                Name: apiclientgo.String("user"),
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
                Verification: &components.Verification{
                    State: components.StateDeprecated,
                    Metadata: &components.VerificationMetadata{
                        LastVerifier: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        Reminders: []components.Reminder{
                            components.Reminder{
                                Assignee: components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Requestor: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                RemindAt: 935874,
                            },
                            components.Reminder{
                                Assignee: components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Requestor: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                RemindAt: 935874,
                            },
                            components.Reminder{
                                Assignee: components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Requestor: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                RemindAt: 935874,
                            },
                        },
                        LastReminder: &components.Reminder{
                            Assignee: components.Person{
                                Name: "George Clooney",
                                ObfuscatedID: "abc123",
                                Metadata: &components.PersonMetadata{
                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                    Title: apiclientgo.String("Actor"),
                                    Department: apiclientgo.String("Movies"),
                                    Email: apiclientgo.String("george@example.com"),
                                    Location: apiclientgo.String("Hollywood, CA"),
                                    Phone: apiclientgo.String("6505551234"),
                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                    DatasourceProfile: []components.DatasourceProfile{
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                    },
                                    QuerySuggestions: &components.QuerySuggestionList{
                                        Suggestions: []components.QuerySuggestion{
                                            components.QuerySuggestion{
                                                Query: "app:github type:pull author:mortimer",
                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                Datasource: apiclientgo.String("github"),
                                            },
                                        },
                                    },
                                    InviteInfo: &components.InviteInfo{
                                        Invites: []components.ChannelInviteInfo{
                                            components.ChannelInviteInfo{},
                                        },
                                    },
                                    Badges: []components.Badge{
                                        components.Badge{
                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                            DisplayName: apiclientgo.String("New hire"),
                                            IconConfig: &components.IconConfig{
                                                Color: apiclientgo.String("#343CED"),
                                                Key: apiclientgo.String("person_icon"),
                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                Name: apiclientgo.String("user"),
                                            },
                                        },
                                    },
                                },
                            },
                            Requestor: &components.Person{
                                Name: "George Clooney",
                                ObfuscatedID: "abc123",
                                Metadata: &components.PersonMetadata{
                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                    Title: apiclientgo.String("Actor"),
                                    Department: apiclientgo.String("Movies"),
                                    Email: apiclientgo.String("george@example.com"),
                                    Location: apiclientgo.String("Hollywood, CA"),
                                    Phone: apiclientgo.String("6505551234"),
                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                    DatasourceProfile: []components.DatasourceProfile{
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                    },
                                    QuerySuggestions: &components.QuerySuggestionList{
                                        Suggestions: []components.QuerySuggestion{
                                            components.QuerySuggestion{
                                                Query: "app:github type:pull author:mortimer",
                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                Datasource: apiclientgo.String("github"),
                                            },
                                        },
                                    },
                                    InviteInfo: &components.InviteInfo{
                                        Invites: []components.ChannelInviteInfo{
                                            components.ChannelInviteInfo{},
                                        },
                                    },
                                    Badges: []components.Badge{
                                        components.Badge{
                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                            DisplayName: apiclientgo.String("New hire"),
                                            IconConfig: &components.IconConfig{
                                                Color: apiclientgo.String("#343CED"),
                                                Key: apiclientgo.String("person_icon"),
                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                Name: apiclientgo.String("user"),
                                            },
                                        },
                                    },
                                },
                            },
                            RemindAt: 346805,
                        },
                        CandidateVerifiers: []components.Person{
                            components.Person{
                                Name: "George Clooney",
                                ObfuscatedID: "abc123",
                            },
                        },
                    },
                },
                Shortcuts: []components.Shortcut{
                    components.Shortcut{
                        InputAlias: "<value>",
                        CreatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                    },
                    components.Shortcut{
                        InputAlias: "<value>",
                        CreatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                    },
                    components.Shortcut{
                        InputAlias: "<value>",
                        CreatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
                CustomData: map[string]components.CustomDataValue{
                    "someCustomField": components.CustomDataValue{},
                },
                ContactPerson: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.String("Actor"),
                        Department: apiclientgo.String("Movies"),
                        Email: apiclientgo.String("george@example.com"),
                        Location: apiclientgo.String("Hollywood, CA"),
                        Phone: apiclientgo.String("6505551234"),
                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                        StartDate: types.MustNewDateFromString("2000-01-23"),
                        DatasourceProfile: []components.DatasourceProfile{
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                        },
                        QuerySuggestions: &components.QuerySuggestionList{
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.String("Mortimer's PRs"),
                                    Datasource: apiclientgo.String("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.String("deployment_name_new_hire"),
                                DisplayName: apiclientgo.String("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.String("#343CED"),
                                    Key: apiclientgo.String("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.String("user"),
                                },
                            },
                        },
                    },
                },
            },
        },
        PageSize: apiclientgo.Int64(100),
        MaxSnippetSize: apiclientgo.Int64(400),
        Query: "vacation policy",
        InputDetails: &components.SearchRequestInputDetails{
            HasCopyPaste: apiclientgo.Bool(true),
        },
        RequestOptions: &components.SearchRequestOptions{
            DatasourceFilter: apiclientgo.String("JIRA"),
            DatasourcesFilter: []string{
                "JIRA",
            },
            QueryOverridesFacetFilters: apiclientgo.Bool(true),
            FacetFilters: []components.FacetFilter{
                components.FacetFilter{
                    FieldName: apiclientgo.String("type"),
                    Values: []components.FacetFilterValue{
                        components.FacetFilterValue{
                            Value: apiclientgo.String("Spreadsheet"),
                            RelationType: components.RelationTypeEquals.ToPointer(),
                        },
                        components.FacetFilterValue{
                            Value: apiclientgo.String("Presentation"),
                            RelationType: components.RelationTypeEquals.ToPointer(),
                        },
                    },
                },
            },
            FacetFilterSets: []components.FacetFilterSet{
                components.FacetFilterSet{
                    Filters: []components.FacetFilter{
                        components.FacetFilter{
                            FieldName: apiclientgo.String("type"),
                            Values: []components.FacetFilterValue{
                                components.FacetFilterValue{
                                    Value: apiclientgo.String("Spreadsheet"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                                components.FacetFilterValue{
                                    Value: apiclientgo.String("Presentation"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                            },
                        },
                    },
                },
                components.FacetFilterSet{
                    Filters: []components.FacetFilter{
                        components.FacetFilter{
                            FieldName: apiclientgo.String("type"),
                            Values: []components.FacetFilterValue{
                                components.FacetFilterValue{
                                    Value: apiclientgo.String("Spreadsheet"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                                components.FacetFilterValue{
                                    Value: apiclientgo.String("Presentation"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                            },
                        },
                    },
                },
            },
            FacetBucketSize: 421489,
            AuthTokens: []components.AuthToken{
                components.AuthToken{
                    AccessToken: "123abc",
                    Datasource: "gmail",
                    Scope: apiclientgo.String("email profile https://www.googleapis.com/auth/gmail.readonly"),
                    TokenType: apiclientgo.String("Bearer"),
                    AuthUser: apiclientgo.String("1"),
                },
            },
        },
        TimeoutMillis: apiclientgo.Int64(5000),
        People: []components.Person{
            components.Person{
                Name: "George Clooney",
                ObfuscatedID: "abc123",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [components.SearchRequest](../../models/components/searchrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.AdminsearchResponse](../../models/operations/adminsearchresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| apierrors.GleanDataError | 403, 422                 | application/json         |
| apierrors.APIError       | 4XX, 5XX                 | \*/\*                    |

## Autocomplete

Retrieve query suggestions, operators and documents for the given partially typed query.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/gleanwork/api-client-go/models/components"
	apiclientgo "github.com/gleanwork/api-client-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(components.Security{
            APIToken: apiclientgo.String(os.Getenv("GLEAN_API_TOKEN")),
        }),
    )

    res, err := s.Client.Search.Autocomplete(ctx, components.AutocompleteRequest{
        TrackingToken: apiclientgo.String("trackingToken"),
        Query: apiclientgo.String("San Fra"),
        Datasource: apiclientgo.String("GDRIVE"),
        ResultSize: apiclientgo.Int64(10),
        AuthTokens: []components.AuthToken{
            components.AuthToken{
                AccessToken: "123abc",
                Datasource: "gmail",
                Scope: apiclientgo.String("email profile https://www.googleapis.com/auth/gmail.readonly"),
                TokenType: apiclientgo.String("Bearer"),
                AuthUser: apiclientgo.String("1"),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AutocompleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.AutocompleteRequest](../../models/components/autocompleterequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.AutocompleteResponse](../../models/operations/autocompleteresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RetrieveFeed

The personalized feed/home includes different types of contents including suggestions, recents, calendar events and many more.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/gleanwork/api-client-go/models/components"
	apiclientgo "github.com/gleanwork/api-client-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(components.Security{
            APIToken: apiclientgo.String(os.Getenv("GLEAN_API_TOKEN")),
        }),
    )

    res, err := s.Client.Search.RetrieveFeed(ctx, components.FeedRequest{
        TimeoutMillis: apiclientgo.Int64(5000),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.FeedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [components.FeedRequest](../../models/components/feedrequest.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |

### Response

**[*operations.FeedResponse](../../models/operations/feedresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Recommendations

Retrieve recommended documents for the given URL or Glean Document ID.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/gleanwork/api-client-go/models/components"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(components.Security{
            APIToken: apiclientgo.String(os.Getenv("GLEAN_API_TOKEN")),
        }),
    )

    res, err := s.Client.Search.Recommendations(ctx, &components.RecommendationsRequest{
        SourceDocument: &components.Document{
            Metadata: &components.DocumentMetadata{
                Datasource: apiclientgo.String("datasource"),
                ObjectType: apiclientgo.String("Feature Request"),
                Container: apiclientgo.String("container"),
                ParentID: apiclientgo.String("JIRA_EN-1337"),
                MimeType: apiclientgo.String("mimeType"),
                DocumentID: apiclientgo.String("documentId"),
                CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                Author: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    RelatedDocuments: []components.RelatedDocuments{
                        components.RelatedDocuments{
                            QuerySuggestion: &components.QuerySuggestion{
                                Query: "app:github type:pull author:mortimer",
                                SearchProviderInfo: &components.SearchProviderInfo{
                                    Name: apiclientgo.String("Google"),
                                    SearchLinkURLTemplate: apiclientgo.String("https://www.google.com/search?q={query}&hl=en"),
                                },
                                Label: apiclientgo.String("Mortimer's PRs"),
                                Datasource: apiclientgo.String("github"),
                                RequestOptions: &components.SearchRequestOptions{
                                    DatasourceFilter: apiclientgo.String("JIRA"),
                                    DatasourcesFilter: []string{
                                        "JIRA",
                                    },
                                    QueryOverridesFacetFilters: apiclientgo.Bool(true),
                                    FacetFilters: []components.FacetFilter{
                                        components.FacetFilter{
                                            FieldName: apiclientgo.String("type"),
                                            Values: []components.FacetFilterValue{
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.String("Spreadsheet"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.String("Presentation"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                            },
                                        },
                                    },
                                    FacetFilterSets: []components.FacetFilterSet{
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.String("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.String("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.String("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    FacetBucketSize: 711201,
                                    AuthTokens: []components.AuthToken{
                                        components.AuthToken{
                                            AccessToken: "123abc",
                                            Datasource: "gmail",
                                            Scope: apiclientgo.String("email profile https://www.googleapis.com/auth/gmail.readonly"),
                                            TokenType: apiclientgo.String("Bearer"),
                                            AuthUser: apiclientgo.String("1"),
                                        },
                                    },
                                },
                                Ranges: []components.TextRange{
                                    components.TextRange{
                                        StartIndex: 707124,
                                    },
                                    components.TextRange{
                                        StartIndex: 707124,
                                    },
                                },
                                InputDetails: &components.SearchRequestInputDetails{
                                    HasCopyPaste: apiclientgo.Bool(true),
                                },
                            },
                            Results: []components.SearchResult{
                                components.SearchResult{
                                    Title: apiclientgo.String("title"),
                                    URL: "https://example.com/foo/bar",
                                    NativeAppURL: apiclientgo.String("slack://foo/bar"),
                                    Snippets: []components.SearchResultSnippet{
                                        components.SearchResultSnippet{
                                            Snippet: "snippet",
                                            MimeType: apiclientgo.String("mimeType"),
                                        },
                                    },
                                },
                            },
                        },
                    },
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.String("Actor"),
                        Department: apiclientgo.String("Movies"),
                        Email: apiclientgo.String("george@example.com"),
                        Location: apiclientgo.String("Hollywood, CA"),
                        Phone: apiclientgo.String("6505551234"),
                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                        StartDate: types.MustNewDateFromString("2000-01-23"),
                        DatasourceProfile: []components.DatasourceProfile{
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                        },
                        QuerySuggestions: &components.QuerySuggestionList{
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.String("Mortimer's PRs"),
                                    Datasource: apiclientgo.String("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                            },
                        },
                        CustomFields: []components.CustomFieldData{
                            components.CustomFieldData{
                                Label: "<value>",
                                Values: []components.CustomFieldValue{
                                    components.CreateCustomFieldValueCustomFieldValueStr(
                                        components.CustomFieldValueStr{},
                                    ),
                                },
                            },
                            components.CustomFieldData{
                                Label: "<value>",
                                Values: []components.CustomFieldValue{
                                    components.CreateCustomFieldValueCustomFieldValueStr(
                                        components.CustomFieldValueStr{},
                                    ),
                                },
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.String("deployment_name_new_hire"),
                                DisplayName: apiclientgo.String("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.String("#343CED"),
                                    Key: apiclientgo.String("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.String("user"),
                                },
                            },
                        },
                    },
                },
                Owner: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.String("Actor"),
                        Department: apiclientgo.String("Movies"),
                        Email: apiclientgo.String("george@example.com"),
                        Location: apiclientgo.String("Hollywood, CA"),
                        Phone: apiclientgo.String("6505551234"),
                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                        StartDate: types.MustNewDateFromString("2000-01-23"),
                        DatasourceProfile: []components.DatasourceProfile{
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                        },
                        QuerySuggestions: &components.QuerySuggestionList{
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.String("Mortimer's PRs"),
                                    Datasource: apiclientgo.String("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.String("deployment_name_new_hire"),
                                DisplayName: apiclientgo.String("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.String("#343CED"),
                                    Key: apiclientgo.String("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.String("user"),
                                },
                            },
                        },
                    },
                },
                MentionedPeople: []components.Person{
                    components.Person{
                        Name: "George Clooney",
                        ObfuscatedID: "abc123",
                    },
                },
                Components: []string{
                    "Backend",
                    "Networking",
                },
                Status: apiclientgo.String("[\"Done\"]"),
                Pins: []components.PinDocument{
                    components.PinDocument{
                        AudienceFilters: []components.FacetFilter{
                            components.FacetFilter{
                                FieldName: apiclientgo.String("type"),
                                Values: []components.FacetFilterValue{
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Spreadsheet"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Presentation"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                },
                            },
                        },
                        DocumentID: "<id>",
                        Attribution: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                    },
                    components.PinDocument{
                        AudienceFilters: []components.FacetFilter{
                            components.FacetFilter{
                                FieldName: apiclientgo.String("type"),
                                Values: []components.FacetFilterValue{
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Spreadsheet"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Presentation"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                },
                            },
                        },
                        DocumentID: "<id>",
                        Attribution: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                    },
                    components.PinDocument{
                        AudienceFilters: []components.FacetFilter{
                            components.FacetFilter{
                                FieldName: apiclientgo.String("type"),
                                Values: []components.FacetFilterValue{
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Spreadsheet"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Presentation"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                },
                            },
                        },
                        DocumentID: "<id>",
                        Attribution: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
                AssignedTo: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.String("Actor"),
                        Department: apiclientgo.String("Movies"),
                        Email: apiclientgo.String("george@example.com"),
                        Location: apiclientgo.String("Hollywood, CA"),
                        Phone: apiclientgo.String("6505551234"),
                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                        StartDate: types.MustNewDateFromString("2000-01-23"),
                        DatasourceProfile: []components.DatasourceProfile{
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                        },
                        QuerySuggestions: &components.QuerySuggestionList{
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.String("Mortimer's PRs"),
                                    Datasource: apiclientgo.String("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.String("deployment_name_new_hire"),
                                DisplayName: apiclientgo.String("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.String("#343CED"),
                                    Key: apiclientgo.String("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.String("user"),
                                },
                            },
                        },
                    },
                },
                UpdatedBy: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.String("Actor"),
                        Department: apiclientgo.String("Movies"),
                        Email: apiclientgo.String("george@example.com"),
                        Location: apiclientgo.String("Hollywood, CA"),
                        Phone: apiclientgo.String("6505551234"),
                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                        StartDate: types.MustNewDateFromString("2000-01-23"),
                        DatasourceProfile: []components.DatasourceProfile{
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                        },
                        QuerySuggestions: &components.QuerySuggestionList{
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.String("Mortimer's PRs"),
                                    Datasource: apiclientgo.String("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.String("deployment_name_new_hire"),
                                DisplayName: apiclientgo.String("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.String("#343CED"),
                                    Key: apiclientgo.String("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.String("user"),
                                },
                            },
                        },
                    },
                },
                Collections: []components.Collection{
                    components.Collection{
                        Name: "<value>",
                        Description: "hence why at epic only supposing",
                        AddedRoles: []components.UserRoleSpecification{
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleVerifier,
                            },
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleVerifier,
                            },
                        },
                        RemovedRoles: []components.UserRoleSpecification{
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleVerifier,
                            },
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleVerifier,
                            },
                        },
                        AudienceFilters: []components.FacetFilter{
                            components.FacetFilter{
                                FieldName: apiclientgo.String("type"),
                                Values: []components.FacetFilterValue{
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Spreadsheet"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Presentation"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                },
                            },
                        },
                        ID: 253796,
                        Creator: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        Items: []components.CollectionItem{
                            components.CollectionItem{
                                CollectionID: 94361,
                                CreatedBy: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Shortcut: &components.Shortcut{
                                    InputAlias: "<value>",
                                    CreatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    UpdatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    Roles: []components.UserRoleSpecification{
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleVerifier,
                                        },
                                    },
                                },
                                ItemType: components.CollectionItemItemTypeURL,
                            },
                        },
                    },
                },
                Interactions: &components.DocumentInteractions{
                    Reacts: []components.Reaction{
                        components.Reaction{
                            Reactors: []components.Person{
                                components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                },
                            },
                        },
                        components.Reaction{
                            Reactors: []components.Person{
                                components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                },
                            },
                        },
                        components.Reaction{
                            Reactors: []components.Person{
                                components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                },
                            },
                        },
                    },
                    Shares: []components.Share{
                        components.Share{
                            NumDaysAgo: 652391,
                            Sharer: &components.Person{
                                Name: "George Clooney",
                                ObfuscatedID: "abc123",
                                Metadata: &components.PersonMetadata{
                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                    Title: apiclientgo.String("Actor"),
                                    Department: apiclientgo.String("Movies"),
                                    Email: apiclientgo.String("george@example.com"),
                                    Location: apiclientgo.String("Hollywood, CA"),
                                    Phone: apiclientgo.String("6505551234"),
                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                    DatasourceProfile: []components.DatasourceProfile{
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                    },
                                    QuerySuggestions: &components.QuerySuggestionList{
                                        Suggestions: []components.QuerySuggestion{
                                            components.QuerySuggestion{
                                                Query: "app:github type:pull author:mortimer",
                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                Datasource: apiclientgo.String("github"),
                                            },
                                        },
                                    },
                                    InviteInfo: &components.InviteInfo{
                                        Invites: []components.ChannelInviteInfo{
                                            components.ChannelInviteInfo{},
                                        },
                                    },
                                    Badges: []components.Badge{
                                        components.Badge{
                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                            DisplayName: apiclientgo.String("New hire"),
                                            IconConfig: &components.IconConfig{
                                                Color: apiclientgo.String("#343CED"),
                                                Key: apiclientgo.String("person_icon"),
                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                Name: apiclientgo.String("user"),
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
                Verification: &components.Verification{
                    State: components.StateDeprecated,
                    Metadata: &components.VerificationMetadata{
                        LastVerifier: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        Reminders: []components.Reminder{
                            components.Reminder{
                                Assignee: components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Requestor: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                RemindAt: 611121,
                            },
                            components.Reminder{
                                Assignee: components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Requestor: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                RemindAt: 611121,
                            },
                            components.Reminder{
                                Assignee: components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Requestor: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                RemindAt: 611121,
                            },
                        },
                        LastReminder: &components.Reminder{
                            Assignee: components.Person{
                                Name: "George Clooney",
                                ObfuscatedID: "abc123",
                                Metadata: &components.PersonMetadata{
                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                    Title: apiclientgo.String("Actor"),
                                    Department: apiclientgo.String("Movies"),
                                    Email: apiclientgo.String("george@example.com"),
                                    Location: apiclientgo.String("Hollywood, CA"),
                                    Phone: apiclientgo.String("6505551234"),
                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                    DatasourceProfile: []components.DatasourceProfile{
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                    },
                                    QuerySuggestions: &components.QuerySuggestionList{
                                        Suggestions: []components.QuerySuggestion{
                                            components.QuerySuggestion{
                                                Query: "app:github type:pull author:mortimer",
                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                Datasource: apiclientgo.String("github"),
                                            },
                                        },
                                    },
                                    InviteInfo: &components.InviteInfo{
                                        Invites: []components.ChannelInviteInfo{
                                            components.ChannelInviteInfo{},
                                        },
                                    },
                                    Badges: []components.Badge{
                                        components.Badge{
                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                            DisplayName: apiclientgo.String("New hire"),
                                            IconConfig: &components.IconConfig{
                                                Color: apiclientgo.String("#343CED"),
                                                Key: apiclientgo.String("person_icon"),
                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                Name: apiclientgo.String("user"),
                                            },
                                        },
                                    },
                                },
                            },
                            Requestor: &components.Person{
                                Name: "George Clooney",
                                ObfuscatedID: "abc123",
                                Metadata: &components.PersonMetadata{
                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                    Title: apiclientgo.String("Actor"),
                                    Department: apiclientgo.String("Movies"),
                                    Email: apiclientgo.String("george@example.com"),
                                    Location: apiclientgo.String("Hollywood, CA"),
                                    Phone: apiclientgo.String("6505551234"),
                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                    DatasourceProfile: []components.DatasourceProfile{
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                    },
                                    QuerySuggestions: &components.QuerySuggestionList{
                                        Suggestions: []components.QuerySuggestion{
                                            components.QuerySuggestion{
                                                Query: "app:github type:pull author:mortimer",
                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                Datasource: apiclientgo.String("github"),
                                            },
                                        },
                                    },
                                    InviteInfo: &components.InviteInfo{
                                        Invites: []components.ChannelInviteInfo{
                                            components.ChannelInviteInfo{},
                                        },
                                    },
                                    Badges: []components.Badge{
                                        components.Badge{
                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                            DisplayName: apiclientgo.String("New hire"),
                                            IconConfig: &components.IconConfig{
                                                Color: apiclientgo.String("#343CED"),
                                                Key: apiclientgo.String("person_icon"),
                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                Name: apiclientgo.String("user"),
                                            },
                                        },
                                    },
                                },
                            },
                            RemindAt: 148513,
                        },
                        CandidateVerifiers: []components.Person{
                            components.Person{
                                Name: "George Clooney",
                                ObfuscatedID: "abc123",
                            },
                        },
                    },
                },
                Shortcuts: []components.Shortcut{
                    components.Shortcut{
                        InputAlias: "<value>",
                        CreatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                    },
                    components.Shortcut{
                        InputAlias: "<value>",
                        CreatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
                CustomData: map[string]components.CustomDataValue{
                    "someCustomField": components.CustomDataValue{},
                },
                ContactPerson: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.String("Actor"),
                        Department: apiclientgo.String("Movies"),
                        Email: apiclientgo.String("george@example.com"),
                        Location: apiclientgo.String("Hollywood, CA"),
                        Phone: apiclientgo.String("6505551234"),
                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                        StartDate: types.MustNewDateFromString("2000-01-23"),
                        DatasourceProfile: []components.DatasourceProfile{
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                        },
                        QuerySuggestions: &components.QuerySuggestionList{
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.String("Mortimer's PRs"),
                                    Datasource: apiclientgo.String("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.String("deployment_name_new_hire"),
                                DisplayName: apiclientgo.String("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.String("#343CED"),
                                    Key: apiclientgo.String("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.String("user"),
                                },
                            },
                        },
                    },
                },
            },
        },
        PageSize: apiclientgo.Int64(100),
        MaxSnippetSize: apiclientgo.Int64(400),
        RequestOptions: &components.RecommendationsRequestOptions{
            FacetFilterSets: []components.FacetFilterSet{
                components.FacetFilterSet{
                    Filters: []components.FacetFilter{
                        components.FacetFilter{
                            FieldName: apiclientgo.String("type"),
                            Values: []components.FacetFilterValue{
                                components.FacetFilterValue{
                                    Value: apiclientgo.String("Spreadsheet"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                                components.FacetFilterValue{
                                    Value: apiclientgo.String("Presentation"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                            },
                        },
                    },
                },
                components.FacetFilterSet{
                    Filters: []components.FacetFilter{
                        components.FacetFilter{
                            FieldName: apiclientgo.String("type"),
                            Values: []components.FacetFilterValue{
                                components.FacetFilterValue{
                                    Value: apiclientgo.String("Spreadsheet"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                                components.FacetFilterValue{
                                    Value: apiclientgo.String("Presentation"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                            },
                        },
                    },
                },
                components.FacetFilterSet{
                    Filters: []components.FacetFilter{
                        components.FacetFilter{
                            FieldName: apiclientgo.String("type"),
                            Values: []components.FacetFilterValue{
                                components.FacetFilterValue{
                                    Value: apiclientgo.String("Spreadsheet"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                                components.FacetFilterValue{
                                    Value: apiclientgo.String("Presentation"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                            },
                        },
                    },
                },
            },
            Context: &components.Document{
                Metadata: &components.DocumentMetadata{
                    Datasource: apiclientgo.String("datasource"),
                    ObjectType: apiclientgo.String("Feature Request"),
                    Container: apiclientgo.String("container"),
                    ParentID: apiclientgo.String("JIRA_EN-1337"),
                    MimeType: apiclientgo.String("mimeType"),
                    DocumentID: apiclientgo.String("documentId"),
                    CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                    UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                    Author: &components.Person{
                        Name: "George Clooney",
                        ObfuscatedID: "abc123",
                        Metadata: &components.PersonMetadata{
                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                            Title: apiclientgo.String("Actor"),
                            Department: apiclientgo.String("Movies"),
                            Email: apiclientgo.String("george@example.com"),
                            Location: apiclientgo.String("Hollywood, CA"),
                            Phone: apiclientgo.String("6505551234"),
                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                            StartDate: types.MustNewDateFromString("2000-01-23"),
                            DatasourceProfile: []components.DatasourceProfile{
                                components.DatasourceProfile{
                                    Datasource: "github",
                                    Handle: "<value>",
                                },
                                components.DatasourceProfile{
                                    Datasource: "github",
                                    Handle: "<value>",
                                },
                            },
                            QuerySuggestions: &components.QuerySuggestionList{
                                Suggestions: []components.QuerySuggestion{
                                    components.QuerySuggestion{
                                        Query: "app:github type:pull author:mortimer",
                                        Label: apiclientgo.String("Mortimer's PRs"),
                                        Datasource: apiclientgo.String("github"),
                                    },
                                },
                            },
                            InviteInfo: &components.InviteInfo{
                                Invites: []components.ChannelInviteInfo{
                                    components.ChannelInviteInfo{},
                                },
                            },
                            Badges: []components.Badge{
                                components.Badge{
                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                    DisplayName: apiclientgo.String("New hire"),
                                    IconConfig: &components.IconConfig{
                                        Color: apiclientgo.String("#343CED"),
                                        Key: apiclientgo.String("person_icon"),
                                        IconType: components.IconTypeGlyph.ToPointer(),
                                        Name: apiclientgo.String("user"),
                                    },
                                },
                            },
                        },
                    },
                    Owner: &components.Person{
                        Name: "George Clooney",
                        ObfuscatedID: "abc123",
                        Metadata: &components.PersonMetadata{
                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                            Title: apiclientgo.String("Actor"),
                            Department: apiclientgo.String("Movies"),
                            Email: apiclientgo.String("george@example.com"),
                            Location: apiclientgo.String("Hollywood, CA"),
                            Phone: apiclientgo.String("6505551234"),
                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                            StartDate: types.MustNewDateFromString("2000-01-23"),
                            DatasourceProfile: []components.DatasourceProfile{
                                components.DatasourceProfile{
                                    Datasource: "github",
                                    Handle: "<value>",
                                },
                                components.DatasourceProfile{
                                    Datasource: "github",
                                    Handle: "<value>",
                                },
                                components.DatasourceProfile{
                                    Datasource: "github",
                                    Handle: "<value>",
                                },
                            },
                            QuerySuggestions: &components.QuerySuggestionList{
                                Suggestions: []components.QuerySuggestion{
                                    components.QuerySuggestion{
                                        Query: "app:github type:pull author:mortimer",
                                        Label: apiclientgo.String("Mortimer's PRs"),
                                        Datasource: apiclientgo.String("github"),
                                    },
                                },
                            },
                            InviteInfo: &components.InviteInfo{
                                Invites: []components.ChannelInviteInfo{
                                    components.ChannelInviteInfo{},
                                },
                            },
                            Badges: []components.Badge{
                                components.Badge{
                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                    DisplayName: apiclientgo.String("New hire"),
                                    IconConfig: &components.IconConfig{
                                        Color: apiclientgo.String("#343CED"),
                                        Key: apiclientgo.String("person_icon"),
                                        IconType: components.IconTypeGlyph.ToPointer(),
                                        Name: apiclientgo.String("user"),
                                    },
                                },
                            },
                        },
                    },
                    Components: []string{
                        "Backend",
                        "Networking",
                    },
                    Status: apiclientgo.String("[\"Done\"]"),
                    AssignedTo: &components.Person{
                        Name: "George Clooney",
                        ObfuscatedID: "abc123",
                        Metadata: &components.PersonMetadata{
                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                            Title: apiclientgo.String("Actor"),
                            Department: apiclientgo.String("Movies"),
                            Email: apiclientgo.String("george@example.com"),
                            Location: apiclientgo.String("Hollywood, CA"),
                            Phone: apiclientgo.String("6505551234"),
                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                            StartDate: types.MustNewDateFromString("2000-01-23"),
                            DatasourceProfile: []components.DatasourceProfile{
                                components.DatasourceProfile{
                                    Datasource: "github",
                                    Handle: "<value>",
                                },
                                components.DatasourceProfile{
                                    Datasource: "github",
                                    Handle: "<value>",
                                },
                            },
                            QuerySuggestions: &components.QuerySuggestionList{
                                Suggestions: []components.QuerySuggestion{
                                    components.QuerySuggestion{
                                        Query: "app:github type:pull author:mortimer",
                                        Label: apiclientgo.String("Mortimer's PRs"),
                                        Datasource: apiclientgo.String("github"),
                                    },
                                },
                            },
                            InviteInfo: &components.InviteInfo{
                                Invites: []components.ChannelInviteInfo{
                                    components.ChannelInviteInfo{},
                                },
                            },
                            Badges: []components.Badge{
                                components.Badge{
                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                    DisplayName: apiclientgo.String("New hire"),
                                    IconConfig: &components.IconConfig{
                                        Color: apiclientgo.String("#343CED"),
                                        Key: apiclientgo.String("person_icon"),
                                        IconType: components.IconTypeGlyph.ToPointer(),
                                        Name: apiclientgo.String("user"),
                                    },
                                },
                            },
                        },
                    },
                    UpdatedBy: &components.Person{
                        Name: "George Clooney",
                        ObfuscatedID: "abc123",
                        Metadata: &components.PersonMetadata{
                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                            Title: apiclientgo.String("Actor"),
                            Department: apiclientgo.String("Movies"),
                            Email: apiclientgo.String("george@example.com"),
                            Location: apiclientgo.String("Hollywood, CA"),
                            Phone: apiclientgo.String("6505551234"),
                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                            StartDate: types.MustNewDateFromString("2000-01-23"),
                            DatasourceProfile: []components.DatasourceProfile{
                                components.DatasourceProfile{
                                    Datasource: "github",
                                    Handle: "<value>",
                                },
                            },
                            QuerySuggestions: &components.QuerySuggestionList{
                                Suggestions: []components.QuerySuggestion{
                                    components.QuerySuggestion{
                                        Query: "app:github type:pull author:mortimer",
                                        Label: apiclientgo.String("Mortimer's PRs"),
                                        Datasource: apiclientgo.String("github"),
                                    },
                                },
                            },
                            InviteInfo: &components.InviteInfo{
                                Invites: []components.ChannelInviteInfo{
                                    components.ChannelInviteInfo{},
                                },
                            },
                            Badges: []components.Badge{
                                components.Badge{
                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                    DisplayName: apiclientgo.String("New hire"),
                                    IconConfig: &components.IconConfig{
                                        Color: apiclientgo.String("#343CED"),
                                        Key: apiclientgo.String("person_icon"),
                                        IconType: components.IconTypeGlyph.ToPointer(),
                                        Name: apiclientgo.String("user"),
                                    },
                                },
                            },
                        },
                    },
                    Interactions: &components.DocumentInteractions{
                        Reacts: []components.Reaction{
                            components.Reaction{
                                Reactors: []components.Person{
                                    components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                    },
                                },
                            },
                            components.Reaction{
                                Reactors: []components.Person{
                                    components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                    },
                                },
                            },
                            components.Reaction{
                                Reactors: []components.Person{
                                    components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                    },
                                },
                            },
                        },
                        Shares: []components.Share{
                            components.Share{
                                NumDaysAgo: 652391,
                                Sharer: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                    Verification: &components.Verification{
                        State: components.StateDeprecated,
                        Metadata: &components.VerificationMetadata{
                            LastVerifier: &components.Person{
                                Name: "George Clooney",
                                ObfuscatedID: "abc123",
                                Metadata: &components.PersonMetadata{
                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                    Title: apiclientgo.String("Actor"),
                                    Department: apiclientgo.String("Movies"),
                                    Email: apiclientgo.String("george@example.com"),
                                    Location: apiclientgo.String("Hollywood, CA"),
                                    Phone: apiclientgo.String("6505551234"),
                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                    DatasourceProfile: []components.DatasourceProfile{
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                    },
                                    QuerySuggestions: &components.QuerySuggestionList{
                                        Suggestions: []components.QuerySuggestion{
                                            components.QuerySuggestion{
                                                Query: "app:github type:pull author:mortimer",
                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                Datasource: apiclientgo.String("github"),
                                            },
                                        },
                                    },
                                    InviteInfo: &components.InviteInfo{
                                        Invites: []components.ChannelInviteInfo{
                                            components.ChannelInviteInfo{},
                                        },
                                    },
                                    Badges: []components.Badge{
                                        components.Badge{
                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                            DisplayName: apiclientgo.String("New hire"),
                                            IconConfig: &components.IconConfig{
                                                Color: apiclientgo.String("#343CED"),
                                                Key: apiclientgo.String("person_icon"),
                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                Name: apiclientgo.String("user"),
                                            },
                                        },
                                    },
                                },
                            },
                            Reminders: []components.Reminder{
                                components.Reminder{
                                    Assignee: components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    Requestor: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    RemindAt: 611121,
                                },
                                components.Reminder{
                                    Assignee: components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    Requestor: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    RemindAt: 611121,
                                },
                                components.Reminder{
                                    Assignee: components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    Requestor: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    RemindAt: 611121,
                                },
                            },
                            LastReminder: &components.Reminder{
                                Assignee: components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Requestor: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                RemindAt: 148513,
                            },
                            CandidateVerifiers: []components.Person{
                                components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                },
                            },
                        },
                    },
                    CustomData: map[string]components.CustomDataValue{
                        "someCustomField": components.CustomDataValue{},
                    },
                    ContactPerson: &components.Person{
                        Name: "George Clooney",
                        ObfuscatedID: "abc123",
                        Metadata: &components.PersonMetadata{
                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                            Title: apiclientgo.String("Actor"),
                            Department: apiclientgo.String("Movies"),
                            Email: apiclientgo.String("george@example.com"),
                            Location: apiclientgo.String("Hollywood, CA"),
                            Phone: apiclientgo.String("6505551234"),
                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                            StartDate: types.MustNewDateFromString("2000-01-23"),
                            DatasourceProfile: []components.DatasourceProfile{
                                components.DatasourceProfile{
                                    Datasource: "github",
                                    Handle: "<value>",
                                },
                                components.DatasourceProfile{
                                    Datasource: "github",
                                    Handle: "<value>",
                                },
                                components.DatasourceProfile{
                                    Datasource: "github",
                                    Handle: "<value>",
                                },
                            },
                            QuerySuggestions: &components.QuerySuggestionList{
                                Suggestions: []components.QuerySuggestion{
                                    components.QuerySuggestion{
                                        Query: "app:github type:pull author:mortimer",
                                        Label: apiclientgo.String("Mortimer's PRs"),
                                        Datasource: apiclientgo.String("github"),
                                    },
                                },
                            },
                            InviteInfo: &components.InviteInfo{
                                Invites: []components.ChannelInviteInfo{
                                    components.ChannelInviteInfo{},
                                },
                            },
                            Badges: []components.Badge{
                                components.Badge{
                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                    DisplayName: apiclientgo.String("New hire"),
                                    IconConfig: &components.IconConfig{
                                        Color: apiclientgo.String("#343CED"),
                                        Key: apiclientgo.String("person_icon"),
                                        IconType: components.IconTypeGlyph.ToPointer(),
                                        Name: apiclientgo.String("user"),
                                    },
                                },
                            },
                        },
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResultsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.RecommendationsRequest](../../models/components/recommendationsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.RecommendationsResponse](../../models/operations/recommendationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Query

Retrieve results from the index for the given query and filters.

### Example Usage

```go
package main

import(
	"context"
	"os"
	"github.com/gleanwork/api-client-go/models/components"
	apiclientgo "github.com/gleanwork/api-client-go"
	"github.com/gleanwork/api-client-go/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := apiclientgo.New(
        apiclientgo.WithSecurity(components.Security{
            APIToken: apiclientgo.String(os.Getenv("GLEAN_API_TOKEN")),
        }),
    )

    res, err := s.Client.Search.Query(ctx, &components.SearchRequest{
        TrackingToken: apiclientgo.String("trackingToken"),
        SourceDocument: &components.Document{
            Metadata: &components.DocumentMetadata{
                Datasource: apiclientgo.String("datasource"),
                ObjectType: apiclientgo.String("Feature Request"),
                Container: apiclientgo.String("container"),
                ParentID: apiclientgo.String("JIRA_EN-1337"),
                MimeType: apiclientgo.String("mimeType"),
                DocumentID: apiclientgo.String("documentId"),
                CreateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                UpdateTime: types.MustNewTimeFromString("2000-01-23T04:56:07.000Z"),
                Author: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    RelatedDocuments: []components.RelatedDocuments{
                        components.RelatedDocuments{
                            QuerySuggestion: &components.QuerySuggestion{
                                Query: "app:github type:pull author:mortimer",
                                SearchProviderInfo: &components.SearchProviderInfo{
                                    Name: apiclientgo.String("Google"),
                                    SearchLinkURLTemplate: apiclientgo.String("https://www.google.com/search?q={query}&hl=en"),
                                },
                                Label: apiclientgo.String("Mortimer's PRs"),
                                Datasource: apiclientgo.String("github"),
                                RequestOptions: &components.SearchRequestOptions{
                                    DatasourceFilter: apiclientgo.String("JIRA"),
                                    DatasourcesFilter: []string{
                                        "JIRA",
                                    },
                                    QueryOverridesFacetFilters: apiclientgo.Bool(true),
                                    FacetFilters: []components.FacetFilter{
                                        components.FacetFilter{
                                            FieldName: apiclientgo.String("type"),
                                            Values: []components.FacetFilterValue{
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.String("Spreadsheet"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.String("Presentation"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                            },
                                        },
                                    },
                                    FacetFilterSets: []components.FacetFilterSet{
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.String("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    FacetBucketSize: 718804,
                                    AuthTokens: []components.AuthToken{
                                        components.AuthToken{
                                            AccessToken: "123abc",
                                            Datasource: "gmail",
                                            Scope: apiclientgo.String("email profile https://www.googleapis.com/auth/gmail.readonly"),
                                            TokenType: apiclientgo.String("Bearer"),
                                            AuthUser: apiclientgo.String("1"),
                                        },
                                    },
                                },
                                Ranges: []components.TextRange{
                                    components.TextRange{
                                        StartIndex: 337360,
                                    },
                                    components.TextRange{
                                        StartIndex: 337360,
                                    },
                                },
                                InputDetails: &components.SearchRequestInputDetails{
                                    HasCopyPaste: apiclientgo.Bool(true),
                                },
                            },
                            Results: []components.SearchResult{
                                components.SearchResult{
                                    Title: apiclientgo.String("title"),
                                    URL: "https://example.com/foo/bar",
                                    NativeAppURL: apiclientgo.String("slack://foo/bar"),
                                    Snippets: []components.SearchResultSnippet{
                                        components.SearchResultSnippet{
                                            Snippet: "snippet",
                                            MimeType: apiclientgo.String("mimeType"),
                                        },
                                    },
                                },
                            },
                        },
                        components.RelatedDocuments{
                            QuerySuggestion: &components.QuerySuggestion{
                                Query: "app:github type:pull author:mortimer",
                                SearchProviderInfo: &components.SearchProviderInfo{
                                    Name: apiclientgo.String("Google"),
                                    SearchLinkURLTemplate: apiclientgo.String("https://www.google.com/search?q={query}&hl=en"),
                                },
                                Label: apiclientgo.String("Mortimer's PRs"),
                                Datasource: apiclientgo.String("github"),
                                RequestOptions: &components.SearchRequestOptions{
                                    DatasourceFilter: apiclientgo.String("JIRA"),
                                    DatasourcesFilter: []string{
                                        "JIRA",
                                    },
                                    QueryOverridesFacetFilters: apiclientgo.Bool(true),
                                    FacetFilters: []components.FacetFilter{
                                        components.FacetFilter{
                                            FieldName: apiclientgo.String("type"),
                                            Values: []components.FacetFilterValue{
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.String("Spreadsheet"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.String("Presentation"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                            },
                                        },
                                    },
                                    FacetFilterSets: []components.FacetFilterSet{
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.String("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    FacetBucketSize: 718804,
                                    AuthTokens: []components.AuthToken{
                                        components.AuthToken{
                                            AccessToken: "123abc",
                                            Datasource: "gmail",
                                            Scope: apiclientgo.String("email profile https://www.googleapis.com/auth/gmail.readonly"),
                                            TokenType: apiclientgo.String("Bearer"),
                                            AuthUser: apiclientgo.String("1"),
                                        },
                                    },
                                },
                                Ranges: []components.TextRange{
                                    components.TextRange{
                                        StartIndex: 337360,
                                    },
                                    components.TextRange{
                                        StartIndex: 337360,
                                    },
                                },
                                InputDetails: &components.SearchRequestInputDetails{
                                    HasCopyPaste: apiclientgo.Bool(true),
                                },
                            },
                            Results: []components.SearchResult{
                                components.SearchResult{
                                    Title: apiclientgo.String("title"),
                                    URL: "https://example.com/foo/bar",
                                    NativeAppURL: apiclientgo.String("slack://foo/bar"),
                                    Snippets: []components.SearchResultSnippet{
                                        components.SearchResultSnippet{
                                            Snippet: "snippet",
                                            MimeType: apiclientgo.String("mimeType"),
                                        },
                                    },
                                },
                            },
                        },
                        components.RelatedDocuments{
                            QuerySuggestion: &components.QuerySuggestion{
                                Query: "app:github type:pull author:mortimer",
                                SearchProviderInfo: &components.SearchProviderInfo{
                                    Name: apiclientgo.String("Google"),
                                    SearchLinkURLTemplate: apiclientgo.String("https://www.google.com/search?q={query}&hl=en"),
                                },
                                Label: apiclientgo.String("Mortimer's PRs"),
                                Datasource: apiclientgo.String("github"),
                                RequestOptions: &components.SearchRequestOptions{
                                    DatasourceFilter: apiclientgo.String("JIRA"),
                                    DatasourcesFilter: []string{
                                        "JIRA",
                                    },
                                    QueryOverridesFacetFilters: apiclientgo.Bool(true),
                                    FacetFilters: []components.FacetFilter{
                                        components.FacetFilter{
                                            FieldName: apiclientgo.String("type"),
                                            Values: []components.FacetFilterValue{
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.String("Spreadsheet"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                                components.FacetFilterValue{
                                                    Value: apiclientgo.String("Presentation"),
                                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                                },
                                            },
                                        },
                                    },
                                    FacetFilterSets: []components.FacetFilterSet{
                                        components.FacetFilterSet{
                                            Filters: []components.FacetFilter{
                                                components.FacetFilter{
                                                    FieldName: apiclientgo.String("type"),
                                                    Values: []components.FacetFilterValue{
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Spreadsheet"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                        components.FacetFilterValue{
                                                            Value: apiclientgo.String("Presentation"),
                                                            RelationType: components.RelationTypeEquals.ToPointer(),
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    FacetBucketSize: 718804,
                                    AuthTokens: []components.AuthToken{
                                        components.AuthToken{
                                            AccessToken: "123abc",
                                            Datasource: "gmail",
                                            Scope: apiclientgo.String("email profile https://www.googleapis.com/auth/gmail.readonly"),
                                            TokenType: apiclientgo.String("Bearer"),
                                            AuthUser: apiclientgo.String("1"),
                                        },
                                    },
                                },
                                Ranges: []components.TextRange{
                                    components.TextRange{
                                        StartIndex: 337360,
                                    },
                                    components.TextRange{
                                        StartIndex: 337360,
                                    },
                                },
                                InputDetails: &components.SearchRequestInputDetails{
                                    HasCopyPaste: apiclientgo.Bool(true),
                                },
                            },
                            Results: []components.SearchResult{
                                components.SearchResult{
                                    Title: apiclientgo.String("title"),
                                    URL: "https://example.com/foo/bar",
                                    NativeAppURL: apiclientgo.String("slack://foo/bar"),
                                    Snippets: []components.SearchResultSnippet{
                                        components.SearchResultSnippet{
                                            Snippet: "snippet",
                                            MimeType: apiclientgo.String("mimeType"),
                                        },
                                    },
                                },
                            },
                        },
                    },
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.String("Actor"),
                        Department: apiclientgo.String("Movies"),
                        Email: apiclientgo.String("george@example.com"),
                        Location: apiclientgo.String("Hollywood, CA"),
                        Phone: apiclientgo.String("6505551234"),
                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                        StartDate: types.MustNewDateFromString("2000-01-23"),
                        DatasourceProfile: []components.DatasourceProfile{
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                        },
                        QuerySuggestions: &components.QuerySuggestionList{
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.String("Mortimer's PRs"),
                                    Datasource: apiclientgo.String("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                                components.ChannelInviteInfo{},
                            },
                        },
                        CustomFields: []components.CustomFieldData{
                            components.CustomFieldData{
                                Label: "<value>",
                                Values: []components.CustomFieldValue{
                                    components.CreateCustomFieldValueCustomFieldValueStr(
                                        components.CustomFieldValueStr{},
                                    ),
                                    components.CreateCustomFieldValueCustomFieldValueStr(
                                        components.CustomFieldValueStr{},
                                    ),
                                },
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.String("deployment_name_new_hire"),
                                DisplayName: apiclientgo.String("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.String("#343CED"),
                                    Key: apiclientgo.String("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.String("user"),
                                },
                            },
                        },
                    },
                },
                Owner: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.String("Actor"),
                        Department: apiclientgo.String("Movies"),
                        Email: apiclientgo.String("george@example.com"),
                        Location: apiclientgo.String("Hollywood, CA"),
                        Phone: apiclientgo.String("6505551234"),
                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                        StartDate: types.MustNewDateFromString("2000-01-23"),
                        DatasourceProfile: []components.DatasourceProfile{
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                        },
                        QuerySuggestions: &components.QuerySuggestionList{
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.String("Mortimer's PRs"),
                                    Datasource: apiclientgo.String("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                                components.ChannelInviteInfo{},
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.String("deployment_name_new_hire"),
                                DisplayName: apiclientgo.String("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.String("#343CED"),
                                    Key: apiclientgo.String("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.String("user"),
                                },
                            },
                        },
                    },
                },
                MentionedPeople: []components.Person{
                    components.Person{
                        Name: "George Clooney",
                        ObfuscatedID: "abc123",
                    },
                },
                Components: []string{
                    "Backend",
                    "Networking",
                },
                Status: apiclientgo.String("[\"Done\"]"),
                Pins: []components.PinDocument{
                    components.PinDocument{
                        AudienceFilters: []components.FacetFilter{
                            components.FacetFilter{
                                FieldName: apiclientgo.String("type"),
                                Values: []components.FacetFilterValue{
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Spreadsheet"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Presentation"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                },
                            },
                        },
                        DocumentID: "<id>",
                        Attribution: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                    },
                    components.PinDocument{
                        AudienceFilters: []components.FacetFilter{
                            components.FacetFilter{
                                FieldName: apiclientgo.String("type"),
                                Values: []components.FacetFilterValue{
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Spreadsheet"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Presentation"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                },
                            },
                        },
                        DocumentID: "<id>",
                        Attribution: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                    },
                    components.PinDocument{
                        AudienceFilters: []components.FacetFilter{
                            components.FacetFilter{
                                FieldName: apiclientgo.String("type"),
                                Values: []components.FacetFilterValue{
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Spreadsheet"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Presentation"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                },
                            },
                        },
                        DocumentID: "<id>",
                        Attribution: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
                AssignedTo: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.String("Actor"),
                        Department: apiclientgo.String("Movies"),
                        Email: apiclientgo.String("george@example.com"),
                        Location: apiclientgo.String("Hollywood, CA"),
                        Phone: apiclientgo.String("6505551234"),
                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                        StartDate: types.MustNewDateFromString("2000-01-23"),
                        DatasourceProfile: []components.DatasourceProfile{
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                        },
                        QuerySuggestions: &components.QuerySuggestionList{
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.String("Mortimer's PRs"),
                                    Datasource: apiclientgo.String("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                                components.ChannelInviteInfo{},
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.String("deployment_name_new_hire"),
                                DisplayName: apiclientgo.String("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.String("#343CED"),
                                    Key: apiclientgo.String("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.String("user"),
                                },
                            },
                        },
                    },
                },
                UpdatedBy: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.String("Actor"),
                        Department: apiclientgo.String("Movies"),
                        Email: apiclientgo.String("george@example.com"),
                        Location: apiclientgo.String("Hollywood, CA"),
                        Phone: apiclientgo.String("6505551234"),
                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                        StartDate: types.MustNewDateFromString("2000-01-23"),
                        DatasourceProfile: []components.DatasourceProfile{
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                        },
                        QuerySuggestions: &components.QuerySuggestionList{
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.String("Mortimer's PRs"),
                                    Datasource: apiclientgo.String("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                                components.ChannelInviteInfo{},
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.String("deployment_name_new_hire"),
                                DisplayName: apiclientgo.String("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.String("#343CED"),
                                    Key: apiclientgo.String("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.String("user"),
                                },
                            },
                        },
                    },
                },
                Collections: []components.Collection{
                    components.Collection{
                        Name: "<value>",
                        Description: "incidentally provided bonfire furiously besides whose aw smoggy until following",
                        AddedRoles: []components.UserRoleSpecification{
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleEditor,
                            },
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleEditor,
                            },
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleEditor,
                            },
                        },
                        RemovedRoles: []components.UserRoleSpecification{
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleViewer,
                            },
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleViewer,
                            },
                        },
                        AudienceFilters: []components.FacetFilter{
                            components.FacetFilter{
                                FieldName: apiclientgo.String("type"),
                                Values: []components.FacetFilterValue{
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Spreadsheet"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Presentation"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                },
                            },
                        },
                        ID: 709012,
                        Creator: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        Items: []components.CollectionItem{
                            components.CollectionItem{
                                CollectionID: 94240,
                                CreatedBy: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Shortcut: &components.Shortcut{
                                    InputAlias: "<value>",
                                    CreatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    UpdatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    Roles: []components.UserRoleSpecification{
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                    },
                                },
                                ItemType: components.CollectionItemItemTypeText,
                            },
                            components.CollectionItem{
                                CollectionID: 94240,
                                CreatedBy: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Shortcut: &components.Shortcut{
                                    InputAlias: "<value>",
                                    CreatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    UpdatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    Roles: []components.UserRoleSpecification{
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                    },
                                },
                                ItemType: components.CollectionItemItemTypeText,
                            },
                            components.CollectionItem{
                                CollectionID: 94240,
                                CreatedBy: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Shortcut: &components.Shortcut{
                                    InputAlias: "<value>",
                                    CreatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    UpdatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    Roles: []components.UserRoleSpecification{
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                    },
                                },
                                ItemType: components.CollectionItemItemTypeText,
                            },
                        },
                    },
                    components.Collection{
                        Name: "<value>",
                        Description: "incidentally provided bonfire furiously besides whose aw smoggy until following",
                        AddedRoles: []components.UserRoleSpecification{
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleEditor,
                            },
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleEditor,
                            },
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleEditor,
                            },
                        },
                        RemovedRoles: []components.UserRoleSpecification{
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleViewer,
                            },
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleViewer,
                            },
                        },
                        AudienceFilters: []components.FacetFilter{
                            components.FacetFilter{
                                FieldName: apiclientgo.String("type"),
                                Values: []components.FacetFilterValue{
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Spreadsheet"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Presentation"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                },
                            },
                        },
                        ID: 709012,
                        Creator: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        Items: []components.CollectionItem{
                            components.CollectionItem{
                                CollectionID: 94240,
                                CreatedBy: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Shortcut: &components.Shortcut{
                                    InputAlias: "<value>",
                                    CreatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    UpdatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    Roles: []components.UserRoleSpecification{
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                    },
                                },
                                ItemType: components.CollectionItemItemTypeText,
                            },
                            components.CollectionItem{
                                CollectionID: 94240,
                                CreatedBy: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Shortcut: &components.Shortcut{
                                    InputAlias: "<value>",
                                    CreatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    UpdatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    Roles: []components.UserRoleSpecification{
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                    },
                                },
                                ItemType: components.CollectionItemItemTypeText,
                            },
                            components.CollectionItem{
                                CollectionID: 94240,
                                CreatedBy: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Shortcut: &components.Shortcut{
                                    InputAlias: "<value>",
                                    CreatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    UpdatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    Roles: []components.UserRoleSpecification{
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                    },
                                },
                                ItemType: components.CollectionItemItemTypeText,
                            },
                        },
                    },
                    components.Collection{
                        Name: "<value>",
                        Description: "incidentally provided bonfire furiously besides whose aw smoggy until following",
                        AddedRoles: []components.UserRoleSpecification{
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleEditor,
                            },
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleEditor,
                            },
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleEditor,
                            },
                        },
                        RemovedRoles: []components.UserRoleSpecification{
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleViewer,
                            },
                            components.UserRoleSpecification{
                                Person: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Role: components.UserRoleViewer,
                            },
                        },
                        AudienceFilters: []components.FacetFilter{
                            components.FacetFilter{
                                FieldName: apiclientgo.String("type"),
                                Values: []components.FacetFilterValue{
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Spreadsheet"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                    components.FacetFilterValue{
                                        Value: apiclientgo.String("Presentation"),
                                        RelationType: components.RelationTypeEquals.ToPointer(),
                                    },
                                },
                            },
                        },
                        ID: 709012,
                        Creator: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        Items: []components.CollectionItem{
                            components.CollectionItem{
                                CollectionID: 94240,
                                CreatedBy: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Shortcut: &components.Shortcut{
                                    InputAlias: "<value>",
                                    CreatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    UpdatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    Roles: []components.UserRoleSpecification{
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                    },
                                },
                                ItemType: components.CollectionItemItemTypeText,
                            },
                            components.CollectionItem{
                                CollectionID: 94240,
                                CreatedBy: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Shortcut: &components.Shortcut{
                                    InputAlias: "<value>",
                                    CreatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    UpdatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    Roles: []components.UserRoleSpecification{
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                    },
                                },
                                ItemType: components.CollectionItemItemTypeText,
                            },
                            components.CollectionItem{
                                CollectionID: 94240,
                                CreatedBy: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Shortcut: &components.Shortcut{
                                    InputAlias: "<value>",
                                    CreatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    UpdatedBy: &components.Person{
                                        Name: "George Clooney",
                                        ObfuscatedID: "abc123",
                                        Metadata: &components.PersonMetadata{
                                            Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                            Title: apiclientgo.String("Actor"),
                                            Department: apiclientgo.String("Movies"),
                                            Email: apiclientgo.String("george@example.com"),
                                            Location: apiclientgo.String("Hollywood, CA"),
                                            Phone: apiclientgo.String("6505551234"),
                                            PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                            StartDate: types.MustNewDateFromString("2000-01-23"),
                                            DatasourceProfile: []components.DatasourceProfile{
                                                components.DatasourceProfile{
                                                    Datasource: "github",
                                                    Handle: "<value>",
                                                },
                                            },
                                            QuerySuggestions: &components.QuerySuggestionList{
                                                Suggestions: []components.QuerySuggestion{
                                                    components.QuerySuggestion{
                                                        Query: "app:github type:pull author:mortimer",
                                                        Label: apiclientgo.String("Mortimer's PRs"),
                                                        Datasource: apiclientgo.String("github"),
                                                    },
                                                },
                                            },
                                            InviteInfo: &components.InviteInfo{
                                                Invites: []components.ChannelInviteInfo{
                                                    components.ChannelInviteInfo{},
                                                    components.ChannelInviteInfo{},
                                                },
                                            },
                                            Badges: []components.Badge{
                                                components.Badge{
                                                    Key: apiclientgo.String("deployment_name_new_hire"),
                                                    DisplayName: apiclientgo.String("New hire"),
                                                    IconConfig: &components.IconConfig{
                                                        Color: apiclientgo.String("#343CED"),
                                                        Key: apiclientgo.String("person_icon"),
                                                        IconType: components.IconTypeGlyph.ToPointer(),
                                                        Name: apiclientgo.String("user"),
                                                    },
                                                },
                                            },
                                        },
                                    },
                                    Roles: []components.UserRoleSpecification{
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                        components.UserRoleSpecification{
                                            Person: &components.Person{
                                                Name: "George Clooney",
                                                ObfuscatedID: "abc123",
                                                Metadata: &components.PersonMetadata{
                                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                                    Title: apiclientgo.String("Actor"),
                                                    Department: apiclientgo.String("Movies"),
                                                    Email: apiclientgo.String("george@example.com"),
                                                    Location: apiclientgo.String("Hollywood, CA"),
                                                    Phone: apiclientgo.String("6505551234"),
                                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                                    DatasourceProfile: []components.DatasourceProfile{
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                        components.DatasourceProfile{
                                                            Datasource: "github",
                                                            Handle: "<value>",
                                                        },
                                                    },
                                                    QuerySuggestions: &components.QuerySuggestionList{
                                                        Suggestions: []components.QuerySuggestion{
                                                            components.QuerySuggestion{
                                                                Query: "app:github type:pull author:mortimer",
                                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                                Datasource: apiclientgo.String("github"),
                                                            },
                                                        },
                                                    },
                                                    InviteInfo: &components.InviteInfo{
                                                        Invites: []components.ChannelInviteInfo{
                                                            components.ChannelInviteInfo{},
                                                            components.ChannelInviteInfo{},
                                                        },
                                                    },
                                                    Badges: []components.Badge{
                                                        components.Badge{
                                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                                            DisplayName: apiclientgo.String("New hire"),
                                                            IconConfig: &components.IconConfig{
                                                                Color: apiclientgo.String("#343CED"),
                                                                Key: apiclientgo.String("person_icon"),
                                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                                Name: apiclientgo.String("user"),
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                            Role: components.UserRoleAnswerModerator,
                                        },
                                    },
                                },
                                ItemType: components.CollectionItemItemTypeText,
                            },
                        },
                    },
                },
                Interactions: &components.DocumentInteractions{
                    Reacts: []components.Reaction{
                        components.Reaction{
                            Reactors: []components.Person{
                                components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                },
                            },
                        },
                        components.Reaction{
                            Reactors: []components.Person{
                                components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                },
                            },
                        },
                    },
                    Shares: []components.Share{
                        components.Share{
                            NumDaysAgo: 211330,
                            Sharer: &components.Person{
                                Name: "George Clooney",
                                ObfuscatedID: "abc123",
                                Metadata: &components.PersonMetadata{
                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                    Title: apiclientgo.String("Actor"),
                                    Department: apiclientgo.String("Movies"),
                                    Email: apiclientgo.String("george@example.com"),
                                    Location: apiclientgo.String("Hollywood, CA"),
                                    Phone: apiclientgo.String("6505551234"),
                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                    DatasourceProfile: []components.DatasourceProfile{
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                    },
                                    QuerySuggestions: &components.QuerySuggestionList{
                                        Suggestions: []components.QuerySuggestion{
                                            components.QuerySuggestion{
                                                Query: "app:github type:pull author:mortimer",
                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                Datasource: apiclientgo.String("github"),
                                            },
                                        },
                                    },
                                    InviteInfo: &components.InviteInfo{
                                        Invites: []components.ChannelInviteInfo{
                                            components.ChannelInviteInfo{},
                                            components.ChannelInviteInfo{},
                                        },
                                    },
                                    Badges: []components.Badge{
                                        components.Badge{
                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                            DisplayName: apiclientgo.String("New hire"),
                                            IconConfig: &components.IconConfig{
                                                Color: apiclientgo.String("#343CED"),
                                                Key: apiclientgo.String("person_icon"),
                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                Name: apiclientgo.String("user"),
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
                Verification: &components.Verification{
                    State: components.StateVerified,
                    Metadata: &components.VerificationMetadata{
                        LastVerifier: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        Reminders: []components.Reminder{
                            components.Reminder{
                                Assignee: components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Requestor: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                RemindAt: 43921,
                            },
                            components.Reminder{
                                Assignee: components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                Requestor: &components.Person{
                                    Name: "George Clooney",
                                    ObfuscatedID: "abc123",
                                    Metadata: &components.PersonMetadata{
                                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                        Title: apiclientgo.String("Actor"),
                                        Department: apiclientgo.String("Movies"),
                                        Email: apiclientgo.String("george@example.com"),
                                        Location: apiclientgo.String("Hollywood, CA"),
                                        Phone: apiclientgo.String("6505551234"),
                                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                        StartDate: types.MustNewDateFromString("2000-01-23"),
                                        DatasourceProfile: []components.DatasourceProfile{
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                            components.DatasourceProfile{
                                                Datasource: "github",
                                                Handle: "<value>",
                                            },
                                        },
                                        QuerySuggestions: &components.QuerySuggestionList{
                                            Suggestions: []components.QuerySuggestion{
                                                components.QuerySuggestion{
                                                    Query: "app:github type:pull author:mortimer",
                                                    Label: apiclientgo.String("Mortimer's PRs"),
                                                    Datasource: apiclientgo.String("github"),
                                                },
                                            },
                                        },
                                        InviteInfo: &components.InviteInfo{
                                            Invites: []components.ChannelInviteInfo{
                                                components.ChannelInviteInfo{},
                                                components.ChannelInviteInfo{},
                                            },
                                        },
                                        Badges: []components.Badge{
                                            components.Badge{
                                                Key: apiclientgo.String("deployment_name_new_hire"),
                                                DisplayName: apiclientgo.String("New hire"),
                                                IconConfig: &components.IconConfig{
                                                    Color: apiclientgo.String("#343CED"),
                                                    Key: apiclientgo.String("person_icon"),
                                                    IconType: components.IconTypeGlyph.ToPointer(),
                                                    Name: apiclientgo.String("user"),
                                                },
                                            },
                                        },
                                    },
                                },
                                RemindAt: 43921,
                            },
                        },
                        LastReminder: &components.Reminder{
                            Assignee: components.Person{
                                Name: "George Clooney",
                                ObfuscatedID: "abc123",
                                Metadata: &components.PersonMetadata{
                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                    Title: apiclientgo.String("Actor"),
                                    Department: apiclientgo.String("Movies"),
                                    Email: apiclientgo.String("george@example.com"),
                                    Location: apiclientgo.String("Hollywood, CA"),
                                    Phone: apiclientgo.String("6505551234"),
                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                    DatasourceProfile: []components.DatasourceProfile{
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                    },
                                    QuerySuggestions: &components.QuerySuggestionList{
                                        Suggestions: []components.QuerySuggestion{
                                            components.QuerySuggestion{
                                                Query: "app:github type:pull author:mortimer",
                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                Datasource: apiclientgo.String("github"),
                                            },
                                        },
                                    },
                                    InviteInfo: &components.InviteInfo{
                                        Invites: []components.ChannelInviteInfo{
                                            components.ChannelInviteInfo{},
                                            components.ChannelInviteInfo{},
                                        },
                                    },
                                    Badges: []components.Badge{
                                        components.Badge{
                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                            DisplayName: apiclientgo.String("New hire"),
                                            IconConfig: &components.IconConfig{
                                                Color: apiclientgo.String("#343CED"),
                                                Key: apiclientgo.String("person_icon"),
                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                Name: apiclientgo.String("user"),
                                            },
                                        },
                                    },
                                },
                            },
                            Requestor: &components.Person{
                                Name: "George Clooney",
                                ObfuscatedID: "abc123",
                                Metadata: &components.PersonMetadata{
                                    Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                    Title: apiclientgo.String("Actor"),
                                    Department: apiclientgo.String("Movies"),
                                    Email: apiclientgo.String("george@example.com"),
                                    Location: apiclientgo.String("Hollywood, CA"),
                                    Phone: apiclientgo.String("6505551234"),
                                    PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                    StartDate: types.MustNewDateFromString("2000-01-23"),
                                    DatasourceProfile: []components.DatasourceProfile{
                                        components.DatasourceProfile{
                                            Datasource: "github",
                                            Handle: "<value>",
                                        },
                                    },
                                    QuerySuggestions: &components.QuerySuggestionList{
                                        Suggestions: []components.QuerySuggestion{
                                            components.QuerySuggestion{
                                                Query: "app:github type:pull author:mortimer",
                                                Label: apiclientgo.String("Mortimer's PRs"),
                                                Datasource: apiclientgo.String("github"),
                                            },
                                        },
                                    },
                                    InviteInfo: &components.InviteInfo{
                                        Invites: []components.ChannelInviteInfo{
                                            components.ChannelInviteInfo{},
                                            components.ChannelInviteInfo{},
                                        },
                                    },
                                    Badges: []components.Badge{
                                        components.Badge{
                                            Key: apiclientgo.String("deployment_name_new_hire"),
                                            DisplayName: apiclientgo.String("New hire"),
                                            IconConfig: &components.IconConfig{
                                                Color: apiclientgo.String("#343CED"),
                                                Key: apiclientgo.String("person_icon"),
                                                IconType: components.IconTypeGlyph.ToPointer(),
                                                Name: apiclientgo.String("user"),
                                            },
                                        },
                                    },
                                },
                            },
                            RemindAt: 973534,
                        },
                        CandidateVerifiers: []components.Person{
                            components.Person{
                                Name: "George Clooney",
                                ObfuscatedID: "abc123",
                            },
                        },
                    },
                },
                Shortcuts: []components.Shortcut{
                    components.Shortcut{
                        InputAlias: "<value>",
                        CreatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                    },
                    components.Shortcut{
                        InputAlias: "<value>",
                        CreatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                    },
                    components.Shortcut{
                        InputAlias: "<value>",
                        CreatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                        UpdatedBy: &components.Person{
                            Name: "George Clooney",
                            ObfuscatedID: "abc123",
                            Metadata: &components.PersonMetadata{
                                Type: components.PersonMetadataTypeFullTime.ToPointer(),
                                Title: apiclientgo.String("Actor"),
                                Department: apiclientgo.String("Movies"),
                                Email: apiclientgo.String("george@example.com"),
                                Location: apiclientgo.String("Hollywood, CA"),
                                Phone: apiclientgo.String("6505551234"),
                                PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                                StartDate: types.MustNewDateFromString("2000-01-23"),
                                DatasourceProfile: []components.DatasourceProfile{
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                    components.DatasourceProfile{
                                        Datasource: "github",
                                        Handle: "<value>",
                                    },
                                },
                                QuerySuggestions: &components.QuerySuggestionList{
                                    Suggestions: []components.QuerySuggestion{
                                        components.QuerySuggestion{
                                            Query: "app:github type:pull author:mortimer",
                                            Label: apiclientgo.String("Mortimer's PRs"),
                                            Datasource: apiclientgo.String("github"),
                                        },
                                    },
                                },
                                InviteInfo: &components.InviteInfo{
                                    Invites: []components.ChannelInviteInfo{
                                        components.ChannelInviteInfo{},
                                        components.ChannelInviteInfo{},
                                    },
                                },
                                Badges: []components.Badge{
                                    components.Badge{
                                        Key: apiclientgo.String("deployment_name_new_hire"),
                                        DisplayName: apiclientgo.String("New hire"),
                                        IconConfig: &components.IconConfig{
                                            Color: apiclientgo.String("#343CED"),
                                            Key: apiclientgo.String("person_icon"),
                                            IconType: components.IconTypeGlyph.ToPointer(),
                                            Name: apiclientgo.String("user"),
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
                CustomData: map[string]components.CustomDataValue{
                    "someCustomField": components.CustomDataValue{},
                },
                ContactPerson: &components.Person{
                    Name: "George Clooney",
                    ObfuscatedID: "abc123",
                    Metadata: &components.PersonMetadata{
                        Type: components.PersonMetadataTypeFullTime.ToPointer(),
                        Title: apiclientgo.String("Actor"),
                        Department: apiclientgo.String("Movies"),
                        Email: apiclientgo.String("george@example.com"),
                        Location: apiclientgo.String("Hollywood, CA"),
                        Phone: apiclientgo.String("6505551234"),
                        PhotoURL: apiclientgo.String("https://example.com/george.jpg"),
                        StartDate: types.MustNewDateFromString("2000-01-23"),
                        DatasourceProfile: []components.DatasourceProfile{
                            components.DatasourceProfile{
                                Datasource: "github",
                                Handle: "<value>",
                            },
                        },
                        QuerySuggestions: &components.QuerySuggestionList{
                            Suggestions: []components.QuerySuggestion{
                                components.QuerySuggestion{
                                    Query: "app:github type:pull author:mortimer",
                                    Label: apiclientgo.String("Mortimer's PRs"),
                                    Datasource: apiclientgo.String("github"),
                                },
                            },
                        },
                        InviteInfo: &components.InviteInfo{
                            Invites: []components.ChannelInviteInfo{
                                components.ChannelInviteInfo{},
                                components.ChannelInviteInfo{},
                            },
                        },
                        Badges: []components.Badge{
                            components.Badge{
                                Key: apiclientgo.String("deployment_name_new_hire"),
                                DisplayName: apiclientgo.String("New hire"),
                                IconConfig: &components.IconConfig{
                                    Color: apiclientgo.String("#343CED"),
                                    Key: apiclientgo.String("person_icon"),
                                    IconType: components.IconTypeGlyph.ToPointer(),
                                    Name: apiclientgo.String("user"),
                                },
                            },
                        },
                    },
                },
            },
        },
        PageSize: apiclientgo.Int64(100),
        MaxSnippetSize: apiclientgo.Int64(400),
        Query: "vacation policy",
        InputDetails: &components.SearchRequestInputDetails{
            HasCopyPaste: apiclientgo.Bool(true),
        },
        RequestOptions: &components.SearchRequestOptions{
            DatasourceFilter: apiclientgo.String("JIRA"),
            DatasourcesFilter: []string{
                "JIRA",
            },
            QueryOverridesFacetFilters: apiclientgo.Bool(true),
            FacetFilters: []components.FacetFilter{
                components.FacetFilter{
                    FieldName: apiclientgo.String("type"),
                    Values: []components.FacetFilterValue{
                        components.FacetFilterValue{
                            Value: apiclientgo.String("Spreadsheet"),
                            RelationType: components.RelationTypeEquals.ToPointer(),
                        },
                        components.FacetFilterValue{
                            Value: apiclientgo.String("Presentation"),
                            RelationType: components.RelationTypeEquals.ToPointer(),
                        },
                    },
                },
            },
            FacetFilterSets: []components.FacetFilterSet{
                components.FacetFilterSet{
                    Filters: []components.FacetFilter{
                        components.FacetFilter{
                            FieldName: apiclientgo.String("type"),
                            Values: []components.FacetFilterValue{
                                components.FacetFilterValue{
                                    Value: apiclientgo.String("Spreadsheet"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                                components.FacetFilterValue{
                                    Value: apiclientgo.String("Presentation"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                            },
                        },
                    },
                },
                components.FacetFilterSet{
                    Filters: []components.FacetFilter{
                        components.FacetFilter{
                            FieldName: apiclientgo.String("type"),
                            Values: []components.FacetFilterValue{
                                components.FacetFilterValue{
                                    Value: apiclientgo.String("Spreadsheet"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                                components.FacetFilterValue{
                                    Value: apiclientgo.String("Presentation"),
                                    RelationType: components.RelationTypeEquals.ToPointer(),
                                },
                            },
                        },
                    },
                },
            },
            FacetBucketSize: 400611,
            AuthTokens: []components.AuthToken{
                components.AuthToken{
                    AccessToken: "123abc",
                    Datasource: "gmail",
                    Scope: apiclientgo.String("email profile https://www.googleapis.com/auth/gmail.readonly"),
                    TokenType: apiclientgo.String("Bearer"),
                    AuthUser: apiclientgo.String("1"),
                },
            },
        },
        TimeoutMillis: apiclientgo.Int64(5000),
        People: []components.Person{
            components.Person{
                Name: "George Clooney",
                ObfuscatedID: "abc123",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SearchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [components.SearchRequest](../../models/components/searchrequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |
| `opts`                                                               | [][operations.Option](../../models/operations/option.md)             | :heavy_minus_sign:                                                   | The options for this request.                                        |

### Response

**[*operations.SearchResponse](../../models/operations/searchresponse.md), error**

### Errors

| Error Type               | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| apierrors.GleanDataError | 403, 422                 | application/json         |
| apierrors.APIError       | 4XX, 5XX                 | \*/\*                    |