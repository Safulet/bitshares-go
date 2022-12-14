package bitshares

import (
	"github.com/Safulet/bitshares-go/types"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	"time"
)

const (
	testNet = "wss://node.testnet.bitshares.eu"
	mainNet = "wss://bitshares.openledger.info/ws"
)

func TestClient(t *testing.T) {
	t.Run("valid ws url", func(t *testing.T) {
		_, err := NewClient(mainNet)
		require.NoError(t, err)
	})

	t.Run("invalid ws url", func(t *testing.T) {
		_, err := NewClient("wss://invalid")
		require.Error(t, err)
	})
}

func TestClient_Transfer(t *testing.T) {
	client, err := NewClient(testNet)
	require.Nil(t, err)

	cali4888arr, err := client.Database.LookupAccounts("cali4889", 2)
	require.Nil(t, err)

	log.Println(cali4888arr["cali4889"])

	cali4889ID := cali4888arr["cali4889"]
	cali4890ID := cali4888arr["cali4890"]

	assets, err := client.Database.LookupAssetSymbols("TEST")
	require.Nil(t, err)

	cali4889IDActiveKey := "5JiTY3m9u1iPfoKsZdn18pnf26XvX2WnXFJckSiSaiUniNVzxLn"
	from := cali4889ID
	to := cali4890ID
	amount := types.AssetAmount{
		AssetID: assets[0].ID,
		Amount:  1000,
	}
	fee := types.AssetAmount{
		AssetID: assets[0].ID,
		Amount:  0,
	}

	require.NoError(t, client.Transfer(cali4889IDActiveKey, from, to, amount, fee))
}

func TestClient_LimitOrderCreate(t *testing.T) {
	client, err := NewClient(testNet)
	require.Nil(t, err)

	cali4889arr, err := client.Database.LookupAccounts("cali4889", 1)
	require.Nil(t, err)
	cali4889ID := cali4889arr["cali4889"]

	sellAsset, err := client.Database.LookupAssetSymbols("TEST")
	require.NoError(t, err)

	buyAsset, err := client.Database.LookupAssetSymbols("PEG.FAKEUSD")
	require.NoError(t, err)

	amSell := types.AssetAmount{
		Amount:  100,
		AssetID: sellAsset[0].ID,
	}
	minBuy := types.AssetAmount{
		Amount:  10,
		AssetID: buyAsset[0].ID,
	}

	fee := types.AssetAmount{
		AssetID: sellAsset[0].ID,
	}

	expiration := 40 * time.Hour

	cali4889IDActiveKey := "5JiTY3m9u1iPfoKsZdn18pnf26XvX2WnXFJckSiSaiUniNVzxLn"

	id, err := client.LimitOrderCreate(cali4889IDActiveKey, cali4889ID, fee, amSell, minBuy, expiration, false)
	require.NoError(t, err)

	orderID := types.MustParseObjectID(id)

	err = client.LimitOrderCancel(cali4889IDActiveKey, cali4889ID, orderID, fee)
	require.NoError(t, err)
}
