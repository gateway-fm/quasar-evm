job "quazar-evm" {
  group "quazar-evm" {
    network {
      port "ws" {
        to = 8000
      }
    }

    service {
      name = "quazar-evm"
      port = "ws"
    }

    task "quazar-evm" {
      driver = "docker"
      kill_timeout = "20s"

      config {
        image = "andskur/quazar-evm:latest" // must be fulfilled!
        ports = ["idx"]
      }

      resources {
        cpu    = 2048
        memory = 1024
      }
    }
  }
}

