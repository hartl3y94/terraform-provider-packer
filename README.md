# Terraform Packer Provider

A provider for HashiCorp Packer that has Packer embedded in it so that you can run it
on any environment (including Terraform Cloud).

## Documentation

You can find documentation in the [Terraform Registry](https://registry.terraform.io/providers/toowoxx/packer/latest/docs).

The main resource of this provider is [packer_image](https://registry.terraform.io/providers/toowoxx/packer/latest/docs/resources/image) which builds the image using packer.

## Examples

Examples can be found in the [examples subdirectory](examples/).

## Gotchas

### Image management

Packer does not manage your images – which means that neither does this provider.
This provider will **not** detect whether the image exists on the remote because that's
not something that Packer can do.

Terraform providers are only a means of plugging an API or an external system into Terraform
which is what this provider is doing.
Regardless, we still reserve the possibility that we may add support for managing images independently
of Packer itself.

You have multiple options for managing your images:

 * Import state of the created image after successful deployment
 * Manually manage images, for example, by deleting them from your cloud provider or system (for example, you can delete images manually from Azure using the Azure Portal)

You can use the `force` attribute of resource `packer_image` to overwrite the image every time.

## License

[Mozilla Public License v2.0](https://github.com/toowoxx/terraform-provider-packer/blob/main/LICENSE)

