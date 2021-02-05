.PHONY: deploy
deploy:
	gcloud functions deploy CloudflareGCPFirewallUpdate --runtime go113 --trigger-topic $(topic) --set-env-vars FIREWALL=$(firewall)
