package test

import (
	"CP-ABE-Blockchain/tools"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset/kvrwset"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric-protos-go/peer"
	"testing"
)

func TestGetBlockData(t *testing.T) {
	ledgerClient, sdk, err := tools.GetLedgerClient()
	defer sdk.Close()
	if err != nil {
		t.Error(err)
		return
	}
	response, err := ledgerClient.QueryInfo()
	if err != nil {
		t.Error(err)
		return
	}
	//区块基本信息
	blockHeight := response.BCI.Height
	block, err := ledgerClient.QueryBlock(blockHeight - 1)
	if err != nil {
		t.Error(err)
		return
	}
	// fmt.Printf("block.Data.GetData(): %v\n", block.Data.GetData())
	fmt.Printf("block.Header.GetNumber(): %v\n", block.Header.GetNumber())
	//des, i := block.Descriptor()
	//fmt.Printf("des: %x\n", des)
	//fmt.Printf("i: %v\n", i)
	fmt.Print("=====================\n")

	for _, v := range block.Data.GetData() {
		var envelope common.Envelope

		if err = proto.Unmarshal(v, &envelope); err != nil {
			t.Error(err)
			return
		}
		envelope.GetPayload()

		//获得payload
		var payload common.Payload
		err = proto.Unmarshal(envelope.GetPayload(), &payload)
		if err != nil {
			t.Error(err)
			return
		}

		//获得channelHeader
		var channelHeader common.ChannelHeader
		err = proto.Unmarshal(payload.Header.GetChannelHeader(), &channelHeader)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Printf("channelHeader.ChannelId: %v\n", channelHeader.ChannelId)
		fmt.Printf("channelHeader.TxId: %v\n", channelHeader.TxId)
		fmt.Printf("channelHeader.Timestamp: %v\n", channelHeader.Timestamp)
		fmt.Printf("channelHeader.Type: %v\n", common.HeaderType_name[channelHeader.Type])

		//签名header
		signatureHeader := &common.SignatureHeader{}
		err := proto.Unmarshal(payload.Header.SignatureHeader, signatureHeader)
		if err != nil {
			t.Error(err)
			return
		}
		creator := &msp.SerializedIdentity{}
		err = proto.Unmarshal(signatureHeader.Creator, creator)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Printf("creator.Mspid: %v\n", creator.Mspid)

		//获得交易数据
		var transAction peer.Transaction
		err = proto.Unmarshal(payload.GetData(), &transAction)
		if err != nil {
			t.Error(err)
			return
		}
		//多个transactionAction
		for _, transactionAction := range transAction.GetActions() {

			var chaincodeActionPayload peer.ChaincodeActionPayload
			err = proto.Unmarshal(transactionAction.GetPayload(), &chaincodeActionPayload)
			if err != nil {
				t.Error(err)
				return
			}
			//获得背书节点名称
			for _, e := range chaincodeActionPayload.Action.Endorsements {
				endorser := &msp.SerializedIdentity{}
				err = proto.Unmarshal(e.Endorser, endorser)
				if err != nil {
					t.Error(err)
					return
				}
				fmt.Printf("endorser.Mspid: %v\n", endorser.Mspid)
			}

			//ChaincodeProposalPayload 解析输入 链码，参数
			var chaincodeProposalPayload peer.ChaincodeProposalPayload
			err = proto.Unmarshal(chaincodeActionPayload.GetChaincodeProposalPayload(), &chaincodeProposalPayload)
			if err != nil {
				t.Error(err)
				return
			}
			var input peer.ChaincodeInvocationSpec
			err = proto.Unmarshal(chaincodeProposalPayload.GetInput(), &input)
			if err != nil {
				t.Error(err)
				return
			}
			//fmt.Printf("input.ChaincodeSpec.ChaincodeId.Name: %v\n", input.ChaincodeSpec.ChaincodeId.Name)
			//fmt.Printf("input.ChaincodeSpec.ChaincodeId.Version: %v\n", input.ChaincodeSpec.ChaincodeId.Version)
			for i, v := range input.ChaincodeSpec.Input.Args {
				fmt.Printf("第%v个参数，值：%s\n", i, v)
			}

			//一次交易的内部输入输出
			var proposalResponsePayload peer.ProposalResponsePayload
			err = proto.Unmarshal(chaincodeActionPayload.Action.ProposalResponsePayload, &proposalResponsePayload)
			if err != nil {
				t.Error(err)
				return
			}
			fmt.Printf("proposalResponsePayload.ProposalHash: %x\n", proposalResponsePayload.ProposalHash)
			var chaincodeAction peer.ChaincodeAction
			err = proto.Unmarshal(proposalResponsePayload.Extension, &chaincodeAction)
			if err != nil {
				t.Error(err)
				return
			}
			fmt.Printf("chaincodeAction.ChaincodeId.Name: %v\n", chaincodeAction.ChaincodeId.Name)
			fmt.Printf("chaincodeAction.ChaincodeId.Version: %v\n", chaincodeAction.ChaincodeId.Version)
			var txrwset rwset.TxReadWriteSet
			err = proto.Unmarshal(chaincodeAction.Results, &txrwset)
			if err != nil {
				t.Error(err)
				return
			}
			fmt.Printf("txrwset.DataModel.String(): %v\n", txrwset.DataModel.String())
			for i, v := range txrwset.NsRwset {
				fmt.Printf("i: %v\n", i)
				fmt.Printf("v.Namespace: %v\n", v.Namespace)
				var kvset kvrwset.KVRWSet
				err = proto.Unmarshal(v.Rwset, &kvset)
				if err != nil {
					t.Error(err)
					return
				}
				fmt.Printf("kvset.Reads: %v\n", kvset.Reads)
				fmt.Printf("kvset.Writes: %v\n", kvset.Writes)
			}
		}

	}
	fmt.Print("-----success------\n")

}

func TestBasicGetAllAssets(t *testing.T) {
	result, err := tools.QueryChaincode("basic", "GetAllAssets")
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %s \n", result)
}

func TestBasicTransfer(t *testing.T) {
	result, _, err := tools.ExecuteChaincode("basic", "TransferAsset", "asset6", "Christopher")
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %s \n", result)
}
