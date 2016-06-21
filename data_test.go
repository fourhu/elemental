package elemental

import "fmt"

const (
	ListAttributeNameID          AttributeSpecificationNameKey = "list/ID"
	ListAttributeNameDescription AttributeSpecificationNameKey = "list/description"
	ListAttributeNameName        AttributeSpecificationNameKey = "list/name"
	ListAttributeNameParentID    AttributeSpecificationNameKey = "list/parentID"
	ListAttributeNameParentType  AttributeSpecificationNameKey = "list/parentType"
)

// ListIdentity represents the Identity of the object
var ListIdentity = Identity{
	Name:     "list",
	Category: "lists",
}

// ListsList represents a list of Lists
type ListsList []*List

// List represents the model of a list
type List struct {
	ID          string `json:"ID,omitempty" cql:"id,omitempty"`
	Description string `json:"description,omitempty" cql:"description,omitempty"`
	Name        string `json:"name,omitempty" cql:"name,omitempty"`
	ParentID    string `json:"parentID,omitempty" cql:"parentid,omitempty"`
	ParentType  string `json:"parentType,omitempty" cql:"parenttype,omitempty"`
}

// NewList returns a new *List
func NewList() *List {

	return &List{}
}

// Identity returns the Identity of the object.
func (o *List) Identity() Identity {

	return ListIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *List) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *List) SetIdentifier(ID string) {

	o.ID = ID
}

// Validate valides the current information stored into the structure.
func (o *List) Validate() Errors {

	errors := Errors{}

	if err := ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o List) SpecificationForAttribute(name AttributeSpecificationNameKey) AttributeSpecification {

	return ListAttributesMap[name]
}

var ListAttributesMap = map[AttributeSpecificationNameKey]AttributeSpecification{
	ListAttributeNameID: AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	ListAttributeNameDescription: AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	ListAttributeNameName: AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	ListAttributeNameParentID: AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Format:         "free",
		Name:           "parentID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	ListAttributeNameParentType: AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "parentType",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
}

const (
	TaskAttributeNameID          AttributeSpecificationNameKey = "task/ID"
	TaskAttributeNameDescription AttributeSpecificationNameKey = "task/description"
	TaskAttributeNameName        AttributeSpecificationNameKey = "task/name"
	TaskAttributeNameParentID    AttributeSpecificationNameKey = "task/parentID"
	TaskAttributeNameParentType  AttributeSpecificationNameKey = "task/parentType"
	TaskAttributeNameStatus      AttributeSpecificationNameKey = "task/status"
)

// TaskStatusValue represents the possible values for attribute "status".
type TaskStatusValue string

const (
	TaskStatusDone     TaskStatusValue = "DONE"
	TaskStatusProgress TaskStatusValue = "PROGRESS"
	TaskStatusTodo     TaskStatusValue = "TODO"
)

// TaskIdentity represents the Identity of the object
var TaskIdentity = Identity{
	Name:     "task",
	Category: "tasks",
}

// TasksList represents a list of Tasks
type TasksList []*Task

// Task represents the model of a task
type Task struct {
	ID          string          `json:"ID,omitempty" cql:"id,omitempty"`
	Description string          `json:"description,omitempty" cql:"description,omitempty"`
	Name        string          `json:"name,omitempty" cql:"name,omitempty"`
	ParentID    string          `json:"parentID,omitempty" cql:"parentid,omitempty"`
	ParentType  string          `json:"parentType,omitempty" cql:"parenttype,omitempty"`
	Status      TaskStatusValue `json:"status,omitempty" cql:"status,omitempty"`
}

// NewTask returns a new *Task
func NewTask() *Task {

	return &Task{
		Status: "TODO",
	}
}

// Identity returns the Identity of the object.
func (o *Task) Identity() Identity {

	return TaskIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Task) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Task) SetIdentifier(ID string) {

	o.ID = ID
}

// Validate valides the current information stored into the structure.
func (o *Task) Validate() Errors {

	errors := Errors{}

	if err := ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := ValidateStringInList("status", string(o.Status), []string{"DONE", "PROGRESS", "TODO"}); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o Task) SpecificationForAttribute(name AttributeSpecificationNameKey) AttributeSpecification {

	return TaskAttributesMap[name]
}

var TaskAttributesMap = map[AttributeSpecificationNameKey]AttributeSpecification{
	TaskAttributeNameID: AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	TaskAttributeNameDescription: AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	TaskAttributeNameName: AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	TaskAttributeNameParentID: AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Format:         "free",
		Name:           "parentID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	TaskAttributeNameParentType: AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "parentType",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	TaskAttributeNameStatus: AttributeSpecification{
		AllowedChoices: []string{"DONE", "PROGRESS", "TODO"},
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
}

var UnmarshalableListIdentity = Identity{Name: "list", Category: "lists"}

type UnmarshalableList struct {
	List
}

func NewUnmarshalableList() *UnmarshalableList {
	return &UnmarshalableList{List: List{}}
}

func (o *UnmarshalableList) Identity() Identity { return UnmarshalableListIdentity }

func (o *UnmarshalableList) UnmarshalJSON([]byte) error {
	return fmt.Errorf("error unmarshalling")
}

func (o *UnmarshalableList) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("error marshalling")
}

func (o *UnmarshalableList) Validate() Errors { return nil }
