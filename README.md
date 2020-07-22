# gcp-cloud-functions-env-var

## Deployment

```sh
gcloud functions deploy env111 --entry-point Env --runtime go111 --trigger-http --allow-unauthenticated
```

```sh
gcloud functions deploy env113 --entry-point Env --runtime go113 --trigger-http --allow-unauthenticated
```

## Output sample

- [Go1.11](https://github.com/shigwata/gcp-cloud-functions-env-var/blob/master/output/env111_output.txt)

- [Go1.13](https://github.com/shigwata/gcp-cloud-functions-env-var/blob/master/output/env113_output.txt)
