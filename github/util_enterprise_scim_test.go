package github

import (
	"testing"
	"time"

	gh "github.com/google/go-github/v83/github"
)

func TestFlattenEnterpriseSCIMMeta(t *testing.T) {
	t.Run("returns nil for nil input", func(t *testing.T) {
		result := flattenEnterpriseSCIMMeta(nil)
		if result != nil {
			t.Fatalf("expected nil, got %v", result)
		}
	})

	t.Run("returns resource_type only when no optional fields", func(t *testing.T) {
		meta := &gh.SCIMEnterpriseMeta{
			ResourceType: "User",
		}
		result := flattenEnterpriseSCIMMeta(meta)
		if len(result) != 1 {
			t.Fatalf("expected 1 element, got %d", len(result))
		}
		m := result[0].(map[string]any)
		if m["resource_type"] != "User" {
			t.Fatalf("expected resource_type 'User', got %v", m["resource_type"])
		}
		if _, ok := m["created"]; ok {
			t.Fatal("expected 'created' to be absent")
		}
		if _, ok := m["last_modified"]; ok {
			t.Fatal("expected 'last_modified' to be absent")
		}
		if _, ok := m["location"]; ok {
			t.Fatal("expected 'location' to be absent")
		}
	})

	t.Run("returns all fields when fully populated", func(t *testing.T) {
		now := time.Date(2025, 1, 15, 10, 0, 0, 0, time.UTC)
		meta := &gh.SCIMEnterpriseMeta{
			ResourceType: "Group",
			Created:      &gh.Timestamp{Time: now},
			LastModified: &gh.Timestamp{Time: now.Add(time.Hour)},
			Location:     gh.Ptr("https://api.github.com/scim/v2/enterprises/test/Groups/123"),
		}
		result := flattenEnterpriseSCIMMeta(meta)
		if len(result) != 1 {
			t.Fatalf("expected 1 element, got %d", len(result))
		}
		m := result[0].(map[string]any)
		if m["resource_type"] != "Group" {
			t.Fatalf("expected resource_type 'Group', got %v", m["resource_type"])
		}
		if m["created"] == "" {
			t.Fatal("expected 'created' to be non-empty")
		}
		if m["last_modified"] == "" {
			t.Fatal("expected 'last_modified' to be non-empty")
		}
		if m["location"] != "https://api.github.com/scim/v2/enterprises/test/Groups/123" {
			t.Fatalf("expected location URL, got %v", m["location"])
		}
	})
}

func TestFlattenEnterpriseSCIMGroupMembers(t *testing.T) {
	t.Run("returns empty slice for nil input", func(t *testing.T) {
		result := flattenEnterpriseSCIMGroupMembers(nil)
		if len(result) != 0 {
			t.Fatalf("expected empty slice, got %d elements", len(result))
		}
	})

	t.Run("returns members with all fields", func(t *testing.T) {
		members := []*gh.SCIMEnterpriseDisplayReference{
			{
				Value:   "user-1",
				Ref:     gh.Ptr("https://api.github.com/scim/v2/enterprises/test/Users/user-1"),
				Display: gh.Ptr("Alice"),
			},
			{
				Value: "user-2",
			},
		}
		result := flattenEnterpriseSCIMGroupMembers(members)
		if len(result) != 2 {
			t.Fatalf("expected 2 members, got %d", len(result))
		}

		first := result[0].(map[string]any)
		if first["value"] != "user-1" {
			t.Fatalf("expected value 'user-1', got %v", first["value"])
		}
		if first["ref"] != "https://api.github.com/scim/v2/enterprises/test/Users/user-1" {
			t.Fatalf("expected ref URL, got %v", first["ref"])
		}
		if first["display_name"] != "Alice" {
			t.Fatalf("expected display_name 'Alice', got %v", first["display_name"])
		}

		second := result[1].(map[string]any)
		if second["value"] != "user-2" {
			t.Fatalf("expected value 'user-2', got %v", second["value"])
		}
		if _, ok := second["ref"]; ok {
			t.Fatal("expected 'ref' to be absent")
		}
		if _, ok := second["display_name"]; ok {
			t.Fatal("expected 'display_name' to be absent")
		}
	})
}

func TestFlattenEnterpriseSCIMGroup(t *testing.T) {
	t.Run("returns all fields", func(t *testing.T) {
		group := &gh.SCIMEnterpriseGroupAttributes{
			ID:          gh.Ptr("group-123"),
			ExternalID:  gh.Ptr("ext-456"),
			DisplayName: gh.Ptr("Engineering"),
			Schemas:     []string{"urn:ietf:params:scim:schemas:core:2.0:Group"},
			Members: []*gh.SCIMEnterpriseDisplayReference{
				{Value: "user-1", Display: gh.Ptr("Alice")},
			},
		}
		result := flattenEnterpriseSCIMGroup(group)
		if result["id"] != "group-123" {
			t.Fatalf("expected id 'group-123', got %v", result["id"])
		}
		if result["external_id"] != "ext-456" {
			t.Fatalf("expected external_id 'ext-456', got %v", result["external_id"])
		}
		if result["display_name"] != "Engineering" {
			t.Fatalf("expected display_name 'Engineering', got %v", result["display_name"])
		}
		members := result["members"].([]any)
		if len(members) != 1 {
			t.Fatalf("expected 1 member, got %d", len(members))
		}
	})

	t.Run("handles nil optional fields", func(t *testing.T) {
		group := &gh.SCIMEnterpriseGroupAttributes{
			Schemas: []string{"urn:ietf:params:scim:schemas:core:2.0:Group"},
		}
		result := flattenEnterpriseSCIMGroup(group)
		if _, ok := result["id"]; ok {
			t.Fatal("expected 'id' to be absent")
		}
		if _, ok := result["external_id"]; ok {
			t.Fatal("expected 'external_id' to be absent")
		}
		if _, ok := result["display_name"]; ok {
			t.Fatal("expected 'display_name' to be absent")
		}
	})
}

func TestFlattenEnterpriseSCIMUserName(t *testing.T) {
	t.Run("returns nil for nil input", func(t *testing.T) {
		result := flattenEnterpriseSCIMUserName(nil)
		if result != nil {
			t.Fatalf("expected nil, got %v", result)
		}
	})

	t.Run("returns required fields only", func(t *testing.T) {
		name := &gh.SCIMEnterpriseUserName{
			GivenName:  "John",
			FamilyName: "Doe",
		}
		result := flattenEnterpriseSCIMUserName(name)
		if len(result) != 1 {
			t.Fatalf("expected 1 element, got %d", len(result))
		}
		m := result[0].(map[string]any)
		if m["given_name"] != "John" {
			t.Fatalf("expected given_name 'John', got %v", m["given_name"])
		}
		if m["family_name"] != "Doe" {
			t.Fatalf("expected family_name 'Doe', got %v", m["family_name"])
		}
		if _, ok := m["formatted"]; ok {
			t.Fatal("expected 'formatted' to be absent")
		}
		if _, ok := m["middle_name"]; ok {
			t.Fatal("expected 'middle_name' to be absent")
		}
	})

	t.Run("returns all fields when populated", func(t *testing.T) {
		name := &gh.SCIMEnterpriseUserName{
			GivenName:  "John",
			FamilyName: "Doe",
			Formatted:  gh.Ptr("John M. Doe"),
			MiddleName: gh.Ptr("M."),
		}
		result := flattenEnterpriseSCIMUserName(name)
		if len(result) != 1 {
			t.Fatalf("expected 1 element, got %d", len(result))
		}
		m := result[0].(map[string]any)
		if m["formatted"] != "John M. Doe" {
			t.Fatalf("expected formatted 'John M. Doe', got %v", m["formatted"])
		}
		if m["middle_name"] != "M." {
			t.Fatalf("expected middle_name 'M.', got %v", m["middle_name"])
		}
	})
}

func TestFlattenEnterpriseSCIMUserEmails(t *testing.T) {
	t.Run("returns empty slice for nil input", func(t *testing.T) {
		result := flattenEnterpriseSCIMUserEmails(nil)
		if len(result) != 0 {
			t.Fatalf("expected empty slice, got %d elements", len(result))
		}
	})

	t.Run("returns emails with all fields", func(t *testing.T) {
		emails := []*gh.SCIMEnterpriseUserEmail{
			{Value: "alice@example.com", Type: "work", Primary: true},
			{Value: "alice@personal.com", Type: "home", Primary: false},
		}
		result := flattenEnterpriseSCIMUserEmails(emails)
		if len(result) != 2 {
			t.Fatalf("expected 2 emails, got %d", len(result))
		}

		first := result[0].(map[string]any)
		if first["value"] != "alice@example.com" {
			t.Fatalf("expected value 'alice@example.com', got %v", first["value"])
		}
		if first["type"] != "work" {
			t.Fatalf("expected type 'work', got %v", first["type"])
		}
		if first["primary"] != true {
			t.Fatalf("expected primary true, got %v", first["primary"])
		}

		second := result[1].(map[string]any)
		if second["primary"] != false {
			t.Fatalf("expected primary false, got %v", second["primary"])
		}
	})
}

func TestFlattenEnterpriseSCIMUserRoles(t *testing.T) {
	t.Run("returns empty slice for nil input", func(t *testing.T) {
		result := flattenEnterpriseSCIMUserRoles(nil)
		if len(result) != 0 {
			t.Fatalf("expected empty slice, got %d elements", len(result))
		}
	})

	t.Run("returns roles with all fields", func(t *testing.T) {
		roles := []*gh.SCIMEnterpriseUserRole{
			{
				Value:   "user",
				Display: gh.Ptr("User"),
				Type:    gh.Ptr("default"),
				Primary: gh.Ptr(true),
			},
		}
		result := flattenEnterpriseSCIMUserRoles(roles)
		if len(result) != 1 {
			t.Fatalf("expected 1 role, got %d", len(result))
		}
		role := result[0].(map[string]any)
		if role["value"] != "user" {
			t.Fatalf("expected value 'user', got %v", role["value"])
		}
		if role["display"] != "User" {
			t.Fatalf("expected display 'User', got %v", role["display"])
		}
		if role["type"] != "default" {
			t.Fatalf("expected type 'default', got %v", role["type"])
		}
		if role["primary"] != true {
			t.Fatalf("expected primary true, got %v", role["primary"])
		}
	})

	t.Run("handles nil optional fields", func(t *testing.T) {
		roles := []*gh.SCIMEnterpriseUserRole{
			{Value: "admin"},
		}
		result := flattenEnterpriseSCIMUserRoles(roles)
		if len(result) != 1 {
			t.Fatalf("expected 1 role, got %d", len(result))
		}
		role := result[0].(map[string]any)
		if _, ok := role["display"]; ok {
			t.Fatal("expected 'display' to be absent")
		}
		if _, ok := role["type"]; ok {
			t.Fatal("expected 'type' to be absent")
		}
		if _, ok := role["primary"]; ok {
			t.Fatal("expected 'primary' to be absent")
		}
	})
}

func TestFlattenEnterpriseSCIMUser(t *testing.T) {
	t.Run("returns all fields", func(t *testing.T) {
		user := &gh.SCIMEnterpriseUserAttributes{
			ID:          gh.Ptr("scim-user-123"),
			UserName:    "alice",
			DisplayName: "Alice Wonderland",
			ExternalID:  "ext-789",
			Active:      true,
			Schemas:     []string{"urn:ietf:params:scim:schemas:core:2.0:User"},
			Name: &gh.SCIMEnterpriseUserName{
				GivenName:  "Alice",
				FamilyName: "Wonderland",
			},
			Emails: []*gh.SCIMEnterpriseUserEmail{
				{Value: "alice@example.com", Type: "work", Primary: true},
			},
			Roles: []*gh.SCIMEnterpriseUserRole{
				{Value: "user"},
			},
		}
		result := flattenEnterpriseSCIMUser(user)
		if result["id"] != "scim-user-123" {
			t.Fatalf("expected id 'scim-user-123', got %v", result["id"])
		}
		if result["user_name"] != "alice" {
			t.Fatalf("expected user_name 'alice', got %v", result["user_name"])
		}
		if result["display_name"] != "Alice Wonderland" {
			t.Fatalf("expected display_name 'Alice Wonderland', got %v", result["display_name"])
		}
		if result["external_id"] != "ext-789" {
			t.Fatalf("expected external_id 'ext-789', got %v", result["external_id"])
		}
		if result["active"] != true {
			t.Fatalf("expected active true, got %v", result["active"])
		}
		nameSlice := result["name"].([]any)
		if len(nameSlice) != 1 {
			t.Fatalf("expected 1 name entry, got %d", len(nameSlice))
		}
		emailsSlice := result["emails"].([]any)
		if len(emailsSlice) != 1 {
			t.Fatalf("expected 1 email entry, got %d", len(emailsSlice))
		}
		rolesSlice := result["roles"].([]any)
		if len(rolesSlice) != 1 {
			t.Fatalf("expected 1 role entry, got %d", len(rolesSlice))
		}
	})

	t.Run("handles nil ID", func(t *testing.T) {
		user := &gh.SCIMEnterpriseUserAttributes{
			UserName:    "bob",
			DisplayName: "Bob Builder",
			ExternalID:  "ext-000",
			Active:      false,
			Schemas:     []string{"urn:ietf:params:scim:schemas:core:2.0:User"},
		}
		result := flattenEnterpriseSCIMUser(user)
		if _, ok := result["id"]; ok {
			t.Fatal("expected 'id' to be absent")
		}
		if result["user_name"] != "bob" {
			t.Fatalf("expected user_name 'bob', got %v", result["user_name"])
		}
		if result["active"] != false {
			t.Fatalf("expected active false, got %v", result["active"])
		}
	})
}
