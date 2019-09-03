package api

import "github.com/gogo/protobuf/proto"

func (m *AuthReply) String() string {
	return proto.MarshalTextString(m)
}
func (m *TokenReq) String() string {
	return proto.MarshalTextString(m)
}
