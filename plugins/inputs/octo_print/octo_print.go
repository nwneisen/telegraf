package octoprint

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type OctoPrint struct {
	source string `toml:"source"`
}

func (op *OctoPrint) Description() string {
	return "a demo plugin"
}

func (op *OctoPrint) SampleConfig() string {
	return `
  ## Indicate if everything is fine
  ok = true
`
}

func (op *OctoPrint) Init() error {
	return nil
}

func (op *OctoPrint) Gather(acc telegraf.Accumulator) error {
	if s.Ok {
		acc.AddFields("state", map[string]interface{}{"value": "pretty good"}, nil)
	} else {
		acc.AddFields("state", map[string]interface{}{"value": "not great"}, nil)
	}

	return nil
}

func init() {
	inputs.Add("octo-print", func() telegraf.Input { return &Simple{} })
}
