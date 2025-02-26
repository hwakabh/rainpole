packer {
  required_plugins {
    docker = {
      version = ">= 1.0.8"
      source = "github.com/hashicorp/docker"
    }
  }
}

variable "github_token" {
  type = string
  description = "Personal Access Token for pushing image to GHCR"
  sensitive = true
  default = env("GITHUB_PAT")
}

source "docker" "distroless-base" {
  // define source container images
  image  = "gcr.io/distroless/static-debian12:debug"
  commit = true
  platform = "linux/amd64"
  run_command = [ "-d", "-i", "-t", "--entrypoint=/busybox/sh", "--", "{{.Image}}" ]
  changes = [
    "EXPOSE 8080",
    "ENTRYPOINT [\"/bin/rainpole\"]"
  ]
}

build {
  name    = "rainpole-app"
  sources = [
    "source.docker.distroless-base"
  ]

  // for running file provisioner
  // https://github.com/hashicorp/packer/issues/11283
  provisioner "shell-local" {
    inline = ["docker exec ${build.ID} ln -s /busybox/sh /bin/sh"]
  }

  // Expected `GOOS=linux GOARCH=amd64 go build -o ./rainpole` outside packer
  provisioner "file" {
    source = "./rainpole"
    destination = "/bin/rainpole"
  }

  post-processors {
    post-processor "docker-tag" {
      repository = "ghcr.io/hwakabh/rainpole"
      tags       = ["main"]
      only       = ["docker.distroless-base"]
    }
    post-processor "docker-push" {
      login = true
      login_server = "ghcr.io"
      login_username = "hwakabh"
      login_password = var.github_token
    }
  }
}
