package transcoders

// Protocol ...
type Protocol struct {
	Name       string
	Transcoder Transcoder
}

// Protocols ...
var Protocols = []Protocol{
	Protocol{"ip4", TranscoderIP4},
	Protocol{"ip6", TranscoderIP6},
	// ....
}
