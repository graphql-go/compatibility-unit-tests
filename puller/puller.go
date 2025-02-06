package puller

type Puller struct {
}

type PullerResult struct {
}

type PullerParams struct {
}

func (p *Puller) Pull(p *PullerParams) *PullerResult {
  return &PullerResult{}
}
