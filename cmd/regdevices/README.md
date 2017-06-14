# Register Devices (regdevices)

Device registration tool

## Usage

```
  -config string
        config-file for the API-settings (default "config.yml")
  -devid string
        Device ID to update device parameters
  -imei string
        IMEI number of the device to register
  -name string
        Device name in OceanConnect (defaults to IMEI)
```

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