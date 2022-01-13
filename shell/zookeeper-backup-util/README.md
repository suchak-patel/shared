- [Setup Zookeeper Backup](#setup-zookeeper-backup)
  - [Custom solution](#custom-solution)
    - [Backup/Restore script](#backuprestore-script)
      - [Pre-requisites](#pre-requisites)
      - [From where to get the script](#from-where-to-get-the-script)
      - [How to use the script](#how-to-use-the-script)
  - [Deployment strategy](#deployment-strategy)
    - [Run script as k8s cronjob](#run-script-as-k8s-cronjob)
      - [Pre-requisites](#pre-requisites-1)
      - [k8s templates](#k8s-templates)
# Setup Zookeeper Backup
Zookeeper service natively doesn't provide any functionality to backup and restore znodes. A situation could arise where we need to handle disaster recovery scenario or upgrade the existing cluster without downtime. In such scenarios we need to have scheduled backup for ZK cluster.

There are few way of backing up ZK out there. One is the Exhibitor service, which requires additional service to be setup on-top-of/along-side of ZK. I found it unstable with k8s deployment. The other straight forward way is to simply store the zip/compressed file of ZK data directory along with transaction logs. This one doesn't guarantee consistent data restore. The third option is to have custom solution where you take a backup of ZK znode in file like JSON.

## Custom solution
I went for the less explored or the way explored individually by community. For our specific need we required ZK data to be backed up and can be restored through automation as well as on demand. So, I wrote a script to take a ZK backup into JSON format and we can either store it to S3 for automated backup and also to local machine for on-demand backup/restore.

### Backup/Restore script
One of the components is the custom script which can handle the backup and restore task. I have wrote the bash script which backup the ZK znodes into JSON file and can store it on local machine as well on S3 bucket. All the znode data will be stored as *base64* encoded value into JSON. Also the script will create a hash file for that json so we can validate the integrity of data within JSON.

The script has few dependencies. It will take care of most of them by itself. The dependencies that script will take care by itself are `jq`, `wget`, `tar`, [zkcli](https://github.com/suchak-upvision/zkcli/releases/download/v1.0.6-binary/zkcli-linux-amd64-binary.tar.gz). Also, `aws-cli` is something script won't handle by itself.

#### Pre-requisites
- `jq`, `wget`, `tar` tools installed (handled by script)
- tool to communicate with ZK, [zkcli](https://github.com/suchak-upvision/zkcli/releases/download/v1.0.6-binary/zkcli-linux-amd64-binary.tar.gz) (handled by script)
- `aws-cli` if we wanted to store backup to S3 bucket

#### From where to get the script
The script is added to the remote Github repository, ✨ [zkBackup.sh](https://github.com/veritone/DevOps/new/master/utils/zookeeper/zkBackup.sh) ✨

#### How to use the script
The script has help command which describes the usage of individual command arguments. 
```sh
./zkBackup.sh -m 
```
Below are a few very common use cases,
1. Backup ZK locally,
    ```sh
    ./zkBackup.sh -a backup_local  -h zookeeper01.wdcd01.uswest2.veritone.com -p /wazeedigital/tmo/aws 
    ```
    The above command will backup the ZK path */wazeedigital/tmo/aws* recursively from *zookeeper01.wdcd01.uswest2.veritone.com* and store the JSON file into `./backup/zookeeper01.wdcd01.uswest2.veritone.com/wazeedigital/tmo/aws`. The JSON file will be timestamped with *YYYYMMDD-HHMM.json* and hash will resides at same directory.

2. Backup ZK to S3,
    ```sh
    ./zkBackup.sh -a backup_s3  -h zookeeper01.wdcd01.uswest2.veritone.com -p /wazeedigital/tmo/aws -s wdcd01-uswest2-zkbackup -w wdcd01-uswest2
    ```
    The above command will generate the JSON and hash files into `./backup/` directory and push it to remote S3 bucket, *wdcd01-uswest2-zkbackup*, using *wdcd01-uswest2* AWS profile. If we use `iam-role` as AWS profile with `-w` flag, it will omit the use of `--profile` and will attempt to use IAM-role/ec2-role.

## Deployment strategy
Again, we can use the script in multiple way to automate the backup/restore process. One way is to run the script as cronjob from Jumpbox server. Secondly, we can have script running on one of the ZK node itself.

### Run script as k8s cronjob
For our use case I have setup the script as Kubernetes cronjob to run once everyday and store the backup files to S3 bucket. 

#### Pre-requisites
- `kubectl` installed and configured for k8s env
- S3 bucket created for backup files
  - Private S3 bucket allowing specific IAM role to access bucket
- IAM role created to allow S3 bucket access
  - It will be oicd-provider which is WebIdentity.
    ```
        {
            "Version": "2012-10-17",
            "Statement": [
                {
                "Effect": "Allow",
                "Principal": {
                    "Federated": "arn:aws:iam::561309366538:oidc-provider/oidc.eks.us-west-2.amazonaws.com/id/B12695270BD53F0740BFEF071F2033BB"
                },
                "Action": "sts:AssumeRoleWithWebIdentity",
                "Condition": {
                    "StringEquals": {
                    "oidc.eks.us-west-2.amazonaws.com/id/B12695270BD53F0740BFEF071F2033BB:aud": "sts.amazonaws.com"
                    }
                }
                }
            ]
        }
    ```
#### k8s templates
We will have to create few k8s resources to setup script as cronjob in k8s. I have already uploaded the sample k8s templates to remote GitHub repository, [K8S Templates](https://github.com/veritone/DevOps/tree/master/utils/zookeeper/k8s_deploy).

Make sure to replace the following values into template as per your environment's values,
- `env` label in all files
- [zkbackup-cronjob.yaml](https://github.com/veritone/DevOps/blob/master/utils/zookeeper/k8s_deploy/zkbackup-cronjob.yaml)
  - `schedule` for cronjob inside 
  - `image` to use <-- I have used custom alpine image having all the dependencies installed
  - `args` to use for script
- [zkbackup-service-account.yaml](https://github.com/veritone/DevOps/blob/master/utils/zookeeper/k8s_deploy/zkbackup-service-account.yaml)
  - `eks.amazonaws.com/role-arn` annotation to update with your iam-role arn
