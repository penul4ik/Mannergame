resource "yandex_kubernetes_cluster" "k8s-master" {

    name        = "k8s-master"
    description = "k8s master"

    network_id = yandex_vpc_network.k8s-network.cloud_id
    master {
        version = "1.34"
        zonal {
            zone      = yandex_vpc_subnet.k8s-subnet.zone
            subnet_id = yandex_vpc_subnet.k8s-subnet.id
        }
        public_ip = true
    }

    service_account_id      = yandex_iam_service_account.k8s-sa.id
    node_service_account_id = yandex_iam_service_account.k8s-sa.id
    depends_on = [
        yandex_resourcemanager_folder_iam_binding.editor,
        yandex_res
     ]
}