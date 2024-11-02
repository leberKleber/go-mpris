package mpris

func msgOrEmpty(err error) string {
	if err == nil {
		return ""
	}

	return err.Error()

}
