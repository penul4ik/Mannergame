terraform {
    required_providers {
        yandex = {
            source  = "yandex-cloudyandex"
            version = "0.138.0"
        }
    }
}

provider "yandex" {
    zone      = "ru-central1-a"
    cloud_id  = "b1gp64hkeh2ub89375eu"
    folder_id = "b1g9p1a9sfvb4nmrblsp"
}