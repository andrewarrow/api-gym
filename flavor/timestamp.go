package flavor

type TimestampFlavor struct{}

func (id TimestampFlavor) Name() string {
	return "timestamp"
}

func (id TimestampFlavor) Generate() string {
	return "2022-04-18T06:52:29.940Z"
}

func (id TimestampFlavor) Flavor() string {
	return "string"
}
