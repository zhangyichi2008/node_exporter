package collector

import (
	"os/exec"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type rsyslogProcessCollector struct {
	processName string
	logger      log.Logger
}

func init() {
	registerCollector("rsyslog", defaultEnabled, NewRsyslogProcessCollector)
}

func NewRsyslogProcessCollector(logger log.Logger) (Collector, error) {
	return &rsyslogProcessCollector{
		processName: "rsyslog",
		logger:      logger,
	}, nil
}

func (c *rsyslogProcessCollector) Update(ch chan<- prometheus.Metric) error {
	upDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, c.processName, "up"),
		"Value is 1 if rsyslog process is 'up', 0 otherwise.",
		[]string{"process_name"},
		nil,
	)
	upValue := 0.0
	alive := checkSyslogProcess(c.processName)
	if alive {
		upValue = 1.0
	}

	ch <- prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, upValue, c.processName)
	return nil
}

// checkSyslogProcess checks if a process with the given name is alive
// returns true if alive, false otherwise
func checkSyslogProcess(name string) bool {
	cmd := exec.Command("pgrep", name)
	err := cmd.Run()
	return err == nil
}
