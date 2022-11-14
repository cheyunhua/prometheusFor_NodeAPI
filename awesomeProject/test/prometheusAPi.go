package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/prometheus/common/config"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

func ExampleAPI_query() {
	client, err := api.NewClient(api.Config{
		Address: "http://demo.robustperception.io:9090",
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, warnings, err := v1api.Query(ctx, "up", time.Now(), v1.WithTimeout(5*time.Second))
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	fmt.Printf("Result:\n%v\n", result)
}

func ExampleAPI_queryRange() {
	client, err := api.NewClient(api.Config{
		Address: "http://192.168.100.103:9090",
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r := v1.Range{
		Start: time.Now().Add(-1),
		End:   time.Now(),
		Step:  time.Minute,
	}
	result, warnings, err := v1api.QueryRange(ctx, "node_memory_MemAvailable_bytes / node_memory_MemTotal_bytes * 100", r, v1.WithTimeout(5*time.Second))
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}

	fmt.Printf("Result:\n%v\n", result)

}

type userAgentRoundTripper struct {
	name string
	rt   http.RoundTripper
}

// RoundTrip implements the http.RoundTripper interface.
func (u userAgentRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	if r.UserAgent() == "" {
		// The specification of http.RoundTripper says that it shouldn't mutate
		// the request so make a copy of req.Header since this is all that is
		// modified.
		r2 := new(http.Request)
		*r2 = *r
		r2.Header = make(http.Header)
		for k, s := range r.Header {
			r2.Header[k] = s
		}
		r2.Header.Set("User-Agent", u.name)
		r = r2
	}
	return u.rt.RoundTrip(r)
}

func ExampleAPI_queryRangeWithUserAgent() {
	client, err := api.NewClient(api.Config{
		Address:      "http://demo.robustperception.io:9090",
		RoundTripper: userAgentRoundTripper{name: "Client-Golang", rt: api.DefaultRoundTripper},
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r := v1.Range{
		Start: time.Now().Add(-time.Hour),
		End:   time.Now(),
		Step:  time.Minute,
	}
	result, warnings, err := v1api.QueryRange(ctx, "rate(prometheus_tsdb_head_samples_appended_total[5m])", r)
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	fmt.Printf("Result:\n%v\n", result)
}
func StringToMap(data string, spilt string) map[string]string {
	list := strings.Split(data, spilt)
	return ListToMap(list)
}

func ListToMap(data []string) map[string]string {
	m := make(map[string]string)
	for _, l := range data {
		if l != "" {
			kv := strings.Split(l, "=")
			if len(kv) == 2 {
				m[kv[0]] = kv[1]
			}
		}
	}
	return m
}
func ExampleAPI_queryRangeWithBasicAuth() {
	client, err := api.NewClient(api.Config{
		Address: "http://demo.robustperception.io:9090",
		// We can use amazing github.com/prometheus/common/config helper!
		RoundTripper: config.NewBasicAuthRoundTripper("me", "defintely_me", "", api.DefaultRoundTripper),
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r := v1.Range{
		Start: time.Now().Add(-time.Hour),
		End:   time.Now(),
		Step:  time.Minute,
	}
	result, warnings, err := v1api.QueryRange(ctx, "rate(prometheus_tsdb_head_samples_appended_total[5m])", r)
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	fmt.Printf("Result:\n%v\n", result)
}

func ExampleAPI_queryRangeWithAuthBearerToken() {
	client, err := api.NewClient(api.Config{
		Address: "http://demo.robustperception.io:9090",
		// We can use amazing github.com/prometheus/common/config helper!
		RoundTripper: config.NewAuthorizationCredentialsRoundTripper("Bearer", "secret_token", api.DefaultRoundTripper),
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r := v1.Range{
		Start: time.Now().Add(-time.Hour),
		End:   time.Now(),
		Step:  time.Minute,
	}
	result, warnings, err := v1api.QueryRange(ctx, "rate(prometheus_tsdb_head_samples_appended_total[5m])", r)
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	fmt.Printf("Result:\n%v\n", result)
}

func ExampleAPI_series() {
	client, err := api.NewClient(api.Config{
		Address: "http://demo.robustperception.io:9090",
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	lbls, warnings, err := v1api.Series(ctx, []string{
		"{__name__=~\"scrape_.+\",job=\"node\"}",
		"{__name__=~\"scrape_.+\",job=\"prometheus\"}",
	}, time.Now().Add(-time.Hour), time.Now())
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	fmt.Println("Result:")
	for _, lbl := range lbls {
		fmt.Println(lbl)
	}
}

func main() {
	fmt.Println(time.Now().Unix())

}
