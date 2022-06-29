package swfheader

type CompressionLevel uint8

const (
	CompressionNone CompressionLevel = 'F'
	CompressionZlib CompressionLevel = 'C'
	CompressionLzma CompressionLevel = 'Z'
)
