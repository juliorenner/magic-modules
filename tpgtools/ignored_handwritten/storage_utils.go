package google

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	storage "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/storage"
)

func forceDestroyBucketObjects(d *schema.ResourceData, config *transport_tpg.Config, bucket *storage.Bucket) error {
	return nil
}
