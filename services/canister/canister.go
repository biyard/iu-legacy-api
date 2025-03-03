package canister

import (
	"context"
	"fmt"

	"biyard.co/common/utils/httpclient"
)

type CanisterAPI struct {
	ctx context.Context
	cli *httpclient.HTTPClient
}

func NewCanisterAPI(ctx context.Context, endpoint string) (*CanisterAPI, error) {
	cli, err := httpclient.NewHttpClient(httpclient.ConstantBackoff(3, 1), httpclient.Limiter(1000), httpclient.Headers(map[string][]string{}), httpclient.Endpoint(endpoint))

	if err != nil {
		return nil, err
	}

	return &CanisterAPI{
		ctx: ctx,
		cli: cli,
	}, nil
}

type ProposalSummary struct {
	ID           uint    `json:"id"`
	Proposer     string  `json:"proposer"`
	ProposalType uint    `json:"proposal_type"`
	Title        string  `json:"title"`
	Status       string  `json:"status"`
	Result       *string `json:"result"`
	Votes        uint    `json:"votes"`
	Deadline     int64   `json:"deadline"`
}

type ProposalDetail struct {
	ID           uint   `json:"id"`
	Description  string `json:"description"`
	ExternalLink string `json:"external_link"`
	CreatedAt    int64  `json:"created_at"`
	Accepts      uint   `json:"accepts"`
	Rejects      uint   `json:"rejects"`
	Abstains     uint   `json:"abstains"`
	VotingPowers uint   `json:"voting_powers"`

	// skip metadata parsing
}

type Proposal struct {
	Summary ProposalSummary `json:"summary"`
	Detail  ProposalDetail  `json:"detail"`
}

func (r *CanisterAPI) GetProposal(ctx context.Context, proposalID uint) (*Proposal, error) {
	res := &Proposal{}
	_, err := r.cli.Get(ctx, fmt.Sprintf("/dao/proposals/%v", proposalID), nil, res)

	return res, err
}
