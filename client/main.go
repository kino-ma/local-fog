package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"local-fog/core"
	"local-fog/core/types"
	"local-fog/core/utils"
	"log"
	"os"
	"strconv"
	"time"
)

const cloudHostName string = "cloud"
const testDuration = 10 * time.Minute
const testInterval = time.Second / 24

const imageFileName = "data/sample.jpeg"

var imageBin []byte

func init() {
	bin, err := ioutil.ReadFile(imageFileName)
	if err != nil {
		err = fmt.Errorf("failed to read image file: %w", err)
		log.Fatal(err)
	}
	imageBin = bin
}

type result struct {
	host            string
	startTime       time.Time
	requestDuration time.Duration
	overallDuration time.Duration
	success         bool
}

func main() {
	results := make([]result, 0, testDuration/testInterval)

	timeout := time.After(testDuration)
	ticker := time.NewTicker(testInterval)
loop:
	for {
		select {
		case <-ticker.C:
			go func() {
				s := time.Now()
				nodes, err := core.Discover(1)
				if err != nil {
					log.Printf("failed to discover: %v", err)
					return
				}
				host := chooseHost(nodes, 0)

				consumer, err := core.Connect(host, core.DEFAULT_PORT)

				if err != nil {
					log.Fatalf("failed to connec to the server: %v", err)
				}

				_, eReq, err := call(&consumer, &types.CallRequest{
					AppId: 2,
					Body:  imageBin,
				})

				eAll := time.Since(s)
				log.Printf("overall: %s", eAll)

				r := result{host, s, eReq, eAll, err == nil}
				results = append(results, r)
			}()
		case <-timeout:
			break loop
		}
	}

	f, err := os.Create("/log/log.csv")
	if err != nil {
		log.Fatalf("failed to open log.csv: %v", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	headers := (&result{}).GetHeaders()
	if err := w.Write(headers); err != nil {
		log.Fatalf("failed to write csv header: %v", err)
	}

	for _, r := range results {
		values := r.GetValues()
		if err := w.Write(values); err != nil {
			log.Fatalf("failed to write csv row: %v", err)
		}
	}
}

func chooseHost(ns []*types.NodeInfoWrapper, i int) string {
	if len(ns) < i+1 {
		return cloudHostName
	}

	node := ns[i]
	addr := utils.Uint32ToIp((node.AddrV4))
	log.Printf("discovered: %+v", addr)
	return addr.String()
}

func (r *result) GetHeaders() []string {
	return []string{"host", "startTime", "requestDuration", "overallDuration", "succes"}
}

func (r *result) GetValues() []string {
	h := r.host
	s := r.startTime.Format(time.RFC3339)
	rd := strconv.FormatInt(r.requestDuration.Microseconds(), 10)
	od := strconv.FormatInt(r.overallDuration.Microseconds(), 10)
	sc := strconv.FormatBool(r.success)
	return []string{h, s, rd, od, sc}
}
