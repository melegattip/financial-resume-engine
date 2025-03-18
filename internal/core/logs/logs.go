package logs

type Tags map[string]interface{}

type LogMessage struct {
	Message string
}

func (l LogMessage) GetMessage() string {
	return l.Message
}

var (
	ErrorLoadingConfiguration = LogMessage{
		Message: "Error loading configuration",
	}
)
