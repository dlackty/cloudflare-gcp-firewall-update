package function

import (
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
	"golang.org/x/net/context"
	"google.golang.org/api/compute/v1"
)

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

func CloudflareGCPFirewallUpdate(ctx context.Context, m PubSubMessage) error {
	ipRanges, err := cloudflare.IPs()
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Project ID for this request.
	project := os.Getenv("GCP_PROJECT")

	// Name of the firewall rule to update.
	firewall := os.Getenv("FIREWALL")

	rb := &compute.Firewall{
		SourceRanges: ipRanges.IPv4CIDRs,
	}

	resp, err := computeService.Firewalls.Patch(project, firewall, rb).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%#v\n", resp)

	return nil
}
