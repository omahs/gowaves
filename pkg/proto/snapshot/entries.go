package snapshot

import (
	"bytes"
	"io"
	"sort"

	"github.com/wavesplatform/gowaves/pkg/crypto"
)

type HashEntryType byte

const (
	WavesBalance = HashEntryType(iota) // TODO: is it ok that keys enum starts with 0?
	AssetBalance
	DataEntry
	AccountScript
	AssetScript
	LeaseBalance
	LeaseStatus
	Sponsorship
	Alias
	VolumeAndFee
	StaticAssetInfo
	AssetReissuability
	AssetNameDescription
)

func (t HashEntryType) CreateSimpleEntry(key []byte, value []byte) SimpleHashEntry {
	return SimpleHashEntry{
		entryType: t,
		key:       key,
		value:     value,
	}
}

type SimpleHashEntry struct {
	entryType HashEntryType
	key       []byte
	value     []byte
}

func (e SimpleHashEntry) EntryType() HashEntryType {
	return e.entryType
}

func (e SimpleHashEntry) Key() []byte {
	return e.key
}

func (e SimpleHashEntry) WriteTo(w io.Writer) (n int64, err error) {
	// write entry type
	if written, err := w.Write([]byte{byte(e.entryType)}); err != nil {
		return n + int64(written), err
	} else {
		n += int64(written)
	}
	// write key
	if written, err := w.Write(e.key); err != nil {
		return n + int64(written), err
	} else {
		n += int64(written)
	}
	// write value
	if written, err := w.Write(e.value); err != nil {
		return n + int64(written), err
	} else {
		n += int64(written)
	}
	// return result
	return n, nil
}

type HashEntry interface {
	Key() []byte
	EntryType() HashEntryType
	io.WriterTo
}

type HashEntries []HashEntry

func (s HashEntries) Len() int {
	return len(s)
}

func (s HashEntries) Less(i, j int) bool {
	return s[i].EntryType() < s[j].EntryType() && (bytes.Compare(s[i].Key(), s[j].Key()) == -1)
}

func (s HashEntries) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s HashEntries) Sort() {
	sort.Sort(s) // TODO: which sort should be used: stable or not?
}

func (s HashEntries) IsSorted() bool {
	return sort.IsSorted(s)
}

func (s HashEntries) CalculateHash() (crypto.Digest, error) {
	if !s.IsSorted() { // TODO: I think IsSorted check here is unnecessary
		s.Sort()
	}
	h, err := crypto.NewFastHash()
	if err != nil {
		return crypto.Digest{}, err
	}
	for i := range s {
		if _, err := s[i].WriteTo(h); err != nil {
			return crypto.Digest{}, err
		}
	}
	var d crypto.Digest
	h.Sum(d[:0])
	return d, nil
}
