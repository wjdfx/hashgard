package issue

type QueryIssueParams struct {
	IssueID string
}

// creates a new instance of QueryIssueParams
func NewQueryIssueParams(IssueID string) QueryIssueParams {
	return QueryIssueParams{
		IssueID: IssueID,
	}
}
