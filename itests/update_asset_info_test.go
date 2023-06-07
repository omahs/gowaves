package itests

import (
	"testing"

	"github.com/stretchr/testify/suite"
	f "github.com/wavesplatform/gowaves/itests/fixtures"
	"github.com/wavesplatform/gowaves/itests/testdata"
	utl "github.com/wavesplatform/gowaves/itests/utilities"
	"github.com/wavesplatform/gowaves/itests/utilities/issue_utilities"
	"github.com/wavesplatform/gowaves/itests/utilities/update_asset_info_utilities"
)

type UpdateAssetInfoTxSuite struct {
	f.BaseSuite
}

func (suite *UpdateAssetInfoTxSuite) TestUpdateAssetInfoTxReissuableTokenPositive() {
	versions := update_asset_info_utilities.GetVersions(&suite.BaseSuite)
	issue_versions := issue_utilities.GetVersions(&suite.BaseSuite)
	waitForTx := true
	for _, v := range versions {
		for _, iv := range issue_versions {
			reissuable := testdata.GetCommonIssueData(&suite.BaseSuite).Reissuable
			itx := issue_utilities.IssueSendWithTestData(&suite.BaseSuite, reissuable, iv, waitForTx)
			tdmatrix := testdata.GetUpdateAssetInfoPositiveDataMatrix(&suite.BaseSuite, itx.TxID)
			for name, td := range tdmatrix {
				height := utl.GetHeight(&suite.BaseSuite)
				caseName := utl.GetTestcaseNameWithVersion(name, v)
				suite.Run(caseName, func() {
					//***wait n blocks***
					utl.WaitForHeight(&suite.BaseSuite, height+1)

					tx, actualDiffBalanceInWaves, actualDiffBalanceInAsset := update_asset_info_utilities.SendUpdateAssetInfoTxAndGetDiffBalances(
						&suite.BaseSuite, td, v, waitForTx)
					errMsg := caseName + "Updating Asset Info tx: " + tx.TxID.String()

					utl.TxInfoCheck(suite.T(), tx.WtErr.ErrWtGo, tx.WtErr.ErrWtScala, errMsg)
					utl.WavesDiffBalanceCheck(suite.T(), td.Expected.WavesDiffBalance, actualDiffBalanceInWaves.BalanceInWavesGo,
						actualDiffBalanceInWaves.BalanceInWavesScala, errMsg)
					utl.AssetDiffBalanceCheck(suite.T(), td.Expected.AssetDiffBalance, actualDiffBalanceInAsset.BalanceInAssetGo,
						actualDiffBalanceInAsset.BalanceInAssetScala, errMsg)

					assetDetailsGo, assetDetailsScala := utl.GetAssetInfoGrpc(&suite.BaseSuite, itx.TxID)
					utl.AssetNameCheck(suite.T(), td.AssetName, assetDetailsGo.GetName(), assetDetailsScala.GetName(), errMsg)
					utl.AssetDescriptionCheck(suite.T(), td.AssetDesc, assetDetailsGo.GetDescription(),
						assetDetailsScala.GetDescription(), errMsg)
				})
			}
		}
	}
}

func (suite *UpdateAssetInfoTxSuite) TestUpdateAssetInfoTxNFTPositive() {
	versions := update_asset_info_utilities.GetVersions(&suite.BaseSuite)
	issue_versions := issue_utilities.GetVersions(&suite.BaseSuite)
	waitForTx := true
	for _, v := range versions {
		for _, iv := range issue_versions {
			nft := testdata.GetCommonIssueData(&suite.BaseSuite).NFT
			itx := issue_utilities.IssueSendWithTestData(&suite.BaseSuite, nft, iv, waitForTx)
			tdmatrix := testdata.GetUpdateAssetInfoPositiveDataMatrix(&suite.BaseSuite, itx.TxID)
			for name, td := range tdmatrix {
				height := utl.GetHeight(&suite.BaseSuite)
				caseName := utl.GetTestcaseNameWithVersion(name, v)
				suite.Run(caseName, func() {
					//***wait n blocks***
					utl.WaitForHeight(&suite.BaseSuite, height+1)

					tx, actualDiffBalanceInWaves, actualDiffBalanceInAsset := update_asset_info_utilities.SendUpdateAssetInfoTxAndGetDiffBalances(
						&suite.BaseSuite, td, v, waitForTx)
					errMsg := caseName + "Updating Asset Info tx: " + tx.TxID.String()

					utl.TxInfoCheck(suite.T(), tx.WtErr.ErrWtGo, tx.WtErr.ErrWtScala, errMsg)
					utl.WavesDiffBalanceCheck(suite.T(), td.Expected.WavesDiffBalance, actualDiffBalanceInWaves.BalanceInWavesGo,
						actualDiffBalanceInWaves.BalanceInWavesScala, errMsg)
					utl.AssetDiffBalanceCheck(suite.T(), td.Expected.AssetDiffBalance, actualDiffBalanceInAsset.BalanceInAssetGo,
						actualDiffBalanceInAsset.BalanceInAssetScala, errMsg)

					assetDetailsGo, assetDetailsScala := utl.GetAssetInfoGrpc(&suite.BaseSuite, itx.TxID)
					utl.AssetNameCheck(suite.T(), td.AssetName, assetDetailsGo.GetName(), assetDetailsScala.GetName(), errMsg)
					utl.AssetDescriptionCheck(suite.T(), td.AssetDesc, assetDetailsGo.GetDescription(),
						assetDetailsScala.GetDescription(), errMsg)
				})
			}
		}
	}
}

func (suite *UpdateAssetInfoTxSuite) TestUpdateAssetInfoTxSmartAssetPositive() {
	versions := update_asset_info_utilities.GetVersions(&suite.BaseSuite)
	issue_versions := issue_utilities.GetVersions(&suite.BaseSuite)
	waitForTx := true
	for _, v := range versions {
		for _, iv := range issue_versions {
			smart := testdata.GetCommonIssueData(&suite.BaseSuite).Smart
			itx := issue_utilities.IssueSendWithTestData(&suite.BaseSuite, smart, iv, waitForTx)
			tdmatrix := testdata.GetUpdateSmartAssetInfoPositiveDataMatrix(&suite.BaseSuite, itx.TxID)
			for name, td := range tdmatrix {
				height := utl.GetHeight(&suite.BaseSuite)
				caseName := utl.GetTestcaseNameWithVersion(name, v)
				suite.Run(caseName, func() {
					//***wait n blocks***
					utl.WaitForHeight(&suite.BaseSuite, height+1)

					tx, actualDiffBalanceInWaves, actualDiffBalanceInAsset := update_asset_info_utilities.SendUpdateAssetInfoTxAndGetDiffBalances(
						&suite.BaseSuite, td, v, waitForTx)
					errMsg := caseName + "Updating Asset Info tx: " + tx.TxID.String()

					utl.TxInfoCheck(suite.T(), tx.WtErr.ErrWtGo, tx.WtErr.ErrWtScala, errMsg)
					utl.WavesDiffBalanceCheck(suite.T(), td.Expected.WavesDiffBalance, actualDiffBalanceInWaves.BalanceInWavesGo,
						actualDiffBalanceInWaves.BalanceInWavesScala, errMsg)
					utl.AssetDiffBalanceCheck(suite.T(), td.Expected.AssetDiffBalance, actualDiffBalanceInAsset.BalanceInAssetGo,
						actualDiffBalanceInAsset.BalanceInAssetScala, errMsg)

					assetDetailsGo, assetDetailsScala := utl.GetAssetInfoGrpc(&suite.BaseSuite, itx.TxID)
					utl.AssetNameCheck(suite.T(), td.AssetName, assetDetailsGo.GetName(), assetDetailsScala.GetName(), errMsg)
					utl.AssetDescriptionCheck(suite.T(), td.AssetDesc, assetDetailsGo.GetDescription(),
						assetDetailsScala.GetDescription(), errMsg)
				})
			}
		}
	}
}

func (suite *UpdateAssetInfoTxSuite) TestUpdateAssetInfoTxReissuableTokenNegative() {
	versions := update_asset_info_utilities.GetVersions(&suite.BaseSuite)
	issue_versions := issue_utilities.GetVersions(&suite.BaseSuite)
	waitForTx := true
	for _, v := range versions {
		for _, iv := range issue_versions {
			reissuable := testdata.GetCommonIssueData(&suite.BaseSuite).Reissuable
			itx := issue_utilities.IssueSendWithTestData(&suite.BaseSuite, reissuable, iv, waitForTx)
			tdmatrix := testdata.GetUpdateAssetInfoNegativeDataMatrix(&suite.BaseSuite, itx.TxID)
			for name, td := range tdmatrix {
				height := utl.GetHeight(&suite.BaseSuite)
				caseName := utl.GetTestcaseNameWithVersion(name, v)
				initAssetDetailsGo, initAssetDetailsScala := utl.GetAssetInfoGrpc(&suite.BaseSuite, itx.TxID)
				suite.Run(caseName, func() {
					//***wait n blocks***
					utl.WaitForHeight(&suite.BaseSuite, height+1)
					tx, actualDiffBalanceInWaves, actualDiffBalanceInAsset := update_asset_info_utilities.SendUpdateAssetInfoTxAndGetDiffBalances(
						&suite.BaseSuite, td, v, !waitForTx)
					errMsg := caseName + "Updating Asset Info tx: " + tx.TxID.String()

					utl.ErrorMessageCheck(suite.T(), td.Expected.ErrGoMsg, td.Expected.ErrScalaMsg, tx.WtErr.ErrWtGo,
						tx.WtErr.ErrWtScala, errMsg)
					utl.WavesDiffBalanceCheck(suite.T(), td.Expected.WavesDiffBalance, actualDiffBalanceInWaves.BalanceInWavesGo,
						actualDiffBalanceInWaves.BalanceInWavesScala, errMsg)
					utl.AssetDiffBalanceCheck(suite.T(), td.Expected.AssetDiffBalance, actualDiffBalanceInAsset.BalanceInAssetGo,
						actualDiffBalanceInAsset.BalanceInAssetScala, errMsg)

					assetDetailsGo, assetDetailsScala := utl.GetAssetInfoGrpc(&suite.BaseSuite, itx.TxID)
					utl.AssetNameCheck(suite.T(), initAssetDetailsGo.GetName(), assetDetailsGo.GetName(), assetDetailsScala.GetName(), errMsg)
					utl.AssetNameCheck(suite.T(), initAssetDetailsScala.GetName(), assetDetailsGo.GetName(), assetDetailsScala.GetName(), errMsg)
					utl.AssetDescriptionCheck(suite.T(), initAssetDetailsGo.GetDescription(), assetDetailsGo.GetDescription(),
						assetDetailsScala.GetDescription(), errMsg)
					utl.AssetDescriptionCheck(suite.T(), initAssetDetailsScala.GetDescription(), assetDetailsGo.GetDescription(),
						assetDetailsScala.GetDescription(), errMsg)
				})
			}
		}
	}
}

func (suite *UpdateAssetInfoTxSuite) TestUpdateAssetInfoTxNFTNegative() {
	versions := update_asset_info_utilities.GetVersions(&suite.BaseSuite)
	issue_versions := issue_utilities.GetVersions(&suite.BaseSuite)
	waitForTx := true
	for _, v := range versions {
		for _, iv := range issue_versions {
			nft := testdata.GetCommonIssueData(&suite.BaseSuite).NFT
			itx := issue_utilities.IssueSendWithTestData(&suite.BaseSuite, nft, iv, waitForTx)
			tdmatrix := testdata.GetUpdateAssetInfoNegativeDataMatrix(&suite.BaseSuite, itx.TxID)
			for name, td := range tdmatrix {
				height := utl.GetHeight(&suite.BaseSuite)
				caseName := utl.GetTestcaseNameWithVersion(name, v)
				initAssetDetailsGo, initAssetDetailsScala := utl.GetAssetInfoGrpc(&suite.BaseSuite, itx.TxID)
				suite.Run(caseName, func() {
					//***wait n blocks***
					utl.WaitForHeight(&suite.BaseSuite, height+1)
					tx, actualDiffBalanceInWaves, actualDiffBalanceInAsset := update_asset_info_utilities.SendUpdateAssetInfoTxAndGetDiffBalances(
						&suite.BaseSuite, td, v, !waitForTx)
					errMsg := caseName + "Updating Asset Info tx: " + tx.TxID.String()

					utl.ErrorMessageCheck(suite.T(), td.Expected.ErrGoMsg, td.Expected.ErrScalaMsg, tx.WtErr.ErrWtGo,
						tx.WtErr.ErrWtScala, errMsg)
					utl.WavesDiffBalanceCheck(suite.T(), td.Expected.WavesDiffBalance, actualDiffBalanceInWaves.BalanceInWavesGo,
						actualDiffBalanceInWaves.BalanceInWavesScala, errMsg)
					utl.AssetDiffBalanceCheck(suite.T(), td.Expected.AssetDiffBalance, actualDiffBalanceInAsset.BalanceInAssetGo,
						actualDiffBalanceInAsset.BalanceInAssetScala, errMsg)

					assetDetailsGo, assetDetailsScala := utl.GetAssetInfoGrpc(&suite.BaseSuite, itx.TxID)
					utl.AssetNameCheck(suite.T(), initAssetDetailsGo.GetName(), assetDetailsGo.GetName(), assetDetailsScala.GetName(), errMsg)
					utl.AssetNameCheck(suite.T(), initAssetDetailsScala.GetName(), assetDetailsGo.GetName(), assetDetailsScala.GetName(), errMsg)
					utl.AssetDescriptionCheck(suite.T(), initAssetDetailsGo.GetDescription(), assetDetailsGo.GetDescription(),
						assetDetailsScala.GetDescription(), errMsg)
					utl.AssetDescriptionCheck(suite.T(), initAssetDetailsScala.GetDescription(), assetDetailsGo.GetDescription(),
						assetDetailsScala.GetDescription(), errMsg)
				})
			}
		}
	}
}

func (suite *UpdateAssetInfoTxSuite) TestUpdateAssetInfoTxSmartAssetNegative() {
	versions := update_asset_info_utilities.GetVersions(&suite.BaseSuite)
	issue_versions := issue_utilities.GetVersions(&suite.BaseSuite)
	waitForTx := true
	for _, v := range versions {
		for _, iv := range issue_versions {
			smart := testdata.GetCommonIssueData(&suite.BaseSuite).Smart
			itx := issue_utilities.IssueSendWithTestData(&suite.BaseSuite, smart, iv, waitForTx)
			tdmatrix := testdata.GetUpdateSmartAssetInfoNegativeDataMatrix(&suite.BaseSuite, itx.TxID)
			for name, td := range tdmatrix {
				height := utl.GetHeight(&suite.BaseSuite)
				caseName := utl.GetTestcaseNameWithVersion(name, v)
				initAssetDetailsGo, initAssetDetailsScala := utl.GetAssetInfoGrpc(&suite.BaseSuite, itx.TxID)
				suite.Run(caseName, func() {
					//***wait n blocks***
					utl.WaitForHeight(&suite.BaseSuite, height+1)
					tx, actualDiffBalanceInWaves, actualDiffBalanceInAsset := update_asset_info_utilities.SendUpdateAssetInfoTxAndGetDiffBalances(
						&suite.BaseSuite, td, v, !waitForTx)
					errMsg := caseName + "Updating Asset Info tx: " + tx.TxID.String()

					utl.ErrorMessageCheck(suite.T(), td.Expected.ErrGoMsg, td.Expected.ErrScalaMsg, tx.WtErr.ErrWtGo,
						tx.WtErr.ErrWtScala, errMsg)
					utl.WavesDiffBalanceCheck(suite.T(), td.Expected.WavesDiffBalance, actualDiffBalanceInWaves.BalanceInWavesGo,
						actualDiffBalanceInWaves.BalanceInWavesScala, errMsg)
					utl.AssetDiffBalanceCheck(suite.T(), td.Expected.AssetDiffBalance, actualDiffBalanceInAsset.BalanceInAssetGo,
						actualDiffBalanceInAsset.BalanceInAssetScala, errMsg)

					assetDetailsGo, assetDetailsScala := utl.GetAssetInfoGrpc(&suite.BaseSuite, itx.TxID)
					utl.AssetNameCheck(suite.T(), initAssetDetailsGo.GetName(), assetDetailsGo.GetName(), assetDetailsScala.GetName(), errMsg)
					utl.AssetNameCheck(suite.T(), initAssetDetailsScala.GetName(), assetDetailsGo.GetName(), assetDetailsScala.GetName(), errMsg)
					utl.AssetDescriptionCheck(suite.T(), initAssetDetailsGo.GetDescription(), assetDetailsGo.GetDescription(),
						assetDetailsScala.GetDescription(), errMsg)
					utl.AssetDescriptionCheck(suite.T(), initAssetDetailsScala.GetDescription(), assetDetailsGo.GetDescription(),
						assetDetailsScala.GetDescription(), errMsg)
				})
			}
		}
	}
}

func (suite *UpdateAssetInfoTxSuite) TestUpdateAssetInfoTxWithoutWaitingNegative() {
	versions := update_asset_info_utilities.GetVersions(&suite.BaseSuite)
	issue_versions := issue_utilities.GetVersions(&suite.BaseSuite)
	waitForTx := true
	for _, v := range versions {
		for _, iv := range issue_versions {
			reissuable := testdata.GetCommonIssueData(&suite.BaseSuite).Reissuable
			itx := issue_utilities.IssueSendWithTestData(&suite.BaseSuite, reissuable, iv, waitForTx)
			name := "Updating Asset Info without waiting"
			tdstruct := testdata.GetUpdateAssetInfoWithoutWaitingNegativeData(&suite.BaseSuite, itx.TxID)
			for _, td := range tdstruct {
				caseName := utl.GetTestcaseNameWithVersion(name, v)
				initAssetDetailsGo, initAssetDetailsScala := utl.GetAssetInfoGrpc(&suite.BaseSuite, itx.TxID)
				suite.Run(caseName, func() {
					tx, actualDiffBalanceInWaves, actualDiffBalanceInAsset := update_asset_info_utilities.SendUpdateAssetInfoTxAndGetDiffBalances(
						&suite.BaseSuite, td, v, !waitForTx)
					errMsg := caseName + "Updating Asset Info tx: " + tx.TxID.String()

					utl.ErrorMessageCheck(suite.T(), td.Expected.ErrGoMsg, td.Expected.ErrScalaMsg, tx.WtErr.ErrWtGo,
						tx.WtErr.ErrWtScala, errMsg)
					utl.WavesDiffBalanceCheck(suite.T(), td.Expected.WavesDiffBalance, actualDiffBalanceInWaves.BalanceInWavesGo,
						actualDiffBalanceInWaves.BalanceInWavesScala, errMsg)
					utl.AssetDiffBalanceCheck(suite.T(), td.Expected.AssetDiffBalance, actualDiffBalanceInAsset.BalanceInAssetGo,
						actualDiffBalanceInAsset.BalanceInAssetScala, errMsg)

					assetDetailsGo, assetDetailsScala := utl.GetAssetInfoGrpc(&suite.BaseSuite, itx.TxID)
					utl.AssetNameCheck(suite.T(), initAssetDetailsGo.GetName(), assetDetailsGo.GetName(), assetDetailsScala.GetName(), errMsg)
					utl.AssetNameCheck(suite.T(), initAssetDetailsScala.GetName(), assetDetailsGo.GetName(), assetDetailsScala.GetName(), errMsg)
					utl.AssetDescriptionCheck(suite.T(), initAssetDetailsGo.GetDescription(), assetDetailsGo.GetDescription(),
						assetDetailsScala.GetDescription(), errMsg)
					utl.AssetDescriptionCheck(suite.T(), initAssetDetailsScala.GetDescription(), assetDetailsGo.GetDescription(),
						assetDetailsScala.GetDescription(), errMsg)
				})
			}
		}
	}
}

func TestUpdateAssetInfoTxSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UpdateAssetInfoTxSuite))
}
