package experience

import (
	"context"
	"strings"

	"biyard.co/common/logger"
	"github.com/klaytn/klaytn/accounts/abi"
)

func BiyardExperienceAbi() *abi.ABI {
	ABI := `[
		{
			"inputs": [
			  {
				"internalType": "address",
				"name": "stateAddr",
				"type": "address"
			  }
			],
			"stateMutability": "nonpayable",
			"type": "constructor"
		  },
		  {
			"anonymous": false,
			"inputs": [
			  {
				"indexed": false,
				"internalType": "address",
				"name": "tokenIDContractAddress",
				"type": "address"
			  },
			  {
				"indexed": false,
				"internalType": "address",
				"name": "nftContractAddress",
				"type": "address"
			  },
			  {
				"indexed": false,
				"internalType": "address",
				"name": "addr",
				"type": "address"
			  },
			  {
				"indexed": false,
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			  },
			  {
				"indexed": false,
				"internalType": "address",
				"name": "sender",
				"type": "address"
			  },
			  {
				"indexed": false,
				"internalType": "uint256",
				"name": "tokenID",
				"type": "uint256"
			  },
			  {
				"indexed": false,
				"internalType": "uint256",
				"name": "currentTime",
				"type": "uint256"
			  }
			],
			"name": "AddAccountEXPEvent",
			"type": "event"
		  },
		  {
			"anonymous": false,
			"inputs": [
			  {
				"indexed": false,
				"internalType": "address",
				"name": "contractAddress",
				"type": "address"
			  },
			  {
				"indexed": false,
				"internalType": "uint256",
				"name": "tokenID",
				"type": "uint256"
			  },
			  {
				"indexed": false,
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			  },
			  {
				"indexed": false,
				"internalType": "address",
				"name": "owner",
				"type": "address"
			  },
			  {
				"indexed": false,
				"internalType": "uint256",
				"name": "tokenLevel",
				"type": "uint256"
			  },
			  {
				"indexed": false,
				"internalType": "uint256",
				"name": "tokenExp",
				"type": "uint256"
			  },
			  {
				"indexed": false,
				"internalType": "uint256",
				"name": "currentTime",
				"type": "uint256"
			  }
			],
			"name": "AddNFTEXPEvent",
			"type": "event"
		  },
		  {
			"anonymous": false,
			"inputs": [
			  {
				"indexed": true,
				"internalType": "address",
				"name": "previousOwner",
				"type": "address"
			  },
			  {
				"indexed": true,
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			  }
			],
			"name": "OwnershipTransferred",
			"type": "event"
		  },
		  {
			"anonymous": false,
			"inputs": [
			  {
				"indexed": false,
				"internalType": "address",
				"name": "newVersion",
				"type": "address"
			  }
			],
			"name": "UpgradeEvent",
			"type": "event"
		  },
		  {
			"anonymous": false,
			"inputs": [
			  {
				"indexed": false,
				"internalType": "address",
				"name": "state",
				"type": "address"
			  }
			],
			"name": "UpgradeHookEvent",
			"type": "event"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "contractAddress",
				"type": "address"
			  },
			  {
				"internalType": "uint256",
				"name": "tokenID",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			  }
			],
			"name": "addNFTEXP",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "contractAddress",
				"type": "address"
			  },
			  {
				"internalType": "uint256",
				"name": "tokenID",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256[]",
				"name": "amounts",
				"type": "uint256[]"
			  }
			],
			"name": "addNFTEXPs",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "addr",
				"type": "address"
			  }
			],
			"name": "delTrustedParty",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "bytes",
				"name": "functionSignature",
				"type": "bytes"
			  }
			],
			"name": "executeCode",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "nftContractAddress",
				"type": "address"
			  }
			],
			"name": "getAccountEXP",
			"outputs": [
			  {
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "contractAddress",
				"type": "address"
			  },
			  {
				"internalType": "uint256",
				"name": "tokenID",
				"type": "uint256"
			  }
			],
			"name": "getMaxLevel",
			"outputs": [
			  {
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "contractAddress",
				"type": "address"
			  },
			  {
				"internalType": "uint256",
				"name": "tokenID",
				"type": "uint256"
			  }
			],
			"name": "getNFTEXP",
			"outputs": [
			  {
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "contractAddress",
				"type": "address"
			  },
			  {
				"internalType": "uint256",
				"name": "tokenID",
				"type": "uint256"
			  }
			],
			"name": "getRarity",
			"outputs": [
			  {
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "contractAddress",
				"type": "address"
			  },
			  {
				"internalType": "uint256",
				"name": "tokenID",
				"type": "uint256"
			  }
			],
			"name": "getUniqueCheck",
			"outputs": [
			  {
				"internalType": "bool",
				"name": "",
				"type": "bool"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [],
			"name": "getVersion",
			"outputs": [
			  {
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "uint256",
				"name": "version",
				"type": "uint256"
			  }
			],
			"name": "getVersionMaxLevelAddressList",
			"outputs": [
			  {
				"components": [
				  {
					"internalType": "address",
					"name": "addr",
					"type": "address"
				  },
				  {
					"internalType": "uint256",
					"name": "timestamp",
					"type": "uint256"
				  },
				  {
					"internalType": "uint256",
					"name": "nftNum",
					"type": "uint256"
				  }
				],
				"internalType": "struct MaxLevelAddress[]",
				"name": "",
				"type": "tuple[]"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "addr",
				"type": "address"
			  }
			],
			"name": "isTrusted",
			"outputs": [
			  {
				"internalType": "bool",
				"name": "",
				"type": "bool"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [],
			"name": "listTurstedParties",
			"outputs": [
			  {
				"internalType": "address[]",
				"name": "",
				"type": "address[]"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [],
			"name": "owner",
			"outputs": [
			  {
				"internalType": "address",
				"name": "",
				"type": "address"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [],
			"name": "renounceOwnership",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "contractAddress",
				"type": "address"
			  },
			  {
				"internalType": "uint256[]",
				"name": "tokenIds",
				"type": "uint256[]"
			  },
			  {
				"internalType": "uint256[]",
				"name": "evolutions",
				"type": "uint256[]"
			  }
			],
			"name": "setBulkEvolution",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "contractAddress",
				"type": "address"
			  },
			  {
				"internalType": "uint256[]",
				"name": "tokenIDs",
				"type": "uint256[]"
			  },
			  {
				"internalType": "uint256[]",
				"name": "levels",
				"type": "uint256[]"
			  }
			],
			"name": "setBulkLevel",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "contractAddress",
				"type": "address"
			  },
			  {
				"internalType": "uint256[]",
				"name": "tokenIDs",
				"type": "uint256[]"
			  },
			  {
				"internalType": "uint256[]",
				"name": "rarities",
				"type": "uint256[]"
			  }
			],
			"name": "setBulkRarity",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "uint256",
				"name": "level",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "experience",
				"type": "uint256"
			  }
			],
			"name": "setExceptionalEXP",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "uint256",
				"name": "x",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "y",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "z",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "v",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "s",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "t",
				"type": "uint256"
			  }
			],
			"name": "setGeneralEXP",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "uint256",
				"name": "maxLevel",
				"type": "uint256"
			  }
			],
			"name": "setMaxLevel",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "contractAddress",
				"type": "address"
			  },
			  {
				"internalType": "uint256",
				"name": "tokenID",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "maxLevel",
				"type": "uint256"
			  }
			],
			"name": "setNFTMaxLevel",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "contractAddress",
				"type": "address"
			  },
			  {
				"internalType": "uint256",
				"name": "tokenID",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "rarity",
				"type": "uint256"
			  }
			],
			"name": "setRarity",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "addr",
				"type": "address"
			  }
			],
			"name": "setTrustedParty",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "contractAddress",
				"type": "address"
			  },
			  {
				"internalType": "uint256",
				"name": "tokenID",
				"type": "uint256"
			  },
			  {
				"internalType": "bool",
				"name": "check",
				"type": "bool"
			  }
			],
			"name": "setUniqueCheck",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "contractAddress",
				"type": "address"
			  },
			  {
				"internalType": "uint256",
				"name": "tokenID",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256[]",
				"name": "levels",
				"type": "uint256[]"
			  },
			  {
				"internalType": "uint256[]",
				"name": "exps",
				"type": "uint256[]"
			  }
			],
			"name": "setUniqueNFTEXP",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "uint256",
				"name": "version",
				"type": "uint256"
			  }
			],
			"name": "setVersion",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "uint256",
				"name": "elevel",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "level",
				"type": "uint256"
			  },
			  {
				"internalType": "address",
				"name": "contractAddress",
				"type": "address"
			  },
			  {
				"internalType": "uint256",
				"name": "tokenID",
				"type": "uint256"
			  }
			],
			"name": "simulateExperiencePoints",
			"outputs": [
			  {
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "newOwner",
				"type": "address"
			  }
			],
			"name": "transferOwnership",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "newVersion",
				"type": "address"
			  }
			],
			"name": "upgrade",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "state",
				"type": "address"
			  }
			],
			"name": "upgradeHook",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "stateAddr",
				"type": "address"
			  },
			  {
				"internalType": "address",
				"name": "stateAddrV2",
				"type": "address"
			  }
			],
			"stateMutability": "nonpayable",
			"type": "constructor"
		  },
		  {
			"anonymous": false,
			"inputs": [
			  {
				"indexed": false,
				"internalType": "address",
				"name": "senderAddress",
				"type": "address"
			  },
			  {
				"indexed": false,
				"internalType": "address",
				"name": "userAddress",
				"type": "address"
			  },
			  {
				"indexed": false,
				"internalType": "uint256",
				"name": "exp",
				"type": "uint256"
			  }
			],
			"name": "AddAccountEXPEvent",
			"type": "event"
		  },
		  {
			"anonymous": false,
			"inputs": [
			  {
				"indexed": false,
				"internalType": "address",
				"name": "userAddress",
				"type": "address"
			  },
			  {
				"indexed": false,
				"internalType": "string",
				"name": "missionName",
				"type": "string"
			  },
			  {
				"indexed": false,
				"internalType": "uint256",
				"name": "progressDate",
				"type": "uint256"
			  },
			  {
				"indexed": false,
				"internalType": "uint256",
				"name": "exp",
				"type": "uint256"
			  }
			],
			"name": "ClaimAccountEXPEvent",
			"type": "event"
		  },
		  {
			"anonymous": false,
			"inputs": [
			  {
				"indexed": false,
				"internalType": "address",
				"name": "userAddress",
				"type": "address"
			  },
			  {
				"indexed": false,
				"internalType": "string[]",
				"name": "missionName",
				"type": "string[]"
			  },
			  {
				"indexed": false,
				"internalType": "uint256[]",
				"name": "progressDate",
				"type": "uint256[]"
			  },
			  {
				"indexed": false,
				"internalType": "uint256[]",
				"name": "exp",
				"type": "uint256[]"
			  }
			],
			"name": "ClaimAccountEvent",
			"type": "event"
		  },
		  {
			"anonymous": false,
			"inputs": [
			  {
				"indexed": false,
				"internalType": "address",
				"name": "userAddress",
				"type": "address"
			  },
			  {
				"indexed": false,
				"internalType": "string",
				"name": "missionName",
				"type": "string"
			  },
			  {
				"indexed": false,
				"internalType": "uint256",
				"name": "tokenNum",
				"type": "uint256"
			  },
			  {
				"indexed": false,
				"internalType": "uint256",
				"name": "progressDate",
				"type": "uint256"
			  },
			  {
				"indexed": false,
				"internalType": "uint256",
				"name": "exp",
				"type": "uint256"
			  }
			],
			"name": "DistributeAccountEvent",
			"type": "event"
		  },
		  {
			"anonymous": false,
			"inputs": [
			  {
				"indexed": false,
				"internalType": "address",
				"name": "senderAddress",
				"type": "address"
			  },
			  {
				"indexed": false,
				"internalType": "address",
				"name": "userAddress",
				"type": "address"
			  },
			  {
				"indexed": false,
				"internalType": "uint256",
				"name": "exp",
				"type": "uint256"
			  }
			],
			"name": "UpdateAccountEXPEvent",
			"type": "event"
		  },
		  {
			"anonymous": false,
			"inputs": [
			  {
				"indexed": false,
				"internalType": "address",
				"name": "newVersion",
				"type": "address"
			  }
			],
			"name": "UpgradeEvent",
			"type": "event"
		  },
		  {
			"anonymous": false,
			"inputs": [
			  {
				"indexed": false,
				"internalType": "address",
				"name": "state",
				"type": "address"
			  }
			],
			"name": "UpgradeHookEvent",
			"type": "event"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "userAddress",
				"type": "address"
			  },
			  {
				"components": [
				  {
					"internalType": "string",
					"name": "key",
					"type": "string"
				  },
				  {
					"internalType": "string",
					"name": "missionInfo",
					"type": "string"
				  },
				  {
					"internalType": "uint256",
					"name": "exp",
					"type": "uint256"
				  },
				  {
					"internalType": "uint256",
					"name": "progressDate",
					"type": "uint256"
				  },
				  {
					"internalType": "uint256",
					"name": "startDate",
					"type": "uint256"
				  },
				  {
					"internalType": "uint256",
					"name": "endDate",
					"type": "uint256"
				  }
				],
				"internalType": "struct ExperienceActivityData",
				"name": "data",
				"type": "tuple"
			  }
			],
			"name": "addAccountActivities",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "userAddress",
				"type": "address"
			  },
			  {
				"internalType": "uint256",
				"name": "exp",
				"type": "uint256"
			  }
			],
			"name": "addAccountEXP",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [],
			"name": "claimAccountEXPs",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "addr",
				"type": "address"
			  }
			],
			"name": "delTrustedParty",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "nftContractAddress",
				"type": "address"
			  },
			  {
				"internalType": "uint256",
				"name": "nftNum",
				"type": "uint256"
			  },
			  {
				"internalType": "uint256",
				"name": "e",
				"type": "uint256"
			  }
			],
			"name": "distributeAccountEXP",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "bytes",
				"name": "functionSignature",
				"type": "bytes"
			  }
			],
			"name": "executeCode",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [],
			"name": "forwarder",
			"outputs": [
			  {
				"internalType": "address",
				"name": "",
				"type": "address"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "sender",
				"type": "address"
			  }
			],
			"name": "getAccountEXP",
			"outputs": [
			  {
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [],
			"name": "getExperienceAddress",
			"outputs": [
			  {
				"internalType": "address",
				"name": "",
				"type": "address"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "userAddress",
				"type": "address"
			  }
			],
			"name": "getProfile",
			"outputs": [
			  {
				"components": [
				  {
					"internalType": "address",
					"name": "userAddress",
					"type": "address"
				  },
				  {
					"internalType": "uint256",
					"name": "pendingEXP",
					"type": "uint256"
				  },
				  {
					"internalType": "uint256",
					"name": "totalEXP",
					"type": "uint256"
				  },
				  {
					"internalType": "string",
					"name": "imageUrl",
					"type": "string"
				  },
				  {
					"internalType": "string",
					"name": "profileName",
					"type": "string"
				  },
				  {
					"internalType": "uint256",
					"name": "discordId",
					"type": "uint256"
				  },
				  {
					"internalType": "string",
					"name": "data",
					"type": "string"
				  }
				],
				"internalType": "struct Profile",
				"name": "",
				"type": "tuple"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [],
			"name": "getStateAddress",
			"outputs": [
			  {
				"internalType": "address",
				"name": "",
				"type": "address"
			  },
			  {
				"internalType": "address",
				"name": "",
				"type": "address"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [],
			"name": "getVersion",
			"outputs": [
			  {
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "addr",
				"type": "address"
			  }
			],
			"name": "isTrusted",
			"outputs": [
			  {
				"internalType": "bool",
				"name": "",
				"type": "bool"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [],
			"name": "listAccountActivities",
			"outputs": [
			  {
				"components": [
				  {
					"internalType": "string",
					"name": "key",
					"type": "string"
				  },
				  {
					"internalType": "string",
					"name": "missionInfo",
					"type": "string"
				  },
				  {
					"internalType": "uint256",
					"name": "exp",
					"type": "uint256"
				  },
				  {
					"internalType": "uint256",
					"name": "progressDate",
					"type": "uint256"
				  },
				  {
					"internalType": "uint256",
					"name": "startDate",
					"type": "uint256"
				  },
				  {
					"internalType": "uint256",
					"name": "endDate",
					"type": "uint256"
				  }
				],
				"internalType": "struct ExperienceActivityData[]",
				"name": "",
				"type": "tuple[]"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [],
			"name": "listAccountEXPActivities",
			"outputs": [
			  {
				"components": [
				  {
					"internalType": "string",
					"name": "key",
					"type": "string"
				  },
				  {
					"internalType": "string",
					"name": "missionInfo",
					"type": "string"
				  },
				  {
					"internalType": "uint256",
					"name": "exp",
					"type": "uint256"
				  },
				  {
					"internalType": "uint256",
					"name": "progressDate",
					"type": "uint256"
				  },
				  {
					"internalType": "uint256",
					"name": "startDate",
					"type": "uint256"
				  },
				  {
					"internalType": "uint256",
					"name": "endDate",
					"type": "uint256"
				  }
				],
				"internalType": "struct ExperienceActivityData[]",
				"name": "",
				"type": "tuple[]"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [],
			"name": "listClaimAccountActivities",
			"outputs": [
			  {
				"components": [
				  {
					"internalType": "string",
					"name": "key",
					"type": "string"
				  },
				  {
					"internalType": "string",
					"name": "missionInfo",
					"type": "string"
				  },
				  {
					"internalType": "uint256",
					"name": "exp",
					"type": "uint256"
				  },
				  {
					"internalType": "uint256",
					"name": "progressDate",
					"type": "uint256"
				  },
				  {
					"internalType": "uint256",
					"name": "startDate",
					"type": "uint256"
				  },
				  {
					"internalType": "uint256",
					"name": "endDate",
					"type": "uint256"
				  }
				],
				"internalType": "struct ExperienceActivityData[]",
				"name": "",
				"type": "tuple[]"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [],
			"name": "listTurstedParties",
			"outputs": [
			  {
				"internalType": "address[]",
				"name": "",
				"type": "address[]"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [],
			"name": "owner",
			"outputs": [
			  {
				"internalType": "address",
				"name": "",
				"type": "address"
			  }
			],
			"stateMutability": "view",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "experienceAddress",
				"type": "address"
			  }
			],
			"name": "setExperienceAddress",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "forwarder",
				"type": "address"
			  }
			],
			"name": "setForwarder",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "stateAddress",
				"type": "address"
			  },
			  {
				"internalType": "address",
				"name": "stateAddressV2",
				"type": "address"
			  }
			],
			"name": "setStateAddress",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "addr",
				"type": "address"
			  }
			],
			"name": "setTrustedParty",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "uint256",
				"name": "version",
				"type": "uint256"
			  }
			],
			"name": "setVersion",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  },
		  {
			"inputs": [
			  {
				"internalType": "address",
				"name": "userAddress",
				"type": "address"
			  },
			  {
				"internalType": "uint256",
				"name": "exp",
				"type": "uint256"
			  }
			],
			"name": "updateAccountEXP",
			"outputs": [],
			"stateMutability": "nonpayable",
			"type": "function"
		  }
	  ]`

	parsed, err := abi.JSON(strings.NewReader(ABI))

	if err != nil {
		logger.Criticalf(context.TODO(), "json parsing error: %+v", err)
	}

	return &parsed
}
