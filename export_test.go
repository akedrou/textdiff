package textdiff

func LineEdits(src string, edits []Edit) ([]Edit, error) {
	return lineEdits(src, edits)
}
