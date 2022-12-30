// Code generated by ent, DO NOT EDIT.

package authroles

import (
	"fmt"
)

const (
	// Label holds the string label denoting the authroles type in the database.
	Label = "auth_roles"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// EdgeToken holds the string denoting the token edge name in mutations.
	EdgeToken = "token"
	// Table holds the table name of the authroles in the database.
	Table = "auth_roles"
	// TokenTable is the table that holds the token relation/edge.
	TokenTable = "auth_roles"
	// TokenInverseTable is the table name for the AuthTokens entity.
	// It exists in this package in order to avoid circular dependency with the "authtokens" package.
	TokenInverseTable = "auth_tokens"
	// TokenColumn is the table column denoting the token relation/edge.
	TokenColumn = "auth_tokens_roles"
)

// Columns holds all SQL columns for authroles fields.
var Columns = []string{
	FieldID,
	FieldRole,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "auth_roles"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"auth_tokens_roles",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Role defines the type for the "role" enum field.
type Role string

// RoleUser is the default value of the Role enum.
const DefaultRole = RoleUser

// Role values.
const (
	RoleAdmin       Role = "admin"
	RoleUser        Role = "user"
	RoleAttachments Role = "attachments"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleAdmin, RoleUser, RoleAttachments:
		return nil
	default:
		return fmt.Errorf("authroles: invalid enum value for role field: %q", r)
	}
}