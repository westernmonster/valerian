package conf

// Ecode *ecode.Config
type Config struct {
	DC *DC
}

// DC data center.
type DC struct {
	Num  int
	Desc string
}
