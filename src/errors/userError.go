package apierrors

type UserAlreadyExistsError struct {
}

type UserNotFoundError struct {
}

func (e UserNotFoundError) Error() string {
	return "User not found"
}

func (e UserAlreadyExistsError) Error() string {
	return "User already exists"
}
