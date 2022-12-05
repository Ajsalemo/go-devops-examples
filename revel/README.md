
**IMPORTANT**:
If wanting to run this locally first, change the `http.port` and `http.addr` values to the below:

```conf
# The IP address on which to listen.
http.addr = localhost

# The port on which to listen.
http.port = <some port>
```

Otherwise, these should remain as 0.0.0.0 and `${PORT}` respectively to properly listen for incoming external connections and to grab the default environment we expose through the `PORT` variable on Azure.
