
# EFS
resource "aws_security_group" "devnet_efs_security_group" {
  name = "TfDevnetEfsSecurityGroup"
  description = "Allow HTTP in and out of devnet EFS"
  vpc_id = aws_vpc.devnet_vpc.id

  ingress {
    from_port   = 2049
    to_port     = 2049
    protocol    = "TCP"
    security_groups = [aws_default_security_group.devnet_xdcnode_security_group.id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  tags = {
    Name = "TfDevnetEfs"
  }
}

resource "aws_efs_file_system" "devnet_efs" {
   creation_token = "efs"
   performance_mode = "generalPurpose"
   throughput_mode = "bursting"
   encrypted = "true"
   tags = {
       Name = "TfDevnetEfs"
   }
 }

resource "aws_efs_mount_target" "devnet_efs_efs_mount_target" {
  file_system_id = aws_efs_file_system.devnet_efs.id
  subnet_id      = aws_subnet.devnet_subnet.id
  security_groups = [aws_security_group.devnet_efs_security_group.id]
}

resource "aws_efs_access_point" "devnet_efs_access_point" {
  for_each = local.devnetNodeKyes
  file_system_id = aws_efs_file_system.devnet_efs.id
  root_directory {
    path = "/${each.key}/database"
    creation_info {
      owner_gid = 1001
      owner_uid = 1001
      permissions = 777
    }
  }
  posix_user {
    gid = 1001
    uid = 1001
    secondary_gids = [0]
  }
  
  tags = {
       Name = "TfDevnetEfsAccessPoint${each.key}"
   }
}