package requests

import LogEntry "userLoggingProject/internal/features/logs/entity"

type AddLogRequest struct {
	UserId string
	Log    LogEntry.LogEntry
}
