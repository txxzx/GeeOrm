# GeeOrm
用Go实现简易的数据库操作，参照Gorm Xorm



表名(table name) —— 结构体名(struct name)
字段名和字段类型 —— 成员变量和类型。
额外的约束条件(例如非空、主键等) —— 成员变量的Tag（Go 语言通过 Tag 实现，Java、Python 等语言通过注解实现）
