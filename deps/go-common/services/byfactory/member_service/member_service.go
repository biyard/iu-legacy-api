package byfactory

import (
	"context"
	"fmt"
	"math/big"

	"biyard.co/common/base"
	memberContract "biyard.co/common/contracts/member"
	profileContract "biyard.co/common/contracts/profile"
	StandardMTContract "biyard.co/common/contracts/standard-mt"
	standardTokenContract "biyard.co/common/contracts/standard-token"
	"biyard.co/common/logger"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/common"
)

type MemberService struct {
	clientEth      *client.Client
	klaytnEndpoint string
	senderAddress  string
}

type MemberDt struct {
	TokenIds      []string
	TokenUris     []string
	TokenBalances []string
	IsMember      bool
	Group         string
	Permission    string
	Balance       string
	Profile       string
}

func NewMemberService(endpoint, sender string) *MemberService {
	clientEth, clientErr := client.Dial(endpoint)
	if clientErr != nil {
		logger.Critical(context.TODO(), "load endpoint failed: %+v", clientErr)
	}

	return &MemberService{
		clientEth:      clientEth,
		klaytnEndpoint: endpoint,
		senderAddress:  sender,
	}
}

func (r *MemberService) ListGroups(context *base.Context, memberAddress common.Address) ([]string, []*big.Int, error) {
	memberInstance, memberErr := memberContract.NewMember(memberAddress, r.clientEth)

	if memberErr != nil {
		return nil, nil, memberErr
	}

	dt, err := memberInstance.ListGroups(&bind.CallOpts{
		From: common.HexToAddress(r.senderAddress),
	})
	if err != nil {
		return nil, nil, err
	}

	return dt.GroupNames, dt.Permissions, nil
}

func (r *MemberService) GetMember(context *base.Context, memberAddress common.Address, userAddress common.Address, standardTokenAddr common.Address, profileAddr common.Address, standardMTAddr common.Address) (*MemberDt, error) {
	isMember, groupName, permission, err := getMember(r.clientEth, memberAddress, r.senderAddress, userAddress)
	if err != nil {
		return nil, err
	}

	if !isMember {
		return &MemberDt{
			TokenIds:      nil,
			TokenUris:     nil,
			TokenBalances: nil,
			IsMember:      false,
			Group:         "",
			Permission:    "",
			Balance:       "",
			Profile:       "",
		}, nil
	}

	balance, err := getTokenBalance(r.clientEth, standardTokenAddr, r.senderAddress, userAddress)
	if err != nil {
		return nil, err
	}

	profile, err := getProfile(r.clientEth, profileAddr, r.senderAddress, userAddress)
	if err != nil {
		return nil, err
	}

	ids, uris, tokenBalances, err := getSbts(r.clientEth, standardMTAddr, r.senderAddress, userAddress)
	if err != nil {
		return nil, err
	}

	return &MemberDt{
		TokenIds:      ids,
		TokenUris:     uris,
		TokenBalances: tokenBalances,
		IsMember:      isMember,
		Group:         groupName,
		Permission:    permission,
		Balance:       fmt.Sprintf("%s", balance),
		Profile:       profile,
	}, nil
}

func getMember(clientEth *client.Client, memberAddress common.Address, senderAddress string, userAddress common.Address) (bool, string, string, error) {
	memberInstance, memberErr := memberContract.NewMember(memberAddress, clientEth)

	if memberErr != nil {
		return false, "", "", memberErr
	}

	dt, err := memberInstance.IsMember(&bind.CallOpts{
		From: common.HexToAddress(senderAddress),
	}, userAddress)
	if err != nil || !dt {
		return false, "", "", err
	}

	groups, err := memberInstance.ListGroups(&bind.CallOpts{
		From: common.HexToAddress(senderAddress),
	})
	if err != nil {
		return true, "", "", err
	}

	group, err := memberInstance.GetGroup(&bind.CallOpts{
		From: common.HexToAddress(senderAddress),
	}, userAddress)
	if err != nil {
		return true, "", "", err
	}

	permission := ""

	for i, v := range groups.GroupNames {
		if v == group {
			permission = fmt.Sprintf("%s", groups.Permissions[i])
			break
		}
	}

	return true, group, permission, nil
}

func getSbts(clientEth *client.Client, standardMtAddr common.Address, senderAddress string, userAddress common.Address) ([]string, []string, []string, error) {
	standardMtInstance, tokenErr := StandardMTContract.NewStandardMT(standardMtAddr, clientEth)

	if tokenErr != nil {
		return nil, nil, nil, tokenErr
	}

	ids := []string{"1", "2", "3"}

	tokenUris := []string{}
	tokenBalances := []string{}
	tokenIds := []string{}

	for _, v := range ids {
		tokenId := new(big.Int)
		tokenId, _ = tokenId.SetString(v, 10)

		uri, _ := standardMtInstance.GetTokenUri(&bind.CallOpts{
			From: common.HexToAddress(senderAddress),
		}, tokenId)

		value, _ := standardMtInstance.BalanceOf(&bind.CallOpts{
			From: common.HexToAddress(senderAddress),
		}, userAddress, tokenId)

		id := fmt.Sprintf("%s", v)

		tokenUris = append(tokenUris, uri)
		tokenBalances = append(tokenBalances, fmt.Sprintf("%s", value))
		tokenIds = append(tokenIds, id)
	}

	return tokenIds, tokenUris, tokenBalances, nil
}

func getTokenBalance(clientEth *client.Client, standardTokenAddr common.Address, senderAddress string, userAddress common.Address) (*big.Int, error) {
	standardInstance, tokenErr := standardTokenContract.NewStandardToken(standardTokenAddr, clientEth)

	if tokenErr != nil {
		return nil, tokenErr
	}

	return standardInstance.BalanceOf(&bind.CallOpts{
		From: common.HexToAddress(senderAddress),
	}, userAddress)
}

func getProfile(clientEth *client.Client, profileAddr common.Address, senderAddress string, userAddress common.Address) (string, error) {
	profileInstance, profileErr := profileContract.NewProfile(profileAddr, clientEth)

	if profileErr != nil {
		return "", profileErr
	}

	return profileInstance.GetProfile(&bind.CallOpts{
		From: common.HexToAddress(senderAddress),
	}, userAddress)
}
