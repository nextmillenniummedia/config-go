package params

// Params for field contained in tag
// Example:
//
//	config:"require,format=url,required,enum=one|two|three,doc='This field needed for service'"
type Params struct {
	// Format for field
	Format string
	// Field is required or not
	Required bool
	// List for enum
	Enum []string
	// Documentation
	Doc string
}
