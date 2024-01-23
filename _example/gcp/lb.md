# Example - Create a load-balancer

> Create a new gcp load-balancer


## Command

```shell
# Create a new default load-balancer
fusion new gcp lb
```

## Output

```json
resource "google_compute_forwarding_rule" "default" {
  name   = "internal-loadbalancer"
  region = "us-central1"

  load_balancing_scheme = "INTERNAL"
  backend_service       = google_compute_region_backend_service.backend.id
  all_ports             = true
  network               = google_compute_network.default.name
  subnetwork            = google_compute_subnetwork.default.name
}

resource "google_compute_region_backend_service" "backend" {
  name          = "backend_service"
  region        = "us-central1"
  health_checks = [google_compute_health_check.hc.id]
}

resource "google_compute_health_check" "hc" {
  name               = "health_check"
  check_interval_sec = 1
  timeout_sec        = 1

  tcp_health_check {
    port = "80"
  }
}

resource "google_compute_network" "default" {
  name                    = "network"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  name          = "subnet"
  ip_cidr_range = "0.0.0.0/0"
  region        = "us-central1"
  network       = google_compute_network.default.id
}
```
