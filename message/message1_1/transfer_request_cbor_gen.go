// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package message1_1

import (
	"fmt"
	"io"

	datatransfer "github.com/filecoin-project/go-data-transfer"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf

func (t *transferRequest1_1) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{170}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.BCid (cid.Cid) (struct)
	if len("BCid") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"BCid\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("BCid"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("BCid")); err != nil {
		return err
	}

	if t.BCid == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCidBuf(scratch, w, *t.BCid); err != nil {
			return xerrors.Errorf("failed to write cid field t.BCid: %w", err)
		}
	}

	// t.Type (uint64) (uint64)
	if len("Type") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Type\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Type"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Type")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.Type)); err != nil {
		return err
	}

	// t.Paus (bool) (bool)
	if len("Paus") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Paus\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Paus"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Paus")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.Paus); err != nil {
		return err
	}

	// t.Part (bool) (bool)
	if len("Part") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Part\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Part"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Part")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.Part); err != nil {
		return err
	}

	// t.Pull (bool) (bool)
	if len("Pull") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Pull\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Pull"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Pull")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.Pull); err != nil {
		return err
	}

	// t.Stor (typegen.Deferred) (struct)
	if len("Stor") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Stor\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Stor"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Stor")); err != nil {
		return err
	}

	if err := t.Stor.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Vouch (typegen.Deferred) (struct)
	if len("Vouch") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Vouch\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Vouch"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Vouch")); err != nil {
		return err
	}

	if err := t.Vouch.MarshalCBOR(w); err != nil {
		return err
	}

	// t.VTyp (datatransfer.TypeIdentifier) (string)
	if len("VTyp") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"VTyp\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("VTyp"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("VTyp")); err != nil {
		return err
	}

	if len(t.VTyp) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.VTyp was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.VTyp))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.VTyp)); err != nil {
		return err
	}

	// t.XferID (uint64) (uint64)
	if len("XferID") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"XferID\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("XferID"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("XferID")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.XferID)); err != nil {
		return err
	}

	// t.RestartChannel (datatransfer.ChannelID) (struct)
	if len("RestartChannel") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"RestartChannel\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("RestartChannel"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("RestartChannel")); err != nil {
		return err
	}

	if err := t.RestartChannel.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *transferRequest1_1) UnmarshalCBOR(r io.Reader) error {
	*t = transferRequest1_1{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("transferRequest1_1: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.BCid (cid.Cid) (struct)
		case "BCid":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(br)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.BCid: %w", err)
					}

					t.BCid = &c
				}

			}
			// t.Type (uint64) (uint64)
		case "Type":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.Type = uint64(extra)

			}
			// t.Paus (bool) (bool)
		case "Paus":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.Paus = false
			case 21:
				t.Paus = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}
			// t.Part (bool) (bool)
		case "Part":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.Part = false
			case 21:
				t.Part = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}
			// t.Pull (bool) (bool)
		case "Pull":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.Pull = false
			case 21:
				t.Pull = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}
			// t.Stor (typegen.Deferred) (struct)
		case "Stor":

			{

				t.Stor = new(cbg.Deferred)

				if err := t.Stor.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("failed to read deferred field: %w", err)
				}
			}
			// t.Vouch (typegen.Deferred) (struct)
		case "Vouch":

			{

				t.Vouch = new(cbg.Deferred)

				if err := t.Vouch.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("failed to read deferred field: %w", err)
				}
			}
			// t.VTyp (datatransfer.TypeIdentifier) (string)
		case "VTyp":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.VTyp = datatransfer.TypeIdentifier(sval)
			}
			// t.XferID (uint64) (uint64)
		case "XferID":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.XferID = uint64(extra)

			}
			// t.RestartChannel (datatransfer.ChannelID) (struct)
		case "RestartChannel":

			{

				if err := t.RestartChannel.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.RestartChannel: %w", err)
				}

			}

		default:
			return fmt.Errorf("unknown struct field %d: '%s'", i, name)
		}
	}

	return nil
}
