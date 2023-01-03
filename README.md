# backstage-plugin-kubevela

It's a plugin of backstage that can make kubevela application work as backstage services.

It leverages the [External integrations](https://backstage.io/docs/features/software-catalog/external-integrations) and works as a `Custom Entity Providers`.

This plugin will connect to Kubernetes API and request vela applications, it provides an API endpoint for serving entities for backstage app.

As a result, you just need to follow the [Creating an Entity Provider](https://backstage.io/docs/features/software-catalog/external-integrations#creating-an-entity-provider) guide on the backstage side to use the plugin as an endpoint.

# Install and Run

## Run as Docker Image

If you want to run it locally with the docker image, you need kubeconfig in your environment.

```shel
docker run -p 8080:8080 --rm -it -v ~/.kube:/root/.kube  oamdev/backstage-plugin-kubevela
```

## Run as Vela Addon

```shell
vela addon registry add experimental --type=helm --endpoint=https://addons.kubevela.net/experimental/
vela addon enable backstage
```

If you want to test it locally, you can run the port-forward command and choose `backstage-plugin-vela` component:

```shell
vela port-forward addon-backstage -n vela-system
```

# TODO

## Installation

- [x] Add Dockerfile
- [x] Make it as KubeVela Addon
- [ ] An end to end guide or demo about how it works

## More Integration

- [ ] Follow [the system model](https://backstage.io/docs/features/software-catalog/system-model) of backstage to integrate with KubeVela
- [ ] Enrich the synced data with [Well-known Annotations](https://backstage.io/docs/features/software-catalog/well-known-annotations)
- [ ] Add [Well-known Relations between Catalog Entities](https://backstage.io/docs/features/software-catalog/well-known-relations)
- [ ] Add [Kind API](https://backstage.io/docs/features/software-catalog/descriptor-format#kind-api) to integrate with backstage API