package passport

type UniqueId interface {
	Generate() (string, error)
	Hash(mask int) int
}
