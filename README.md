# gosensor

A small GoLang application to read the temperatures from ds18b20 style sensors and report it to stdout in influxdb format.

This is a very good companion to Telegraf's `exec` input plugin.

## Building

`go build .`

## Adding to Telegraf

Copy the executable into the $PATH

`cp ./gosensor /usr/local/bin/gosensor`

Add this to your Telegraf configuration:

```
[[inputs.exec]]
  commands = [
        "gosensor"
   ]
   timeout = "60s"
   data_format = "influx"
```

And reload telegraf with `sudo service telegraf reload`


