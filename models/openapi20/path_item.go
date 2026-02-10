package openapi20

// PathItem describes operations available on a single path.
// https://swagger.io/specification/v2/#path-item-object
type PathItem struct {
	Node // embedded - provides VendorExtensions and Trix

	ref        string
	get        *Operation
	put        *Operation
	post       *Operation
	delete     *Operation
	options    *Operation
	head       *Operation
	patch      *Operation
	parameters []*ParameterRef
}

func (pi *PathItem) Ref() string                 { return pi.ref }
func (pi *PathItem) Get() *Operation             { return pi.get }
func (pi *PathItem) Put() *Operation             { return pi.put }
func (pi *PathItem) Post() *Operation            { return pi.post }
func (pi *PathItem) Delete() *Operation          { return pi.delete }
func (pi *PathItem) Options() *Operation         { return pi.options }
func (pi *PathItem) Head() *Operation            { return pi.head }
func (pi *PathItem) Patch() *Operation           { return pi.patch }
func (pi *PathItem) Parameters() []*ParameterRef { return pi.parameters }

// NewPathItem creates a new PathItem instance.
func NewPathItem(
	ref string,
	get, put, post, del, options, head, patch *Operation,
	parameters []*ParameterRef,
) *PathItem {
	return &PathItem{
		ref: ref,
		get: get, put: put, post: post, delete: del,
		options: options, head: head, patch: patch,
		parameters: parameters,
	}
}
