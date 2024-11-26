package client

import "github.com/muulinCorp/interlib/scada-layout/pb"

type SmartDefrostLayout interface {
	GetGroups() []*pb.GetSmartDefrostResponse_GroupDetail
}

func newSmartDefrost(resp *pb.GetSmartDefrostResponse) SmartDefrostLayout {
	m := make(map[string]*pb.GetSmartDefrostResponse_GroupDetail)
	for _, g := range resp.Groups {
		m[g.Id] = g
	}
	return &smartDefrost{
		groups:   resp.Groups,
		groupMap: m,
	}
}

type smartDefrost struct {
	groups   []*pb.GetSmartDefrostResponse_GroupDetail
	groupMap map[string]*pb.GetSmartDefrostResponse_GroupDetail
}

func (s *smartDefrost) GetGroups() []*pb.GetSmartDefrostResponse_GroupDetail {
	return s.groups
}
