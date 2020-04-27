// Auto-generated to Go types and interfaces using avdl-compiler v1.4.9 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/incoming-share.avdl

package keybase1

import (
	"fmt"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
	"time"
)

type IncomingShareType int

const (
	IncomingShareType_FILE  IncomingShareType = 0
	IncomingShareType_TEXT  IncomingShareType = 1
	IncomingShareType_IMAGE IncomingShareType = 2
	IncomingShareType_VIDEO IncomingShareType = 3
)

func (o IncomingShareType) DeepCopy() IncomingShareType { return o }

var IncomingShareTypeMap = map[string]IncomingShareType{
	"FILE":  0,
	"TEXT":  1,
	"IMAGE": 2,
	"VIDEO": 3,
}

var IncomingShareTypeRevMap = map[IncomingShareType]string{
	0: "FILE",
	1: "TEXT",
	2: "IMAGE",
	3: "VIDEO",
}

func (e IncomingShareType) String() string {
	if v, ok := IncomingShareTypeRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type IncomingShareItem struct {
	Type          IncomingShareType `codec:"type" json:"type"`
	OriginalPath  *string           `codec:"originalPath,omitempty" json:"originalPath,omitempty"`
	OriginalSize  *int              `codec:"originalSize,omitempty" json:"originalSize,omitempty"`
	ScaledPath    *string           `codec:"scaledPath,omitempty" json:"scaledPath,omitempty"`
	ScaledSize    *int              `codec:"scaledSize,omitempty" json:"scaledSize,omitempty"`
	ThumbnailPath *string           `codec:"thumbnailPath,omitempty" json:"thumbnailPath,omitempty"`
	Content       *string           `codec:"content,omitempty" json:"content,omitempty"`
}

func (o IncomingShareItem) DeepCopy() IncomingShareItem {
	return IncomingShareItem{
		Type: o.Type.DeepCopy(),
		OriginalPath: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.OriginalPath),
		OriginalSize: (func(x *int) *int {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.OriginalSize),
		ScaledPath: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.ScaledPath),
		ScaledSize: (func(x *int) *int {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.ScaledSize),
		ThumbnailPath: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.ThumbnailPath),
		Content: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.Content),
	}
}

type IncomingShareCompressPreference int

const (
	IncomingShareCompressPreference_ORIGINAL   IncomingShareCompressPreference = 0
	IncomingShareCompressPreference_COMPRESSED IncomingShareCompressPreference = 1
)

func (o IncomingShareCompressPreference) DeepCopy() IncomingShareCompressPreference { return o }

var IncomingShareCompressPreferenceMap = map[string]IncomingShareCompressPreference{
	"ORIGINAL":   0,
	"COMPRESSED": 1,
}

var IncomingShareCompressPreferenceRevMap = map[IncomingShareCompressPreference]string{
	0: "ORIGINAL",
	1: "COMPRESSED",
}

func (e IncomingShareCompressPreference) String() string {
	if v, ok := IncomingShareCompressPreferenceRevMap[e]; ok {
		return v
	}
	return fmt.Sprintf("%v", int(e))
}

type GetIncomingShareItemsArg struct {
}

type GetCompressPreferenceArg struct {
}

type SetCompressPreferenceArg struct {
	Preference IncomingShareCompressPreference `codec:"preference" json:"preference"`
}

type IncomingShareInterface interface {
	GetIncomingShareItems(context.Context) ([]IncomingShareItem, error)
	GetCompressPreference(context.Context) (IncomingShareCompressPreference, error)
	SetCompressPreference(context.Context, IncomingShareCompressPreference) error
}

func IncomingShareProtocol(i IncomingShareInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.incomingShare",
		Methods: map[string]rpc.ServeHandlerDescription{
			"getIncomingShareItems": {
				MakeArg: func() interface{} {
					var ret [1]GetIncomingShareItemsArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.GetIncomingShareItems(ctx)
					return
				},
			},
			"getCompressPreference": {
				MakeArg: func() interface{} {
					var ret [1]GetCompressPreferenceArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.GetCompressPreference(ctx)
					return
				},
			},
			"setCompressPreference": {
				MakeArg: func() interface{} {
					var ret [1]SetCompressPreferenceArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]SetCompressPreferenceArg)
					if !ok {
						err = rpc.NewTypeError((*[1]SetCompressPreferenceArg)(nil), args)
						return
					}
					err = i.SetCompressPreference(ctx, typedArgs[0].Preference)
					return
				},
			},
		},
	}
}

type IncomingShareClient struct {
	Cli rpc.GenericClient
}

func (c IncomingShareClient) GetIncomingShareItems(ctx context.Context) (res []IncomingShareItem, err error) {
	err = c.Cli.Call(ctx, "keybase.1.incomingShare.getIncomingShareItems", []interface{}{GetIncomingShareItemsArg{}}, &res, 0*time.Millisecond)
	return
}

func (c IncomingShareClient) GetCompressPreference(ctx context.Context) (res IncomingShareCompressPreference, err error) {
	err = c.Cli.Call(ctx, "keybase.1.incomingShare.getCompressPreference", []interface{}{GetCompressPreferenceArg{}}, &res, 0*time.Millisecond)
	return
}

func (c IncomingShareClient) SetCompressPreference(ctx context.Context, preference IncomingShareCompressPreference) (err error) {
	__arg := SetCompressPreferenceArg{Preference: preference}
	err = c.Cli.Call(ctx, "keybase.1.incomingShare.setCompressPreference", []interface{}{__arg}, nil, 0*time.Millisecond)
	return
}
