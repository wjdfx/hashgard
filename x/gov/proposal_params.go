package gov

type ProposalParam struct {
	Subspace string `json:"subspace"`
	Key		 string	`json:"key"`
	Value	 string	`json:"value"`
}

type ProposalParams []ProposalParam
