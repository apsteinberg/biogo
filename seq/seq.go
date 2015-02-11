package seq

type Sequence struct {
	Id, Name    string
	Seq         []byte
	Annotations map[string]string
}
