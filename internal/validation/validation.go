package validation

func Validate(validateFn func(v *Validator)) map[string][]string {
	vv := NewValidator()
	validateFn(vv)

	if !vv.IsValid() {
		return vv.Errors
	}

	return nil
}
