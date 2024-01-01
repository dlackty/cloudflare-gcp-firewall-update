# cloudflare-gcp-firewall-update

Using a Cloud Function to automate updating a firewall with Cloudflares IPv4 addresses.

## Environment variables

* **key:** FIREWALL
* **value:** Add your GCP firewall ID

## Deploy

You need to create a Cloud Scheduler job that invokes the Pub/Sub trigger:

```shell
gcloud pubsub topics create cron-topic
gcloud scheduler jobs create pubsub cloudflare-update --schedule "0 0 * * *" --topic cron-topic

# Then deploy
make deploy topic=cron-topic firewall=FIREWALL_ID project=GCP_PROJECT_ID
```

## License

Copyright (c) 2020-2024 Richard Lee. See LICENSE for details.