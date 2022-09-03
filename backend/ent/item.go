// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/hay-kot/content/backend/ent/group"
	"github.com/hay-kot/content/backend/ent/item"
	"github.com/hay-kot/content/backend/ent/location"
)

// Item is the model entity for the Item schema.
type Item struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Notes holds the value of the "notes" field.
	Notes string `json:"notes,omitempty"`
	// SerialNumber holds the value of the "serial_number" field.
	SerialNumber string `json:"serial_number,omitempty"`
	// ModelNumber holds the value of the "model_number" field.
	ModelNumber string `json:"model_number,omitempty"`
	// Manufacturer holds the value of the "manufacturer" field.
	Manufacturer string `json:"manufacturer,omitempty"`
	// PurchaseTime holds the value of the "purchase_time" field.
	PurchaseTime time.Time `json:"purchase_time,omitempty"`
	// PurchaseFrom holds the value of the "purchase_from" field.
	PurchaseFrom string `json:"purchase_from,omitempty"`
	// PurchasePrice holds the value of the "purchase_price" field.
	PurchasePrice float64 `json:"purchase_price,omitempty"`
	// PurchaseReceiptID holds the value of the "purchase_receipt_id" field.
	PurchaseReceiptID uuid.UUID `json:"purchase_receipt_id,omitempty"`
	// SoldTime holds the value of the "sold_time" field.
	SoldTime time.Time `json:"sold_time,omitempty"`
	// SoldTo holds the value of the "sold_to" field.
	SoldTo string `json:"sold_to,omitempty"`
	// SoldPrice holds the value of the "sold_price" field.
	SoldPrice float64 `json:"sold_price,omitempty"`
	// SoldReceiptID holds the value of the "sold_receipt_id" field.
	SoldReceiptID uuid.UUID `json:"sold_receipt_id,omitempty"`
	// SoldNotes holds the value of the "sold_notes" field.
	SoldNotes string `json:"sold_notes,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ItemQuery when eager-loading is set.
	Edges          ItemEdges `json:"edges"`
	group_items    *uuid.UUID
	location_items *uuid.UUID
}

// ItemEdges holds the relations/edges for other nodes in the graph.
type ItemEdges struct {
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// Location holds the value of the location edge.
	Location *Location `json:"location,omitempty"`
	// Fields holds the value of the fields edge.
	Fields []*ItemField `json:"fields,omitempty"`
	// Label holds the value of the label edge.
	Label []*Label `json:"label,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ItemEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[0] {
		if e.Group == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// LocationOrErr returns the Location value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ItemEdges) LocationOrErr() (*Location, error) {
	if e.loadedTypes[1] {
		if e.Location == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: location.Label}
		}
		return e.Location, nil
	}
	return nil, &NotLoadedError{edge: "location"}
}

// FieldsOrErr returns the Fields value or an error if the edge
// was not loaded in eager-loading.
func (e ItemEdges) FieldsOrErr() ([]*ItemField, error) {
	if e.loadedTypes[2] {
		return e.Fields, nil
	}
	return nil, &NotLoadedError{edge: "fields"}
}

// LabelOrErr returns the Label value or an error if the edge
// was not loaded in eager-loading.
func (e ItemEdges) LabelOrErr() ([]*Label, error) {
	if e.loadedTypes[3] {
		return e.Label, nil
	}
	return nil, &NotLoadedError{edge: "label"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Item) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case item.FieldPurchasePrice, item.FieldSoldPrice:
			values[i] = new(sql.NullFloat64)
		case item.FieldName, item.FieldDescription, item.FieldNotes, item.FieldSerialNumber, item.FieldModelNumber, item.FieldManufacturer, item.FieldPurchaseFrom, item.FieldSoldTo, item.FieldSoldNotes:
			values[i] = new(sql.NullString)
		case item.FieldCreatedAt, item.FieldUpdatedAt, item.FieldPurchaseTime, item.FieldSoldTime:
			values[i] = new(sql.NullTime)
		case item.FieldID, item.FieldPurchaseReceiptID, item.FieldSoldReceiptID:
			values[i] = new(uuid.UUID)
		case item.ForeignKeys[0]: // group_items
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case item.ForeignKeys[1]: // location_items
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Item", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Item fields.
func (i *Item) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case item.FieldID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[j])
			} else if value != nil {
				i.ID = *value
			}
		case item.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case item.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case item.FieldName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[j])
			} else if value.Valid {
				i.Name = value.String
			}
		case item.FieldDescription:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[j])
			} else if value.Valid {
				i.Description = value.String
			}
		case item.FieldNotes:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notes", values[j])
			} else if value.Valid {
				i.Notes = value.String
			}
		case item.FieldSerialNumber:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial_number", values[j])
			} else if value.Valid {
				i.SerialNumber = value.String
			}
		case item.FieldModelNumber:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field model_number", values[j])
			} else if value.Valid {
				i.ModelNumber = value.String
			}
		case item.FieldManufacturer:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field manufacturer", values[j])
			} else if value.Valid {
				i.Manufacturer = value.String
			}
		case item.FieldPurchaseTime:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field purchase_time", values[j])
			} else if value.Valid {
				i.PurchaseTime = value.Time
			}
		case item.FieldPurchaseFrom:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field purchase_from", values[j])
			} else if value.Valid {
				i.PurchaseFrom = value.String
			}
		case item.FieldPurchasePrice:
			if value, ok := values[j].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field purchase_price", values[j])
			} else if value.Valid {
				i.PurchasePrice = value.Float64
			}
		case item.FieldPurchaseReceiptID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field purchase_receipt_id", values[j])
			} else if value != nil {
				i.PurchaseReceiptID = *value
			}
		case item.FieldSoldTime:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field sold_time", values[j])
			} else if value.Valid {
				i.SoldTime = value.Time
			}
		case item.FieldSoldTo:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sold_to", values[j])
			} else if value.Valid {
				i.SoldTo = value.String
			}
		case item.FieldSoldPrice:
			if value, ok := values[j].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field sold_price", values[j])
			} else if value.Valid {
				i.SoldPrice = value.Float64
			}
		case item.FieldSoldReceiptID:
			if value, ok := values[j].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field sold_receipt_id", values[j])
			} else if value != nil {
				i.SoldReceiptID = *value
			}
		case item.FieldSoldNotes:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sold_notes", values[j])
			} else if value.Valid {
				i.SoldNotes = value.String
			}
		case item.ForeignKeys[0]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field group_items", values[j])
			} else if value.Valid {
				i.group_items = new(uuid.UUID)
				*i.group_items = *value.S.(*uuid.UUID)
			}
		case item.ForeignKeys[1]:
			if value, ok := values[j].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field location_items", values[j])
			} else if value.Valid {
				i.location_items = new(uuid.UUID)
				*i.location_items = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryGroup queries the "group" edge of the Item entity.
func (i *Item) QueryGroup() *GroupQuery {
	return (&ItemClient{config: i.config}).QueryGroup(i)
}

// QueryLocation queries the "location" edge of the Item entity.
func (i *Item) QueryLocation() *LocationQuery {
	return (&ItemClient{config: i.config}).QueryLocation(i)
}

// QueryFields queries the "fields" edge of the Item entity.
func (i *Item) QueryFields() *ItemFieldQuery {
	return (&ItemClient{config: i.config}).QueryFields(i)
}

// QueryLabel queries the "label" edge of the Item entity.
func (i *Item) QueryLabel() *LabelQuery {
	return (&ItemClient{config: i.config}).QueryLabel(i)
}

// Update returns a builder for updating this Item.
// Note that you need to call Item.Unwrap() before calling this method if this Item
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Item) Update() *ItemUpdateOne {
	return (&ItemClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Item entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Item) Unwrap() *Item {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Item is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Item) String() string {
	var builder strings.Builder
	builder.WriteString("Item(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(i.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(i.Description)
	builder.WriteString(", ")
	builder.WriteString("notes=")
	builder.WriteString(i.Notes)
	builder.WriteString(", ")
	builder.WriteString("serial_number=")
	builder.WriteString(i.SerialNumber)
	builder.WriteString(", ")
	builder.WriteString("model_number=")
	builder.WriteString(i.ModelNumber)
	builder.WriteString(", ")
	builder.WriteString("manufacturer=")
	builder.WriteString(i.Manufacturer)
	builder.WriteString(", ")
	builder.WriteString("purchase_time=")
	builder.WriteString(i.PurchaseTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("purchase_from=")
	builder.WriteString(i.PurchaseFrom)
	builder.WriteString(", ")
	builder.WriteString("purchase_price=")
	builder.WriteString(fmt.Sprintf("%v", i.PurchasePrice))
	builder.WriteString(", ")
	builder.WriteString("purchase_receipt_id=")
	builder.WriteString(fmt.Sprintf("%v", i.PurchaseReceiptID))
	builder.WriteString(", ")
	builder.WriteString("sold_time=")
	builder.WriteString(i.SoldTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("sold_to=")
	builder.WriteString(i.SoldTo)
	builder.WriteString(", ")
	builder.WriteString("sold_price=")
	builder.WriteString(fmt.Sprintf("%v", i.SoldPrice))
	builder.WriteString(", ")
	builder.WriteString("sold_receipt_id=")
	builder.WriteString(fmt.Sprintf("%v", i.SoldReceiptID))
	builder.WriteString(", ")
	builder.WriteString("sold_notes=")
	builder.WriteString(i.SoldNotes)
	builder.WriteByte(')')
	return builder.String()
}

// Items is a parsable slice of Item.
type Items []*Item

func (i Items) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}