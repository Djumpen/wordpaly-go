package entities

type Dictionary struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	CreatedAt   string `db:"created_at" json:",omitempty"`
}

type DictionaryResp struct {
	ID          int
	Title       string
	Description string
}
