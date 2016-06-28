package main

import (
	"fmt"
	"runtime"
	"time"

	"k8s.io/kubernetes/pkg/util"

	"github.com/spf13/pflag"

	"github.com/dongyiyang/k8sconnection/pkg/conntrack"
	"github.com/dongyiyang/k8sconnection/pkg/flowcollector"
	"github.com/dongyiyang/k8sconnection/pkg/k8sconnector"
	"github.com/dongyiyang/k8sconnection/pkg/server"
	"github.com/dongyiyang/k8sconnection/pkg/transactioncounter"

	"github.com/golang/glog"
)

func main() {
	k8scon, err := createK8sConnector()
	if err != nil {
		glog.Fatalf("Cannot create require connection monitor: %s", err)
	}
	conntrack.SetK8sConnector(k8scon)

	transactionCounter := transactioncounter.NewTransactionCounter(k8scon)
	flowCollector := flowcollector.NewFlowCollector(k8scon)

	go server.ListenAndServeProxyServer(transactionCounter, flowCollector)

	c, err := conntrack.New()
	if err != nil {
		panic(err)
	}
	for range time.Tick(1 * time.Second) {

		glog.Infof("~~~~~~~~~~~~~~~~   Transaction Counter	~~~~~~~~~~~~~~~~~~~~")
		transactionCounter.ProcessConntrackConnections(c)

		fmt.Println()
		glog.Infof("----------------   Flow Collector	------------------------")

		fmt.Println()
		fmt.Println()

		flowCollector.TrackFlow()
	}
}

func createK8sConnector() (*k8sconnector.K8sConnector, error) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	s := k8sconnector.NewK8sConnectorBuilder().AddFlags(pflag.CommandLine)

	util.InitFlags()
	util.InitLogs()
	defer util.FlushLogs()

	// monitor, err := s.Build(pflag.CommandLine.Args())
	monitor, err := s.Build()
	if err != nil {
		return nil, err
	}

	return monitor, nil
}

func count(connCounterMap map[string]int64, src string) map[string]int64 {
	count, exist := connCounterMap[src]
	if !exist {
		count = 0
	}
	connCounterMap[src] = count + 1
	return connCounterMap
}
