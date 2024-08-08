package constant

type Dialect string

func (d Dialect) String() string {
	return string(d)
}

const (
	PostgresDialect Dialect = "postgres"
)

type TableName string

func (t TableName) String() string {
	return string(t)
}

const (
	CategoriesTableName   TableName = "categories"
	MenuTableName         TableName = "menu"
	Transaction0TableName TableName = "transaction0"
	Transaction1TableName TableName = "transaction1"
	PaymentTableName      TableName = "payment"
	UsersTableName        TableName = "users"
	RoleTableName         TableName = "role"
)

type DateTimeFormat string

func (d DateTimeFormat) String() string {
	return string(d)
}

type RegexFormat string

func (d RegexFormat) String() string {
	return string(d)
}

type RoleId int

func (d RoleId) Int() int {
	return int(d)
}

const (
	AdminRoleIdUser   RoleId = 1
	RegularRoleIdUser RoleId = 2
)

type Action string

func (d Action) String() string {
	return string(d)
}

const (
	CreateActionUser Action = "create"
	ReadActionUser   Action = "read"
	UpdateActionUser Action = "update"
	DeleteActionUser Action = "delete"
)
