package repository

type LogRepository interface {
	SaveLogs(body []byte) error
}
