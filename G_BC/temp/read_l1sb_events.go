package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	l1cdm "mantle/test/contracts"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type L1cdmFailedRelayedMessage struct {
	MsgHash [32]byte
}

type L1cdmInitialized struct {
	Version uint8
}

type L1cdmMessageAllowed struct {
	XDomainCalldataHash [32]byte
}

type L1cdmMessageBlocked struct {
	XDomainCalldataHash [32]byte
}

type L1cdmOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
}

type L1cdmPaused struct {
	Account common.Address
}

type L1cdmRelayedMessage struct {
	MsgHash [32]byte
}

type L1cdmSentMessage struct {
	Target       common.Address
	Sender       common.Address
	Message      []byte
	MessageNonce *big.Int
	GasLimit     *big.Int
}

type L1cdmUnpaused struct {
	Account common.Address
}

func main1() {
	client, err := ethclient.Dial("wss://goerli.infura.io/ws/v3/a07ee340688643dd98ed571bfc1672fb")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol (ZRX) token address
	contractAddress := common.HexToAddress("0x7Bfe603647d5380ED3909F6f87580D0Af1B228B4")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(8424408),
		ToBlock:   big.NewInt(8424508),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(l1cdm.L1cdmABI)))
	if err != nil {
		log.Fatal(err)
	}

	FailedRelayedMessageSig := []byte("FailedRelayedMessage(bytes32)")
	FailedRelayedMessageSigHash := crypto.Keccak256Hash(FailedRelayedMessageSig)

	InitializedSig := []byte("Initialized(uint8)")
	InitializedSigHash := crypto.Keccak256Hash(InitializedSig)

	MessageAllowedSig := []byte("MessageAllowed(bytes32)")
	MessageAllowedSigHash := crypto.Keccak256Hash(MessageAllowedSig)

	MessageBlockedSig := []byte("MessageBlocked(bytes32)")
	MessageBlockedSigHash := crypto.Keccak256Hash(MessageBlockedSig)

	OwnershipTransferredSig := []byte("OwnershipTransferred(address,address)")
	OwnershipTransferredSigHash := crypto.Keccak256Hash(OwnershipTransferredSig)

	PausedSig := []byte("Paused(address)")
	PausedSigHash := crypto.Keccak256Hash(PausedSig)

	RelayedMessageSig := []byte("RelayedMessage(bytes32)")
	RelayedMessageSigHash := crypto.Keccak256Hash(RelayedMessageSig)

	SentMessageSig := []byte("SentMessage(address,address,bytes,uint256,uint256)")
	SentMessageSigHash := crypto.Keccak256Hash(SentMessageSig)

	UnpausedSig := []byte("Unpaused(address)")
	UnpausedSigHash := crypto.Keccak256Hash(UnpausedSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case FailedRelayedMessageSigHash.Hex():
			fmt.Printf("Log Name: FailedRelayedMessage\n")
			logdata, err := contractAbi.Unpack("FailedRelayedMessage", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(logdata)
		case InitializedSigHash.Hex():
			fmt.Printf("Log Name: Initialized\n")
			logdata, err := contractAbi.Unpack("Initialized", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(logdata)
		case MessageAllowedSigHash.Hex():
			fmt.Printf("Log Name: MessageAllowed\n")
			logdata, err := contractAbi.Unpack("MessageAllowed", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(logdata)
		case MessageBlockedSigHash.Hex():
			fmt.Printf("Log Name: MessageBlocked\n")
			logdata, err := contractAbi.Unpack("MessageBlocked", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(logdata)
		case OwnershipTransferredSigHash.Hex():
			fmt.Printf("Log Name: OwnershipTransferred\n")
			logdata, err := contractAbi.Unpack("OwnershipTransferred", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(logdata)
		case PausedSigHash.Hex():
			fmt.Printf("Log Name: Paused\n")
			logdata, err := contractAbi.Unpack("Paused", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(logdata)
		case RelayedMessageSigHash.Hex():
			fmt.Printf("Log Name: RelayedMessage\n")
			logdata, err := contractAbi.Unpack("RelayedMessage", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(logdata)
		case SentMessageSigHash.Hex():
			fmt.Printf("Log Name: SentMessage\n")
			logdata, err := contractAbi.Unpack("SentMessage", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(logdata)
		case UnpausedSigHash.Hex():
			fmt.Printf("Log Name: Unpaused\n")
			logdata, err := contractAbi.Unpack("Unpaused", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(logdata)
		}

		fmt.Printf("\n\n")
	}
}
