package github

func buildProjectV2ID(parts ...string) (string, error) {
	escaped := append([]string(nil), parts...)
	for index := range len(escaped) - 1 {
		escaped[index] = escapeIDPart(escaped[index])
	}
	return buildID(escaped...)
}

func parseProjectV2ID2(id string) (string, string, error) {
	first, second, err := parseID2(id)
	if err != nil {
		return "", "", err
	}
	return unescapeIDPart(first), second, nil
}

func parseProjectV2ID3(id string) (string, string, string, error) {
	first, second, third, err := parseID3(id)
	if err != nil {
		return "", "", "", err
	}
	return unescapeIDPart(first), unescapeIDPart(second), third, nil
}
