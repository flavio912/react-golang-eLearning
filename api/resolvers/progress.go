package resolvers

import "gitlab.codesigned.co.uk/ttc-heathrow/ttc-project/admin-react/api/gentypes"

type ProgressResolver struct {
	Progress gentypes.Progress
}

func (p *ProgressResolver) Total() int32 {
	return int32(p.Progress.Total)
}
func (p *ProgressResolver) Completed() int32 {
	return int32(p.Progress.Completed)
}
func (p *ProgressResolver) Percent() float64 {
	if p.Progress.Total <= 0 {
		return 0
	}

	return float64(p.Progress.Completed / p.Progress.Total)
}
