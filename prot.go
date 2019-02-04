// Auto-generated by avdl-compiler v1.3.29 (https://github.com/keybase/node-avdl-compiler)
//   Input file: prot.avdl

package flip

import (
	"errors"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

type Time int64

func (o Time) DeepCopy() Time {
	return o
}

type GameID []byte

func (o GameID) DeepCopy() GameID {
	return (func(x []byte) []byte {
		if x == nil {
			return nil
		}
		return append([]byte{}, x...)
	})(o)
}

type UID []byte

func (o UID) DeepCopy() UID {
	return (func(x []byte) []byte {
		if x == nil {
			return nil
		}
		return append([]byte{}, x...)
	})(o)
}

type DeviceID []byte

func (o DeviceID) DeepCopy() DeviceID {
	return (func(x []byte) []byte {
		if x == nil {
			return nil
		}
		return append([]byte{}, x...)
	})(o)
}

type Start struct {
	RegistrationEndsAt   Time           `codec:"registrationEndsAt" json:"registrationEndsAt"`
	CommitmentPeriodMsec int64          `codec:"commitmentPeriodMsec" json:"commitmentPeriodMsec"`
	RevealPeriodMsec     int64          `codec:"revealPeriodMsec" json:"revealPeriodMsec"`
	Params               FlipParameters `codec:"params" json:"params"`
}

func (o Start) DeepCopy() Start {
	return Start{
		RegistrationEndsAt:   o.RegistrationEndsAt.DeepCopy(),
		CommitmentPeriodMsec: o.CommitmentPeriodMsec,
		RevealPeriodMsec:     o.RevealPeriodMsec,
		Params:               o.Params.DeepCopy(),
	}
}

type UserDevice struct {
	User   UID      `codec:"user" json:"user"`
	Device DeviceID `codec:"device" json:"device"`
}

func (o UserDevice) DeepCopy() UserDevice {
	return UserDevice{
		User:   o.User.DeepCopy(),
		Device: o.Device.DeepCopy(),
	}
}

type RegistrationComplete struct {
	Player []UserDevice `codec:"player" json:"player"`
}

func (o RegistrationComplete) DeepCopy() RegistrationComplete {
	return RegistrationComplete{
		Player: (func(x []UserDevice) []UserDevice {
			if x == nil {
				return nil
			}
			ret := make([]UserDevice, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Player),
	}
}

type Register struct {
}

func (o Register) DeepCopy() Register {
	return Register{}
}

type FlipType int

const (
	FlipType_INTS    FlipType = 1
	FlipType_SHUFFLE FlipType = 2
)

func (o FlipType) DeepCopy() FlipType { return o }

var FlipTypeMap = map[string]FlipType{
	"INTS":    1,
	"SHUFFLE": 2,
}

var FlipTypeRevMap = map[FlipType]string{
	1: "INTS",
	2: "SHUFFLE",
}

func (e FlipType) String() string {
	if v, ok := FlipTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type IntType int

const (
	IntType_FIXED IntType = 1
	IntType_BIG   IntType = 2
	IntType_BOOL  IntType = 3
)

func (o IntType) DeepCopy() IntType { return o }

var IntTypeMap = map[string]IntType{
	"FIXED": 1,
	"BIG":   2,
	"BOOL":  3,
}

var IntTypeRevMap = map[IntType]string{
	1: "FIXED",
	2: "BIG",
	3: "BOOL",
}

func (e IntType) String() string {
	if v, ok := IntTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type FlipParametersInt struct {
	T__     IntType `codec:"t" json:"t"`
	Big__   *[]byte `codec:"big,omitempty" json:"big,omitempty"`
	Fixed__ *int64  `codec:"fixed,omitempty" json:"fixed,omitempty"`
}

func (o *FlipParametersInt) T() (ret IntType, err error) {
	switch o.T__ {
	case IntType_BIG:
		if o.Big__ == nil {
			err = errors.New("unexpected nil value for Big__")
			return ret, err
		}
	case IntType_FIXED:
		if o.Fixed__ == nil {
			err = errors.New("unexpected nil value for Fixed__")
			return ret, err
		}
	}
	return o.T__, nil
}

func (o FlipParametersInt) Big() (res []byte) {
	if o.T__ != IntType_BIG {
		panic("wrong case accessed")
	}
	if o.Big__ == nil {
		return
	}
	return *o.Big__
}

func (o FlipParametersInt) Fixed() (res int64) {
	if o.T__ != IntType_FIXED {
		panic("wrong case accessed")
	}
	if o.Fixed__ == nil {
		return
	}
	return *o.Fixed__
}

func NewFlipParametersIntWithBig(v []byte) FlipParametersInt {
	return FlipParametersInt{
		T__:   IntType_BIG,
		Big__: &v,
	}
}

func NewFlipParametersIntWithFixed(v int64) FlipParametersInt {
	return FlipParametersInt{
		T__:     IntType_FIXED,
		Fixed__: &v,
	}
}

func NewFlipParametersIntWithBool() FlipParametersInt {
	return FlipParametersInt{
		T__: IntType_BOOL,
	}
}

func (o FlipParametersInt) DeepCopy() FlipParametersInt {
	return FlipParametersInt{
		T__: o.T__.DeepCopy(),
		Big__: (func(x *[]byte) *[]byte {
			if x == nil {
				return nil
			}
			tmp := (func(x []byte) []byte {
				if x == nil {
					return nil
				}
				return append([]byte{}, x...)
			})((*x))
			return &tmp
		})(o.Big__),
		Fixed__: (func(x *int64) *int64 {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.Fixed__),
	}
}

type FlipParameters struct {
	T__       FlipType             `codec:"t" json:"t"`
	Ints__    *[]FlipParametersInt `codec:"ints,omitempty" json:"ints,omitempty"`
	Shuffle__ *int64               `codec:"shuffle,omitempty" json:"shuffle,omitempty"`
}

func (o *FlipParameters) T() (ret FlipType, err error) {
	switch o.T__ {
	case FlipType_INTS:
		if o.Ints__ == nil {
			err = errors.New("unexpected nil value for Ints__")
			return ret, err
		}
	case FlipType_SHUFFLE:
		if o.Shuffle__ == nil {
			err = errors.New("unexpected nil value for Shuffle__")
			return ret, err
		}
	}
	return o.T__, nil
}

func (o FlipParameters) Ints() (res []FlipParametersInt) {
	if o.T__ != FlipType_INTS {
		panic("wrong case accessed")
	}
	if o.Ints__ == nil {
		return
	}
	return *o.Ints__
}

func (o FlipParameters) Shuffle() (res int64) {
	if o.T__ != FlipType_SHUFFLE {
		panic("wrong case accessed")
	}
	if o.Shuffle__ == nil {
		return
	}
	return *o.Shuffle__
}

func NewFlipParametersWithInts(v []FlipParametersInt) FlipParameters {
	return FlipParameters{
		T__:    FlipType_INTS,
		Ints__: &v,
	}
}

func NewFlipParametersWithShuffle(v int64) FlipParameters {
	return FlipParameters{
		T__:       FlipType_SHUFFLE,
		Shuffle__: &v,
	}
}

func (o FlipParameters) DeepCopy() FlipParameters {
	return FlipParameters{
		T__: o.T__.DeepCopy(),
		Ints__: (func(x *[]FlipParametersInt) *[]FlipParametersInt {
			if x == nil {
				return nil
			}
			tmp := (func(x []FlipParametersInt) []FlipParametersInt {
				if x == nil {
					return nil
				}
				ret := make([]FlipParametersInt, len(x))
				for i, v := range x {
					vCopy := v.DeepCopy()
					ret[i] = vCopy
				}
				return ret
			})((*x))
			return &tmp
		})(o.Ints__),
		Shuffle__: (func(x *int64) *int64 {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.Shuffle__),
	}
}

type Stage int

const (
	Stage_START                 Stage = 1
	Stage_REGISTER              Stage = 2
	Stage_REGISTRATION_COMPLETE Stage = 3
	Stage_COMMITMENT            Stage = 4
	Stage_REVEAL                Stage = 5
)

func (o Stage) DeepCopy() Stage { return o }

var StageMap = map[string]Stage{
	"START":                 1,
	"REGISTER":              2,
	"REGISTRATION_COMPLETE": 3,
	"COMMITMENT":            4,
	"REVEAL":                5,
}

var StageRevMap = map[Stage]string{
	1: "START",
	2: "REGISTER",
	3: "REGISTRATION_COMPLETE",
	4: "COMMITMENT",
	5: "REVEAL",
}

func (e Stage) String() string {
	if v, ok := StageRevMap[e]; ok {
		return v
	}
	return ""
}

type Secret [32]byte

func (o Secret) DeepCopy() Secret {
	var ret Secret
	copy(ret[:], o[:])
	return ret
}

type GameMessageBody struct {
	S__                    Stage                 `codec:"s" json:"s"`
	Start__                *Start                `codec:"start,omitempty" json:"start,omitempty"`
	RegistrationComplete__ *RegistrationComplete `codec:"registrationComplete,omitempty" json:"registrationComplete,omitempty"`
	Commitment__           *Secret               `codec:"commitment,omitempty" json:"commitment,omitempty"`
	Reveal__               *Secret               `codec:"reveal,omitempty" json:"reveal,omitempty"`
}

func (o *GameMessageBody) S() (ret Stage, err error) {
	switch o.S__ {
	case Stage_START:
		if o.Start__ == nil {
			err = errors.New("unexpected nil value for Start__")
			return ret, err
		}
	case Stage_REGISTRATION_COMPLETE:
		if o.RegistrationComplete__ == nil {
			err = errors.New("unexpected nil value for RegistrationComplete__")
			return ret, err
		}
	case Stage_COMMITMENT:
		if o.Commitment__ == nil {
			err = errors.New("unexpected nil value for Commitment__")
			return ret, err
		}
	case Stage_REVEAL:
		if o.Reveal__ == nil {
			err = errors.New("unexpected nil value for Reveal__")
			return ret, err
		}
	}
	return o.S__, nil
}

func (o GameMessageBody) Start() (res Start) {
	if o.S__ != Stage_START {
		panic("wrong case accessed")
	}
	if o.Start__ == nil {
		return
	}
	return *o.Start__
}

func (o GameMessageBody) RegistrationComplete() (res RegistrationComplete) {
	if o.S__ != Stage_REGISTRATION_COMPLETE {
		panic("wrong case accessed")
	}
	if o.RegistrationComplete__ == nil {
		return
	}
	return *o.RegistrationComplete__
}

func (o GameMessageBody) Commitment() (res Secret) {
	if o.S__ != Stage_COMMITMENT {
		panic("wrong case accessed")
	}
	if o.Commitment__ == nil {
		return
	}
	return *o.Commitment__
}

func (o GameMessageBody) Reveal() (res Secret) {
	if o.S__ != Stage_REVEAL {
		panic("wrong case accessed")
	}
	if o.Reveal__ == nil {
		return
	}
	return *o.Reveal__
}

func NewGameMessageBodyWithStart(v Start) GameMessageBody {
	return GameMessageBody{
		S__:     Stage_START,
		Start__: &v,
	}
}

func NewGameMessageBodyWithRegistrationComplete(v RegistrationComplete) GameMessageBody {
	return GameMessageBody{
		S__: Stage_REGISTRATION_COMPLETE,
		RegistrationComplete__: &v,
	}
}

func NewGameMessageBodyWithCommitment(v Secret) GameMessageBody {
	return GameMessageBody{
		S__:          Stage_COMMITMENT,
		Commitment__: &v,
	}
}

func NewGameMessageBodyWithReveal(v Secret) GameMessageBody {
	return GameMessageBody{
		S__:      Stage_REVEAL,
		Reveal__: &v,
	}
}

func (o GameMessageBody) DeepCopy() GameMessageBody {
	return GameMessageBody{
		S__: o.S__.DeepCopy(),
		Start__: (func(x *Start) *Start {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Start__),
		RegistrationComplete__: (func(x *RegistrationComplete) *RegistrationComplete {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.RegistrationComplete__),
		Commitment__: (func(x *Secret) *Secret {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Commitment__),
		Reveal__: (func(x *Secret) *Secret {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Reveal__),
	}
}

type GameMessage struct {
	GameID GameID          `codec:"gameID" json:"gameID"`
	Body   GameMessageBody `codec:"body" json:"body"`
}

func (o GameMessage) DeepCopy() GameMessage {
	return GameMessage{
		GameID: o.GameID.DeepCopy(),
		Body:   o.Body.DeepCopy(),
	}
}

type FlipInterface interface {
}

func FlipProtocol(i FlipInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "flip.flip",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type FlipClient struct {
	Cli rpc.GenericClient
}
