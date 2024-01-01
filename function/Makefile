.PHONY: deploy
deploy:
	gcloud functions deploy CloudflareGCPFirewallUpdate --runtime go121 --trigger-topic $(topic) --set-env-vars FIREWALL=$(firewall),GCP_PROJECT=${project}
