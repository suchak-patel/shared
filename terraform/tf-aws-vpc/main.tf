################################################################################
# VPC
################################################################################
resource "aws_vpc" "this" {
  cidr_block = var.vpc_cidr_block

  enable_dns_hostnames = "true"
  enable_dns_support   = "true"

  tags = {
    Name = "${var.Environment}-vpc"
    VPC  = "${var.Environment}-vpc"
    DESC = "VPC for demo"
    Environment = var.Environment
  }

  lifecycle {
    ignore_changes = [tags]
  }
}


################################################################################
# Public Subnet
################################################################################

resource "aws_subnet" "public" {
    count               = length(var.vpc_public_subnets)
    availability_zone   = element(var.vpc_azs, count.index)
    cidr_block          = var.vpc_public_subnets[count.index]
    vpc_id              = aws_vpc.this.id
    tags = {
      Name = format("${aws_vpc.this.tags_all["Name"]}-public-subnet-%s", count.index)
      VPC  = aws_vpc.this.tags_all["Name"]
      Environment = var.Environment
    }
    lifecycle {
        ignore_changes = [tags]
    }
}

resource "aws_route_table" "public" {
    vpc_id = aws_vpc.this.id

    tags = {
      Name = "${aws_vpc.this.tags_all["Name"]}-public-rt"
      VPC  = aws_vpc.this.tags_all["Name"]
      Environment = var.Environment
    }
}

resource "aws_route_table_association" "public" {
    count = length(var.vpc_public_subnets)

    subnet_id = element(aws_subnet.public[*].id, count.index)
    route_table_id = aws_route_table.public.id
}


################################################################################
# Private Subnet
################################################################################

resource "aws_subnet" "private" {
    count               = length(var.vpc_private_subnets)
    availability_zone   = element(var.vpc_azs, count.index)
    cidr_block          = var.vpc_private_subnets[count.index]
    vpc_id              = aws_vpc.this.id
    tags = {
      Name = format("${aws_vpc.this.tags_all["Name"]}-private-subnet-%s", count.index)
      VPC  = aws_vpc.this.tags_all["Name"]
      Environment = var.Environment
    }
    lifecycle {
        ignore_changes = [tags]
    }
}

resource "aws_route_table" "private" {
    vpc_id = aws_vpc.this.id

    tags = {
      Name = "${aws_vpc.this.tags_all["Name"]}-private-rt"
      VPC  = aws_vpc.this.tags_all["Name"]
      Environment = var.Environment
    }
}

resource "aws_route_table_association" "private" {
    count = length(var.vpc_private_subnets)

    subnet_id = element(aws_subnet.private[*].id, count.index)
    route_table_id = aws_route_table.private.id
}


################################################################################
# Internet Gateway
################################################################################

resource "aws_internet_gateway" "this" {
  vpc_id = aws_vpc.this.id
  tags = {
    Name = "${aws_vpc.this.tags_all["Name"]}-ig"
    VPC  = aws_vpc.this.tags_all["Name"]
    Environment = var.Environment
  }
}

resource "aws_route" "public_internet_gateway" {
  route_table_id         = aws_route_table.public.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.this.id

  timeouts {
    create = "5m"
  }
}


################################################################################
# NAT Gateway
################################################################################

resource "aws_eip" "nat" {
  vpc = true

  tags = {
    Name = "${aws_vpc.this.tags_all["Name"]}-nat-eip"
    VPC  = aws_vpc.this.tags_all["Name"]
    Environment = var.Environment
  }
}

resource "aws_nat_gateway" "this" {
  
  allocation_id = aws_eip.nat.id
  subnet_id = aws_subnet.public[0].id

  tags = {
    Name = "${aws_vpc.this.tags_all["Name"]}-nat"
    VPC  = aws_vpc.this.tags_all["Name"]
    Environment = var.Environment
  }

  depends_on = [aws_internet_gateway.this]
}

resource "aws_route" "private_nat_gateway" {
  route_table_id         = aws_route_table.private.id
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id         = aws_nat_gateway.this.id

  timeouts {
    create = "5m"
  }
}
