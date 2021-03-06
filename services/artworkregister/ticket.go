package artworkregister

// Ticket represents artwork registration ticket.
type Ticket struct {
	Image                    []byte
	Name                     string
	Description              *string
	Keywords                 *string
	SeriesName               *string
	IssuedCopies             int
	YoutubeURL               *string
	ArtistPastelID           string
	ArtistPastelIDPassphrase string
	ArtistName               string
	ArtistWebsiteURL         *string
	SpendableAddress         string
	MaximumFee               float64
}
