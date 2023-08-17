package papershamir

func formatIntoBlocks(data string) string {
	for i := 4; i < len(data); i += 5 {
		data = data[:i] + " " + data[i:]
	}

	return data
}
