package ports

type DbPort interface {
	CloseDbConnection()
	AddToHistory(ans int32, op string) error
}
