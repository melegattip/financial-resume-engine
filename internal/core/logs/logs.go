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
	ErrorCreatingTransaction = LogMessage{
		Message: "Error creating transaction",
	}
	ErrorListingTransactions = LogMessage{
		Message: "Error listing transactions",
	}
	ErrorGettingTransaction = LogMessage{
		Message: "Error getting transaction",
	}
	ErrorUpdatingTransaction = LogMessage{
		Message: "Error updating transaction",
	}
	ErrorCreatingCategory = LogMessage{
		Message: "Error creating category",
	}
	ErrorListingCategories = LogMessage{
		Message: "Error listing categories",
	}
	ErrorUpdatingCategory = LogMessage{
		Message: "Error updating category",
	}
	ErrorDeletingCategory = LogMessage{
		Message: "Error deleting category",
	}
	ErrorGeneratingReport = LogMessage{
		Message: "Error generating financial report",
	}
)
