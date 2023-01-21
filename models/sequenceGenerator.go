package models

func InitSequenceGenerator() error {
	return db.Exec("CREATE SEQUENCE IF NOT EXISTS num_gen_seq OWNED BY NONE").Error
}

func nextInSequence() (uint, error) {
	var num uint
	err := db.Raw("SELECT nextval('num_gen_seq') as num").Scan(&num).Error
	return num, err
}
