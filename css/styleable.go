package css

// Styleable defines an interface for styleable elements.
type Styleable interface {
	// ID returns the id of the element.
	ID() string
	// Name returns the name of the element.
	Name() string
	// Classes returns all classes of the element.
	HasClass(c string) bool
	// ElementStyle returns style values which are directly bound this an instance
	ElementStyle() Style

	// Parent returns the styleable parent of the element or nil if it has no parent.
	Parent() Styleable

	Children() []Styleable
}

// ClassMap is a helper to manage classes for elements.
type ClassMap map[string]struct{}

// Add a style to the classmap if it is not already defined.
func (cm ClassMap) Add(s string) {
	cm[s] = struct{}{}
}

// Remove the style from the classmap if it is defined.
func (cm ClassMap) Remove(s string) {
	delete(cm, s)
}

// HasClass checks if the styleable has a specific class
func (cm ClassMap) HasClass(c string) bool {
	_, ok := cm[c]
	return ok
}

// IDAndClasses is a helper for implementing simple parts of the Styleable interface.
type IDAndClasses struct {
	id      string
	classes ClassMap
	es      Style
}

// ElementStyle returns style values which are directly bound this an instance
func (s *IDAndClasses) ElementStyle() Style {
	if s.es == nil {
		s.es = make(Style)
	}
	return s.es
}

// Sets a style value for this instance
func (s *IDAndClasses) SetProperty(p *Property, v interface{}) {
	if s.es == nil {
		s.es = make(Style)
	}
	s.es[p] = v
}

// AddClass adds a class to the element
func (s *IDAndClasses) AddClass(name string) {
	if s.classes == nil {
		s.classes = make(ClassMap)
	}
	s.classes.Add(name)
}

// RemoveClass removes the class from the element.
func (s *IDAndClasses) RemoveClass(name string) {
	if s.classes != nil {
		s.classes.Remove(name)
	}
}

// HasClass checks if the styleable has a specific class
func (s *IDAndClasses) HasClass(c string) bool {
	return s.classes.HasClass(c)
}

// SetID sets the hopefully unique id of the element.
func (s *IDAndClasses) SetID(id string) {
	s.id = id
}

// ID returns the currently set id of the element.
func (s *IDAndClasses) ID() string {
	return s.id
}
