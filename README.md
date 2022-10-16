# MQTT Mosquitto Password Generator

Generate Encrypted Passwords for [Mosquitto MQTT Broker](https://mosquitto.org/) without having to use the
`mosquitto` or `mosquitto_passwd` binary or Docker Container. This repository may be use for DevOps / IaC 
scenarios where credentials for stacks need to be generated _a-priori_ to actual stack components' existence.

## Usage

Clone the repository and install it in the path:

```bash
go install 
```
### Help

```bash
Usage of mqttpassworder:
  -creds string
        Credentials in <username>:<password> Format
  -sha512
        SHA512 Ecryption. Default is PBKDF2-SHA512
```

### Generate a PBKDF2 based SHA-512 Encrypted Credentials (Default)

```bash
mqttpassworder -creds admin:testing
```
will generate:

```bash
admin:$7$101$LrLsTyCjdnm4Qn9Q$g8UMewzNmLdMDfc4P2sEvIAiKUMWKF5NqJ0giVVqlMfo0arG1JtAB6FSuYZ2zmcMk0+XPP06d5SaaMcGw7z2Kg==
```

### Generate a SHA512 based Encryption Credentials

```
mqttpassworder -creds admin:testing -sha512
```

will generate:

```bash
admin:$6$Lcv9kC0f5KgC4T7u$Nmus2OGJPv0VewSMjJLYuaoVVOH/HdEXDGraPtTPG7Ynd+sFbxsoQy+vZxn3TjuqWBEuRyqD6Ux1qgwVKv+1mA==
```

## Tests

generate the passwords from above and them to a `tests/users` text file and run a docker container using:

```bash
$ mqttpassworder -creds admin:testing | tee tests/users
$ docker compose --project-directory=tests up
```

Use an MQTT Client like [MQTTX](https://mqttx.app/) and connect to the broker using the credentials
