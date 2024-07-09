package client

import (
	"errors"

	"github.com/muulinCorp/interlib/scada-layout/pb"
)

type ScenarioReader interface {
	GetScenario(scenarioId string) (*ScenarioInfo, error)
	GetElementScenarios(elementId string) ([]*ScenarioOption, error)
}

func newScenarioReader(resp *pb.GetScenarioResponse) ScenarioReader {
	reader := &scenarioReaderImpl{
		elements:  make(map[string][]*pb.GetScenarioResponse_ScenarioDetail),
		scenarios: make(map[string]*pb.GetScenarioResponse_ScenarioDetail),
	}
	for _, s := range resp.Elements {
		scenarioMap := make([]*pb.GetScenarioResponse_ScenarioDetail, len(s.Scenarios))
		for i, d := range s.Scenarios {
			reader.scenarios[d.ScenarioId] = d
			scenarioMap[i] = d
		}
		reader.elements[s.ElementId] = scenarioMap
	}
	return reader
}

type scenarioReaderImpl struct {
	elements  map[string][]*pb.GetScenarioResponse_ScenarioDetail
	scenarios map[string]*pb.GetScenarioResponse_ScenarioDetail
}

type ScenarioOption struct {
	Key   string
	Value string
}

func (s *scenarioReaderImpl) GetElementScenarios(elementId string) ([]*ScenarioOption, error) {
	elements, ok := s.elements[elementId]
	if !ok {
		return nil, errors.New("element not found: " + elementId)
	}
	scenarios := make([]*ScenarioOption, len(elements))
	for i, e := range elements {
		scenarios[i] = &ScenarioOption{
			Key:   e.ScenarioId,
			Value: e.Desc,
		}
	}
	return scenarios, nil
}

type ScenarioInfo struct {
	Desc         string
	DelayMinutes int32
	Actions      []*fieldAction
}

type fieldAction struct {
	Field string
	Value float64
}

func (s *scenarioReaderImpl) GetScenario(scenarioId string) (*ScenarioInfo, error) {
	scenario, ok := s.scenarios[scenarioId]
	if !ok {
		return nil, errors.New("scenario not found: " + scenarioId)
	}
	actions := make([]*fieldAction, len(scenario.FieldActions))
	for i, a := range scenario.FieldActions {
		actions[i] = &fieldAction{
			Field: a.Field,
			Value: a.Value,
		}
	}
	return &ScenarioInfo{
		Desc:         scenario.Desc,
		DelayMinutes: scenario.DelayMinutes,
		Actions:      actions,
	}, nil
}
