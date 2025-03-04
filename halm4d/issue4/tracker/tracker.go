package tracker

type Track struct {
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	ReleaseDate string `json:"release_date"`
}

func Convert(tracks []Track) [][]string {
	var csvTracks [][]string
	for _, v := range tracks {
		csvTracks = append(csvTracks, []string{
			v.Title,
			v.Type,
			v.Description,
			v.Genre,
			v.ReleaseDate,
		})
	}
	return csvTracks
}
