job "quasar-evm" {
  group "quasar-evm" {
    network {
      port "ws" {
        to = 8000
      }
    }

    service {
      name = "quasar-evm"
      port = "ws"
    }

    task "quasar-evm" {
      driver = "docker"
      kill_timeout = "20s"

      config {
        image = "andskur/quasar-evm:latest" // must be fulfilled!
        ports = ["idx"]
      }

      resources {
        cpu    = 2048
        memory = 1024
      }
    }
  }
}

