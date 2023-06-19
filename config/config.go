package config

type db struct {
	DSN string
}

var DB = db{
	DSN: "postgres://alpha_animal:pa55word@localhost:5432/animals",
}
