package client

import (
	govclient "github.com/joshklop/monomer-cosmos-sdk/x/gov/client"
	"github.com/joshklop/monomer-cosmos-sdk/x/params/client/cli"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd)
