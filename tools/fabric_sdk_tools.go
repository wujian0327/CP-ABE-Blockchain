package tools

import (
	"crypto/sha256"
	"encoding/asn1"
	"fmt"
	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"math/big"
)

const (
	ConfigPath = "config/config.yaml"
	ChannelID  = "mychannel"
	OrgName    = "Org1"
	OrgAdmin   = "Admin"
)

// 调用Execute链码
func ExecuteChaincode(chaincodeId string, functionName string, args ...string) ([]byte, error) {
	client, sdk, err := GetChannelClient()
	defer sdk.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Printf("------------invoke chaincode: %v,%v,%v------------------\n", chaincodeId, functionName, args)
	byteArgs := make([][]byte, len(args))
	for i, v := range args {
		byteArgs[i] = []byte(v)
	}
	response, err := client.Execute(channel.Request{ChaincodeID: chaincodeId, Fcn: functionName, Args: byteArgs},
		channel.WithRetry(retry.DefaultChannelOpts),
		channel.WithTargetEndpoints(),
	)
	if err != nil {
		fmt.Printf("invoke to query funds: %s", err)
		return nil, err
	}
	fmt.Printf("invoke chaincode success,status:%v,Payload: %v\n", response.ChaincodeStatus, string(response.Payload))
	return response.Payload, nil
}

// 调用Query链码
func QueryChaincode(chaincodeId string, functionName string, args ...string) ([]byte, error) {
	client, sdk, err := GetChannelClient()
	defer sdk.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Printf("------------query chaincode: %v,%v,%v------------------\n", chaincodeId, functionName, args)
	byteArgs := make([][]byte, len(args))
	for i, v := range args {
		byteArgs[i] = []byte(v)
	}
	response, err := client.Query(channel.Request{ChaincodeID: chaincodeId, Fcn: functionName, Args: byteArgs},
		channel.WithRetry(retry.DefaultChannelOpts),
		channel.WithTargetEndpoints(),
	)
	if err != nil {
		fmt.Printf("Failed to query funds: %s", err)
		return nil, err
	}
	fmt.Printf("query chaincode success,status:%v,Payload: %v\n", response.ChaincodeStatus, string(response.Payload))
	return response.Payload, nil
}

// 获得ledgerClient
func GetLedgerClient() (*ledger.Client, *fabsdk.FabricSDK, error) {
	sdk := SDKInit()
	clientChannelContext := sdk.ChannelContext(ChannelID, fabsdk.WithUser(OrgAdmin), fabsdk.WithOrg(OrgName))

	// Ledger client
	ledgerClient, err := ledger.New(clientChannelContext)
	if err != nil {
		err := fmt.Errorf("failed to get ledgerClient: %s", err)
		return nil, nil, err
	}
	return ledgerClient, sdk, nil
}

// 获得contextClient对象
func GetContextClient() (*resmgmt.Client, *fabsdk.FabricSDK, error) {
	sdk := SDKInit()

	context := sdk.Context(fabsdk.WithUser(OrgAdmin), fabsdk.WithOrg(OrgName))
	resmgmtClient, err := resmgmt.New(context)
	if err != nil {
		err := fmt.Errorf("failed to get resmgmtClient: %s", err)
		return nil, nil, err
	}
	return resmgmtClient, sdk, nil
}

// 获得eventClient对象
func GetEventClient() (*event.Client, *fabsdk.FabricSDK, error) {
	sdk := SDKInit()

	clientChannelContext := sdk.ChannelContext(ChannelID, fabsdk.WithUser(OrgAdmin), fabsdk.WithOrg(OrgName))

	eventClient, err := event.New(clientChannelContext, event.WithBlockEvents())

	if err != nil {
		err := fmt.Errorf("failed to get eventClient: %s", err)
		return nil, nil, err
	}
	return eventClient, sdk, nil
}

// 获得ChannelClient对象
func GetChannelClient() (*channel.Client, *fabsdk.FabricSDK, error) {
	sdk := SDKInit()

	clientChannelContext := sdk.ChannelContext(ChannelID, fabsdk.WithUser(OrgAdmin), fabsdk.WithOrg(OrgName))
	client, err := channel.New(clientChannelContext)
	if err != nil {
		fmt.Printf("Failed to channel: %s", err)
		return nil, nil, err
	}
	return client, sdk, nil
}

// 获得fabricSDK
// 加载配置文件
func SDKInit() *fabsdk.FabricSDK {
	configProvider := config.FromFile(ConfigPath)
	sdk, err := fabsdk.New(configProvider)
	if err != nil {
		fmt.Printf("Failed to create new SDK: %s", err)
		return nil
	}
	return sdk
}

type asn1Header struct {
	Number       *big.Int
	PreviousHash []byte
	DataHash     []byte
}

func BlockHeaderBytes(b *cb.BlockHeader) []byte {
	asn1Header := asn1Header{
		PreviousHash: b.PreviousHash,
		DataHash:     b.DataHash,
		Number:       new(big.Int).SetUint64(b.Number),
	}
	result, err := asn1.Marshal(asn1Header)
	if err != nil {
		panic(err)
	}
	return result
}

// 计算当前hash
func BlockHeaderHash(b *cb.BlockHeader) []byte {
	sum := sha256.Sum256(BlockHeaderBytes(b))
	return sum[:]
}
