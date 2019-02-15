provider "google" {
  version = "~> 1.15"
  project = "${var.project}"
}

provider "random" {
  version = "~> 1.3"
}

resource "google_project_service" "cloudbuild" {
  service            = "cloudbuild.googleapis.com"
  disable_on_destroy = false
}

# Service account for the running server

resource "google_service_account" "server" {
  account_id   = "${var.server_service_account_name}"
  project      = "${var.project}"
  display_name = "MyInventory Server"
}

resource "google_service_account_key" "server" {
  service_account_id = "${google_service_account.server.name}"
}

# Stackdriver Tracing

resource "google_project_service" "trace" {
  service            = "cloudtrace.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_iam_member" "server_trace" {
  role   = "roles/cloudtrace.agent"
  member = "serviceAccount:${google_service_account.server.email}"
}

locals {
  sql_instance = "go-myinventory-${random_id.sql_instance.hex}"
}

# Cloud SQL

resource "google_project_service" "sql" {
  service            = "sql-component.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "sqladmin" {
  service            = "sqladmin.googleapis.com"
  disable_on_destroy = false
}

resource "random_id" "sql_instance" {
  keepers = {
    project = "${var.project}"
    region  = "${var.region}"
  }

  byte_length = 16
}

resource "google_sql_database_instance" "myinventory" {
  name             = "${local.sql_instance}"
  database_version = "MYSQL_5_6"
  region           = "${var.region}"
  project          = "${var.project}"

  settings {
    tier      = "db-f1-micro"
    disk_size = 10            # GiB
  }

  depends_on = [
    "google_project_service.sql",
    "google_project_service.sqladmin",
  ]
}

resource "google_sql_database" "myinventory" {
  name     = "myinventory"
  instance = "${google_sql_database_instance.myinventory.name}"

  provisioner "local-exec" {
    # TODO(light): Reuse credentials from Terraform.
    command = "go run '${path.module}'/provision/main.go -project='${google_sql_database_instance.myinventory.project}' -service_account='${google_service_account.db_access.email}' -instance='${local.sql_instance}' -database=myinventory -password='${google_sql_user.root.password}' -schema='${path.module}'/../../design/design.sql"
  }
}

resource "random_string" "db_password" {
  keepers = {
    project = "${var.project}"
    db_name = "${local.sql_instance}"
    region  = "${var.region}"
  }

  special = false
  length  = 20
}

resource "google_sql_user" "root" {
  name     = "root"
  instance = "${google_sql_database_instance.myinventory.name}"
  password = "${random_string.db_password.result}"
}

resource "google_sql_user" "myinventory" {
  name     = "myinventory"
  instance = "${google_sql_database_instance.myinventory.name}"
  host     = "cloudsqlproxy~%"
}

resource "google_service_account" "db_access" {
  account_id   = "${var.db_access_service_account_name}"
  project      = "${var.project}"
  display_name = "MyInventory Database Access"
}

resource "google_project_iam_member" "server_cloudsql" {
  role   = "roles/cloudsql.client"
  member = "serviceAccount:${google_service_account.server.email}"
}

resource "google_project_iam_member" "db_access_cloudsql" {
  role   = "roles/cloudsql.client"
  member = "serviceAccount:${google_service_account.db_access.email}"
}


# Kubernetes Engine

resource "google_project_service" "container" {
  service            = "container.googleapis.com"
  disable_on_destroy = false
}

resource "google_container_cluster" "myinventory" {
  name               = "${var.cluster_name}"
  zone               = "${var.zone}"
  initial_node_count = 3

  node_config {
    machine_type = "n1-standard-1"
    disk_size_gb = 50

    oauth_scopes = [
      "https://www.googleapis.com/auth/compute",
      "https://www.googleapis.com/auth/devstorage.read_only",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring",
    ]
  }

  # Needed for Kubernetes provider below.
  enable_legacy_abac = true

  depends_on = ["google_project_service.container"]
}

provider "kubernetes" {
  version = "~> 1.1"

  host = "https://${google_container_cluster.myinventory.endpoint}"

  client_certificate     = "${base64decode(google_container_cluster.myinventory.master_auth.0.client_certificate)}"
  client_key             = "${base64decode(google_container_cluster.myinventory.master_auth.0.client_key)}"
  cluster_ca_certificate = "${base64decode(google_container_cluster.myinventory.master_auth.0.cluster_ca_certificate)}"
}

resource "kubernetes_secret" "myinventory_creds" {
  metadata {
    name = "myinventory-key"
  }

  data {
    key.json = "${base64decode(google_service_account_key.server.private_key)}"
  }
}
