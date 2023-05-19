package state

import (
	"github.com/pkg/errors"
	"github.com/wavesplatform/gowaves/pkg/proto"
)

type snapshotApplier struct {
	balances *balances
	aliases  *aliases
	assets   *assets
	//scriptsStorage    *scriptsStorage
	//scriptsComplexity *scriptsComplexity
}

var _ = (&snapshotApplier{}).applyWavesBalance // TODO: remove it, need for linter for now

func (a *snapshotApplier) applyWavesBalance(blockID proto.BlockID, snapshot WavesBalanceSnapshot) error {
	addrID := snapshot.Address.ID()
	profile, err := a.balances.wavesBalance(addrID)
	if err != nil {
		return errors.Wrapf(err, "failed to get waves balance profile for address %q", snapshot.Address.String())
	}
	newProfile := profile
	newProfile.balance = snapshot.Balance
	value := newWavesValue(profile, newProfile)
	if err := a.balances.setWavesBalance(addrID, value, blockID); err != nil {
		return errors.Wrapf(err, "failed to get set balance profile for address %q", snapshot.Address.String())
	}
	return nil
}

var _ = (&snapshotApplier{}).applyLeaseBalance // TODO: remove it, need for linter for now

func (a *snapshotApplier) applyLeaseBalance(blockID proto.BlockID, snapshot LeaseBalanceSnapshot) error {
	addrID := snapshot.Address.ID()
	profile, err := a.balances.wavesBalance(addrID)
	if err != nil {
		return errors.Wrapf(err, "failed to get waves balance profile for address %q", snapshot.Address.String())
	}
	newProfile := profile
	newProfile.leaseIn = int64(snapshot.LeaseIn)
	newProfile.leaseOut = int64(snapshot.LeaseOut)
	value := newWavesValue(profile, newProfile)
	if err := a.balances.setWavesBalance(addrID, value, blockID); err != nil {
		return errors.Wrapf(err, "failed to get set balance profile for address %q", snapshot.Address.String())
	}
	return nil
}

var _ = (&snapshotApplier{}).applyAssetBalance // TODO: remove it, need for linter for now

func (a *snapshotApplier) applyAssetBalance(blockID proto.BlockID, snapshot AssetBalanceSnapshot) error {
	addrID := snapshot.Address.ID()
	assetID := proto.AssetIDFromDigest(snapshot.AssetID)
	return a.balances.setAssetBalance(addrID, assetID, snapshot.Balance, blockID)
}

var _ = (&snapshotApplier{}).applyAlias // TODO: remove it, need for linter for now

func (a *snapshotApplier) applyAlias(blockID proto.BlockID, snapshot AliasSnapshot) error {
	return a.aliases.createAlias(snapshot.Alias.Alias, snapshot.Address, blockID)
}

var _ = (&snapshotApplier{}).applyStaticAssetInfo // TODO: remove it, need for linter for now

func (a *snapshotApplier) applyStaticAssetInfo(blockID proto.BlockID, snapshot StaticAssetInfoSnapshot) error {
	assetID := proto.AssetIDFromDigest(snapshot.AssetID)
	info := &assetInfo{
		assetConstInfo: assetConstInfo{
			tail:                 proto.DigestTail(snapshot.AssetID),
			issuer:               snapshot.IssuerPublicKey,
			decimals:             snapshot.Decimals,
			issueHeight:          0, // TODO: add info?
			issueSequenceInBlock: 0, // TODO: add info?
		},
		assetChangeableInfo: assetChangeableInfo{}, // TODO: add info?
	}
	return a.assets.issueAsset(assetID, info, blockID)
}

var _ = (&snapshotApplier{}).applyAssetDescription // TODO: remove it, need for linter for now

func (a *snapshotApplier) applyAssetDescription(blockID proto.BlockID, snapshot AssetDescriptionSnapshot) error {
	change := &assetInfoChange{
		newName:        snapshot.AssetName,
		newDescription: snapshot.AssetDescription,
		newHeight:      snapshot.ChangeHeight,
	}
	return a.assets.updateAssetInfo(snapshot.AssetID, change, blockID)
}

var _ = (&snapshotApplier{}).applyAssetVolume // TODO: remove it, need for linter for now

func (a *snapshotApplier) applyAssetVolume(blockID proto.BlockID, snapshot AssetVolumeSnapshot) error {
	assetID := proto.AssetIDFromDigest(snapshot.AssetID)
	info, err := a.assets.newestAssetInfo(assetID)
	if err != nil {
		return errors.Wrapf(err, "failed to get newest asset info for asset %q", snapshot.AssetID.String())
	}
	info.assetChangeableInfo.reissuable = snapshot.IsReissuable
	info.assetChangeableInfo.quantity = snapshot.TotalQuantity
	return a.assets.storeAssetInfo(assetID, info, blockID)
}

//func (a *snapshotApplier) applyAssetScript(blockID proto.BlockID, snapshot AssetScriptSnapshot) error {
//	var pk crypto.PublicKey // TODO: stub for now, can be omitted for assets. any change will cause state incompatibility for now
//	_ = snapshot.Complexity // this field is not necessary for us because we do parsing for each script
//	err := a.scriptsStorage.setAssetScript(snapshot.AssetID, snapshot.Script, pk, blockID)
//
//	scriptsComplexity.saveComplexitiesForAsset()
//}
