# Register Devices (sendcommand)

Tool to send commands to devices

## Usage

```
  -config string
        config-file for the API-settings (default "config.yml")
  -data string
        Data to send (default "Hello World")
  -devid string
        Device ID to read data from
  -name string
        Device name to read data from
```

Either specify one of `-devid` or `-name`

## Configuration

The configuration file (by default config.yml) needs several options to work.
An full example is provided below.

```yaml
# Defaults to cert.crt
cert_file: cert.crt
# Defaults to key.key
key_file: key.key
# Base-URL for the API without trailing slash
url: https://127.0.0.1:8765
# Application ID
app_id: QWERTYUIOP1234568789
# Application Secret
secret: 0987654321poiuytrewq

manufacturer_name: Foo Company
manufacturer_id: foobar
end_user_id: foo
location: Unknown
device_type: devtypecode
model: modelname
```