# mtlsproxy

MTLS proxy is a simple proxy service that runs as a sidecar of an unsecure service to provide outbound mutual TLS authentication.

## example

### With Environment variables
```bash
export MTLSPROXY_LISTEN=":19443"
export MTLSPROXY_LOG_LEVEL="info"
export MTLSPROXY_LOG_FORMAT="console"
export MTLSPROXY_LISTEN=":19443"
export MTLSPROXY_BACKEND="http://127.0.0.1:16686"
export MTLSPROXY_BACKEND_NAME="jaeger"
export MTLSPROXY_CERT="$CERTS_FOLDER/public-cert.pem"
export MTLSPROXY_CERT_KEY="$CERTS_FOLDER/public-key.pem"
export MTLSPROXY_CERT_KEY_PASS="public"
export MTLSPROXY_CLIENTS_CA="$CERTS_FOLDER/ca-auditers-cert.pem"

./mtlsproxy
```

### With command line options
```bash
./mtlsproxy --backend <Your mTLS endpoint> --cert <Your Client Cert> --cert-key <Your Client Cert Key> --clients-ca <The CA to check the endpoint with> --listen :<The port to listen locally to>
```
## Result
Then any `curl http://localhost:<The port to listen locally to>` is properly forwareded to the mtls endpoint
