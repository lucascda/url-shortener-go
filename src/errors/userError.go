package apierrors

type UserAlreadyExistsError struct {
}

func (e UserAlreadyExistsError) Error() string {
	return "User already exists"
}
