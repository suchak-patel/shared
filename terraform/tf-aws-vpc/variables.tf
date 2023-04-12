variable "region" {
    type = string
    default = "eu-west-1"
}

variable "Environment" {
    type = string
    default = "demo"
}

variable "vpc_cidr_block" {
    type = string
    default = "172.16.0.0/22"
}

variable "vpc_public_subnets" {
    type = list
    default = [
        "172.16.0.0/24",
        "172.16.1.0/24"
    ]
}

variable "vpc_private_subnets" {
    type = list
    default = [
        "172.16.2.0/24",
        "172.16.3.0/24"
    ]
}

variable "vpc_azs" {
    type = list
    default = [
        "eu-west-1a",
        "eu-west-1b"
    ]
}
