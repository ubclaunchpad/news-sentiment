package db

func appendIfMissing(sources []Source, source Source) []Source {
	for _, s := range sources {
		if s.URL == source.URL {
			return sources
		}
	}
	return append(sources, source)
}
