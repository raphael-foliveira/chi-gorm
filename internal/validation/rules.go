package validation

type Validator struct {
	Errors map[string][]string
}

func NewValidator() *Validator {
	return &Validator{
		Errors: make(map[string][]string),
	}
}

func (v *Validator) Add(key string, messages ...string) {
	v.Errors[key] = append(v.Errors[key], messages...)
}

func (v *Validator) Set(key string, messages ...string) {
	v.Errors[key] = messages
}

func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) Check(key string, ok bool, message string) {
	if !ok {
		v.Add(key, message)
	}
}
