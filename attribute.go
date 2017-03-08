// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

// An AttributeSpecifiable is the interface an object must implement in order to access specification of its attributes.
type AttributeSpecifiable interface {
	SpecificationForAttribute(string) AttributeSpecification
	AttributeSpecifications() map[string]AttributeSpecification
	Version() float64
}

// UniqueScope is a the type used to define uniqueness.
type UniqueScope int

const (
	// LocallyUnique represents the uniqueness in a particular context.
	LocallyUnique UniqueScope = iota + 1

	// GloballyUnique represents the absolute uniqueness.
	GloballyUnique
)

// An AttributeSpecification represents all the metadata of an attribute.
//
// This information is coming from the Monolithe Specifications.
type AttributeSpecification struct {

	// AllowedChars is a regexp that will be used to validate
	// what value a string attribute can take.
	//
	// This is enforced by elemental.
	AllowedChars string

	// AllowedChoices is a list of possible values for an attribute.
	//
	// This is enforced by elemental.
	AllowedChoices []string

	// Autogenerated defines if the attribute is autogenerated by the server.
	// It can be used in conjunction with ReadOnly.
	//
	// This is not enforced by elemental. You must write your own business logic to honor this.
	Autogenerated bool

	// Availability is reserved for later use.
	Availability string

	// Channel is reserved for later use.
	Channel string

	// CreationOnly defines if the attribute can be set only during creation.
	//
	// This is not enforced by elemental. You must write your own business logic to honor this.
	CreationOnly bool

	// DefaultOrder defines if the attribute is used as the default ordering key.
	//
	// This is not enforced by elemental. You must write your own business logic to honor this.
	DefaultOrder bool

	// DefaultValue holds the default value declared in specification.
	DefaultValue interface{}

	// Deprecated defines if the attribute is deprecated.
	Deprecated bool

	// Description contains the description of the attribute.
	Description string

	// Exposed defines if the attribute is exposed through the north bound API.
	Exposed bool

	// Filterable defines if it is possible to filter based on this attribute.
	//
	// This is not enforced by elemental. You must write your own business logic to honor this.
	Filterable bool

	// ForeignKey defines if the attribute is a foreign key.
	ForeignKey bool

	// Format defines the format of the attribute.
	// Monolithe defines various formats, but it is up to your implementation
	// to decide what values it can take.
	//
	// This is not enforced by elemental. You must write your own business logic to honor this.
	Format string

	// Getter defines if the attribute needs to define a getter method.
	// This is useful if you can to define an Interface based on this attribute.
	Getter bool

	// Identifier defines if the attribute is used the access key from the
	// northbound API.
	Identifier bool

	// Index defines if the attribute is indexed or not.
	//
	// This is not enforced by elemental. You must write your own business logic to honor this.
	Index bool

	// MaxLength defines what is the maximun length of the attribute.
	// This only makes sense if the type is a string.
	//
	// This is enforced by elemental.
	MaxLength uint

	// MaxValue defines what is the maximun value of the attribute.
	// This only makes sense if the type has a numeric type.
	//
	// This is enforced by elemental.
	MaxValue float64

	// MinLength defines what is the minimum length of the attribute.
	// This only makes sense if the type is a string.
	//
	// This is enforced by elemental.
	MinLength uint

	// MinValue defines what is the minimum value of the attribute.
	// This only makes sense if the type has a numeric type.
	//
	// This is enforced by elemental.
	MinValue float64

	// Name defines what is the name of the attribute.
	// This will be the raw Monolithe Specification name, without
	// Go translation.
	Name string

	// Orderable defines if it is possible to order based on the value of this attribute.
	//
	// This is not enforced by elemental. You must write your own business logic to honor this.
	Orderable bool

	// PrimaryKey defines if the attribute is used as a primary key.
	PrimaryKey bool

	// ReadOnly defines if the attribute is read only.
	//
	// This is not enforced by elemental. You must write your own business logic to honor this.
	ReadOnly bool

	// Required defines is the attribute must be set or not.
	//
	// This is enforced by elemental.
	Required bool

	// Setter defines if the attribute needs to define a setter method.
	// This is useful if you can to define an Interface based on this attribute.
	Setter bool

	// Stored defines if the attribute will be stored in the northbound API.
	//
	// If this is true, the backend tags will be generated by Monolithe.
	Stored bool

	// SubType defines the Monolithe Subtype.
	SubType string

	// Transient defines if the attributes is transient or not.
	//
	// This is not enforced by elemental. You must write your own business logic to honor this.
	Transient bool

	// Type defines the raw Monolithe type.
	Type string

	// Unique defines if the value of the attribute must be unique.
	//
	// This is not enforced by elemental. You must write your own business logic to honor this.
	Unique bool

	// UniqueScope defines the scope of the uniqueness of the value of the attribute.
	// This makes sense only if Unique is true.
	//
	// This is not enforced by elemental. You must write your own business logic to honor this.
	UniqueScope UniqueScope
}
