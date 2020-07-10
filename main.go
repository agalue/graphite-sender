// A simplified version of https://github.com/marpaia/graphite-golang
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

// Metric is a struct that defines the relevant properties of a graphite metric
type Metric struct {
	Name      string
	Value     string
	Timestamp int64
}

func (metric Metric) String() string {
	return fmt.Sprintf(
		"%s %s %s",
		metric.Name,
		metric.Value,
		time.Unix(metric.Timestamp, 0).Format("2006-01-02 15:04:05"),
	)
}

// Graphite is a struct that defines the relevant properties of a graphite UDP connection
type Graphite struct {
	Target  string
	Prefix  string
	Timeout time.Duration
	conn    net.Conn
}

// defaultTimeout is the default number of seconds that we're willing to wait before forcing the connection establishment to fail
const defaultTimeout = 5

// Given a Graphite struct, Connect populates the Graphite.conn field with an appropriate TCP connection
func (graphite *Graphite) Connect() error {
	if graphite.Timeout == 0 {
		graphite.Timeout = defaultTimeout * time.Second
	}
	if udpAddr, err := net.ResolveUDPAddr("udp4", graphite.Target); err != nil {
		return err
	} else {
		if conn, err := net.DialUDP("udp4", nil, udpAddr); err != nil {
			return err
		} else {
			graphite.conn = conn
		}
	}
	return nil
}

// Given a Graphite struct, Disconnect closes the Graphite.conn field
func (graphite *Graphite) Disconnect() error {
	err := graphite.conn.Close()
	graphite.conn = nil
	return err
}

// Given a Metric struct, the SendMetric method sends the supplied metric to the Graphite connection that the method is called upon
func (graphite *Graphite) SendMetric(metric Metric) error {
	if metric.Timestamp == 0 {
		metric.Timestamp = time.Now().Unix()
	}
	if graphite.Prefix != "" {
		metric.Name = fmt.Sprintf("%s.%s", graphite.Prefix, metric.Name)
	}
	log.Printf("Sending metric: %s", metric.String())
	_, err := graphite.conn.Write([]byte(fmt.Sprintf("%s %s %d\n", metric.Name, metric.Value, metric.Timestamp)))
	return err
}

func main() {
	graphite := Graphite{}
	flag.StringVar(&graphite.Target, "target", "localhost:2003", "Graphite Target ServerN")
	flag.StringVar(&graphite.Prefix, "prefix", "", "Graphite Prefix")
	frequency := flag.Duration("frequency", 5*time.Second, "Frequency of packet generation")
	flag.Parse()
	if err := graphite.Connect(); err != nil {
		panic(err)
	} else {
		defer graphite.Disconnect()
	}
	rand.Seed(time.Now().Unix())
	log.Printf("Sending Graphite metrics via UDP to %s every %s", graphite.Target, (*frequency).String())
	for {
		metric := Metric{Name: "sample", Value: fmt.Sprintf("%d", rand.Intn(1000))}
		if err := graphite.SendMetric(metric); err != nil {
			log.Printf("ERROR: %v", err)
		}
		time.Sleep(*frequency)
	}
}
