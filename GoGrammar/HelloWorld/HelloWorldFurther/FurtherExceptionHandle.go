package helloworldfurther

type FileHandleError struct {
	error_msg string
}

func (err *FileHandleError) Error() string {
	return err.error_msg
}
