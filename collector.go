package main

import (
	"context"
	"github.com/andanhm/go-prettytime"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

type cmdCollector struct {
	cmdMetric *prometheus.Desc
}

func newCollector() *cmdCollector {
	return &cmdCollector{
		cmdMetric: prometheus.NewDesc("docker_swarm_service_list",
			"docker_swarm_service_list",
			[]string{"ID","NAME","MODE","REPLICAS","IMAGE","PORT","STARTED","UPDATED"}, nil,
		),
	}
}

func (collector *cmdCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.cmdMetric
}

func (collector *cmdCollector) Collect(ch chan<- prometheus.Metric) {
	var port string
	var metricValue float64
	metricValue = 1

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	services, err := cli.ServiceList(ctx, types.ServiceListOptions{})
	if err != nil {
		panic(err)
	}
	for _, service := range services {
		port =""
		for _, p:= range service.Endpoint.Ports {
			port += "["+strconv.Itoa(int(p.TargetPort))+"->"+strconv.Itoa(int(p.PublishedPort))+"/"+string(p.Protocol)+"]"

		}
		ch <- prometheus.MustNewConstMetric(collector.cmdMetric, prometheus.GaugeValue, metricValue, service.ID[:12],service.Spec.Name,"Replicated",strconv.Itoa(int(*service.Spec.Mode.Replicated.Replicas)),service.Spec.Labels["com.docker.stack.image"],port,prettytime.Format(service.Meta.CreatedAt),prettytime.Format(service.Meta.UpdatedAt))
	}

}

