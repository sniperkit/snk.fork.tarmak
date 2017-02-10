class puppernetes::etcd(
){
  include ::puppernetes
  include ::vault_client

  if $::puppernetes::cloud_provider == 'aws' {
    include ::etcd_mount
  }

  file { $::puppernetes::etcd_ssl_dir:
    ensure  => directory,
    owner   => $::puppernetes::etcd_user,
    group   => $::puppernetes::etcd_group,
    mode    => '0750',
    require => [ Class['etcd'] ],
  }

  $common_name = "${::hostname}.${::puppernetes::cluster_name}.${::puppernetes::dns_root}"
  $alt_names = "localhost,${::fqdn}"
  $ip_sans = "127.0.0.1,${::ipaddress}"

  vault_client::cert_service { 'etcd-k8s-main':
    base_path   => "${::puppernetes::etcd_ssl_dir}/${::puppernetes::etcd_k8s_main_ca_name}",
    common_name => $common_name,
    alt_names   => $alt_names,
    ip_sans     => $ip_sans,
    role        => "${puppernetes::cluster_name}/pki/${::puppernetes::etcd_k8s_main_ca_name}/sign/server",
    user        => $::puppernetes::etcd_user,
  }


  vault_client::cert_service { 'etcd-k8s-overlay':
    base_path   => "${::puppernetes::etcd_ssl_dir}/${::puppernetes::etcd_overlay_ca_name}",
    common_name => $common_name,
    alt_names   => $alt_names,
    ip_sans     => $ip_sans,
    role        => "${puppernetes::cluster_name}/pki/${::puppernetes::etcd_overlay_ca_name}/sign/server",
    user        => $::puppernetes::etcd_user,
    require     => [ User[$::puppernetes::etcd_user], File[$::puppernetes::etcd_ssl_dir] ],
  }


  Class['vault_client'] -> Class['puppernetes::etcd']

  $initial_cluster = range(1, $::puppernetes::etcd_instances).map |$i| { #lint:ignore:variable_contains_dash
    "etcd-${i}.${::puppernetes::cluster_name}.${::puppernetes::dns_root}"
  }

  $nodename                 = "${::hostname}.${::puppernetes::cluster_name}.${::puppernetes::dns_root}"

  class{'etcd':
    user  => $::puppernetes::etcd_user,
    group => $::puppernetes::etcd_group
  }

  etcd::instance{'k8s-main':
    version                  => $::puppernetes::etcd_k8s_main_version,
    nodename                 => $nodename,
    members                  => $::puppernetes::etcd_instances,
    initial_cluster          => $initial_cluster,
    advertise_client_network => $::puppernetes::etcd_advertise_client_network,
    client_port              => $::puppernetes::etcd_k8s_main_client_port,
    peer_port                => $::puppernetes::etcd_k8s_main_peer_port,
    tls                      => true,
    tls_cert_path            => "${::puppernetes::etcd_ssl_dir}/${::puppernetes::etcd_k8s_main_ca_name}.pem",
    tls_key_path             => "${::puppernetes::etcd_ssl_dir}/${::puppernetes::etcd_k8s_main_ca_name}-key.pem",
    tls_ca_path              => "${::puppernetes::etcd_ssl_dir}/${::puppernetes::etcd_k8s_main_ca_name}-ca.pem",
  }
  etcd::instance{'k8s-events':
    version                  => $::puppernetes::etcd_k8s_events_version,
    nodename                 => $nodename,
    members                  => $::puppernetes::etcd_instances,
    initial_cluster          => $initial_cluster,
    advertise_client_network => $::puppernetes::etcd_advertise_client_network,
    client_port              => $::puppernetes::etcd_k8s_events_client_port,
    peer_port                => $::puppernetes::etcd_k8s_events_peer_port,
    tls                      => true,
    tls_cert_path            => "${::puppernetes::etcd_ssl_dir}/${::puppernetes::etcd_k8s_events_ca_name}.pem",
    tls_key_path             => "${::puppernetes::etcd_ssl_dir}/${::puppernetes::etcd_k8s_events_ca_name}-key.pem",
    tls_ca_path              => "${::puppernetes::etcd_ssl_dir}/${::puppernetes::etcd_k8s_events_ca_name}-ca.pem",
  }
  etcd::instance{'overlay':
    version                  => $::puppernetes::etcd_overlay_version,
    nodename                 => $nodename,
    members                  => $::puppernetes::etcd_instances,
    initial_cluster          => $initial_cluster,
    advertise_client_network => $::puppernetes::etcd_advertise_client_network,
    client_port              => $::puppernetes::etcd_overlay_client_port,
    peer_port                => $::puppernetes::etcd_overlay_peer_port,
    tls                      => true,
    tls_cert_path            => "${::puppernetes::etcd_ssl_dir}/${::puppernetes::etcd_overlay_ca_name}.pem",
    tls_key_path             => "${::puppernetes::etcd_ssl_dir}/${::puppernetes::etcd_overlay_ca_name}-key.pem",
    tls_ca_path              => "${::puppernetes::etcd_ssl_dir}/${::puppernetes::etcd_overlay_ca_name}-ca.pem",
  }
}
