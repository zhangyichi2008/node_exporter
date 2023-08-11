package collector

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/go-cmd/cmd"

	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

type filebeatProcessCollector struct {
	processName string
	logger      log.Logger
}

func init() {
	registerCollector("filebeat", defaultEnabled, NewFilebeatProcessCollector)
}

func NewFilebeatProcessCollector(logger log.Logger) (Collector, error) {
	return &filebeatProcessCollector{
		processName: "filebeat",
		logger:      logger,
	}, nil
}

func (c *filebeatProcessCollector) Update(ch chan<- prometheus.Metric) error {
	upDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, c.processName, "up"),
		"Value is 1 if filebeat process is 'up', 0 otherwise.",
		[]string{"process_name"},
		nil,
	)
	upValue := 0.0
	alive := checkProcess(c.processName)
	if alive {
		//fmt.Printf("%s is 1\n", c.processName)
		upValue = 1.0
	}

	ch <- prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, upValue, c.processName)
	openfilesDesc := prometheus.NewDesc(
		prometheus.BuildFQName(namespace, c.processName, "openfiles"),
		"Filebeat monitoring log harvester openfiles running.",
		nil,
		nil,
	)
	openfilesValue := checkFilebeatStatus()
	ch <- prometheus.MustNewConstMetric(openfilesDesc, prometheus.GaugeValue, openfilesValue)

	return nil
}

// checkProcess checks if a process with the given name is alive
// returns true if alive, false otherwise
func checkProcess(name string) bool {
	cmd := exec.Command("pgrep", name)
	err := cmd.Run()
	return err == nil
}

//check filebeat openfiles returns type float64
func checkFilebeatStatus() float64 {
	var openFiles float64 = 0.0
	command := `journalctl -u filebeat  -n 100 --no-tail |grep monitoring |tail -1 |awk -F'harvester' '{print $2}' |awk -F'running' '{print $2}' |awk -F':' '{print $2}' |awk -F',' '{print $1}' |awk -F'}' '{print $1}' |grep -v '^$'`
	c := cmd.NewCmd("bash", "-c", command)
	statusChan := c.Start()
	finalStatus := <-statusChan
	if finalStatus.Error != nil {
		fmt.Println("Filebeat_error:", finalStatus.Error)
		return openFiles
	}
	cmdOut := finalStatus.Stdout
	cmdOutErr := finalStatus.Stderr
	if len(cmdOut) == 0 {
		if len(cmdOutErr) != 0 {
			fmt.Println("Filebeat_cmdOutErr:", cmdOutErr[0])
		} else {
			fmt.Println("Filebeat: get monitoring log failed, please check the process.")
		}
	} else {
		outFloat, err := strconv.ParseFloat(cmdOut[0], 64)
		if err != nil {
			fmt.Println("Filebeat_err:", err)
		}
		openFiles = outFloat
	}
	return openFiles
}
