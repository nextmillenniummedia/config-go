package params

// Params for field contained in tag
// Example:
//
//	config:"require,format=url,field=port,required,enum=one|two|three,doc='This field needed for service'"
type Params struct {
	// Field name
	Field string
	// Splitter for array value
	Splitter string
	// Format for field
	Format string
	// Field is required or not
	Required bool
	// List for enum
	Enum []string
	// Documentation
	Doc string
}
