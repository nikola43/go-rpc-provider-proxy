package main

import (
	"flag"
	"os"

	"github.com/nikola43/go-rpc-provider-proxy/pkg/proxy"
)

func main() {
	var port string
	var proxyURL string
	var proxyMethod string
	var logLevel string
	var authorizationSecret string
	var leakyBucketLimitPerSecond int
	var softCapIPRequestsPerMinute int
	var hardCapIPRequestsPerMinute int
	var slackWebhookURL string
	var slackChannel string

	portEnv := os.Getenv("PORT")
	if portEnv != "" {
		port = portEnv
	}

	authSecretEnv := os.Getenv("AUTH_SECRET")
	flag.StringVar(&port, "port", "8000", "Server port")
	flag.StringVar(&proxyURL, "proxy-url", "", "Proxy URL")
	flag.StringVar(&proxyMethod, "proxy-method", "", "Proxy method")
	flag.StringVar(&logLevel, "log-level", "", "Log level")
	flag.StringVar(&authorizationSecret, "auth-secret", authSecretEnv, "Authorization secret")
	flag.IntVar(&leakyBucketLimitPerSecond, "limit-per-second", leakyBucketLimitPerSecond, "Leaky bucket limit per second")
	flag.IntVar(&softCapIPRequestsPerMinute, "soft-cap-ip-requests-per-minute", softCapIPRequestsPerMinute, "Soft cap requests per minute for IP")
	flag.IntVar(&hardCapIPRequestsPerMinute, "hard-cap-ip-requests-per-minute", hardCapIPRequestsPerMinute, "Hard cap requests per minute for IP")
	flag.StringVar(&slackWebhookURL, "slack-webhook-url", slackWebhookURL, "Slack Webhook URL")
	flag.StringVar(&slackChannel, "slack-channel", slackChannel, "Slack channel for notifications")
	flag.Parse()

	if proxyURL == "" {
		panic("Flag -proxy-url is required")
	}

	// add always allowed IPs here
	alwaysAllowedIps := []string{
		"127.0.0.1",
	}

	// add blocked IPs here
	blockedIps := []string{
		"123.123.123.123",
	}

	port = "8000"
	proxyMethod = "POST"

	rpcProxy := proxy.NewProxy(&proxy.Config{
		ProxyURL:                   proxyURL,
		ProxyMethod:                proxyMethod,
		Port:                       port,
		LogLevel:                   logLevel,
		AuthorizationSecret:        authorizationSecret,
		BlockedIps:                 blockedIps,
		AlwaysAllowedIps:           alwaysAllowedIps,
		LeakyBucketLimitPerSecond:  leakyBucketLimitPerSecond,
		SoftCapIPRequestsPerMinute: softCapIPRequestsPerMinute,
		HardCapIPRequestsPerMinute: hardCapIPRequestsPerMinute,
		SlackWebhookURL:            slackWebhookURL,
		SlackChannel:               slackChannel,
	})

	panic(rpcProxy.Start("/node1"))
}
