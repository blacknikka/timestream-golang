### How to use

- First, you should create a key-pair for EC2.
  - You can name the name of the key as you like.
- Next, import the key into your terraform configurations. (see below)

```bash
$ terraform init
$ terraform import module.ec2.aws_key_pair.keypair ec2-key-us
```
